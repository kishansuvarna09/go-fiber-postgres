# Go Fiber PostgreSQL CRUD API

This is a simple RESTful API built with Go, [Fiber](https://gofiber.io/), and [GORM](https://gorm.io/) for performing CRUD (Create, Read, Update, Delete) operations on a `Book` resource. The application is designed to connect to a PostgreSQL database.

## Features

- **Create:** Add new books to the database.
- **Read:** Retrieve a list of all books or a single book by its ID.
- **Delete:** Remove a book from the database by its ID.
- Built with the high-performance Fiber web framework.
- Uses GORM as the ORM for database interactions.
- Configuration managed via environment variables.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (version 1.18 or newer recommended)
- PostgreSQL
- A running PostgreSQL instance. You can also use Docker to run a PostgreSQL container.

## Getting Started

Follow these instructions to get the project running on your local machine.

### 1. Clone the Repository

```bash
git clone <your-repository-url>
cd go-fiber-postgres
```

### 2. Install Dependencies

Download the Go modules required for the project.

```bash
go mod tidy
```

### 3. Configure Environment Variables

The application uses a `.env` file to manage environment variables. Create a `.env` file in the root of the project by copying the example file:

```bash
cp .env.example .env
```

Now, open the `.env` file and update the database connection details to match your local PostgreSQL setup.

### 4. Run the Application

You can run the application using the following command:

```bash
go run main.go
```

The server will start and listen for requests on `http://localhost:8080`.

## API Endpoints

- `POST /api/create_books`: Creates a new book.
- `GET /api/get_books`: Retrieves all books.
- `GET /api/get_books/:id`: Retrieves a single book by its ID.
- `DELETE /api/delete_books/:id`: Deletes a book by its ID.