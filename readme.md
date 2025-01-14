# RBAC (Role-Based Access Control) Project with GoFiber and GORM

This project is an implementation of a Role-Based Access Control (RBAC) system using the GoFiber web framework and GORM as the ORM (Object-Relational Mapping) tool. The system is designed to manage user roles, permissions, and access control for secure and scalable web applications.

## Features

- **Product Management**: CRUD operations for managing Products.
- **Authentication and Authorization**: Secure API endpoints using JWT-based authentication and middleware for role-based access control.
- **Scalable Architecture**: Built with GoFiber for fast and lightweight HTTP server performance.
- **Database Integration**: Powered by GORM for easy database interaction.

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: [GoFiber](https://gofiber.io/)
- **ORM**: [GORM](https://gorm.io/)
- **Authentication**: JWT (JSON Web Tokens)
- **Database**: PostgreSQL
- **Validation**: [Validator](https://github.com/go-playground/validator)

## Prerequisites

Before running this project, ensure you have the following installed:

- Go (v1.18 or later)
- A database (e.g., MySQL, PostgreSQL, or SQLite)
- Postman (optional, for testing APIs)

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/Alpharii/rbac-golang.git
cd rbac-golang
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Configure Environment Variables

Create a `.env` file in the project root directory and add the following:

```env
DATABASE_URL=your_database_url
JWT_SECRET=your_jwt_secret
```

### 4. Run Database Migrations

Make sure the database is created, then run the application to auto-migrate the tables:
```bash
go run main.go
```

### 5. Start the Server
```bash
go run main.go
```
The server will start on the port specified in the `.env` file (default: `localhost:3000`).

## API Endpoints

### Authentication
- **POST** `/login` - User login
- **POST** `/register` - User registration

### Product Management
- **GET** `/products` - List all products               -- must login
- **GET** `/products/:id` - Get a product by ID         -- must login
- **POST** `/products` - Create a new product           -- must login and admin
- **PUT** `/products/:id` - Update a product            -- must login and admin
- **DELETE** `/products/:id` - Delete a product         -- must login and admin

## Middleware

- **JWT Middleware**: Ensures that only authenticated users can access protected routes.
- **RBAC Middleware**: Ensures users have the required roles/permissions to access specific endpoints.

## Project Structure

```plaintext
.
├── main.go            # Entry point of the application
├── config/            # Configuration files
├── controllers/       # Handlers for API endpoints
├── middlewares/       # Middleware for authentication and authorization
├── models/            # Database models
├── routes/            # API route definitions
├── utils/             # Utility functions
```


Enjoy building secure and scalable applications with this RBAC system!

