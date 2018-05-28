#!/usr/bin/env bash

#
# Copyright 2017 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

set -ue
cd "$(dirname "$0")"

export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin
go install omg-cli

if [ -z ${ENV_DIR+X} ]; then
    export ENV_DIR="$PWD/env/pcf"
    echo "ENV_DIR unset, using: ${ENV_DIR}"
fi


mkdir -p ${ENV_DIR}
terraform_output="${ENV_DIR}/terraform_output.json"
terraform_config="${ENV_DIR}/terraform.tfvars"
terraform_state="${ENV_DIR}/terraform.tfstate"
env_config="${ENV_DIR}/config.json"

if [ ! -f $env_config ]; then
    if [ -z ${PIVNET_API_TOKEN+x} ]; then
        echo "PIVNET_API_TOKEN environment required (requires registration)."
        echo "Find the value for 'API TOKEN' at https://network.pivotal.io/users/dashboard/edit-profile"
        echo "and run: export PIVNET_API_TOKEN=<value of 'API TOKEN'> before running this command."
        exit 1
    fi

    if [ -z ${GCP_PROJECT+X} ]; then
        export GCP_PROJECT=$(gcloud config get-value project | xargs echo -n)
        echo "GCP_PROJECT unset, using: ${GCP_PROJECT}"
    fi

    omg-cli generate-config --env-dir="${ENV_DIR}" --pivnet-api-token="${PIVNET_API_TOKEN}" --gcp-project="${GCP_PROJECT}"

    echo ""
    echo "The following settings are defaults:"
    echo ""
    omg-cli source-config --env-dir="${ENV_DIR}"

    echo ""
    echo "Review the settings above. Modify them by editing the file: ${env_config} and re-running this script"
    echo ""
    read -p "Accept defaults (y/n)? " choice

    case "$choice" in
      y|Y );;
      * ) exit 0;;
    esac
fi

mkdir -p ${ENV_DIR}/pks
cp -r pks ${ENV_DIR}
pushd ${ENV_DIR}/pks
    ./download_helm.sh
    ./install_kubectl.sh
popd

set -o allexport
eval $(omg-cli source-config --env-dir="${ENV_DIR}")
set +o allexport

pushd src/omg-tf
    # Verify all prerequisites are ready. This ensures downloads will succeed from PivNet
    # and the project we're deploying to is valid.
    if [ ! -f $terraform_config ]; then
        if [ -z ${PIVNET_ACCEPT_EULA+x} ]; then
            omg-cli review-eulas --env-dir="${ENV_DIR}"
        else
            omg-cli review-eulas --env-dir="${ENV_DIR}" --accept-all
        fi

        omg-cli prepare-project --env-dir="${ENV_DIR}"
        ./init.sh
    fi

    # Setup infrastructure
    gcloud config set project ${PROJECT_ID}
    terraform init
    terraform get
    terraform apply --auto-approve --parallelism=100 -state=${terraform_state} -var-file=${terraform_config}
    terraform output -json -state=${terraform_state} > ${terraform_output}
popd

# Deploy PCF
omg-cli remote --env-dir="${ENV_DIR}" "push-tiles"
omg-cli remote --env-dir="${ENV_DIR}" "deploy $@"
