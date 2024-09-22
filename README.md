# Phonebook API

This is a simple Phonebook API built with Go, using the Gin framework and GORM for database interactions. It allows users to manage contacts through a RESTful API.

## Getting Started

To get started with this project, clone the repository and follow the instructions below.

## Prerequisites

- Go (version 1.19 or higher)
- Docker (for running PostgreSQL)
- PostgreSQL (if running locally without Docker)

## Running the Project

1. **Clone the repository:**
   ```bash
   git clone https://github.com/AtaraZohar/phonebook.git
   cd phonebook
   
2. **Build and run the Docker containers:**
   Make sure you have Docker and Docker Compose installed. Run the following command to start the application:
   ```bash
   docker-compose up --build
Access the API: The API will be available at http://localhost:8080

## API Documentation

```
- GET /contacts   Retrieve a paginated list of contacts.

- GET /contacts/search?param=value   Search for contacts based on query parameters (first name, last name, phone number, address).

- POST /contacts   Create a new contact by sending contact details in the request body.

- PUT /contacts/{id}   Update an existing contact by ID, using the provided contact details in the request body.

- DELETE /contacts/{id}   Delete a contact by its ID.
```

### Get Contacts

- **Endpoint:** `GET /contacts`
- **Query Parameters:**
  - `page`: Page number (default is 1)
  - `limit`: Number of contacts per page (default is 10)
  
- **Response:**
  - `200 OK`: Returns a list of contacts.

### Search Contacts

- **Endpoint:** `GET /contacts/search`
- **Query Parameters:**
  - `first_name`: (optional)
  - `last_name`: (optional)
  - `phone_number`: (optional)
  - `address`: (optional)
  
- **Response:**
  - `200 OK`: Returns a list of matching contacts.
  - `404 Not Found`: If no contacts match the search criteria.
 
### Create Contact

- **Endpoint:** `POST /contacts`
- **Request Body:**
  ```json
  {
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "123-456-7890",
    "address": "123 Elm Street"
  }
- **Response:**
  - `201 Created`: Returns the created contact.

### Update Contact

- **Endpoint:** `PUT /contacts/{id}`

- **Request Body:**
  ```json
  {
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "123-456-7890",
    "address": "123 Elm Street"
  }
- **Response:**
  - `200 OK`: Returns the updated contact.
  - `400 Bad Request`: If the ID is invalid.
  - `404 Not Found`: If the contact does not exist.

### Delete Contact

  - **Endpoint:** `DELETE /contacts/{id}`
  - **Response:**
    - `204 No Content`: If the contact is deleted successfully.
    - `404 Not Found`: If the contact does not exist.




