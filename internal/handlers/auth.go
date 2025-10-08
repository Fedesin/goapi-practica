package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/Fedesin/goapi-practica/internal/db"
	"github.com/Fedesin/goapi-practica/internal/models"
	"github.com/Fedesin/goapi-practica/internal/utils"
)

// RegisterHandler godoc
// @Summary Registro de usuario
// @Description Crea un nuevo usuario con contraseña encriptada
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   usuario body models.Usuario true "Datos del usuario"
// @Success 200 {object} models.Usuario
// @Router /register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Encriptar contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(nuevo.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al encriptar contraseña", http.StatusInternalServerError)
		return
	}
	nuevo.Password = string(hash)

	coll := db.GetCollection("usuarios")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := coll.InsertOne(ctx, nuevo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nuevo.ID = res.InsertedID.(primitive.ObjectID)
	nuevo.Password = "" // no devolver el hash

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nuevo)
}

// LoginHandler godoc
// @Summary Login de usuario
// @Description Verifica credenciales y devuelve token JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   credenciales body models.Usuario true "Email y Password"
// @Success 200 {object} map[string]string
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	coll := db.GetCollection("usuarios")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.Usuario
	err := coll.FindOne(ctx, bson.M{"email": creds.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Verificar hash
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Email, time.Hour)
	if err != nil {
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
