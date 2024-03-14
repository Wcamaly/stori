
provider "aws" {
  
 access_key                  = "test"
  secret_key                  = "test"
  region                      = "us-east-1"
  s3_use_path_style           = false
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    apigateway     = "http://localhost:4566"
    cloudformation = "http://localhost:4566"
    cloudwatch     = "http://localhost:4566"
    dynamodb       = "http://localhost:4566"
    ec2            = "http://localhost:4566"
    es             = "http://localhost:4566"
    elasticache    = "http://localhost:4566"
    firehose       = "http://localhost:4566"
    iam            = "http://localhost:4566"
    kinesis        = "http://localhost:4566"
    lambda         = "http://localhost:4566"
    rds            = "http://localhost:4566"
    redshift       = "http://localhost:4566"
    route53        = "http://localhost:4566"
    s3             = "http://s3.localhost.localstack.cloud:4566"
    secretsmanager = "http://localhost:4566"
    ses            = "http://localhost:4566"
    sns            = "http://localhost:4566"
    sqs            = "http://localhost:4566"
    ssm            = "http://localhost:4566"
    stepfunctions  = "http://localhost:4566"
    sts            = "http://localhost:4566"
  }
}

module "iam_lambda_role" {
  source        = "./modules/iam"
  role_name     = "lambda_execution_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
      },
    ]
  })
  policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole",
    "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
  ]
  custom_policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "ses:SendEmail",
          "ses:SendRawEmail"
        ],
        "Resource": "*"
      },
      {
        "Effect": "Allow",
        "Action": [
          "sqs:SendMessage",
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes"
        ],
        "Resource": "${var.sqs_topic_arn}${var.sqs_topic}"
      },
      {
        "Effect": "Allow",
        "Action": [
          "sns:Publish"
        ],
        "Resource": "${var.sns_topic_arn}${var.sns_topic}"
      },
      {
        "Effect": "Allow",
        "Action": [
          "lambda:InvokeFunction"
        ],
        "Resource": "arn:aws:lambda:${var.region}:${var.account_id}:function:transaction"
      },

    ]
  })
}
module "s3_bucket" {
  source      = "./modules/s3" 
  bucket_name = var.bucket_name
  acl         = "public-read-write"
}


module "lambda_functions" {
  source = "./modules/lambda"

  lambdas = {
    transaction = {
      filename         = "${path.module}/../lambdas/main-transaction.zip"
      handler          = "main"
      runtime          = "go1.x"
      role_arn         = module.iam_lambda_role.role_arn
      environment      = {
       TRANSACTION_SERVICE_URL="http://transaction-service:8080/api/v1"
        USER_SERVICE_URL="http://user-service:8080/api/v1"
        SNS_TOPIC=var.sns_topic
        AWS_REGION=var.region
      }
    },
    email = {
      filename         = "${path.module}/../lambdas/main-email.zip"
      handler          = "main"
      runtime          = "go1.x"
      role_arn         = module.iam_lambda_role.role_arn
      environment      = {
        TRANSACTION_SERVICE_URL="http://transaction-service:8080/api/v1"
        USER_SERVICE_URL="http://user-service:8080/api/v1"
        SNS_TOPIC=var.sns_topic
        AWS_REGION=var.region
      }
    }
  }
}


module "sqs_queue" {
  source      = "./modules/sqs" 
  queue_name  = var.sqs_topic
  lambda_arn = module.lambda_functions.lambda_functions_info["transaction"].arn
  bucket_name = module.s3_bucket.bucket_name
}


resource "aws_s3_bucket_notification" "s3_notification" {
  bucket = module.s3_bucket.bucket_name
  queue {
    queue_arn = module.sqs_queue.sqs_queue_arn  
    events    = ["s3:ObjectCreated:*"]
  }
  depends_on = [module.sqs_queue]
}

/*
module "sns_topic" {
  source    = "./modules/sns"
  topic_name = var.sns_topic
  lambda_arn = module.lambda_functions.lambda_functions_info["email"].arn
} */