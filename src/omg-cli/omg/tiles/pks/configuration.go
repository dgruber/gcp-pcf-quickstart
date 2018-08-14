package pks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"omg-cli/config"
	"omg-cli/omg/tiles"
	"omg-cli/ops_manager"
)

const (
	skipSSLValidation = "true"
)

type Properties struct {
	PksTls         GenerateCertDomainsValue `json:".pivotal-container-service.pks_tls"`
	PKSAPIHostname tiles.Value              `json:".properties.pks_api_hostname"`
	// Configuration Plan1
	P1Selector                          tiles.Value        `json:".properties.plan1_selector"`
	P1SelectorName                      tiles.Value        `json:".properties.plan1_selector.active.name"`
	P1SelectorDescription               tiles.Value        `json:".properties.plan1_selector.active.description"`
	P1SelectorMasterAz                  tiles.AZsValue     `json:".properties.plan1_selector.active.master_az_placement"`
	P1SelectorMasterVMType              tiles.Value        `json:".properties.plan1_selector.active.master_vm_type"`
	P1SelectorMasterDiskType            tiles.Value        `json:".properties.plan1_selector.active.master_persistent_disk_type"`
	P1SelectorWorkerAz                  tiles.AZsValue     `json:".properties.plan1_selector.active.worker_az_placement"`
	P1SelectorWorkerVMType              tiles.Value        `json:".properties.plan1_selector.active.worker_vm_type"`
	P1SelectorWorkerDiskType            tiles.Value        `json:".properties.plan1_selector.active.worker_persistent_disk_type"`
	P1SelectorWorkerInstances           tiles.IntegerValue `json:".properties.plan1_selector.active.worker_instances"`
	P1SelectorErrandVMType              tiles.Value        `json:".properties.plan1_selector.active.errand_vm_type"`
	P1SelectorAddonSpec                 tiles.Value        `json:".properties.plan1_selector.active.addons_spec"`
	P1SelectorAllowPrivContainers       tiles.BooleanValue `json:".properties.plan1_selector.active.allow_privileged_containers"`
	P1SelectorDisableDenyEscalationExec tiles.BooleanValue `json:".properties.plan1_selector.active.disable_deny_escalating_exec"`

	P2Selector tiles.Value `json:".properties.plan2_selector"`
	P3Selector tiles.Value `json:".properties.plan3_selector"`

	K8sCloudProvider                    tiles.Value `json:".properties.cloud_provider"`
	K8sCloudProviderGCPProjectID        tiles.Value `json:".properties.cloud_provider.gcp.project_id"`
	K8sCloudProviderGCPNetwork          tiles.Value `json:".properties.cloud_provider.gcp.network"`
	K8sCloudProviderGCPMasterServiceKey tiles.Value `json:".properties.cloud_provider.gcp.master_service_account"`
	K8sCloudProviderGCPWorkerServiceKey tiles.Value `json:".properties.cloud_provider.gcp.worker_service_account"`

	NetworkSelector tiles.Value `json:".properties.network_selector"`

	Uaa tiles.Value `json:".properties.uaa"`
	// ldap here ...
	UaaPksCliAccessTokenLifetime  tiles.IntegerValue `json:".properties.uaa_pks_cli_access_token_lifetime"`
	UaaPksCliRefreshTokenLifetime tiles.IntegerValue `json:".properties.uaa_pks_cli_refresh_token_lifetime"`

	SyslogMigrationSelector tiles.Value `json:".properties.syslog_migration_selector"`
	TelemetrySelector       tiles.Value `json:".properties.telemetry_selector"`
}

type GenerateCertDomainsValue struct {
	GDs  []string `json:"generate_cert_domains"`
	Pems Certs    `json:"value"`
}

type Certs struct {
	CertPem       string `json:"cert_pem"`
	PrivateKeyPem string `json:"private_key_pem"`
}

type Resources struct {
	PKS tiles.Resource `json:"pivotal-container-service"`
}

func (t *Tile) Configure(envConfig *config.EnvConfig, cfg *config.Config, om *ops_manager.Sdk) error {
	certBytes, err := ioutil.ReadFile("keys/pks.crt")
	if err != nil {
		panic(err)
	}

	keyBytes, err := ioutil.ReadFile("keys/pks.key")
	if err != nil {
		panic(err)
	}

	if err := om.StageProduct(tile.Product); err != nil {
		return err
	}

	//network := tiles.NetworkODBConfig(cfg.ServicesSubnetName, cfg, cfg.DynamicServicesSubnetName)
	network := tiles.NetworkODBConfig(cfg.MgmtSubnetName, cfg, cfg.DynamicServicesSubnetName)

	networkBytes, err := json.Marshal(&network)
	if err != nil {
		return err
	}

	properties := &Properties{
		PksTls: GenerateCertDomainsValue{
			Pems: Certs{CertPem: string(certBytes), PrivateKeyPem: string(keyBytes)},
		},
		PKSAPIHostname:                      tiles.Value{fmt.Sprintf("api.pks.%s", cfg.DnsSuffix)},
		P1Selector:                          tiles.Value{"Plan Active"},
		P1SelectorName:                      tiles.Value{"small"},
		P1SelectorDescription:               tiles.Value{"Small Plan"},
		P1SelectorMasterAz:                  tiles.AZsValue{[]string{cfg.Zone1}},
		P1SelectorMasterVMType:              tiles.Value{"xlarge.mem"},
		P1SelectorMasterDiskType:            tiles.Value{"102400"},
		P1SelectorWorkerAz:                  tiles.AZsValue{[]string{cfg.Zone1, cfg.Zone2, cfg.Zone3}},
		P1SelectorWorkerVMType:              tiles.Value{"large.mem"},
		P1SelectorWorkerDiskType:            tiles.Value{"102400"},
		P1SelectorWorkerInstances:           tiles.IntegerValue{3},
		P1SelectorErrandVMType:              tiles.Value{"micro"},
		P1SelectorAddonSpec:                 tiles.Value{""},
		P1SelectorAllowPrivContainers:       tiles.BooleanValue{true},
		P1SelectorDisableDenyEscalationExec: tiles.BooleanValue{false},
		P2Selector:                          tiles.Value{"Plan Inactive"},
		P3Selector:                          tiles.Value{"Plan Inactive"},
		K8sCloudProvider:                    tiles.Value{"GCP"},
		K8sCloudProviderGCPProjectID:        tiles.Value{cfg.ProjectName},
		K8sCloudProviderGCPNetwork:          tiles.Value{cfg.NetworkName},
		K8sCloudProviderGCPMasterServiceKey: tiles.Value{cfg.PKSAccountEmail},
		K8sCloudProviderGCPWorkerServiceKey: tiles.Value{cfg.PKSAccountEmail},
		NetworkSelector:                     tiles.Value{"flannel"},
		Uaa:                                 tiles.Value{"internal"},
		UaaPksCliAccessTokenLifetime:  tiles.IntegerValue{86400},
		UaaPksCliRefreshTokenLifetime: tiles.IntegerValue{172800},
		SyslogMigrationSelector:       tiles.Value{"disabled"},
		TelemetrySelector:             tiles.Value{"disabled"},
	}

	propertiesBytes, err := json.Marshal(&properties)
	if err != nil {
		return err
	}

	resources := Resources{
		PKS: tiles.Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:pks-api")},
			InternetConnected: false,
			VmTypeId:          "xlarge.mem",
			DiskTypeId:        "102400",
		},
	}

	resourcesBytes, err := json.Marshal(&resources)
	if err != nil {
		return err
	}

	return om.ConfigureProduct(tile.Product.Name, string(networkBytes), string(propertiesBytes), string(resourcesBytes))
}
