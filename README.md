# 🧠 Go REST API – JWT, MongoDB 7.0, Swagger y Docker

Microservicio backend desarrollado en **Go 1.24**, con autenticación **JWT**, persistencia en **MongoDB 7.0**, contraseñas seguras con **bcrypt**, y documentación interactiva **Swagger UI**.  
Listo para correr localmente o con **Docker Compose** en un solo comando.

---

## 🚀 Características principales

- 🧩 Arquitectura modular (`handlers`, `middleware`, `models`, `db`, `utils`)
- 🔐 Autenticación segura con JWT (expiración configurable)
- 💾 Base de datos: **MongoDB 7.0**
- 🔑 Contraseñas hasheadas con **bcrypt**
- 🧭 Documentación Swagger en `/swagger/index.html`
- 🐳 Deploy completo vía **Docker Compose**
- 🧰 Librerías principales:
  - [go-chi/chi](https://github.com/go-chi/chi)
  - [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
  - [joho/godotenv](https://github.com/joho/godotenv)
  - [swaggo/http-swagger](https://github.com/swaggo/http-swagger)
  - [mongo-driver oficial](https://github.com/mongodb/mongo-go-driver)

---

## ⚙️ Instalación local

### 1️⃣ Clonar el repositorio
```bash
git clone https://github.com/Fedesin/goapi-practica.git
cd goapi-practica
```

### 2️⃣ Variables de entorno
Crear `.env` en la raíz del proyecto:
```env
MONGO_URI=mongodb://mongo:27017
MONGO_DB=goapi
JWT_SECRET=claveultrasecreta
```

### 3️⃣ Instalar dependencias
```bash
go mod tidy
```

### 4️⃣ Generar documentación Swagger (si modificaste handlers)
```bash
swag init --parseDependency --parseInternal
```

> ⚠️ Este repo **incluye** la carpeta `docs/` para que Swagger funcione sin pasos extra.  
> Si no la incluís, el import `_ "…/docs"` va a fallar al compilar.

### 5️⃣ Ejecutar la API
```bash
go run main.go
```

Servidor:
```
http://localhost:8080
```

Swagger:
```
http://localhost:8080/swagger/index.html
```

---

## 🐳 Ejecución con Docker

### 1️⃣ Build y run
```bash
docker compose up --build
```

Swagger UI disponible en:  
👉 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

MongoDB disponible en:  
👉 `mongodb://localhost:27017`

### 2️⃣ Ver logs
```bash
docker logs -f goapi
```

### 3️⃣ Detener los servicios
```bash
docker compose down
```

---

## 📡 Endpoints principales

| Método | Endpoint    | Descripción                                   |
|:-----:|-------------|-----------------------------------------------|
| POST  | `/register` | Registro de nuevo usuario (guarda hash bcrypt)|
| POST  | `/login`    | Autenticación y generación de token JWT       |
| GET   | `/perfil`   | Perfil del usuario autenticado *(requiere JWT)* |

---

## 🧠 Ejemplo rápido

**Registro**
```json
POST /register
{
  "nombre": "Fede",
  "email": "fede@unlu.edu.ar",
  "password": "12345678"
}
```

**Login**
```json
POST /login
{
  "email": "fede@unlu.edu.ar",
  "password": "12345678"
}
```

Respuesta:
```json
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." }
```

**Perfil (protegido)**

Header:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

---

## 🧰 Estructura del proyecto

```
goapi-practica/
├── main.go
├── .env
├── Dockerfile
├── docker-compose.yml
├── docs/                  
└── internal/
    ├── db/
    │   └── mongo.go
    ├── handlers/
    │   ├── auth.go
    │   └── profile.go
    ├── middleware/
    │   └── auth.go
    ├── models/
    │   └── user.go
    └── utils/
        └── jwt.go
```

---

## 🧩 Próximas mejoras

- [ ] Validación de email duplicado en `/register`
- [ ] Reglas mínimas de password (8+ caracteres)
- [ ] Refresh tokens y expiración extendida
- [ ] Roles (`admin`, `user`) y permisos por ruta
- [ ] Tests unitarios (`testing`, `httptest`)
- [ ] Logging con `zap` o `logrus`
- [ ] Middleware CORS configurable

---

## 🧑‍💻 Autor

**Federico Simone**  
Linux SysAdmin & Backend Developer  
[fedesimone31@gmail.com](mailto:fedesimone31@gmail.com)

---

## 🏷️ Licencia

Este proyecto se distribuye bajo licencia MIT.
