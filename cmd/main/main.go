package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nk-code-lab/EasyMoney/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterUser(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}