variable "region" {
  description = "The name of the IAM role to be created"
  type        = string
  default = "us-east-1"
}

variable "account_id" {
  description = "The name of the IAM role to be created"
  type        = string
  default = "000000000000"
}

variable "sqs_topic_arn" {
  default = "arn:aws:sqs:us-east-1:000000000000:"
}

variable "sqs_topic" {
  description = "value of the sqs topic"
  type        = string
  default = "S3-SQS-QUEUE"
}

variable "sns_topic_arn" {
  default = "arn:aws:sns:us-east-1:000000000000:"
  
}

variable "sns_topic" {
  description = "value of the sns topic"
  type        = string
  default = "EMAIL-TOPIC"
}

variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string
  default = "stori-bucket"
}