resource "aws_lambda_function" "this" {
  for_each = var.lambdas

  handler       = each.value.handler
  runtime       = each.value.runtime
  role          = each.value.role_arn
  function_name = each.key

  filename         = each.value.filename
  source_code_hash = filebase64sha256(each.value.filename)
  memory_size       = 128
  timeout           = 10

  environment {
    variables = each.value.environment
  }
}

