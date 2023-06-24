# Event-Calendar

When maintaining a calendar of events, it is important to know if an event overlaps with another event.

## Installation

1. Install the project dependencies:

   (Note: This project requires Go 1.19 or higher. If you don't have Golang installed, you can download it from https://go.dev/doc/install)

   ```
   make install
   ```

2. Configure the project:

   ```
   make configure
   ```

   Then edit `config/.env` with your desired configuration values.


3. Start the project:

   ```
   make start
   ```

   The project should now be running at http://localhost:3000.


4. Run unit tests:

   ```
   make test
   ```

5. Run test coverage:

   ```
   make cover
   ```

## Usage

### Endpoints

The API has the following endpoints:

#### POST /v1/event

Given a sequence of events, each having a start and end time, response will contain the sequence of all pairs of overlapping events.

##### Request

```
POST /v1/event
Content-Type: application/json

{
  "events":[[1, 2], [3, 5], [4, 7], [6, 8], [9, 10]]
}
```

##### Response

```
HTTP/1.1 201 Created
Content-Type: application/json

{
    "meta": {
        "offset": 0,
        "limit": 0,
        "record_count": 1
    },
    "records": [
        {
            "overlapping_events": [[1, 2], [3, 8], [9, 10]]
        }
    ]
}
```

### Error Handling

If an error occurs, the API will return a JSON object with an error message:

```
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
    "developer_message": "Key: 'OverlappingRequest.Events' Error:Field validation for 'Events' failed on the 'required' tag",
    "user_message": "Malformed request",
    "status_code": 400
}
```

Possible HTTP status codes for errors include:

- `400 Bad Request` for invalid request data
- `500 Internal Server Error` for server-side errors

## Swagger API Documentation

This project uses Swagger for API documentation. Swagger provides a user-friendly interface for exploring and testing the API.

To access the Swagger page:

1. Start the application if it's not already running.
2. Open a web browser and navigate to `http://localhost:3000/v1/swagger/index.html#/`.
3. The Swagger page should load, displaying a list of available endpoints.

From here, you can explore the available endpoints, see what parameters they require, and test them out.

If you have any questions or issues with the Swagger page, please refer to the API documentation or contact the project maintainers.