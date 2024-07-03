package handlers

import (
	"GoTTH-commerce/internal/templates/layouts"
	"GoTTH-commerce/internal/templates/views"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := views.GuestIndex()

	err := layouts.MenuLayout(c, "My E-commerce using GoTTH stack").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
