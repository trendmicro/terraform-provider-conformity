#for running sample folder
cd ..
cd ..
cd ..
go mod vendor
make install
cd example/profile_settings/values_string_int

terraform init
terraform plan
terraform apply -auto-approve

cd ../../../script/profile_settings/values_string_int
