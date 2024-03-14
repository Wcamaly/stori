resource "aws_sqs_queue" "this" {
  name                      = var.queue_name
  delay_seconds             = var.delay_seconds
  max_message_size          = var.max_message_size
  message_retention_seconds = var.message_retention_seconds
  receive_wait_time_seconds = var.receive_wait_time_seconds
}

resource "aws_lambda_event_source_mapping" "s3_to_lambda" {
  event_source_arn  = aws_sqs_queue.this.arn
  function_name     = var.lambda_arn
  enabled           = true
}


resource "aws_sqs_queue_policy" "s3_to_sqs_policy" {
  queue_url = aws_sqs_queue.this.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal = "*",
      Action = "sqs:SendMessage",
      Resource = aws_sqs_queue.this.arn,
      Condition = {
        ArnLike = {
          "aws:SourceArn" = "arn:aws:s3:*:*:${var.bucket_name}"
        }
      }
    }]
  })
}