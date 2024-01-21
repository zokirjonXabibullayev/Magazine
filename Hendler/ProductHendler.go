package Hendler

import (
	"Dokon/helper"
	"Dokon/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func PostHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllProduct(w, r)

	case "POST":
		CreateProduct(w, r)
		//finished
	case "PUT":
		UpdateProduct(w, r)
		//finished

	case "DELETE":
		DeleteProduct(w, r)
		//finished 13:59 20.01.2024
	}

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var NewProdect models.ProdectModel
	json.NewDecoder(r.Body).Decode(&NewProdect)

	var NewData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &NewData)

	NewProdect.ID = helper.MaxIdProduct(NewData) // shu yerda Id ozgartirish kerak
	NewProdect.Available = true
	NewProdect.CreatedAt = time.Now()
	NewProdect.UptadedAt = time.Now()
	NewData = append(NewData, NewProdect)

	res, _ := json.Marshal(NewData)
	os.WriteFile("db/Prodect.json", res, 0)
	// ekranga chiqarish uchun
	fmt.Println("Prodect Created ", NewProdect.ID)
	fmt.Fprintln(w, "prodect Created ")
	json.NewEncoder(w).Encode(NewProdect)

	//finished 13:24 20.01.2024
}

// starting Update prodect
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var UpdateProdect models.ProdectModel
	json.NewDecoder(r.Body).Decode(&UpdateProdect)

	var NewData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/Prodect.json") 
	json.Unmarshal(ProdectByte, &NewData)

	var ProdectFound bool
	// Update qilinyayotgan userni ID boyicha topih uchun for dan foydalanish
	for i := 0; i < len(NewData); i++ {
		if UpdateProdect.ID == NewData[i].ID {

			if UpdateProdect.ProdectType != "" {
				NewData[i].ProdectType = UpdateProdect.ProdectType

			}
			if UpdateProdect.Name != "" {
				NewData[i].Name = UpdateProdect.Name
			}
			if UpdateProdect.Quantity != 0 {
				NewData[i].Quantity = UpdateProdect.Quantity
			}
			if UpdateProdect.Price != 0 {
				NewData[i].Price = UpdateProdect.Price
			}
			// NewData[i].CreatedAt=NewData[i].CreatedAt

			NewData[i].UptadedAt = time.Now()
			NewData[i].Available = true

			ProdectFound = true
			break
		}

	}

	// Agar Berilgan ID topilmagan bolsa quyIDagi IF ishlaydi
	if ProdectFound == false {
		fmt.Fprintln(w, "Prodect can not found with ID: ", UpdateProdect.ID)
		fmt.Println("Prodect can not found with ID: ", UpdateProdect.ID)
		w.WriteHeader(http.StatusNotFound)
		return

	}

	res, _ := json.Marshal(NewData)
	os.WriteFile("db/Prodect.json", res, 0)
	json.NewEncoder(w).Encode(UpdateProdect)
	// ekranga chiqarish uchun
	fmt.Println("Prodect Created ", UpdateProdect.ID)
	fmt.Fprintln(w, "Prodect Created ")
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var DeleteProdect models.GetModel
	json.NewDecoder(r.Body).Decode(&DeleteProdect)

	var ProdectData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &ProdectData)

	var ProdectFound bool

	for i := 0; i < len(ProdectData); i++ {
		if ProdectData[i].ID == DeleteProdect.ID {
			ProdectData = append(ProdectData[:i], ProdectData[i+1:]...)
			ProdectFound = true
		}

	}

	// endi Berilgan ID topilmagan holarni korib chiqamiz

	if ProdectFound {
		fmt.Println("Prodect deleted with ID: ", DeleteProdect.ID)
		fmt.Fprintln(w, "Prodect deleted with ID: ", DeleteProdect.ID)
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println("Prodect can not found with ID ", DeleteProdect.ID)
		fmt.Fprintln(w, "Prodect can nor found with ID: ", DeleteProdect.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// delete bolgan arrayni endi jsonga o'tkazish kerak
	res, _ := json.Marshal(ProdectData)
	os.WriteFile("db.Prodect.json", res, 0)

	// ekranga User delete qilinganligi haqIDagi malumot chiqarish

}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var NewProdect models.GetModel
	json.NewDecoder(r.Body).Decode(&NewProdect)

	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData)

	// userni malumotlarini jsondan yechib olish
	var ProdectData []models.ProdectModel
	ProdectByte, _ := os.ReadFile("db/Prodect.json")
	json.Unmarshal(ProdectByte, &ProdectData)

	if NewProdect.ID >= 1 {

		for p := 0; p < len(ProdectData); p++ {
			if NewProdect.ID == ProdectData[p].ID {
				fmt.Fprintln(w, "--------------------------------------")
				fmt.Fprintln(w, "Product's ID", ProdectData[p].ID)
				fmt.Fprintln(w, "Product's ProdectType", ProdectData[p].ProdectType)
				fmt.Fprintln(w, "Product's Name", ProdectData[p].Name)
				fmt.Fprintln(w, "Product's Quantity", ProdectData[p].Quantity)
				fmt.Fprintln(w, "Product's Price", ProdectData[p].Price)
				fmt.Fprintln(w, "Product's available", ProdectData[p].Available)
				fmt.Fprintln(w, "Product's CreatedAt", ProdectData[p].CreatedAt)
				fmt.Fprintln(w, "Product's UptadedAt", ProdectData[p].UptadedAt)

			}

		}

	} else {
		for i := 0; i < len(ProdectData); i++ {
			fmt.Fprintln(w, "--------------------------------------")
			fmt.Fprintln(w, "Product's ID ", ProdectData[i].ID)
			fmt.Fprintln(w, "Product's ProdectType  ", ProdectData[i].ProdectType)
			fmt.Fprintln(w, "Product's Name  ", ProdectData[i].Name)
			fmt.Fprintln(w, "Product's Quantity ", ProdectData[i].Quantity)
			fmt.Fprintln(w, "Product's Price ", ProdectData[i].Price)
			fmt.Fprintln(w, "Product's available", ProdectData[i].Available)
			fmt.Fprintln(w, "Product's CreatedAt ", ProdectData[i].CreatedAt)
			fmt.Fprintln(w, "Product's UptadedAt", ProdectData[i].UptadedAt)

		}

	}
}

// finished
