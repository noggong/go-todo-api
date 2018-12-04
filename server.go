package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	//port := 4040

	//http.HandleFunc("/", Index)
	//http.HandleFunc("/todos", TodoIndex)


	//log.Printf("Server starting on port %v\n", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	//db, err := sql.Open("mysql", "noggong:123456@tcp(localhost:3306)/GoTodo")

	//if err != nil {
	//	panic(err.Error())
	//}

	//insert, err := db.Query("insert into todos values (3, 'third', now());")

	//if err != nil {
	//	panic(err.Error())
	//}



	router := NewRouter()
	log.Fatal(http.ListenAndServe(":4040", router))

	//defer insert.Close()
}
