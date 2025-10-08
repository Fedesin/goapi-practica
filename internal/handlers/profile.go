package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Fedesin/goapi-practica/internal/db"
	"github.com/Fedesin/goapi-practica/internal/models"
)

// PerfilHandler godoc
// @Summary Perfil del usuario
// @Description Devuelve información del usuario autenticado
// @Tags Perfil
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} models.Usuario
// @Router /perfil [get]
func PerfilHandler(w http.ResponseWriter, r *http.Request) {
	email, _ := r.Context().Value("email").(string)

	coll := db.GetCollection("usuarios")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.Usuario
	err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	user.Password = "" // no devolvemos el hash

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
