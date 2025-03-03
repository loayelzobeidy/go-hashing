# Golang High-Performance Microservice

## Overview
This project is a high-performance Golang microservice designed to handle CPU-intensive operations, specifically hashing multiple values (up to 1000) using a base text. The microservice is secured with JWT authentication and user claims-based authorization. It leverages **Gin-Gonic** for HTTP routing and **GORM** for ORM operations.

## Features
- **Efficient Hashing:** Handles batch hashing operations with optimized performance.
- **JWT Security:** Ensures secure access using JSON Web Tokens (JWT).
- **User Claims Authorization:** Restricts access based on user roles and permissions.
- **Gin-Gonic:** Fast HTTP framework for handling API requests.
- **GORM:** ORM for seamless database interactions.
- **Scalable & Lightweight:** Built for performance and scalability.

## Tech Stack
- **Golang**
- **Gin-Gonic** (HTTP framework)
- **GORM** (ORM for database operations)
- **JWT-Go** (for authentication)
- **Bcrypt/SHA256** (for secure hashing)

## Installation
### Clone the repository:
```sh
git clone https://github.com/loayelzobeidy/go-hashing.git
cd golang-microservice
```

### Install dependencies:
```sh
go mod tidy
```

### Set up environment variables:
```sh
export POSTGRES_HOST=localhost
export POSTGRES_USERNAME=youruser
export POSTGRES_PASSWORD=yourpass
export POSTGRES_PORT=yourPORT
export JWT_SECRET=your_secret_key
```

### Run the service:
```sh
go run main.go
```

## API Endpoints
### Authentication
- **POST /user/login** - Authenticate user and generate JWT token
- **POST /user/register** - Register a new user

### Hashing Service
- **POST /encrypted/hash** (Authorized users only)
  - Accepts a JSON payload with multiple values to hash
  - Returns the hashed values

## Example Request
```sh
curl -X POST "http://localhost:8080/encrypted/hash" \
     -H "Authorization: Bearer <your_jwt_token>" \
     -d '{"values": ["value1", "value2", ..., "value1000"]}'
```

## Contributing
Feel free to fork this repository, open issues, or submit pull requests.

## License
MIT License