provider "aws" {
  access_key                  = "mock_access_key"
  secret_key                  = "mock_secret_key"
  region                      = "us-east-1"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  #s3_force_path_style         = true

  endpoints {
    dynamodb = "http://localhost:4566"
  }
}

module "user_table" {
  source = "./dynamodb-table"

  table_name = "User"
  hash_key   = "userId"
  range_key  = "email" # Opcional, basado en tu diseño de tabla

  attributes = [
    {
      name = "userId"
      type = "S"
    },
    {
      name = "email"
      type = "S"
    }
  ]
}