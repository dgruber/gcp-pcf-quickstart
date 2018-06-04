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
	PksTls GenerateCertDomainsValue `json:".pivotal-container-service.pks_tls"`
	// Configuration Plan1
	P1Selector                    tiles.Value        `json:".properties.plan1_selector"`
	P1SelectorName                tiles.Value        `json:".properties.plan1_selector.active.name"`
	P1SelectorDescription         tiles.Value        `json:".properties.plan1_selector.active.description"`
	P1SelectorAz                  tiles.Value        `json:".properties.plan1_selector.active.az_placement"`
	P1SelectorAuth                tiles.Value        `json:".properties.plan1_selector.active.authorization_mode"`
	P1SelectorMasterVMType        tiles.Value        `json:".properties.plan1_selector.active.master_vm_type"`
	P1SelectorMasterDiskType      tiles.Value        `json:".properties.plan1_selector.active.master_persistent_disk_type"`
	P1SelectorWorkerVMType        tiles.Value        `json:".properties.plan1_selector.active.worker_vm_type"`
	P1SelectorWorkerDiskType      tiles.Value        `json:".properties.plan1_selector.active.persistent_disk_type"`
	P1SelectorWorkerInstances     tiles.IntegerValue `json:".properties.plan1_selector.active.worker_instances"`
	P1SelectorErrandVMType        tiles.Value        `json:".properties.plan1_selector.active.errand_vm_type"`
	P1SelectorAddonSpec           tiles.Value        `json:".properties.plan1_selector.active.addons_spec"`
	P1SelectorAllowPrivContainers tiles.BooleanValue `json:".properties.plan1_selector.active.allow_privileged_containers"`

	P2Selector tiles.Value `json:".properties.plan2_selector"`
	P3Selector tiles.Value `json:".properties.plan3_selector"`

	K8sCloudProvider                    tiles.Value `json:".properties.cloud_provider"`
	K8sCloudProviderGCPProjectID        tiles.Value `json:".properties.cloud_provider.gcp.project_id"`
	K8sCloudProviderGCPNetwork          tiles.Value `json:".properties.cloud_provider.gcp.network"`
	K8sCloudProviderGCPMasterServiceKey tiles.Value `json:".properties.cloud_provider.gcp.master_service_account_key"`
	K8sCloudProviderGCPWorkerServiceKey tiles.Value `json:".properties.cloud_provider.gcp.worker_service_account_key"`

	NetworkSelector tiles.Value `json:".properties.network_selector"`

	UaaUrl                        tiles.Value        `json:".properties.uaa_url"`
	UaaPksCliAccessTokenLifetime  tiles.IntegerValue `json:".properties.uaa_pks_cli_access_token_lifetime"`
	UaaPksCliRefreshTokenLifetime tiles.IntegerValue `json:".properties.uaa_pks_cli_refresh_token_lifetime"`

	SyslogMigrationSelector tiles.Value `json:".properties.syslog_migration_selector"`
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
		P1Selector:                          tiles.Value{"Plan Active"},
		P1SelectorName:                      tiles.Value{"small"},
		P1SelectorDescription:               tiles.Value{"Small Plan"},
		P1SelectorAz:                        tiles.Value{cfg.Zone1},
		P1SelectorAuth:                      tiles.Value{"rbac"},
		P1SelectorMasterVMType:              tiles.Value{"micro"},
		P1SelectorMasterDiskType:            tiles.Value{"10240"},
		P1SelectorWorkerVMType:              tiles.Value{"micro"},
		P1SelectorWorkerDiskType:            tiles.Value{"10240"},
		P1SelectorWorkerInstances:           tiles.IntegerValue{3},
		P1SelectorErrandVMType:              tiles.Value{"micro"},
		P1SelectorAddonSpec:                 tiles.Value{""},
		P1SelectorAllowPrivContainers:       tiles.BooleanValue{true},
		P2Selector:                          tiles.Value{"Plan Inactive"},
		P3Selector:                          tiles.Value{"Plan Inactive"},
		K8sCloudProvider:                    tiles.Value{"GCP"},
		K8sCloudProviderGCPProjectID:        tiles.Value{cfg.ProjectName},
		K8sCloudProviderGCPNetwork:          tiles.Value{cfg.NetworkName},
		K8sCloudProviderGCPMasterServiceKey: tiles.Value{cfg.PKSAccountKey},
		K8sCloudProviderGCPWorkerServiceKey: tiles.Value{cfg.PKSAccountKey},
		NetworkSelector:                     tiles.Value{"flannel"},
		UaaUrl:                              tiles.Value{fmt.Sprintf("api.pks.%s", cfg.DnsSuffix)},
		UaaPksCliAccessTokenLifetime:  tiles.IntegerValue{86400},
		UaaPksCliRefreshTokenLifetime: tiles.IntegerValue{172800},
		SyslogMigrationSelector:       tiles.Value{"disabled"},
	}

	propertiesBytes, err := json.Marshal(&properties)
	if err != nil {
		return err
	}

	resources := Resources{
		PKS: tiles.Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:pks-api")},
			InternetConnected: false,
		},
	}

	resourcesBytes, err := json.Marshal(&resources)
	if err != nil {
		return err
	}

	return om.ConfigureProduct(tile.Product.Name, string(networkBytes), string(propertiesBytes), string(resourcesBytes))
}
