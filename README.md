# URL Shortener Service

A simple URL shortener service built with Go and deployed on Vercel using serverless functions. This application allows users to shorten long URLs and use custom aliases. It uses Redis for data persistence across serverless function invocations.

---

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Installation and Setup](#installation-and-setup)
  - [1. Clone the Repository](#1-clone-the-repository)
  - [2. Initialize Go Modules](#2-initialize-go-modules)
  - [3. Install Dependencies](#3-install-dependencies)
- [Setting Up Redis](#setting-up-redis)
- [Configuring Vercel](#configuring-vercel)
  - [1. Install Vercel CLI](#1-install-vercel-cli)
  - [2. Log In to Vercel](#2-log-in-to-vercel)
  - [3. Set Environment Variables](#3-set-environment-variables)
- [Deployment](#deployment)
- [Usage](#usage)
  - [Shorten a URL](#shorten-a-url)
    - [With Auto-Generated Alias](#with-auto-generated-alias)
    - [With Custom Alias](#with-custom-alias)
  - [Redirect to Original URL](#redirect-to-original-url)
- [Examples](#examples)
- [Testing](#testing)
- [Project Notes](#project-notes)
  - [Data Persistence](#data-persistence)
  - [Serverless Function Limits](#serverless-function-limits)
  - [Performance Considerations](#performance-considerations)

---

## Features

- **Shorten Long URLs**: Convert long URLs into short, manageable aliases.
- **Custom Aliases**: Users can specify custom aliases for their URLs.
- **Persistent Storage**: Uses Redis to store URL mappings persistently.
- **Serverless Deployment**: Deployed on Vercel using Go serverless functions.
- **Input Validation**: Validates URLs and aliases for correctness.
- **Error Handling**: Provides clear error messages and HTTP status codes.

---

## Prerequisites

- **Go (Latest Version)**: Ensure Go 1.21 or higher is installed.
  ```bash
  go version
  ```
- **Redis Instance**: A Redis database accessible over the internet.
  - You can use services like [Redis Cloud](https://redis.com/redis-enterprise-cloud/) or [Upstash](https://upstash.com/) for serverless Redis instances.
- **Vercel Account**: Sign up for a free account at [vercel.com](https://vercel.com/signup).
- **Vercel CLI**: Install Vercel CLI globally.
  ```bash
  npm install -g vercel
  ```
  Note: Requires Node.js and npm.
- **Git**: For version control and deploying the project.
  ```bash
  git --version
  ```

---

## Project Structure

```
url-shortener-vercel/
├── api/
│   ├── redirect/
│         └── redirect.go
│   └── shorten.go
├── go.mod
├── go.sum
├── vercel.json
└── README.md
```

- **api/**: Contains the serverless function files.
  - **shorten.go**: Handles URL shortening requests.
  - **redirect.go**: Handles redirection to the original URL.
- **vercel.json**: Configuration file for Vercel deployment.
- **go.mod** and **go.sum**: Go module files for dependency management.

---

## Installation and Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/url-shortener-vercel.git
cd url-shortener-vercel
```

### 2. Initialize Go Modules

If you didn't clone from a repository with a `go.mod` file, initialize the module:

```bash
go mod init github.com/yourusername/url-shortener-vercel
```

### 3. Install Dependencies

Download the required Go packages:

```bash
go get ./...
```

---

## Setting Up Redis

You'll need access to a Redis instance to store URL mappings. You can use a cloud-hosted Redis service:

- **Option 1**: [Redis Cloud](https://redis.com/redis-enterprise-cloud/)
- **Option 2**: [Upstash](https://upstash.com/) (offers a free tier with limited usage)

**Note**: Obtain the Redis connection details:

- **REDIS_HOST**: The address and port of your Redis instance (e.g., `redis-12345.c250.us-east-1-3.ec2.cloud.redislabs.com:12345`)
- **REDIS_PASSWORD**: The password for your Redis instance.

---

## Configuring Vercel

### 1. Install Vercel CLI

```bash
npm install -g vercel
```

### 2. Log In to Vercel

```bash
vercel login
```

Follow the prompts to log in to your Vercel account.

### 3. Set Environment Variables

Set up the environment variables required for your application:

1. **Initialize the Project with Vercel**

   ```bash
   vercel
   ```

   Follow the prompts:

   - **Set up and deploy “url-shortener-vercel”?** `Yes`
   - **Which scope do you want to deploy to?** Select your username or team.
   - **Link to existing project?** `No`
   - **What’s your project’s name?** Press Enter to accept `url-shortener-vercel` or provide a custom name.
   - **In which directory is your code located?** `./` (the current directory)

2. **Set Environment Variables in Vercel Dashboard**

   - Go to your project on the Vercel dashboard.
   - Navigate to **Settings** > **Environment Variables**.
   - Add the following variables:

     | Key             | Value                | Environment | Encrypt |
     |-----------------|----------------------|-------------|---------|
     | `REDIS_HOST`    | Your Redis host      | All         | Yes     |
     | `REDIS_PASSWORD`| Your Redis password  | All         | Yes     |

---

## Deployment

Deploy the application to Vercel:

```bash
vercel --prod
```

After deployment, Vercel will provide a URL like `https://url-shortener-vercel.vercel.app`.

---

## Usage

### **Shorten a URL**

#### **With Auto-Generated Alias**

Send a `POST` request to `/shorten` with the JSON payload containing the `url`.

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.example.com"}' \
https://url-shortener-vercel.vercel.app/shorten
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
https://url-shortener-vercel.vercel.app/shorten
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
curl -I https://url-shortener-vercel.vercel.app/example
```

**Response Headers:**

```
HTTP/2 302
location: https://www.example.com
```

---

## Examples

### **1. Shortening a URL with Auto-Generated Alias**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "https://www.openai.com"}' \
https://url-shortener-vercel.vercel.app/shorten
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
https://url-shortener-vercel.vercel.app/shorten
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
curl -I https://url-shortener-vercel.vercel.app/openai
```

**Response Headers:**

```
HTTP/2 302
location: https://www.openai.com
```

---

## Testing

### **Invalid URL**

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"url": "invalid-url"}' \
https://url-shortener-vercel.vercel.app/shorten
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
https://url-shortener-vercel.vercel.app/shorten
```

**Response:**

```
Alias already in use
```

**Status Code:** `409 Conflict`

### **Accessing a Non-Existent Alias**

**Request:**

```bash
curl -I https://url-shortener-vercel.vercel.app/nonexistent
```

**Response:**

```
URL not found
```

**Status Code:** `404 Not Found`

---

## Project Notes

### **Data Persistence**

- The application uses Redis for data storage to ensure persistence across serverless function invocations.
- **Important**: Ensure your Redis instance is properly secured and accessible from Vercel's network.

### **Serverless Function Limits**

- Be aware of Vercel's serverless function limits:
  - **Execution Time**: Functions have a maximum execution time (default 10 seconds).
  - **Memory**: Default memory allocation is 1024 MB.
- Adjust these settings in `vercel.json` if necessary.

### **Performance Considerations**

- **Cold Starts**: Serverless functions may have cold starts, leading to initial latency.
- **Connection Reuse**: The Redis client is initialized at the package level to reuse connections where possible.
- **Concurrency**: Redis handles concurrent connections, but monitor your usage to avoid exceeding limits.

---


**Enjoy your new URL Shortener deployed on Vercel!**

For any questions or issues, please open an issue on the project's GitHub repository.

---


**Feel free to contribute to the project or customize it to suit your needs!**

If you encounter any issues or have suggestions for improvements, please create an issue or pull request on the project's GitHub repository.
