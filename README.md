# Go, Fiber, and PostgreSQL Starter Project

This is a starter template for building RESTful APIs in Go using the [Fiber](https://gofiber.io/) web framework, [GORM](https://gorm.io/) as the ORM for a [PostgreSQL](https://www.postgresql.org/) database, and JWT for authentication.

This project provides a solid and structured foundation for building scalable and maintainable web applications and services.

## ✨ Features

- **Fast & Efficient**: Built with Fiber, a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go.
- **ORM Integration**: Uses GORM for elegant and efficient database interactions with PostgreSQL.
- **Environment-based Configuration**: Manages configuration via `.env` files for different environments (development, production, etc.).
- **Structured Project Layout**: Organized for scalability and easy navigation.
- **CRUD Example**: Includes a basic CRUD (Create, Read, Update, Delete) implementation to get you started.

##  Prerequisites

Before you begin, ensure you have the following installed:

- Go (version 1.18 or newer)
- PostgreSQL
- Docker (Optional, for running PostgreSQL in a container)

## 🚀 Getting Started

Follow these steps to get your development environment set up and running.

### 1. Clone the Repository

```bash
git clone https://github.com/kishansuvarna09/go-fiber-postgres.git
cd go-fiber-postgres
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up Environment Variables

Create a `.env` file in the root of the project by copying the example file:

```bash
cp .env.example .env
```

Now, open the `.env` file and update the `DB_URL` with your PostgreSQL connection string.

```env
# .env

# Example for a local PostgreSQL instance
DB_URL="host=localhost user=postgres password=your_password dbname=your_db port=5432 sslmode=disable"
```

### 4. Run the Database

Make sure your PostgreSQL server is running. If you are using Docker, you can start a PostgreSQL container with the following command:

```bash
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

### 5. Run Database Migrations

The application will automatically run database migrations on startup to create the necessary tables based on your GORM models.

### 6. Run the Application

You can now start the Fiber web server:

```bash
go run main.go
```

The server will start, and you should see output indicating it's listening on a specific port (e.g., `http://localhost:3000`).

## 📁 Project Structure

The project follows a standard layout for Go applications to keep the code organized and maintainable.

```
go-fiber-postgres/
├── controllers/   # Handlers for API routes
├── initializers/  # App initialization (DB, Env vars)
├── middleware/    # Custom middleware (e.g., auth)
├── models/        # GORM models and database structs
├── .env           # Environment variables (ignored by git)
├── .env.example   # Example environment variables
├── go.mod         # Go module definitions
├── go.sum         # Go module checksums
└── main.go        # Main application entry point
```

## 🛠️ Technologies Used

- Go
- Fiber
- GORM
- PostgreSQL
- godotenv

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for details.