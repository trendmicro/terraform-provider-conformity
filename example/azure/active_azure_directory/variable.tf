variable "apikey"{
    type    = string
    default = ""
}
variable "region"{
    type    = string
    default = ""
}
variable "name"{
    type    = string
    default = "Azure Active Directory"
}
variable "directory_id"{
    type    = string
    default = "761d49c9-8898-5d35-c4db-ed8582f20dbf"
}
variable "application_id"{
    type    = string
    default = "c5187d37-8de6-5920-99df-4d8eb3f8cc05"
    sensitive=true
}
variable "application_key"{
    type    = string
    default = "kjx9Q~.CeN4.AxZZVvT8qFRmcx9v9HDVBxgA3mc1"
    sensitive=true
}