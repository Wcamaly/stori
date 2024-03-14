variable "domain" {
  description = "The domain to be used for SES"
  type        = string
}

variable "email_addresses" {
  description = "List of email addresses to verify with SES"
  type        = list(string)
  default     = []
}

variable "route53_zone_id" {
  description = "The Route 53 zone ID for creating DKIM records"
  type        = string
}