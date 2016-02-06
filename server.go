package main

import (
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//<script data-main="/reqmod/login_main" language="JavaScript" defer async="true" src="js/r.js"></script>

func Home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {

	routes := mux.NewRouter()
	// http://localhost:8000/api/user/create?user=jaliat&first=Gao&last=Xiang&email=gxbit2006@163.com
	routes.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	// http://localhost:8000/api/user/1
	routes.HandleFunc("/api/user/{id}", GetUser).Methods("GET")

	http.Handle("/", routes)
//	routes.HandleFunc("/", Home)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
//	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages"))))
//	http.HandleFunc("/images/", fileUpload.DownloadPictureAction)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}