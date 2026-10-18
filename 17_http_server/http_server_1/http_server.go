package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	repo *UserRepository
}

type UserRepository struct {
	data map[string]User
}

type User struct {
	ID   string
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("%s - %s", u.ID, u.Name)
}

type Middleware func(http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("-> %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("<- %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}

func RequestIDMIddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = "new_generated_id"
		}
		w.Header().Set("X-Request-Id", requestID)

		ctx := context.WithValue(r.Context(), "request_id", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic recovered: %v", rec)
				writeError(w, http.StatusInternalServerError, "Internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-Id")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Chain(h http.Handler, middleWares ...Middleware) http.Handler {
	for i := len(middleWares) - 1; i >= 0; i-- {
		h = middleWares[i](h)
	}
	return h
}

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age,omitempty" validate:"omitempty,min=0,max=160"`
}

func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		data: make(map[string]User),
	}

	repo.data["1"] = User{ID: "1", Name: "Govnov"}
	repo.data["2"] = User{ID: "2", Name: "Pupkin"}
	return repo
}

func (ur *UserRepository) GetById(id string) (User, error) {
	user, ok := ur.data[id]
	if !ok {
		return User{}, fmt.Errorf("User with id %s not found", id)
	}
	return user, nil
}

func NewUserHandler(repo *UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := h.repo.GetById(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, user)

}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}

func (h *UserHandler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/users", h.CreateUser)
	mux.HandleFunc("GET /api/v1/users", h.ListUsers)
	mux.HandleFunc("GET /api/v1/users/{id}", h.GetUser)
	mux.HandleFunc("DELETE /api/v1/users/{id}", h.DeleteUser)
}

func main() {
	mux := http.NewServeMux()

	handler := Chain(
		mux,
		CORSMiddleware,
		RequestIDMIddleware,
		LoggingMiddleware,
		RecoveryMiddleware,
	)

	repo := NewUserRepository()
	apiHandler := NewUserHandler(repo)

	apiHandler.Routes(mux)

	http.ListenAndServe(":8080", handler)

}
