@echo off

rem how to create a PAT: https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token
rem requires jq & curl on your machine

set download=tf_download
rmdir /s /q %download%
mkdir %download%
cd %download%

set owner=trendmicro
set name=conformity
set repo=terraform-provider-%name%
set version=0.4.2
set token=YOUR_TOKEN
set os_arch=windows_amd64

set tag=v%version%
set artifact_extension=zip
set artifact=%repo%_%version%_%os_arch%.%artifact_extension%

set list_asset_url="https://api.github.com/repos/%owner%/%repo%/releases/tags/%tag%"

set token_header="Authorization: token %token%"

@curl -s -H %token_header% %list_asset_url% | jq -r ".assets[] | select(.name==\"%artifact%\") | .url" > out.tmp
set /p asset_url=<out.tmp

echo download the artifact
curl -vLJO -H %token_header% -H "Accept: application/octet-stream" "%asset_url%"

echo unzip artifact
tar -xf %artifact%

echo Appdate %AppData%

set target_folder="%AppData%\terraform.d\plugins\%owner%.com\cloudone\conformity\%version%\%os_arch%"
echo create folder in terraform cache
mkdir %target_folder%
echo move artifact to terraform plugin folder
move %repo%_%tag%.exe %target_folder%/%repo%.exe

