
#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/group

terraform init
terraform plan
terraform apply -auto-approve

cd ../../script/group