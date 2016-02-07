package main
import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/gorilla/mux"

	"io"
	"io/ioutil"
	"path"
	"strconv"
)

type User struct {
	ID    int 	 "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

// https://github.com/go-sql-driver/mysql/wiki/Examples

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//	r.ParseForm()
	//
	//	if r.Method == "POST" {
	//		for k, v := range r.Form {
	//			fmt.Print("key:", k, "; ")
	//			fmt.Println("val:", strings.Join(v, ""))
	//		}
	//		fmt.Println("username", r.Form["user"])
	//	} else if r.Method == "GET" {
	//
	//	}

	r.ParseForm()

	newUser := User{}

	newUser.Name = r.Form.Get("user")
	newUser.Email = r.Form.Get("email")
	newUser.First = r.Form.Get("first")
	newUser.Last = r.Form.Get("last")
	output, err := json.Marshal(newUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err := sql.Open("mysql", "root:rabbit@/ocelot")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
//	stmtIns, err := db.Prepare("INSERT INTO users (user_nickname, user_first, user_last, user_email)" +
//	" VALUES( ?, ?, ?, ? )")
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
//	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
//
//	res, err = stmtIns.Exec(newUser.Name, newUser.First, newUser.Last, newUser.Email)
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	} else {
//		res.
//	}

	res, err := db.Exec("INSERT INTO users (user_nickname, user_first, user_last, user_email)" +
		" VALUES( ?, ?, ?, ? )", newUser.Name, newUser.First, newUser.Last, newUser.Email)
	userId, lastInsertedError := res.LastInsertId()
	if lastInsertedError != nil {
		panic(lastInsertedError.Error())
		return
	}

	// read upload file
	file, header, err := r.FormFile("avatarFile")
//	err.Error()
	if err != nil {
		log.Println( err.Error())
//		io.WriteString(w, err.Error())
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	filename := path.Join("/tmp/avatars", strconv.FormatInt(userId,10) + path.Ext(header.Filename))
	log.Println(filename)
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return }
	io.WriteString(w, "Successful")


}

func GetUser(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	ReadUser := User{}

	db, err := sql.Open("mysql", "root:rabbit@/ocelot")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT id, user_nickname, user_first, user_last, user_email FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()


	err = stmtOut.QueryRow(id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}

	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "No such user")
	case err != nil:
		log.Fatal(err)
		fmt.Fprintf(w, "Error")
	default:
		output, _ := json.Marshal(ReadUser)
		fmt.Fprintf(w, string(output))
	}
}
