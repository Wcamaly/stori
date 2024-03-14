variable "lambdas" {
  description = "A map containing the lambda functions configurations"
  type = map(object({
    filename         = string
    handler          = string
    runtime          = string
    role_arn         = string
    environment      = map(string)
  }))
}