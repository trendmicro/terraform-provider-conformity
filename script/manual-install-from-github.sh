#!/bin/bash -e

# source: https://gist.github.com/umohi/bfc7ad9a845fc10289c03d532e3d2c2f

# how to create a PAT: https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token

# requires jq on your machine

download="tf_download"
rm -rf ${download} || true
mkdir ${download}
cd ${download}

owner="trendmicro"
name="conformity"
repo="terraform-provider-${name}"
version="0.4.3"
token="YOUR_TOKEN"
os_arch="darwin_amd64"

tag="v${version}"
artifact_extension="zip"
artifact="${repo}_${version}_${os_arch}.${artifact_extension}"

list_asset_url="https://api.github.com/repos/${owner}/${repo}/releases/tags/${tag}"

token_header="Authorization: token ${token}"
asset_url=$(curl -H "${token_header}" "${list_asset_url}" | jq ".assets[] | select(.name==\"${artifact}\") | .url" | sed 's/\"//g')

echo download the artifact
curl -vLJO -H "${token_header}" -H 'Accept: application/octet-stream' \
     "${asset_url}"

echo unzip artifact
unzip ${artifact}

echo move artifact to terraform plugin folder
target_folder="${HOME}/.terraform.d/plugins/${owner}.com/cloudone/conformity/${version}/${os_arch}"
mkdir -p ${target_folder}
mv ${repo}_${tag} ${target_folder}/${repo}
