package main

import (
	"fmt"
	"net/http"
	"Dokon/Hendler"

)

func main (){

	fmt.Println("Server is working... :8080")


	http.HandleFunc("/catagory", Hendler.CatagoryHendler)
	// Catagory yozish boshlandi
	http.HandleFunc("/product", Hendler.PostHendler)
	http.HandleFunc("/manage", Hendler.ManageProductHendler )
	

	http.ListenAndServe(":8080", nil)

}