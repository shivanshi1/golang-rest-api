# Golang REST API with Gin and GORM

This is a simple REST API built using Golang, Gin, and GORM with a PostgreSQL database. It provides CRUD operations for managing users.

## Features

- RESTful API endpoints for user management (Create, Read, Update, Delete)

- Uses Gin as the web framework

- Uses GORM for database interactions

- Loads environment variables using godotenv

- Database migrations with AutoMigrate

## Prerequisites

Before running this project, make sure you have the following installed:

- Go

- PostgreSQL

- Git

Installation

1. **Clone the repository**
   git clone https://github.com/shivanshi1/golang-rest-api.git
   cd golang-rest-api

4. **Create and configure the .env file**
   Copy the example environment file:
   cp .env.example .env

   Then, update .env with your PostgreSQL database credentials:

   DB_HOST=localhost
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=golang_api
   DB_PORT=5432
   SSL_MODE=disable

3. **Install dependencies**
   go mod tidy

5. **Run the Application**
   go run main.go

## API Endpoints

| Method | Endpoint     | Description        |
| ------ | ------------ | ------------------ |
| POST   | /users       | Create a new user  |
| GET    | /users       | Get all users      |
| GET    | /users/:id   | Get a user by ID   |
| PUT    | /users/:id   | Update user by ID  |
| DELETE | /users/:id   | Delete user by ID  |

## Project Structure
.
├── database/
│   ├── db.go         # Database connection and model definition
├── main.go           # Entry point of the application
├── .env.example      # Environment variables example file
├── go.mod            # Go module file
├── go.sum            # Dependencies checksum file
└── README.md         # Project documentation

