resource "conformity_profile" "profile_settings"{

  // add the profile_id , this values needs to be set if you want to overwrite an existing profile during terraform creation
  profile_id = "your-profile-id" 

  name = "conformity-existing-cloud-with-rules"
  // Optional | type : string
  description = "conformity development - rules included"

  included {
    id         = "EC2-001"
    provider   = "aws"
    enabled    = false
    risk_level = "LOW"
    exceptions {
      tags  = ["some_tag"]
    }
  }

included {
    id         = "RTM-002"
    provider   = "aws"
    enabled    = true
    risk_level = "MEDIUM"
    extra_settings {
      name =  "ttl"
      type =  "ttl"
      value = "72"
    }
  }

}

output "profile" {
  value = conformity_profile.profile_settings
}