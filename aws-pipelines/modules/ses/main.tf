resource "aws_ses_domain_identity" "this" {
  domain = var.domain
}

resource "aws_ses_email_identity" "email" {
  for_each = toset(var.email_addresses)

  email = each.key
}

# Opcional: Configura DKIM para el dominio
resource "aws_ses_domain_dkim" "this" {
  domain = aws_ses_domain_identity.this.domain
}

# Ejemplo de registro DKIM en Route 53. Ajustar según necesidades
resource "aws_route53_record" "dkim" {
  for_each = { for d in aws_ses_domain_dkim.this.dkim_tokens : d => d }

  zone_id = var.route53_zone_id
  name    = "${each.value}._domainkey.${var.domain}"
  type    = "CNAME"
  ttl     = "600"
  records = [each.value.dkim_domain]
}