#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/gcp_organisation

terraform init
terraform plan
terraform apply -auto-approve

cd ../../script/gcp_organisation
