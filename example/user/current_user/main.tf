

data "conformity_current_user" "user" {}
output "user_details"{
    value =data.conformity_current_user.user
}