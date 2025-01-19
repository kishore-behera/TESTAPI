# Go REST API with MongoDB

A simple REST API built with Go (Golang) and MongoDB for performing CRUD operations on user data.

## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Postman (for testing the API)

## Project Structure

```
testapi/
├── docker-compose.yml
├── .env
├── main.go
├── models/
│   └── user.go
└── handlers/
    └── user_handler.go
```

## Setup and Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd testapi
```

2. Start MongoDB using Docker:
```bash
docker-compose up -d
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### 1. Create User
- **Method:** POST
- **URL:** `http://localhost:8080/users`
- **Headers:** 
  - Content-Type: application/json
- **Request Body:**
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
}
```
- **Success Response:**
  - Status: 200 OK
  - Returns the created user ID

### 2. Get All Users
- **Method:** GET
- **URL:** `http://localhost:8080/users`
- **Success Response:**
  - Status: 200 OK
  - Returns array of users

### 3. Get Single User
- **Method:** GET
- **URL:** `http://localhost:8080/users/{id}`
- **URL Parameters:** id=[string] (MongoDB ObjectID)
- **Success Response:**
  - Status: 200 OK
  - Returns user object

### 4. Update User
- **Method:** PUT
- **URL:** `http://localhost:8080/users/{id}`
- **URL Parameters:** id=[string] (MongoDB ObjectID)
- **Headers:**
  - Content-Type: application/json
- **Request Body:**
```json
{
    "name": "John Updated",
    "email": "john.updated@example.com",
    "age": 31
}
```
- **Success Response:**
  - Status: 200 OK
  - Returns update result

### 5. Delete User
- **Method:** DELETE
- **URL:** `http://localhost:8080/users/{id}`
- **URL Parameters:** id=[string] (MongoDB ObjectID)
- **Success Response:**
  - Status: 200 OK
  - Returns delete result

## Testing with Postman

1. Open Postman
2. Create a new Collection called "User API"
3. Add new requests for each endpoint:

### Create User Request
1. Create new POST request to `http://localhost:8080/users`
2. Go to Headers tab:
   - Add `Content-Type: application/json`
3. Go to Body tab:
   - Select "raw"
   - Choose "JSON" from dropdown
   - Add user JSON object

### Get Users Request
1. Create new GET request to `http://localhost:8080/users`
2. No additional headers or body needed

### Get Single User Request
1. Create new GET request to `http://localhost:8080/users/{id}`
2. Replace {id} with actual user ID from create response

### Update User Request
1. Create new PUT request to `http://localhost:8080/users/{id}`
2. Add same headers as Create User
3. Add updated user JSON in body

### Delete User Request
1. Create new DELETE request to `http://localhost:8080/users/{id}`
2. Replace {id} with actual user ID

## Environment Variables

The application uses following environment variables (stored in .env):
- `MONGODB_URI`: MongoDB connection string
- `DATABASE_NAME`: Name of the database

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Success
- 404: Resource not found
- 500: Server error

## MongoDB Connection

MongoDB runs in Docker container with following credentials:
- Host: localhost
- Port: 27017
- Username: admin
- Password: password

## Stopping the Application

1. Stop the Go server with Ctrl+C
2. Stop MongoDB container:
```bash
docker-compose down
```

## Data Model

User model structure:
```go
type User struct {
    ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name  string            `json:"name,omitempty" bson:"name,omitempty"`
    Email string            `json:"email,omitempty" bson:"email,omitempty"`
    Age   int               `json:"age,omitempty" bson:"age,omitempty"`
}
```
