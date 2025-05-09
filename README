# ChronoTrack ⏱️
A blazing-fast, secure, and extensible task tracking backend powered by **Go**, **Fiber**, and **JWT**.

---

## 🔥 Features

- 🚀 Ultra-fast HTTP framework using **Fiber**
- 🔐 Secure JWT-based user authentication
- 🧱 Scalable project structure with MVC separation
- 🧠 Password encryption using bcrypt
- 🗃️ SQLite persistence via GORM
- 🛡️ Middleware-protected endpoints
- 🧪 Modular & testable setup
- 📄 RESTful API structure, Swagger-ready

---

## 🧠 Tech Stack

| Layer        | Tech              |
|--------------|-------------------|
| Language     | Go (Golang)       |
| Web Framework| Fiber             |
| ORM          | GORM              |
| Database     | SQLite            |
| Auth         | JWT               |
| Env Config   | godotenv          |

---

## 🏗️ Project Structure

```bash

chrono-track/
├── main.go                  # Entry point
├── controllers/             # Auth and task handlers
├── routes/                  # Route initialization
├── middleware/              # JWT middleware
├── models/                  # DB models (User, Task)
├── database/                # DB connection and migrations
├── .env                     # Configuration variables
└── go.mod / go.sum          # Go modules

````

---

## 🚀 Quickstart

```bash
git clone https://github.com/fzr365/chrono-track.git
cd chrono-track
go mod tidy
go run main.go
````

Then open Postman or curl:

```bash
# Register
curl -X POST localhost:3000/api/register -H "Content-Type: application/json" -d '{"username":"testuser", "password":"123456"}'

# Login (returns JWT)
curl -X POST localhost:3000/api/login -H "Content-Type: application/json" -d '{"username":"testuser", "password":"123456"}'

# Create Task (with token)
curl -X POST localhost:3000/api/tasks -H "Authorization: <JWT>" -H "Content-Type: application/json" -d '{"title":"Test task", "description":"desc"}'
```

---

## 📌 API Routes

| Method | Route           | Description         | Auth |
| ------ | --------------- | ------------------- | ---- |
| POST   | `/api/register` | Register new user   | ❌    |
| POST   | `/api/login`    | Login and get token | ❌    |
| GET    | `/api/tasks`    | List tasks          | ✅    |
| POST   | `/api/tasks`    | Create new task     | ✅    |

---

## 🧭 Roadmap

* [x] JWT Authentication
* [x] Task CRUD
* [ ] Task deadlines + filters (today, this week)
* [ ] Swagger API docs
* [ ] Docker support
* [ ] CI with GitHub Actions
* [ ] Unit tests for controller logic

---

## 📜 License

MIT License

---

## ✨ Credits

Built with ❤️ by [@fzr365](https://github.com/fzr365)

