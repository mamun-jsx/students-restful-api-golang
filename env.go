// Package main (or project documentation)
// This file serves as documentation for the environment, tools, and libraries used in this project.
package docs

/*
Things Used in This Project:

1. Core Language Features:
   - Go (Golang) standard library for HTTP server (`net/http`)
   - Structured logging (`log/slog`) for structured and leveled logging.
   - `os`, `signal`, `syscall`, `context` for graceful shutdown handling.

2. Configuration Management:
   - `github.com/ilyakaznacheev/cleanenv`: Used to parse the `local.yaml` configuration file and map it into the `Config` struct.

3. Request Validation:
   - `github.com/go-playground/validator/v10`: Used to validate the incoming JSON request payloads against struct tags (e.g., `validate:"required,email"`).

4. Project Structure:
   - `cmd/students-api`: Contains the main entry point (`main.go`).
   - `internal/config`: Handles loading environment variables and config files.
   - `internal/student`: Contains HTTP handlers for the student endpoints.
   - `internal/types`: Contains shared data structures/types for the application.
   - `internal/utlis/response`: Contains utility functions to standardize JSON responses.

5. How configuration is passed:
   - The app looks for a `--config` flag or a `CONFIG_PATH` environment variable.
   - It reads a YAML file (e.g. `congig/local.yaml`) with settings for environment type, storage path, and HTTP server details.
*/
