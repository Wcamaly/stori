output "lambda_functions_info" {
  value = { for lambda_id, lambda in aws_lambda_function.this : lambda_id => {
    function_name = lambda.function_name
    arn = lambda.arn
    invoke_arn = lambda.invoke_arn
    last_modified = lambda.last_modified
  }}
  description = "Información de las funciones Lambda creadas"
}