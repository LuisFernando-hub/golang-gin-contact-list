# 📇 Golang Gin Contact List API

A simple and fast RESTful API built with **Go (Golang)** and the **Gin Web Framework** for managing a contact list.
This project demonstrates clean API design, CRUD operations, and basic project structure in Go.

---

## 🚀 Features

* Create, read, update, and delete contacts (CRUD)
* RESTful API structure
* Lightweight and fast using Gin
* JSON-based request/response
* Simple and easy-to-understand architecture (great for learning)

---

## 🛠️ Tech Stack

* **Go (Golang)**
* **Gin Web Framework**
* **In-memory storage or simple persistence (based on implementation)**

---

## 📁 Project Structure

```bash
.
├── main.go            # Application entry point
├── routes/            # API route definitions
├── handlers/       # Request handlers (business logic)
├── models/            # Data structures (Contact model)
├── config/           # Data storage layer (if applicable)
└── go.mod             # Go module file
```

---

## 📦 Installation

### 1. Clone the repository

```bash
git clone https://github.com/LuisFernando-hub/golang-gin-contact-list.git
cd golang-gin-contact-list
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the project

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### ➕ Create Contact

```
POST /contacts
```

**Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "123456789"
}
```

---

### 📋 Get All Contacts

```
GET /contacts
```

---

### 🔍 Get Contact by ID

```
GET /contacts/:id
```

---

### ✏️ Update Contact

```
PUT /contacts/:id
```

---

### ❌ Delete Contact

```
DELETE /contacts/:id
```

---

## 🧠 Purpose of This Project

This project was created to practice:

* Building REST APIs with Go
* Using the Gin framework
* Structuring Go projects properly
* Handling HTTP requests and JSON
* Basic backend architecture concepts

---

## 👨‍💻 Author

Built by [Luis Fernando](https://github.com/LuisFernando-hub)

---

If you want, I can also help you:

* turn this into a production-ready Go structure (DDD-lite)
* add PostgreSQL + GORM
* or dockerize it properly with multi-stage build
