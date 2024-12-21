package controllers

import (
	"fmt"
	"net/http"
	"simple-go-server/pkg/helper"
	"simple-go-server/pkg/structs"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	reqData := new(structs.AddUser)
	err := helper.ValidateBody(r, reqData)

	if err != nil {	
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(reqData.Email)
	fmt.Println(reqData.Name)
	fmt.Println(reqData.Password)
	fmt.Println(reqData.Role)

 	w.Write([]byte("User created"))	
}

func EditUser(w http.ResponseWriter, r *http.Request){
	reqData := new(structs.EditUser)
	err := helper.ValidateBody(r, reqData)

	if err != nil {	
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqData.Email != "" {
		fmt.Println(reqData.Email)
	}
	if reqData.Name != "" {
		fmt.Println(reqData.Name)
	}
	if reqData.Role != "" {
		fmt.Println(reqData.Role)
	}

 	w.Write([]byte("User edited"))	
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	queryJson, _ := helper.QueryToJSON(query)
	queryParams := new(structs.QueryParams)

	err := helper.ValidateQueryParams([]byte(queryJson), queryParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if queryParams.Page != "" {
		fmt.Println(queryParams.Page)
	}
	if queryParams.Order != "" {
		fmt.Println(queryParams.Order)
	}
	if queryParams.OrderBy != "" {
		fmt.Println(queryParams.OrderBy)
	}
	if queryParams.Size != "" {
		fmt.Println(queryParams.Size)
	}

 	w.Write([]byte("User fetched"))	
}