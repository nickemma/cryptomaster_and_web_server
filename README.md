# Exhibition Management API

An API built with Go for managing exhibitions. The API supports creating, retrieving, and managing exhibitions while enforcing strict validation to ensure data integrity.

---

## Features

- **Add New Exhibitions**: Allows users to add exhibitions with a unique auto-incrementing ID.
- **Retrieve All Exhibitions**: Fetch a list of all exhibitions in the system.
- **Retrieve an Exhibition by ID**: Retrieve details of a specific exhibition by its ID.
- **Strict JSON Validation**: Ensures only valid fields defined in the model are accepted, preventing unexpected or malicious data.

---

## API Endpoints

### 1. **Get All Exhibitions**

- **Endpoint**: `/api/data`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "Id": 1,
      "Title": "Coding in Go",
      "Description": "Go is blazingly fast"
    },
    {
      "Id": 2,
      "Title": "Coding in Java",
      "Description": "Java is Amazing"
    }
  ]
  ```

### 2. **Get Exhibition by ID**

- **Endpoint**: `/api/data/{id}`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "Id": 1,
    "Title": "Coding in Go",
    "Description": "Go is blazingly fast"
  }
  ```

### 3. **Add a New Exhibition**

- **Endpoint**: `/api/data/new`
- **Method**: `POST`
- **Request**:
  ```json
  {
    "Title": "Mastering Go",
    "Description": "Deep dive into Go programming"
  }
  ```
- **Response**:
  ```json
  {
    "Id": 4,
    "Title": "Mastering Go",
    "Description": "Deep dive into Go programming"
  }
  ```

### 4. **How to Run**

1. Clone the repository:

   ```bash
   git clone https://github.com/nickemma/exhibition-api.git
   cd exhibition-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the backend server:
   ```bash
   go run main.go | make run
   ```
   The server will start at `http://localhost:8080`.

.
├── apis
│ ├── HandleGet.go # Handles fetching all exhibitions
│ ├── HandleGetById.go # Handles fetching an exhibition by ID
│ ├── HandlePost.go # Handles adding a new exhibition
├── model
│ ├── exhibition.go # Defines the Exhibition model and data operations
├── main.go # Entry point of the application
└── README.md # Documentation

## Technologies Used

- **Programming Language**: Go (Golang)
- **Framework**: Standard Go HTTP package.
- **Data Storage**: In-memory storage (simulated with a slice)

## Improvements and Future Enhancements

- Add persistent storage (e.g., a database like PostgreSQL or MongoDB).
- Implement update and delete functionality for exhibitions.
- Add authentication and authorization for API endpoints.
- Enhance error handling and logging.
- Write more comprehensive tests for validation and edge cases.

## Contribution

- Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
