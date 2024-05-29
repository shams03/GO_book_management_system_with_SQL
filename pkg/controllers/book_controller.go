package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shams03/book_management_system_using_go_sql/pkg/models"
	"github.com/shams03/book_management_system_using_go_sql/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter,r *http.Request){
  newBooks :=models.GetBooks();
	res,_:=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
   w.Write(res);
}

 func GetBookById(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r);
  bookId:=vars["bookId"];
	ID,err :=strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
		bookDetails ,_:=models.GetBookById(ID);
    bookDetailsJson, _ := json.Marshal(bookDetails)
	   
   w.Header().Set("Content-Type","pkglication/json")
	 w.WriteHeader(http.StatusOK)
		w.Write(bookDetailsJson);
	}
 }

 func CreateBook(w http.ResponseWriter,r *http.Request){
	createdBook := &models.Book{}
	utils.ParseBody(r,createdBook)
	b:=createdBook.CreateBook();
	res,_:=json.Marshal(b);
	   
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	 w.Write(res);
 }

 func DeleteBook(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r);
	bId:=vars["bookId"];
  bookId,_:=strconv.ParseInt(bId,0,0);
	book:=models.DeleteBook(bookId)
	  res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	 w.Write(res);
 }

 func UpdateBook(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r);
	bId:=vars["bookId"];
	var NewBook= &models.Book{};  //think why we used & here
	utils.ParseBody(r,NewBook);
	bookId,_:=strconv.ParseInt(bId,0,0);
	book,db:=models.GetBookById(bookId);
	if NewBook.Name!="" && NewBook.Author!=""{
  if book.Name!=NewBook.Name{
		book.Name=NewBook.Name
	}
	if book.Author!=NewBook.Author{
		book.Author=NewBook.Author
	}}
	db.Save(&book);
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	 w.Write(res);
 }
