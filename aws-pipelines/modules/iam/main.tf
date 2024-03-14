resource "aws_iam_role" "this" {
  name               = var.role_name
  assume_role_policy = var.assume_role_policy
}

# Adjunta políticas IAM administradas por AWS al rol
resource "aws_iam_role_policy_attachment" "policy_attachment" {
  for_each = toset(var.policy_arns)

  role       = aws_iam_role.this.name
  policy_arn = each.value
}

# Crea y adjunta políticas IAM en línea al rol
resource "aws_iam_role_policy" "inline_policy" {
  for_each = var.inline_policies

  role   = aws_iam_role.this.name
  name   = each.key
  policy = each.value
}

# Opcional: Crea una política IAM en línea si se proporciona una política personalizada
resource "aws_iam_role_policy" "custom" {
  count  = var.custom_policy != "" ? 1 : 0
  role   = aws_iam_role.this.name
  name   = "${var.role_name}-custom-policy"
  policy = var.custom_policy
}