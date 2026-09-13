# Students RESTful API (Golang)

A simple RESTful API for managing students, built with Go.

## Technologies Used
- **Go Standard Library**: For the core HTTP server (`net/http`) and structured logging (`log/slog`).
- **Config Management**: [cleanenv](https://github.com/ilyakaznacheev/cleanenv) is used to load configuration from YAML files.
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator) is used to validate incoming JSON requests.

## Local Setup

### Prerequisites
- [Go](https://golang.org/dl/) 1.21 or higher installed on your machine.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mamun-jsx/students-restful-api-golang.git
   cd students-restful-api-golang
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Application

1. Make sure you have your configuration file ready (e.g., `congig/local.yaml`). Example structure for your `local.yaml`:
   ```yaml
   env: "local"
   storage_path: "./storage"
   http_server:
     address: "localhost:8080"
   ```

2. Run the application passing the configuration path:
   ```bash
   go run cmd/students-api/main.go --config=congig/local.yaml
   ```

3. The server should now be running on `http://localhost:8080`.

### Testing the API

You can test the student creation endpoint using curl:
```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "age": 22}'
```
