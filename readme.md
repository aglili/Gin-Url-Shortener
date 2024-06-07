# URL Shortener

This is a simple URL shortener application built in Go using the Gin web framework, SQLite for persistent storage, and Redis for caching.

## Features

- Shorten long URLs to shorter, more manageable ones.
- Resolve short URLs to their original long URLs.
- Uses SQLite for persistent storage and Redis for caching.

## Requirements

- Go (version 1.13 or higher)
- SQLite
- Redis

## Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/url-shortener.git
    cd url-shortener
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

3. **Run the application:**

    ```bash
    go run main.go
    ```

4. **Access the application:**

    Open your browser and go to `http://localhost:8080` to access the URL shortener.

## API Endpoints

- `POST /shorten`: Shorten a long URL. Example request body:
  ```json
  {
    "original_url": "https://www.example.com"
  }
- `GET /:short_code` Redirect to Long URL