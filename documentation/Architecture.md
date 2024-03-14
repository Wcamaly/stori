# Stori Solution Architecture Overview

This document provides an overview of the architecture for the Stori Solution, which is designed to provide robust and scalable web application services.

## Architecture Components

- **Web Interface**: This is the front-end user interface of the application. Users interact with this layer to access services provided by the Stori Solution.

- **API Gateway**: Acts as the single entry point to the application from the outside. It routes incoming requests to the appropriate backend services.

- **Backend for Frontend (BFF)**: This layer is responsible for rendering the web interface. It assembles composite responses that are then returned to the user interface.

- **Transaction Service**: This service handles the storage of transaction data related to users. It interacts with the User Service to ensure transactions are associated with the correct user data.

- **User Service**: Responsible for storing and managing user data. This includes authentication, user profiles, and other relevant information.

- **PostgreSQL Database**: This is the relational database service used for storing all user and transaction data. It serves as a persistent storage layer that is accessed by the Transaction and User Services.

- **Event Queue SNS/SQS**: Utilized for managing communication between services. It decouples the services by using a publish/subscribe mechanism for events.

- **Lambda for S3 Uploads**: This AWS Lambda function listens for file uploads to an S3 bucket and processes these files. Upon processing completion, it emits an event to the event queue with details for further actions, such as email notifications.

- **Lambda for Email via SES**: Triggered by an event in the queue, this Lambda function composes and sends an email using AWS Simple Email Service (SES). It retrieves necessary information from the Transaction Service to construct the email.

## Workflow

1. Users interact with the Web interface, which communicates with the Backend for Frontend (BFF).
2. The BFF processes requests and communicates with the Transaction and User Services as needed.
3. User transactions are stored in the PostgreSQL database through the Transaction Service.
4. Files uploaded by users to S3 buckets trigger the Lambda for S3 Uploads, which processes the files and places events in the SNS/SQS queue.
5. The Event Queue processes these events and, if necessary, triggers the Lambda for Email via SES to send out emails.

## Services Description

### Backend for Frontend (BFF)
- Assembles composite responses for the frontend.
- Communicates with Transaction and User Services.

### Transaction Service
- Manages transaction data.
- Stores transaction records in PostgreSQL.

### User Service
- Manages user information.
- Communicates with PostgreSQL to retrieve and store user data.

### PostgreSQL Database
- Provides a managed relational database service.
- Serves as the storage for user and transaction data.

### Event Queue SNS/SQS
- Handles asynchronous messaging between services.
- Provides a scalable and flexible queue system for events.

### Lambda Functions
- **Lambda for S3 Uploads**: Processes file uploads and emits events to the queue.
- **Lambda for Email via SES**: Sends emails based on events received from the queue.

## Security

The API Gateway ensures that all incoming requests are authenticated and authorized before they reach the backend services. Communication between services is secured and encrypted.

## Scalability

The architecture is designed for scalability, with services such as PostgreSQL, SNS/SQS, and Lambda being able to handle high loads and scale automatically with the demand.

For more details on each component and their specific configurations, please refer to the service documentation or the accompanying detailed architecture diagrams.
