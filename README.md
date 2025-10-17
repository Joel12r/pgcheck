# 🐘 pgcheck

A lightweight Go utility for checking PostgreSQL server versions.

`pgcheck` connects to a PostgreSQL instance, retrieves the version using `SELECT version()`, and validates it against a configurable minimum version requirement.  
It’s designed for use across multiple applications to enforce a consistent PostgreSQL version policy.

---

## 🚀 Features

- Connects to PostgreSQL and retrieves version info  
- Validates against a user-defined minimum version (e.g., 16, 17)  
- Can load database credentials from a YAML config file  
- Reusable as both a CLI tool and Go package  

---

## 📦 Installation

```bash
go get github.com/josephmoyenda/pgcheck
