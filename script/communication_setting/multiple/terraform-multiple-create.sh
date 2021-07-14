#for running sample folder
cd ..
cd ..
cd ..
go mod vendor
make install
cd example/communication_setting/multiple

terraform init
terraform plan
terraform apply -auto-approve

cd ../../../script/communication_setting/multiple
