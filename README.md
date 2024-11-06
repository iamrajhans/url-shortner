# URL Shortener Service

A simple URL shortener service built with Go. This application allows users to shorten long URLs and use custom aliases. It provides endpoints to create shortened URLs and to redirect users to the original URLs.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Usage](#usage)
  - [Shorten a URL](#shorten-a-url)
    - [With Auto-Generated Alias](#with-auto-generated-alias)
    - [With Custom Alias](#with-custom-alias)
  - [Redirect to Original URL](#redirect-to-original-url)
- [Examples](#examples)
- [Testing](#testing)
- [Notes](#notes)
- [License](#license)

---

## Features

- Shorten long URLs with auto-generated aliases.
- Use custom aliases for shortened URLs.
- Redirect users from the shortened URL to the original URL.
- Input validation and error handling.

## Prerequisites

- **Go (Latest Version)**: Ensure Go 1.21 or higher is installed. [Download Go](https://go.dev/dl/)
  ```bash
  go version
  ```
  This should output `go version go 1.21.x` (or higher).


## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/iamrajhans/url-shortener.git
   cd url-shortener
   ```

2. **Initialize Go Modules**

   If you didn't clone from a repository with a `go.mod` file, initialize the module:

   ```bash
   go mod init github.com/iamrajhans/url-shortener
   ```

3. **Download Dependencies**

   Since we're using only the standard library, there are no external dependencies to download.

---

## Running the Application

1. **Build the Application**

   You can build the application into an executable binary:

   ```bash
   go build -o url-shortener
   ```

2. **Run the Application**

   ```bash
   ./url-shortener
   ```

   Alternatively, you can run it directly without building:

   ```bash
   go run main.go
   ```

3. **Server Output**

   The server will start on port `8080`:

   ```
   Server started at :8080
   ```

---

## Usage

### **Shorten a URL**

#### **With Auto-Generated Alias**

Send a `POST` request to `/shorten` with the JSON payload containing the `url`.

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.example.com"}' \
http://localhost:8080/shorten
```

**Response:**

```json
{
  "alias": "aB3dE1",
  "url": "https://www.example.com"
}
```

#### **With Custom Alias**

Provide a custom `alias` in the JSON payload.

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.example.com", "alias": "example"}' \
http://localhost:8080/shorten
```

**Response:**

```json
{
  "alias": "example",
  "url": "https://www.example.com"
}
```

### **Redirect to Original URL**

Access the shortened URL in your browser or via `curl`:

```bash
curl -I http://localhost:8080/example
```

**Response Headers:**

```
HTTP/1.1 302 Found
Location: https://www.example.com
```

---

## Examples

### **1. Shortening a URL with Auto-Generated Alias**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.openai.com"}' \
http://localhost:8080/shorten
```

**Sample Response:**

```json
{
  "alias": "XyZ123",
  "url": "https://www.openai.com"
}
```

### **2. Shortening a URL with a Custom Alias**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.openai.com", "alias": "openai"}' \
http://localhost:8080/shorten
```

**Sample Response:**

```json
{
  "alias": "openai",
  "url": "https://www.openai.com"
}
```

### **3. Accessing the Shortened URL**

Open in a web browser or use `curl`:

```bash
curl -I http://localhost:8080/openai
```

**Response Headers:**

```
HTTP/1.1 302 Found
Location: https://www.openai.com
```

---

## Testing

### **Invalid URL**

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "invalid-url"}' \
http://localhost:8080/shorten
```

**Response:**

```
Invalid URL
```

**Status Code:** `400 Bad Request`

### **Alias Already in Use**

If you try to use an alias that's already taken:

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.example.com", "alias": "openai"}' \
http://localhost:8080/shorten
```

**Response:**

```
Alias already in use
```

**Status Code:** `409 Conflict`

### **Accessing a Non-Existent Alias**

**Request:**

```bash
curl -I http://localhost:8080/nonexistent
```

**Response:**

```
URL not found
```

**Status Code:** `404 Not Found`

---

## Notes

- **Data Persistence**: This application uses in-memory storage. All data will be lost when the server restarts. For persistent storage, consider integrating a database like SQLite or Redis.
- **Concurrency**: The application uses `sync.Map` for thread-safe operations.
- **Security**: This is a basic implementation and does not include HTTPS support or advanced security features. For production use, implement HTTPS and additional security measures.
- **Customization**: Feel free to modify the code to add features like user authentication, analytics, or a front-end interface.

---


**Enjoy your new URL Shortener!**

For any questions or issues, please open an issue on the project's GitHub repository.
