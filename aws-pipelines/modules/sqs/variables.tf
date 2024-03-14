variable "queue_name" {
  description = "The name of the SQS queue."
  type        = string
}

variable "delay_seconds" {
  description = "The time in seconds that the delivery of all messages in the queue will be delayed. Defaults to 0."
  type        = number
  default     = 0
}

variable "max_message_size" {
  description = "The limit of how many bytes a message can contain before Amazon SQS rejects it. Defaults to 262144."
  type        = number
  default     = 262144  # 256 KiB
}

variable "message_retention_seconds" {
  description = "The number of seconds Amazon SQS retains a message. Defaults to 345600 (4 days)."
  type        = number
  default     = 345600
}

variable "receive_wait_time_seconds" {
  description = "The time for which a ReceiveMessage call will wait for a message to arrive. Defaults to 0."
  type        = number
  default     = 0
}

variable "lambda_arn" {
  description = "The ARN of the Lambda function to which the SQS queue is connected."
  type        = string
}

variable "bucket_name" {
  description = "The bucket name listener."
  type        = string
}

variable "bucket_arn" {
  description = "value of the bucket arn"
  type        = string
}