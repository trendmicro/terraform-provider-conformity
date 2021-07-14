
#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/user/sso_user

terraform init
terraform plan
terraform apply -auto-approve

cd ../../../script/sso_user