package Hendler

import (
	
	"Dokon/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)




func ManageProductHendler (w http.ResponseWriter, r *http.Request){
switch r.Method {
// POST addProduct
case "POST": 
	AddProduct(w, r)
	//finished
// Remove Product
case "DELETE":
	RemoveProduct(w, r)
default:
	fmt.Fprintln(w, "Incorrect requst ")
	fmt.Println( "Incorrect requst ")
	w.WriteHeader(http.StatusBadRequest)
}
}

func  AddProduct (w http.ResponseWriter, r *http.Request){
	var NewProdect models.ManageProduct
	json.NewDecoder(r.Body).Decode(&NewProdect)

	var NewData []models.ProdectModel 
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &NewData)

	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData) 

	var Found bool
	for i := 0; i < len(NewData); i++ {
		if NewData[i].ID== NewProdect.ProductID {
			for l := 0; l < len(CatagoryData); l++ {
				if CatagoryData[l].ID==NewProdect.CatagoryID {
					CatagoryData[l].UpdatedAt=time.Now()
					CatagoryData[l].Products=append(CatagoryData[l].Products, NewData[i] )
					Found  = true 
					break
				}
				
			}
		}
		
	} 

	if !Found {
		fmt.Fprintln(w,"ID not Found ")
		fmt.Println("Id not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	res, _ := json.Marshal(CatagoryData)
	os.WriteFile("db/Catagory.json", res, 0)

	kal, _ := json.Marshal(NewData)
	os.WriteFile("db.Prodect.json", kal, 0)
	

	fmt.Fprintln(w, "Product Successfully Added")
	fmt.Println("Product Successfully Added")
	w.WriteHeader(http.StatusOK)


}
func RemoveProduct (w http.ResponseWriter, r *http.Request){
	var NewProdect models.ManageProduct
	json.NewDecoder(r.Body).Decode(&NewProdect)

	var NewData []models.ProdectModel 
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &NewData)

	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData) 
	
	var Found bool
	for i := 0; i < len(CatagoryData); i++ {
		if CatagoryData[i].ID==NewProdect.CatagoryID {
			for l := 0; l < len(CatagoryData[i].Products); l++ {
				if  CatagoryData[i].Products[l].ID==NewProdect.ProductID  {
					CatagoryData[i].Products=append(CatagoryData[i].Products[:l],CatagoryData[i].Products[l+1:]... )
					Found=true
					break
				}
				
			}
		}
		
	}
	if Found {
		fmt.Fprintln(w, "Successfully!!!")
	fmt.Println("Successfully")
	w.WriteHeader(http.StatusOK)

		
	} else {
		fmt.Fprintln(w, "ID not Found")
		fmt.Println("ID not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Successfully!!!")
	fmt.Println("Successfully")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(CatagoryData)
	os.WriteFile("db/Catagory.json", res, 0)

	kal, _ := json.Marshal(NewData)
	os.WriteFile("db.Prodect.json", kal, 0)
	

}