# Service Catalog API

This API provides a read-only service catalog that supports listing, filtering, sorting, 
and pagination of services, along with retrieving detailed information about specific 
services and their versions. It is designed to be lightweight, scalable, and easily 
extendable for future use cases.

## Endpoints

- **GET /services**: Retrieves a list of services with support for filtering, sorting, 
  and pagination.
- **GET /services/{id}**: Retrieves details of a specific service by its ID.
- **GET /services/{id}/versions**: Retrieves all versions of a specific service.

## Running the API

1. Clone the repository using `git clone <repository_url>`.
2. Navigate to the project directory with `cd service-catalog`.
3. Run the server using `go run main.go`.
4. Access the API at `http://localhost:8080`.

## Testing

Run unit tests using:

go test ./tests/...

## Future Enhancements

- Add CRUD operations for services.
- Integrate persistent storage with a database.
- Add authentication and authorization for secure access.
