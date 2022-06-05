#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/gcp

rm -rf update.tfvars
terraform init
terraform plan
terraform apply -auto-approve

cat << EOF >> update.tfvars
name = "trendmicro-update"
environment = "Staging-update"
EOF

terraform apply -var-file="update.tfvars" -auto-approve
# terraform apply --var "name=trendmicro-update" --var "environment=Staging-update" -auto-approve
# Uncomment this to destroy the resources
# terraform destroy -auto-approve
# rm -rf .terraform
# rm -rf .terraform.lock.hcl
# rm -rf terraform.tfstate
# rm -rf terraform.tfstate.backup

cd ../../script/gcp
