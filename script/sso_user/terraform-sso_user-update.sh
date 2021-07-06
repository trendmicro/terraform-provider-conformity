#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/user/sso_user

rm -rf update.tfvars
terraform init
terraform plan
terraform apply -auto-approve

cat << EOF >> update.tfvars
# conformity_sso_user
role       = "ADMIN"

# # access_list01 (can be multiple)
# #level can be "NONE" "READONLY" "FULL"
# account01 = "cloud-conformity-account-access"
# level01  = "ADD-LEVEL"

EOF


terraform apply -var-file="update.tfvars" -auto-approve

cd ../../../script/sso_user
