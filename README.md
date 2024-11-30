# Song Library API

Song Library API is a RESTful service for managing songs. It allows users to perform CRUD operations on songs and provides a Swagger interface for documentation and testing.

---

## **Features**

- Add, update, delete, and fetch songs.
- PostgreSQL as the database.
- Swagger documentation for API endpoints.
- Lightweight and scalable using Gin framework.

---

## **Tech Stack**

- **Programming Language:** Go (Golang)
- **Framework:** Gin
- **Database:** PostgreSQL
- **Documentation:** Swagger

---

## **Setup and Installation**

### **Prerequisites**

- [Go](https://golang.org/doc/install) (version 1.22 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/)

### **1. Configure Environment Variables**

Create a `.env` file in the root directory with the following content:

``` env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=song_library
SERVER_PORT=8080
```

### **2. Install Dependencies**

Use the following command to install all required dependencies

``` bash
go mod tidy
```

### **3. Run Tests**

Unit and integration tests can be run using:

``` bash
go test ./...
```

### **4. Generate Swagger Documentation**

Generate Swagger files using:

``` bash
swag init --output ./docs --generalInfo ./cmd/server/main.go 
```

### **5. Run the Application**

Start the server with:

``` bash
go run ./cmd/server/main.go
```

### **Run with Docker**

You can containerize the application and its dependencies using Docker.

#### *1. Build Docker Image*

Run the following command to build the Docker image:

```bash
docker build -t song-library .
```

#### *2. Run the Container*

Start the container with the following command:

```bash
docker run --name song-library -p 8080:8080 --env-file .env song-library
```

---

### **Folder Structure**

```plaintext
.
├── cmd/
│   └── server/
│       └── main.go        # Entry point of the application
├── internal/
│   ├── handler/           # HTTP Handlers
│   ├── router/            # API Routes
│   ├── service/           # Business Logic
│   ├── repository/        # Data Access Layer
│   ├── db/                # Database Connection
│   ├── model/             # Data Models
├── pkg/
│   └── logger/            # Logging Utilities
├── docs/                  # Swagger Documentation
├── go.mod                 # Module Definition
├── go.sum                 # Dependency Checksum
├── .env                   # Environment Variables
└── README.md              # Project Documentation
```

### **API Endpoints**

| Method | Endpoint          | Description               |
|--------|-------------------|---------------------------|
| GET    | /api/v1/songs     | Retrieve all songs        |
| GET    | /api/v1/songs/:id | Retrieve a song by ID     |
| POST   | /api/v1/songs     | Add a new song            |
| PUT    | /api/v1/songs/:id | Update an existing song   |
| DELETE | /api/v1/songs/:id | Delete a song by ID       |

### **Swagger Address**

```text
/swagger/index.html 
```
