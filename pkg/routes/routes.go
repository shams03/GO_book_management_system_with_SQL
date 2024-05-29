package route

import(
	"github.com/shams03/book_management_system_using_go_sql/pkg/controllers"
	"github.com/gorilla/mux"
)

var Route=func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST");
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET");
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET");
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT");
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE");
   

}
