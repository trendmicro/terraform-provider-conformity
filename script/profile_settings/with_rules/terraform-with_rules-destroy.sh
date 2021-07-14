#for running sample folder
cd ..
cd ..
cd ..
go mod vendor
make install
cd example/profile_settings/with_rules

terraform init
terraform plan
terraform destroy -auto-approve

rm -rf .terraform
rm -rf update.tfvars
rm -rf .terraform.lock.hcl
rm -rf terraform.tfstate
rm -rf terraform.tfstate.backup

cd ../../../script/profile_settings/with_rules