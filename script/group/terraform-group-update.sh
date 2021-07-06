#for running sample folder
cd ..
cd ..
go mod vendor
make install
cd example/group

rm -rf update.tfvars
terraform init
terraform plan
terraform apply -auto-approve

cat << EOF >> update.tfvars
name_group1 = "update_group_01"
tag_group1 = ["g1_update_tag1","g1_update_tag2"]

name_group2 = "update_group_02"
tag_group2 = ["g2_update_tag1", "g2_update_tag2"]
EOF


terraform apply -var-file="update.tfvars" -auto-approve

cd ../../script/group
