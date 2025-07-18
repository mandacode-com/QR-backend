# QR Backend Service

QR Backend service is backend API for handling QR code redirection requests. It is designed to be lightweight and efficient, providing a simple interface for redirecting users based on scanned QR codes.

---

## 📁 Project Structure

```
redirect/ # Backend API for QR code redirection
```

---

## 🌐 Services Overview

### 🔁 `qr-redirect`

A backend redirection service that processes scanned QR requests.

- **Tech Stack:** Go, Gin, ent (ORM)
- **Main Features:**
  - Accepts QR code IDs and redirects to the mapped URL
  - Responds with proper HTTP 301/302 status

---
