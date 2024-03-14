# Project Structure Overview

This README outlines the structure and organization of the project's directories and files. Each component plays a specific role in the configuration, deployment, and operation of the application.

## Directory and File Explanation

### `aws-pipelines/`
This directory contains all the necessary Terraform configurations for setting up the AWS infrastructure required for our pipelines.

#### `aws-pipelines/.terraform/`
- Stores Terraform plugin details specific to this project's configuration.

#### `aws-pipelines/modules/`
- Contains reusable Terraform modules for different parts of the infrastructure.

#### `aws-pipelines/main.tf`
- The main Terraform configuration file that ties together modules and defines the AWS resources.

#### `aws-pipelines/variables.tf`
- Defines variables used in the Terraform configurations.

#### `aws-pipelines/terraform.tfstate`
- The Terraform state file, which holds the current state of the resources managed by Terraform.

### `backends/`
This directory holds the backend services of the application, divided into respective service directories.

#### `backends/transaction-service/`
- Contains the code and configurations for the transaction processing service.

#### `backends/user-service/`
- Contains the code and configurations for the user management service.

### `documentation/`
Stores documentation files related to the project, such as design docs, API specifications, and manuals.

### `frontend/`
Contains all the front-end code, assets, and dependencies required for the web application interface.

### `lambdas/`
This directory is home to AWS Lambda functions that handle various backend tasks.

#### `lambdas/email-process/`
- Contains the code for the Lambda function responsible for processing emails.

#### `lambdas/transaction-process/`
- Contains the code for the Lambda function responsible for processing transactions.

#### `lambdas/main-email.zip`
- A packaged Lambda function for email processing ready for deployment.

#### `lambdas/main-transaction.zip`
- A packaged Lambda function for transaction processing ready for deployment.

### `localstack/`
Used for local development, mimicking AWS services on the developer's machine.

### Root Level Files

#### `script.sh`
- A general utility script for the project, purpose specified within.

#### `.gitignore`
- Root-level file for specifying untracked files.

#### `docker-compose.yml`
- Defines and runs the entire application stack using Docker.

#### `nginx.conf`
- The main nginx configuration file used when the application is deployed.

#### `README.MD`
- The file you're currently reading, providing an overview of the project structure.

#### `run.sh`
- A script to execute the application.

## Additional Notes

