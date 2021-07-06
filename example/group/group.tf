resource "conformity_group" "group_1" {
    name = var.name_group1
    tags =  var.tag_group1
}

resource "conformity_group" "group_2" {
    name =var.name_group2
    tags = var.tag_group2
}
