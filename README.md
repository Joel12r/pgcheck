# 🐘 # pgversion

A small Go library for verifying that a connected PostgreSQL instance meets a **minimum required version**.

This library is useful for enforcing version requirements across multiple services or projects, ensuring developers and CI environments use the correct PostgreSQL version.

---

## 🚀 Features

- Connects to PostgreSQL via a standard DSN.  
- Runs `SELECT version()` and extracts the **major version**.  
- Compares it against a **minimum version** you specify.  
- Returns an error if the version is too low.  
- Lightweight and dependency-free (only uses the official `lib/pq` driver).

---

## 📦 Installation

```bash
go get github.com/pgversion
