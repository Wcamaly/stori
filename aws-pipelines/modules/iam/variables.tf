variable "role_name" {
  description = "The name of the IAM role to be created"
  type        = string
}

variable "assume_role_policy" {
  description = "The policy that grants an entity permission to assume the role"
  type        = string
}

variable "policy_arns" {
  description = "List of ARN of the IAM policies to attach to the role"
  type        = list(string)
  default     = []
}

variable "inline_policies" {
  description = "Map of inline policies to attach to the role, where key is the name of the policy, and value is the policy document"
  type        = map(string)
  default     = {}
}
variable "custom_policy" {
  description = "Custom IAM policy for additional permissions."
  type        = string
  default     = ""
}