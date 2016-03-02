package main
import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Phone string
}

type Invoice struct {
	InvoiceCode string
	InvoiceNumber string
	DeptName string
	Sum int
	InvoiceDate string
}

func CreateMongo(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("invoice")

	newInvoice := Invoice{}

	newInvoice.InvoiceCode = r.Form.Get("invoiceCode")
	newInvoice.InvoiceNumber = r.Form.Get("invoiceNum")
	newInvoice.DeptName = r.Form.Get("companyName")
	newInvoice.Sum, err = strconv.Atoi(r.Form.Get("finalPrice"))
	newInvoice.InvoiceDate = r.Form.Get("billingDate")

	err = c.Insert(&newInvoice)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"invoiceCode": "110"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
