package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"umikyoRest/libs"
)

type CreateProductType struct {
	ProductName string    `json:"product_name" bson:"product_name"`
	Quantity    int64     `json:"quantity" bson:"quantity"`
	Price       int64     `json:"price" bson:"price"`
	Category    string    `json:"category" bson:"category"`
	SubCategory *string   `json:"sub,omitempty" bson:"sub,omitempty"`
	Images      *[]string `json:"images,omitempty" bson:"images,omitempty"`
	Description *string   `json:"desc,omitempty" bson:"desc,omitempty"`
	CreatedDate time.Time `json:"created_date" bson:"created_date"`
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var t CreateProductType
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Print(err)
	}
	// to print out as an object
	tJson, _ := json.Marshal(t)
	fmt.Println(string(tJson))
	//
	t.CreatedDate = time.Now()
	obj := libs.Field{FieldName: "Testing"}
	res, err := obj.Insert(t)
	if err != nil {
		libs.SendErrorResponse(w, "Failed to add Data", 405)
		return
	}
	json.NewEncoder(w).Encode(res)
}
