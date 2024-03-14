variable "bucket_name" {
  description = "The name of the bucket. Must be unique across all of AWS"
  type        = string
}

variable "acl" {
  description = "The canned ACL to apply. Defaults to 'private'"
  default     = "private"
  type        = string
}