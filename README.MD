# Backend API for Stand With Dorayaki App

This repository contains the source code for the backend API of the Stand With Dorayaki App. The API is designed to manage and process data related to the Stand With Dorayaki application.

## API Documentation

You can find the API documentation at `/api/swagger/index.html`. This documentation provides comprehensive information about the available endpoints, request and response formats.

## Tech Stack

The backend API is built using the following tech stack:

- **Language:** Golang
- **Web Framework:** Echo
- **ORM:** GORM
- **Database:** PostgreSQL, Google Cloud Storage

## Getting Started

To run the API locally, you need to follow these steps:

1. Clone the repository:

```
git clone https://github.com/kadeksuryam-project-1/stand-with-dorayaki-BE.git
cd stand-with-dorayaki-BE
```

2. Create a `.env` file from the `.env.example` file:
```
cp .env.example .env
```
Update the `.env` file with your PostgreSQL and Google Cloud Storage credentials to allow the app to connect to these services.

3. Install the required dependencies:
```
go mod tidy
```
4. Migrate the database for initial seeding:
```
go run main.go migrate
```
5. Run the API server:

```
go run main.go server
```

The API server will now be running on `http://localhost:8080`. You can access the Swagger documentation at `http://localhost:8080/api/swagger/index.html`.

## Contact

If you have any questions, suggestions, or issues related to the Backend API for Stand With Dorayaki App, feel free to contact me at [kadeksuryam@gmail.com](mailto:kadeksuryam@gmail.com).

Happy coding!
