package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware" // alias para chi
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/Fedesin/goapi-practica/internal/handlers"
	"github.com/Fedesin/goapi-practica/internal/db"
	authmw "github.com/Fedesin/goapi-practica/internal/middleware" // alias para tu middleware JWT
	_ "github.com/Fedesin/goapi-practica/docs"                     // swagger docs
)

// @title API de Usuarios Go
// @version 1.0
// @description API de ejemplo con registro, login y JWT.
// @contact.name Federico Simone
// @contact.email fsimone@unlu.edu.ar
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ingresa el token JWT con el prefijo "Bearer ". Ej: "Bearer eyJhbGciOi..."
// @security BearerAuth
func main() {
	// Carga .env (JWT_SECRET)
	_ = godotenv.Load()
	// Si no existe, los paquetes internos se encargan de fallback.

	_, err := db.Connect()
    if err != nil {
        panic(err)
    }

	r := chi.NewRouter()
	r.Use(chimw.Logger)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Home -> redirige a Swagger
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusSeeOther)
	})

	// Rutas públicas
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	// Rutas protegidas
	r.Group(func(pr chi.Router) {
		pr.Use(authmw.AuthMiddleware) // evitar colisión de nombre con chi/middleware
		pr.Get("/perfil", handlers.PerfilHandler)
	})

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	_ = http.ListenAndServe(":8080", r)
}
