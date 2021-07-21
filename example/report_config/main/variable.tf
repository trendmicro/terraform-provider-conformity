variable "apikey"{
    type    = string
    default = ""
}
variable "region"{
    type    = string
    default = ""
}
// if you want to create organisation-level report config, leave account_id and group_id commented.
variable "account_id" {
    type    = string 
    default = ""
}
// if you want to create group-level report config, uncomment group_id value and provide account ID on the terraform.tfvars
// variable "group_id" {
//     type    = string 
//     default = ""
// }