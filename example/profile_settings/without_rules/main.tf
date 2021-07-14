resource "conformity_profile" "profile_settings"{
  // Optional | type: string
  name = "conformity-without-rules"

  // OPtional | type : string
  description = "conformity guardrail development - without rules"


}

output "profile" {
  value = conformity_profile.profile_settings
}