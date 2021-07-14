
#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/user/user

terraform init
terraform plan
terraform apply -auto-approve

cd ../../../script/user