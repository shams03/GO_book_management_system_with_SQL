package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shams03/book_management_system_using_go_sql/pkg/routes"
)

func main(){
     r :=mux.NewRouter();
		 route.Route(r);
		 http.Handle("/",r)
		 log.Fatal(http.ListenAndServe(":8000",r))
}