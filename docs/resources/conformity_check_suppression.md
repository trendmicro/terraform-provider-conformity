---
page_title= "conformity_check_suppression Resource"
subcategory= "Check"
description= |-
  Allows you to create Custom Check on Conformity. 
---

# Resource `conformity_check_suppression`
Allows you to create Custom Check on Conformity

## Example Usage
```
resource "conformity_check_suppression" "check"{
        account_id="f7ca8f9d-2b4b-7624-289c-d1d2fe23b54d"
        rule_id="EC2-054"
        region="us-1"
        resource_id="sg-016c1348bdc0454a4"
        note="suppression check"
}
```