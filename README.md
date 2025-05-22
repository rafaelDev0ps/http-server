# HTTP Server Lab

This project is a HTTP server implemented in Go. It is designed as a learning exercise to help understand how HTTP servers work at a low level, including handling TCP connections, parsing HTTP requests, and generating HTTP responses.

## Features

- Listens for HTTP requests on TCP port 4221
- Handles multiple endpoints:
  - `/` — Default response
  - `/echo/<text>` — Echoes back the text, supports gzip encoding if requested
  - `/user-agent` — Returns the User-Agent header sent by the client
  - `/files/<filename>` — 
    - `GET`: Returns the contents of a file from `/tmp/`
    - `POST`: Writes the request body to a file in `/tmp/`
- Properly handles HTTP headers, including `Connection: close` and `Accept-Encoding: gzip`
- Returns appropriate HTTP status codes (200, 201, 404, 500)

## Project Structure

```
go.mod
README.md
app/
  main.go                # Entry point, starts the server and handles connections
  controller/
    defaultController.go # Handles requests to `/`
    echoController.go    # Handles `/echo/<text>` endpoint
    filesController.go   # Handles `/files/<filename>` endpoint
    userAgentController.go # Handles `/user-agent` endpoint
  utils/
    utils.go              # Utilities for reading/writing files and compressing content
  request/
    request.go           # HTTP request parsing utilities
  response/
    response.go          # HTTP response formatting utilities
```

## How It Works

- The server listens for TCP connections.
- For each connection, it reads and parses the HTTP request.
- Based on the request path and method, it dispatches to the appropriate controller.
- Controllers build and send HTTP responses, including headers and body.
- Supports gzip compression for echo responses if requested by the client.

## Running the Server

1. Make sure you have Go installed (version 1.24.1 or later).
2. Run the server:

   ```sh
   go run app/main.go -port=8000
   ```

3. The server will listen on `0.0.0.0:4221` by default.

## Example Requests

- **Echo:**
  ```
  curl -v http://localhost:4221/echo/hello
  curl -v -H "Accept-Encoding: gzip" http://localhost:4221/echo/abc
  ```
- **User-Agent:**
  ```
  curl -v http://localhost:4221/user-agent -H "User-Agent: strawberry/strawberry"
  ```
- **File GET:**
  ```
  curl -i -v http://localhost:4221/files/test.json
  ```
- **File POST:**
  ```
  curl -v -X POST --data '{"animal": "gorilla"}' -H "Content-Type: application/json" http://localhost:4221/files/test.json
  curl -v --data "12345" -H "Content-Type: application/octet-stream" http://localhost:4221/files/file_123
  ```

## Learning Goals

- Understand how to work with raw TCP connections in Go
- Learn how HTTP requests and responses are structured
- Practice parsing and generating HTTP headers and bodies
- Implement basic file I/O and content encoding

---

Feel free to explore the code and experiment with adding new features or endpoints!