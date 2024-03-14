resource "aws_sns_topic" "this" {
  name = var.topic_name
}

resource "aws_sns_topic_subscription" "lambda_subscription" {
  topic_arn = aws_sns_topic.this.arn
  protocol  = "lambda"
  endpoint  = var.lambda_arn
}
