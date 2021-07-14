#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/report_config/main

terraform init
terraform plan
terraform apply -auto-approve

cd ../../../script/report_config
