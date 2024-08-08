package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateProductType struct {
	product_name string
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var t CreateProductType
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Print(err)
	}
	// obj := libs.Field{FieldName: "Testing"}
	fmt.Println(t)
	// res, err := obj.Insert(t)
	// if err != nil {
	// 	libs.SendErrorResponse(w, "Failed to add Data", 405)
	// 	return
	// }
	json.NewEncoder(w).Encode("res.InsertedID")
}
