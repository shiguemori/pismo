# Pismo Project

This project is designed to create a suite of APIs using Gin Web Framework and Gorm, the ORM library for Golang. It aims to provide a robust backend solution for financial operations.

## Overview

The Pismo project is structured to support scalable and maintainable API development. It includes a Dockerized PostgreSQL database and utilizes the Gin framework for handling HTTP requests and Gorm for object-relational mapping.

## Prerequisites

Before running this project, you must have the following installed:

- Docker and Docker Compose
- Go programming language
- Git (for version control)

## Getting Started

To get the application up and running, follow these steps:

### 1. Clone the Repository

First, clone the repository to your local machine:

```sh
git clone https://github.com/shiguemori/pismo.git
```

### 2. Start the Database
Use docker-compose to spin up the PostgreSQL database container:

```docker-compose up -d postgres```

### 3. Run the Application
Navigate to the project directory and run the application with the following command:

```go run main.go```

### Development Workflow
During development, you might want to refresh your database and application state. To reset your environment use the following commands:
```
docker-compose down -v
docker-compose up -d postgres
go run main.go
```

### Accessing the API Documentation
Once the application is running, the API documentation can be accessed via Swagger UI at the following URL:

- http://localhost:8080/swagger/index.html

Here you can view the list of available endpoints, their expected request formats, and try out the API calls directly from the browser.

### Environment Variables
For security and configuration management, the application utilizes environment variables. Ensure you have a .env file at the root of your project with the necessary variables set, such as POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_DB.

### Testing
in progress 

#### Contributing
Contributions to the project are welcome. Please follow the standard fork and pull request workflow. Ensure that you write or update tests as necessary and adhere to the existing coding style.

#### Contact Information
- email: shiguemori@hotmail.com
- phone: +55 11 975676977
- LinkedIn: https://www.linkedin.com/in/shiguemori-fullstack-dev/
- Github: https://github.com/shiguemori