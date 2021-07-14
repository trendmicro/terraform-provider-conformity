#for running sample folder
cd ..
cd ..
cd ..
go mod vendor
make install
cd example/profile_settings/existing_profile

terraform init
terraform plan
terraform apply -auto-approve

cd ../../../script/profile_settings/existing_profile
