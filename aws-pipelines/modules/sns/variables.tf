variable "topic_name" {
  description = "The name of the SNS topic"
  type        = string
}

variable "lambda_arn" {
  description = "The ARN of the Lambda function to subscribe to the SNS topic"
  type        = string
}