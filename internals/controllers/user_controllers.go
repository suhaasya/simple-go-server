package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request){
	
 	w.Write([]byte("User created"))	
}

func EditUser(w http.ResponseWriter, r *http.Request){
	
 	w.Write([]byte("User edited"))	
}

func GetUsers(w http.ResponseWriter, r *http.Request){

 	w.Write([]byte("User fetched"))	
}