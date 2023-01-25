# Golang Project Structure

## Overview📦

This project structure is designed for Golang projects and is based on the hexagonal architecture, with some modifications😵.

### `/cmd`

This directory serves as the entry point for the application, containing the main.go file. Keep the code in this directory minimal.

### `/internal`

This directory contains the functional code for the application, including routing and database interactions.

### `/database`

This directory contains the configuration file for the database used in the project. In this example, the database used is PostgreSQL.

### `internal/ports`

This directory contains two files: `repository.go` and `service.go`. These files define the functions of the application, with repository.go responsible for querying the database (Data Access Layer) and service.go containing the business logic.

### `internal/repositories`

This directory contains the ORM (such as gorm) used to query the database on function calls from the service layer.

### `internal/services`

This directory contains functions called by the repository layer and contains the business logic of the application. It is framework agnostic.

### `internal/handler`

This directory contains the application's routes and is not framework agnostic. Functions in this directory can only call the service layer.

### `internal/model`

This directory contains the database models with gorm and json extensions.

### `internal/middleware`

This directory contains the middleware functions used inside the handler directory

### `internal/utils`

This directory contains any repeating code, such as code for generating IDs or verifying tokens.

## Note🎯

This project structure is not suitable for small projects and may become unwieldy for large projects.
