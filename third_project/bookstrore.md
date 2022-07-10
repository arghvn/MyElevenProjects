for create bookstore management API :
1-database-mysql
2-gorm package for connection to database
3-json marshall , unmarshall
4-project structure 
5-gorilla mux

project structure contain two main section , pkg and cmd
pkg : config(APP.GO) ,controllers(book-controllers) ,models(book.go) ,routes(bookstore.routes) ,utils(utils.go)
cmd : main.go

A configuration file, often shortened to config file, defines the parameters, options, settings and preferences applied to operating systems (OSes), infrastructure devices and applications in an IT context. Software and hardware devices can be profoundly complex, supporting myriad options and parameters.

for routes ,
Create folder handlers and create files based on routes i.e apiHandler.go. Add all the methods with base url “/api” to it. To use the functions outside of the package (package is folder ), method names must start with capital letter. To import the files we will use the path which we used while creating the go.

Controllers in Go:
s who pick up Go will be familiar with the MVC model. Models, views, and controllers are great for abstracting away code, but sadly there aren’t many examples of how to use them in Go.

routes :
Get , book , getbooks
post , book , create a book 
Get , book/bookid , get book by id
put , book/bookid , updatebook
delete , book/bookid , delete book

the first parameters above are method
the third parameters above are controllers function

building bookstrore management system
so in terminal , cd third project , go mod init github.com/arghvn/main.go , go get "github.com/jinzhu/gorm" , go get "github.com/jinzhu/gorm/dialects/mysql"
 go get "github.com/gorilla/mux"