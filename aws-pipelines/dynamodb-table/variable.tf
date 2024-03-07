variable "table_name" {
  description = "The name of the DynamoDB table"
  type        = string
}

variable "hash_key" {
  description = "The hash key of the table"
  type        = string
}

variable "range_key" {
  description = "The range key of the table, optional"
  type        = string
  default     = ""
}

variable "attributes" {
  description = "A list of attributes to define in the table"
  type        = list(object({
    name = string
    type = string
  }))
}
