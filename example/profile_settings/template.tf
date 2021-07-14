resource "conformity_profile" "profile"{
  // Optional | type: string
  name = ""

  // Optional | type : string
  description = ""
  
  // Optional | type : string
  profile_id = ""

  // optional
  // can be multiple declaration
  included {

    // optional | type: string
    // id of the rule set
    id   = ""

    // optional | type: string
    type = ""

    // optional | type: bool | default: true
    enabled = bool

    // optional | type: string
    provider = ""
    // optional | type: string
    // value can be: "LOW" "MEDIUM" "HIGH" "VERY_HIGH" "EXTREME"
    risk_level = []

    // optional
    exceptions {

      // optional | type: array of string
      filter_tags = []

      // optional | type: array of string
      resources   = []

      // optional | type: array of string
      tags  = []

    }

    // optional
    // can be multiple declaration
    extra_settings {
      // optional | type: bool
      countries = bool

      // optional | type: bool
      multiple = bool

      // optional | type: string
      name = ""

      // optional | type: bool
      regions = bool 

      // required | type: string
      type = ""

      // optional | type: string
      value =  ""


      // optional | type: list
            values {
                // required | type: bool
                enabled = bool
                // required | type: string
                label   = ""
                // required | type: string
                value   = ""
            }

    }
    
  }

}