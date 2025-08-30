# Personal Blogging Platform API

## 📖 Description

Personal Blogging Platform API is a RESTful API built with Golang (Fiber Framework) for managing a personal blogging system. It supports the full set of CRUD operations (Create, Read, Update, Delete) for blog articles and connects to an SQL database.

### ✨ Features
- 📚 Retrieve all articles (with filters such as publishing date or tag)
- 🔍 Retrieve an article by UID
- ✍️ Create a new article
- 🗑️ Delete an article by UID
- ♻️ Update an article by UID

## 🏗️ API Endpoints

Base Path: /api/v1

| Method | Endpoint        | Description                                           |
| ------ | --------------- | ----------------------------------------------------- |
| GET    | `/articles`     | Get all articles (supports filters such as date, tag) |
| GET    | `/article/:uid` | Get a single article by `uid`                         |
| POST   | `/article`      | Create a new article                                  |
| DELETE | `/article/:uid` | Delete an article by `uid`                            |
| PUT    | `/article/:uid` | Update an article by `uid`                            |

## ⚙️ Installation

1. Clone repository

```bash
git clone https://github.com/saharuth-nu/personal-blog-api-golang.git
cd personal-blog-api-golang
```

2. Install dependencies

```bash
go mod tidy
```

3. Run server

```bash
go run cmd/server/main.go
```

## 📌 Usage

### Get all articles

```bash
curl -X GET http://localhost:8080/api/v1/articles
```

**Response:**

```json
[
  {
    "article_id": "12345",
    "title": "My First Blog",
    "content": "Hello World!",
    "tag": "golang",
    "published_at": "2025-08-20"
  }
]
```

## 🔑 Environment Variables

Create a .env file at the root of the project:

```env
# Example environment variables
# TODO: Add details such as DB connection, App port, etc.
PORT=8080
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
JWT_SECRET=
```

## 📂 Project Structure

```plaintext
personal-blog-api/
├── cmd/
│   └── server/
│       └── main.go         # Entry point
├── config/
│   ├── database            # Database Configuration
│   ├── environment         # Environment Project
├── core/
│   ├── handlers/           # Controllers / Handlers
│   ├── models/             # Database models
│   ├── repositories/       # Data access logic
│   └── services/           # Business logic
├── utils/                  # Utilities / helpers
│   ├── response            # Http Response
├── go.mod
├── go.sum
└── README.md
```
