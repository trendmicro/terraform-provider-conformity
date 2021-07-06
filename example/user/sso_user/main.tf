resource "conformity_sso_user" "user" {

  first_name = var.first_name
  last_name  = var.last_name
  email      = var.email
  role       = var.role

  access_list {
    account = var.account01
    level   = var.level01
    }

  # access_list {
  #   account = var.account02
  #   level   = var.level02
  #   }

  }

  output "user"{
    value = conformity_sso_user.user
  }