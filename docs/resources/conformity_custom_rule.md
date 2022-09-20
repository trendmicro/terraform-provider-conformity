---
page_title: "conformity_custom_rule Resource"
subcategory: "Custom Rules"
description: |-
  Allows you to create Custom Rules on Conformity. 
---

# Resource `conformity_custom_rule`
Allows you to create Custom Rules on Conformity

There are 6 Custom rule that is being tested
  1.Storage Naming for all 3 clouds (aws,azure and gcp)
  2.MongoDB port restriction for all 3 clouds.(aws,azure and gcp)


1.Storage Naming for all 3 clouds (aws,azure and gcp)

## Storage Naming for aws

  ### For Running the Create Custom Rule API  for aws
    1.First you have the aws access so that you can create a bucket for it 
    2.while creating the bucket give appropriate name to bucket 
    3.Run the Create Custom Rule api by giving required value 
        -> resourceType,service and provider
        -> resourceType and service you can find from the documentation of the api
         there was a link for resourceType and services just click that you will get response
         find the appropriate resourceType and services
    4.After giving the reesourceType and services and provider set the rule appropriately and also the condition which will check for the bucket name which you  have created
    5.Run the api 
    
  ### Checking the custom rule
    6.Copy the id from the successfull response
    7.Run the Run Custom Rule api by passing the accountId of your aws account and rule id (which you will get in Create Custom Rule api response)
    8.Copy the same body of Create Custom Rule api 
    9.Run the api you will see your bucket details over there and your rule status whether its failed / Success
  
  ### If you dont know the Bucket Details and also what to pass in path of attributes and rules
    1.You need to run the Run Custom Rule api by providing the aws accoundid and the rule idand also the set the resourceData=true
    2. In body you need to pass 
    ```
    {
          "configuration": {
            "name": "S3 bucket logging enabled",
            "description": "S3 buckets have logging enabled",
            "service": "S3",
            "resourceType": "s3-bucket",
            "attributes": [
              {
                "name": "bucketLogging",
                "path": "data.LoggingEnabled",
                "required": true
              }
            ],
            "rules": [
              {
                "conditions": {
                  "all": [
                    {
                      "value": null,
                      "operator": "notEqual",
                      "fact": "bucketLogging"
                    }
                  ]
                },
                "event": {
                  "type": "Bucket has logging enabled"
                }
              }
            ],
            "severity": "MEDIUM",
            "categories": [
              "security"
            ],
            "provider": "aws",
            "enabled": true
          }
        }
    
    
    ```
  3.Change the resourceType,services and  provider appropriately
  4.Put the same rule as Create Custom Rule Api
  5.Save and send you will see your bucketdetails and status of your rule and the resourceData details in which you can see all the bucket details where you can create a rule in any of the attributes


## Example Usage of Create Custom Rule API for aws
```hcl
AWS Storage Naming-:
        resource "conformity_custom_rule" "example"{
          name= "S3 Bucket Custom Rule"
          description      = "This custom rule ensures S3 buckets follow our best practice updated"
        	remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
        	service          = "S3"
        	resource_type     = "s3-bucket"
        	categories       = ["security"]
        	severity         = "HIGH"
        	cloud_provider   = "aws"
        	enabled          = true
        	attributes {
        		name     = "bucketName"
        		path     = "data.Name"
        		required = true
        	  }
        	  rules {
        	    operation = "all"
        		conditions {
        		  fact     = "bucketName"
        		  operator = "pattern"
        		  value    = "^shunyekaa$"
        		}
        		event_type = "Bucket name is longer than 32 characters"
        	  }
  }

```

## Storage Naming for Azure

  ### For Running the Create Custom Rule API for azure
   
    1.First you have the azure access so that you can create a Storage Accounts in it
    2.while creating the Storage Accounts  give appropriate name 
    3.Run the Create Custom Rule api by giving required value 
        -> resourceType,service and provider
        -> resourceType and service you can find from the documentation of the api
         there was a link for resourceType and services just click that you will get response
         find the appropriate resourceType and services
    4.After giving the resourceType and services and provider set the rule appropriately and also the condition which will check for the Storage Account name which you  have created
    5.Run the api 

  ### Checking the custom rule
    6.Copy the id from the successfull response
    7.Run the Run Custom Rule api by passing the accountId of your azure account and rule id (which you will get in Create Custom Rule api response)
    8.Copy the same body of Create Custom Rule api 
    9.Run the api you will see your Storage account Details over there and your rule status whether its failed / Success
   
  ### If you dont know the Storage Account Details and also what to pass in path of attributes and rules
    1.You need to run the Run Custom Rule api by providing the azure accoundId and the rule id and also the set the resourceData=true
    2. In body you need to pass the same "configuration details as passed above in aws"
    3.Just change the details of configuration (resourceType,services ,provider and rules)
    4.Run it and you can see in response resourceData where you can get your Storage Accounts Details and you can apply rule to any attributes
## Example Usage of Create Custom Rule API for Azure
```
        Azure Storage Naming-:
        resource "conformity_custom_rule" "example"{
          "name": "Azure Storage Account Custom Rule testing",
          "description": "This custom rule ensures Azure Storage Account follow our best practice",
          "service": "StorageAccounts",
          "resourceType": "storage-accounts",
          "categories": [
            "security"
          ],
          "severity": "HIGH",
          "provider": "azure",
          "enabled": true,
          "attributes": [
            {   
              "name": "storageName",
              "path": "data.name",
              "required": false
            }
          ],
          "rules": [
            {
              "conditions": {
                "any": [
                  {
                    "fact": "storageName",
                    "operator": "equal",
                    "value": "customruletesting"
                  }
                ]
              },
              "event": {
                "type": "Storage Account Name should be customruletesting"
              }
            }
            
          ],
          "remediationNotes": "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n"
        }
```
## Storage Naming for GCP
  
  ### For Running the Create Custom Rule API for GCP
   
    1.First you have the gcp access so that you can create a Cloud Storage in it
    2.while creating the Cloud Storage  give appropriate name 
    3.Run the Create Custom Rule api by giving required value 
        -> resourceType,service and provider
        -> resourceType and service you can find from the documentation of the api
         there was a link for resourceType and services just click that you will get response
         find the appropriate resourceType and services
    4.After giving the resourceType and services and provider set the rule appropriately and also the condition which will check for the Cloud Storage name which you  have created
    5.Run the api 

  ### Checking the custom rule
    6.Copy the id from the successfull response
    7.Run the Run Custom Rule api by passing the accountId of your gcp account and rule id (which you will get in Create Custom Rule api response)
    8.Copy the same body of Create Custom Rule api 
    9.Run the api you will see your Cloud Storage Details over there and your rule status whether its Failed / Success
   
  ### If you dont know the Cloud Storage  and also what to pass in path of attributes and rules
    1.You need to run the Run Custom Rule api by providing the gcp accoundId and the rule id and also the set the resourceData=true
    2. In body you need to pass the same "configuration details as passed above in aws"
    3.Just change the details of configuration (resourceType,services ,provider and rules)
    4.Run it and you can see in response resourceData where you can get your Cloud Storage Details and you can apply rule to any attributes

## Example Usage of Create Custom Rule API for GCP
```
GCP:-
resource "conformity_custom_rule" "example"{
  "name": "Gcp Google Cloud  Custom Rule testing",
  "description": "This custom rule ensures Azure Storage Account follow our best practice",
  "service": "CloudStorage",
  "resourceType": "cloudstorage-buckets",
  "categories": [
    "security"
  ],
  "severity": "HIGH",
  "provider": "gcp",
  "enabled": true,
  "attributes": [
    {   
      "name": "cloudStorageName",
      "path": "data.name",
      "required": true
    }
  ],
  "rules": [
    {
      "conditions": {
        "any": [
          {
            "fact": "cloudStorageName",
            "operator": "pattern",
            "value": "^shunyekaa$"
          }
        ]
      },
      "event": {
        "type": "Cloud Storage Bucket Name should be shunyekaa"
      }
    }
    
  ],
  "remediationNotes": "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n"
}

```
2.MongoDB port restriction for all 3 clouds.

  ## Mongo DB Port Restriction  for aws
   ### For Running the Create Custom Rule API  for aws
    1.First you have the aws access so that you can create a Security Group
    2.while creating the Security Group give appropriate name of security group 
        ->Set the Inbound and Outbound rule by giving the port,protocols and other details properly
    3.Run the Create Custom Rule api by giving required value 
        -> resourceType,service and provider
        -> resourceType and service you can find from the documentation of the api
         there was a link for resourceType and services just click that you will get response
         find the appropriate resourceType and services
    4.After giving the reesourceType and services and provider set the rule appropriately and also the condition which will check for the port restriction
    5.Run the api 
    
  ### Checking the custom rule
    6.Copy the id from the successfull response
    7.Run the Run Custom Rule api by passing the accountId of your aws account and rule id (which you will get in Create Custom Rule api response)
    8.Copy the same body of Create Custom Rule api 
    9.Run the api you will see your bucket details over there and your rule status whether its failed / Success
  
  ### If you dont know the Security Group Details and also what to pass in path of attributes and rules
    1.You need to run the Run Custom Rule api by providing the aws accoundid and the rule id and also the set the resourceData=true
    2. In body you need to pass "configuration details as passed above in aws" 
    3.Just change the resourceType,service ,provider and other related details
    4.Run it and you can see in response resourceData where you can get your Security Group Details and you can apply rule to any attributes


## Example Usage of Create Custom Rule API for AWS

```
     resource "conformity_custom_rule" "example" {
          "name": "AWS Security Group for Mongo db port restriction",
          "description": "This custom rule ensures AWS Security Group  follow our best practice",
          "service": "EC2",
          "resourceType": "ec2-securitygroup",
          "categories": [
              "security"
          ],
          "severity": "HIGH",
          "provider": "aws",
          "enabled": true,
          "attributes": [
                  {
                        "name": "portRestriction",
                        "path": "data.IpPermissions[:].FromPort",
                        "required": false
                  },
                  {
                        "name": "IpPermissionCheck",
                        "path": "data.IpPermissions",
                        "required": false
                  }
          ],
          "rules": [
              {
                  "conditions": {
                      "any": [
                          {
                              "path": "$.length",
                              "fact": "IpPermissionCheck",
                              "value": 0,
                              "operator": "equal"
                          },
                          {
                              "value": 27017,
                              "operator": "doesNotContain",
                              "fact": "portRestriction"
                         }
                      ]
                  },
                  "event": {
                      "type": "The Port should not be 27017"
                  }
              }
          ],
          "remediationNotes": "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n"
      }

```

## Mongo DB Port Restriction  for azure

  ## Mongo DB Port Restriction  for azure
   ### For Running the Create Custom Rule API  for azure
    1.First you have the azure access so that you can create a Network Security Group
    2.while creating the Network Security Group give appropriate name of Network security group 
        ->Set the Inbound and Outbound rule by giving the port,protocols and other details properly
    3.Run the Create Custom Rule api by giving required value 
        -> resourceType,service and provider
        -> resourceType and service you can find from the documentation of the api
         there was a link for resourceType and services just click that you will get response
         find the appropriate resourceType and services
    4.After giving the resourceType and services and provider set the rule appropriately for port restriction
    5.Run the api 
    
  ### Checking the custom rule
    6.Copy the id from the successfull response
    7.Run the Run Custom Rule api by passing the accountId of your azure account and rule id (which you will get in Create Custom Rule api response)
    8.Copy the same body of Create Custom Rule api 
    9.Run the api you will see your bucket details over there and your rule status whether its failed / Success
  
  ### If you dont know the Network Security Group Details and also what to pass in path of attributes and rules
    1.You need to run the Run Custom Rule api by providing the azure accoundid and the rule id and also the set the resourceData=true
    2. In body you need to pass "configuration details as passed above in aws" 
    3.Just change the resourceType,service ,provider and other related details
    4.Run it and you can see in response resourceData where you can get your Network Security Group Details and you can apply rule to any attributes

## Example Usage of Create Custom Rule API for Azure

```
          resource "conformity_custom_rule" "example" {
              "name": "Azure Network Security Group Mongo db port Restriction",
              "description": "This custom rule ensures Azure Network Security follow our best practice",
              "service": "Network",
              "resourceType": "network-network-security-groups",
              "categories": [
                  "security"
              ],
              "severity": "HIGH",
              "provider": "azure",
              "enabled": true,
              "attributes": [
                      {
                        "name": "networksecurity",
                        "path": "data.securityRules[:].destinationPortRange",
                        "required": false
                      },
                      {
                        "name": "networksecurityCheck",
                        "path": "data.securityRules",
                        "required": false
                      }
              ],
              "rules": [
                  {
                      "conditions": {
                          "any": [
                                {
                                    "value": "27017",
                                    "operator": "doesNotContain",
                                    "fact": "networksecurity"
                                },
                                {
                                    "path": "$.lenght",
                                    "fact": "networksecurityCheck",
                                    "value": 0,
                                    "operator": "equal"
                                }
                          ]
                      },
                      "event": {
                          "type": "The Port should not be 27017"
                      }
                  }
              ],
              "remediationNotes": "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n"
          }


```


## Mongo DB Port Restriction  for gcp

   ### For Running the Create Custom Rule API  for azure
    1.First you have the gcp access so that you can create a Firewall Rules
    2.while creating the Firewall Rules  give appropriate name of Network security group 
        ->Set the Inbound and Outbound rule by giving the port,protocols and other details properly
    3.Run the Create Custom Rule api by giving required value 
        -> resourceType,service and provider
        -> resourceType and service you can find from the documentation of the api
         there was a link for resourceType and services just click that you will get response
         find the appropriate resourceType and services
    4.After giving the resourceType and services and provider set the rule appropriately for port restriction
    5.Run the api 
    
  ### Checking the custom rule
    6.Copy the id from the successfull response
    7.Run the Run Custom Rule api by passing the accountId of your gcp account and rule id (which you will get in Create Custom Rule api response)
    8.Copy the same body of Create Custom Rule api 
    9.Run the api you will see your bucket details over there and your rule status whether its failed / Success
  
  ### If you dont know the Firewall Rule  Details and also what to pass in path of attributes and rules
    1.You need to run the Run Custom Rule api by providing the azure accoundid and the rule id and also the set the resourceData=true
    2. In body you need to pass "configuration details as passed above in aws" 
    3.Just change the resourceType,service ,provider and other related details
    4.Run it and you can see in response resourceData where you can get your Firewall Rule Details and you can apply rule to any attributes

## Example Usage of Create Custom Rule API for Azure

```
          resource "conformity_custom_rule" "example" {
              "name": "GCP Firewalls Rules Mongo db port restriction",
              "description": "This custom rule ensures GCP Firewalls Rules  follow our best practice",
              "service": "CloudVPC",
              "resourceType": "cloudvpc-firewallrules",
              "categories": [
                  "security"
              ],
              "severity": "HIGH",
              "provider": "gcp",
              "enabled": true,
              "attributes": [
                      {
                        "name": "firewallrules",
                        "path": "data.allowed[:].ports[:]",
                        "required": false
                    },
                    {
                        "name": "firewallrulescheck",
                        "path": "data.allowed[:].ports",
                        "required": false
                    }
              ],
              "rules": [
                  {
                      "conditions": {
                          "any": [
                                {
                                    "value": "27017",
                                    "operator": "doesNotContain",
                                    "fact": "firewallrules"
                                },
                                {
                                    "path": "$.lenght",
                                    "fact": "firewallrulescheck",
                                    "value": 0,
                                    "operator": "equal"
                                }
                          ]
                      },
                      "event": {
                          "type": "The Port should not be 27017"
                      }
                  }
              ],
              "remediationNotes": "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n"
          }


```
## Argument reference

- `name` (String) -  Name of the custom rule.
- `description` (String) -  description of the custom rule.
- `remediation_notes` (String) - remediation_notes of the custom 
- `service` (String) - service  of the custom rule 
- `resource_type` (String) - resource type of the custom rule
- `categories` (Array of String) -  categories of the custom rule. Enum: ["security", "sustainability", "performance-efficiency", "operational-excellence"]
- `severity` (String) - severity of the custom rule. Enum :"LOW","MEDIUM","HIGH","VERY_HIGH","EXTREME"
- `cloud_provider` (String ) -  Name of the cloud provider. Enum: "aws","azure","gcp".
- `enabled` (Bool) - This attributes determines whether this setting enabled or not (true ,false)

- `attributes` List: Can be multiple declaration 
    * `name` (String) BucketName.
    * `path` (String) Path of the Bucket.
    * `required` (String) This  determines whether the attribute is required or not. 

-  `rule` List: Can be multiple declaration
    * `operation` (String) -  operation of the rule. Enum: "any","all"
    * `condition` List: Can be multiple declaration
        * `fact` (String) - BucketName
        * `operation` (String) - pattern
        *  `value` (String) - value  of the operator
    * `event_type` (String) - Message
