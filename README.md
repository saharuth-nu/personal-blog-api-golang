# Personal Blogging Platform API

## 📖 Description

Personal Blogging Platform API เป็นโปรเจค RESTful API ที่พัฒนาโดยใช้ Golang (Fiber Framework) สำหรับจัดการระบบ Blog ส่วนตัว รองรับการทำงาน CRUD (Create, Read, Update, Delete) ของบทความ โดยเชื่อมต่อกับฐานข้อมูล SQL

### ✨ Features
- 📚 ดึงรายการบทความทั้งหมด (พร้อม filter ด้วยวันที่หรือ tag)
- 🔍 ดึงบทความตาม UID
- ✍️ สร้างบทความใหม่
- 🗑️ ลบบทความตาม UID
- ♻️ แก้ไขบทความตาม UID

## 🏗️ API Endpoints

Base Path: /api/v1

| Method | Endpoint        | Description                                               |
| ------ | --------------- | --------------------------------------------------------- |
| GET    | `/articles`     | คืนค่ารายการบทความทั้งหมด (รองรับ filter เช่นวันที่, tag) |
| GET    | `/article/:uid` | คืนค่าบทความเดียวตาม `uid`                                |
| POST   | `/article`      | สร้างบทความใหม่                                           |
| DELETE | `/article/:uid` | ลบบทความตาม `uid`                                         |
| UPDATE | `/article/:uid` | อัพเดทบทความตาม `uid`                                         |

## ⚙️ Installation

1. Clone repository

```bash
git clone https://github.com/saharuth-nu/personal-blog-api-golang.git
cd personal-blog-api-golang
```

2. ติดตั้ง dependency

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

สร้างไฟล์ .env ที่ root ของโปรเจค

```env
# Example environment variables
# TODO: เพิ่มรายละเอียดเช่น DB connection, App port ฯลฯ
PORT=8080
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
```

## 📂 Project Structure

```plaintext
personal-blog-api/
├── cmd/
│   └── server/
│       └── main.go         # Entry point
├── core/
│   ├── handlers/           # Controllers / Handlers
│   ├── models/             # Database models
│   ├── repositories/       # Data access logic
│   └── services/           # Business logic
├── pkg/                    # Utilities / helpers
├── go.mod
├── go.sum
└── README.md
```
