package api

import (
	"encoding/json"
	"net/http"
	"user-service-app/internals/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) RegisterRoutes() {
	http.HandleFunc("/users", h.handleUsers)
}


 
func (h *UserHandler) handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		users := h.userService.ListUsers()
		json.NewEncoder(w).Encode(users)
	


	case http.MethodPost:
		
		var input struct {
			Name string `json: "name"`
			Age int `json: "age"`
			
		}  


		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		user := h.userService.CreateUser(input.Name, input.Age)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)



	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}





}




/*
func handleUSers(h *UserHandler ) (w http.ResponseWriter, r *http.Request){


}

*/