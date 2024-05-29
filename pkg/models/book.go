package models

import(
	"github.com/jinzhu/gorm"
	"github.com/shams03/book_management_system_using_go_sql/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
  Name string `gorm:"json":"name"`
	Author string `gorm:"json":"author"`
}

func Init(){
	config.Connect();
	db :=config.GetDb();
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
   db.NewRecord(b)
	 db.Create(&b)
	 return b;
}

func  GetBooks() []Book{
	var Books []Book;
	db.Find(&Books)
	return Books;
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book;
	db:=db.Where("ID=?",Id).Find(&getBook);
	return &getBook,db;
}

func DeleteBook(Id int64) (Book){
	var book Book;
	db.Where("ID=?",Id).Delete(book)
	return book;
}




