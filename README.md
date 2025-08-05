# Elotus Interview Project

This project contains multiple tasks for the Elotus interview process. It includes Data Structures and Algorithms (DSA) implementations, Authentication Service, and Upload File Service with both command-line interface and Docker support.

## Project Tasks

### Task 1: Data Structures and Algorithms (DSA) ✅
- **Status**: Completed
- **Description**: Implementation of various algorithms with CLI testing interface
- **Commands**: `go run main.go dsa [algorithm]`

### Task 2: Authentication Service ✅
- **Status**: Completed
- **Description**: Authentication server implementation with JWT tokens
- **Commands**: `go run main.go authentication`

### Task 3: Upload File Service ✅
- **Status**: Completed
- **Description**: File upload server with authentication middleware
- **Commands**: `go run main.go upload`

## Prerequisites

- Go 1.16 or higher
- MySQL Database (for authentication and upload services)

## Installation

1. Clone the repository
2. Navigate to the project directory
3. Run `go mod tidy` to install dependencies

## Usage

### Task 1: DSA Algorithms

The project includes three different algorithms that can be tested:

#### 1. Gray Code
Generates Gray code sequences for given values of n.

```bash
go run main.go dsa grayCode
```

**Output Example:**
```
Running DSA test for algorithm: grayCode
=== Gray Code Tests ===
n = 1, Result: [0 1]
n = 2, Result: [0 1 3 2]
n = 3, Result: [0 1 3 2 6 7 5 4]
n = 4, Result: [0 1 3 2 6 7 5 4 12 13 15 14 10 11 9 8]
```

#### 2. Sum of Distances in Tree
Calculates the sum of distances from each node to all other nodes in a tree.

```bash
go run main.go dsa sumOfDistancesInTree
```

**Output Example:**
```
Running DSA test for algorithm: sumOfDistancesInTree
=== Sum of Distances in Tree Tests ===
Test case 1:
n = 6, edges = [[0 1] [0 2] [2 3] [2 4] [2 5]]
Result: [8 12 6 10 10 10]
```

#### 3. Find Length (Longest Common Subarray)
Finds the length of the longest common subarray between two arrays.

```bash
go run main.go dsa findLength
```

**Output Example:**
```
Running DSA test for algorithm: findLength
=== Find Length Tests ===
Test case 1:
nums1 = [1 2 3 2 1], nums2 = [3 2 1 4 7]
Result: 3
```

### Task 2: Authentication Service

The authentication service provides user registration, login, and JWT token management.

#### Running Locally

1. **Setup Database**: Ensure MySQL is running and create the database:
   ```sql
   CREATE DATABASE elotus;
   ```

2. **Configure Environment**: Update `config.yaml` with your database credentials:
   ```yaml
   MYSQL_URI: localhost:3306
   MYSQL_DATABASE: elotus
   MYSQL_USERNAME: your_username
   MYSQL_PASSWORD: your_password
   AUTHEN_SERVER_PORT: "8080"
   JWT_SECRET_KEY: your_jwt_secret
   JWT_REFRESH_SECRET_KEY: your_refresh_secret
   ```

3. **Run Authentication Server**:
   ```bash
   go run main.go authentication
   ```

The authentication server will start on `http://localhost:8080`

#### API Endpoints

- `POST /authentication/register` - User registration
- `POST /authentication/login` - User login
- `POST /authentication/refresh` - Refresh JWT token

#### Example API Calls

**Register a new user:**
```bash
curl --location 'http://localhost:8080/authentication/v1/register' \
--header 'Content-Type: application/json' \
--data '{
    "username":"kien.do",
    "password":"kien.do"
}'
```

**Login:**
```bash
curl --location 'http://localhost:8080/authentication/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"kien.do",
    "password":"kien.do"
}'
```

**Refresh JWT token:**
```bash
curl --location 'http://localhost:8080/authentication/v1/refresh' \
--header 'Content-Type: application/json' \
--data '{
    "refresh_token":"your_refresh_token_here"
}'
```

### Task 3: Upload File Service

The upload service provides file upload functionality with authentication middleware.

#### Running Locally

1. **Setup Database**: Ensure MySQL is running (same database as authentication service)

2. **Configure Environment**: Update `config.yaml` with your database credentials:
   ```yaml
   MYSQL_URI: localhost:3306
   MYSQL_DATABASE: elotus
   MYSQL_USERNAME: your_username
   MYSQL_PASSWORD: your_password
   UPLOAD_SERVER_PORT: "8081"
   JWT_SECRET_KEY: your_jwt_secret
   ```

3. **Run Upload Server**:
   ```bash
   go run main.go upload
   ```

The upload server will start on `http://localhost:8081`

#### API Endpoints

- `POST /upload` - Upload file (requires authentication)


#### Example API Calls

**Upload a file (requires Bearer token from authentication):**
```bash
curl --location 'http://localhost:8081/upload/v1/upload' \
--header 'Authorization: Bearer your_jwt_token_here' \
--form 'data=@"path/to/your/file.png"'
```

**Example with actual token:**
```bash
curl --location 'http://localhost:8081/upload/v1/upload' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImtpZW4uZG8iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNzU0MzY4ODMzLCJuYmYiOjE3NTQzNjc5MzMsImlhdCI6MTc1NDM2NzkzM30.oSoJOhVWPTO-A7qh7CjzBmB0chzRaWUbOoxQzrqqmRs' \
--form 'data=@"/Users/kiendo/Desktop/Screenshot 2025-04-30 at 09.35.30.png"'
```

## Docker Setup (Recommended)

For easy setup and consistent environment, use Docker. See [docker/README.md](docker/README.md) for complete Docker setup instructions.

**Quick Start:**
```bash
cd docker
./run.sh
```

This will start all services:
- MySQL Database on port 3306
- Authentication Server on port 8080  
- Upload Server on port 8081

## Available Commands

### Current Commands
- `go run main.go dsa [algorithm]` - Run DSA algorithm tests
  - `grayCode` - Generates Gray code sequences
  - `sumOfDistancesInTree` - Calculates sum of distances in a tree
  - `findLength` - Finds longest common subarray length
- `go run main.go authentication` - Run authentication server
- `go run main.go upload` - Run upload server

## Error Handling

If you provide an invalid algorithm name, the program will display an error message with the available options:

```bash
go run main.go dsa invalidAlgorithm
```

**Output:**
```
Unknown functions: invalidAlgorithm. Available functions: grayCode, sumOfDistancesInTree, findLength
```

## Project Structure

```
elotus/
├── cmd/
│   ├── authentication/      # Authentication service
│   │   ├── handlers/       # HTTP handlers
│   │   ├── services/       # Business logic
│   │   ├── repositories/   # Data access layer
│   │   └── main.go        # Authentication server entry point
│   ├── upload/             # Upload service
│   │   ├── handlers/       # HTTP handlers
│   │   ├── services/       # Business logic
│   │   ├── repositories/   # Data access layer
│   │   ├── middleware/     # Authentication middleware
│   │   └── main.go        # Upload server entry point
│   ├── dsa/
│   │   └── main.go        # DSA algorithm implementations
│   └── root.go            # CLI command definitions
├── pkg/                   # Shared packages
│   ├── cfg/              # Configuration management
│   ├── db/               # Database connection
│   ├── jwt/              # JWT utilities
│   ├── logger/           # Logging utilities
│   └── server/           # HTTP server utilities
├── docker/               # Docker configuration
│   ├── docker-compose.yaml
│   ├── run.sh
│   └── README.md
├── config.yaml           # Application configuration
├── main.go              # Application entry point
├── go.mod               # Go module file
└── README.md           # This file
```
## Task Progress

- ✅ **Task 1**: DSA Algorithms - Complete
- ✅ **Task 2**: Authentication Service - Complete
- ✅ **Task 3**: Upload Service - Complete 