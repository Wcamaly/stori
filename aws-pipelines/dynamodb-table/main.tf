resource "aws_dynamodb_table" "table" {
  name           = var.table_name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = var.hash_key

  # Se incluye range_key solo si se proporciona un valor no nulo
  range_key      = var.range_key != null ? var.range_key : null

  attribute {
    name = var.hash_key
    type = "S" # Suponiendo que hash_key es siempre un String
  }

  # Definición dinámica para todos los atributos, incluyendo hash_key, range_key, y cualquier otro atributo adicional
  dynamic "attribute" {
    for_each = var.attributes
    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  tags = {
    Name = var.table_name
  }
}