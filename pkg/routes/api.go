package routes

import (
	"github.com/gorilla/mux"
	"github.com/nk-code-lab/EasyMoney/pkg/controllers"
)

var RegisterUser = func(router *mux.Router) {
    router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
    router.HandleFunc("/users/", controllers.GetAllUser).Methods("GET")
    router.HandleFunc("/user/{user_id}", controllers.GetUserByID).Methods("GET")
    router.HandleFunc("/user/{user_id}", controllers.UpdateUser).Methods("PUT")
    router.HandleFunc("/user/{user_id}", controllers.DeleteUser).Methods("DELETE")
}