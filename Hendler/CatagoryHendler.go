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

func CatagoryHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllCatagory(w, r)

	case "POST":
		CreateCatagory(w, r)
		// tugadi
	case "PUT":
		UpdateCatagory(w, r)
		// boshlandi
	case "DELETE":
		DeleteCatagory(w, r)
	}

}
func GetAllCatagory (w http.ResponseWriter, r *http.Request) {
	var NewCatagory models.GetModel
	json.NewDecoder(r.Body).Decode(&NewCatagory)




	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData)

	var ProductData []models.ProdectModel
	ProductByte, _ := os.ReadFile("db/Products.json")
	json.Unmarshal(ProductByte, &ProductData)

	var Found bool
	if NewCatagory.ID >= 1 {
		for i := 0; i < len(CatagoryData); i++ { 
			if CatagoryData[i].ID==NewCatagory.ID {
				fmt.Fprintln(w, "----------------------------------")
				fmt.Fprintln(w, "Catagory's ID: ", CatagoryData[i].ID)
				fmt.Fprintln(w, "Catagory's Name: ", CatagoryData[i].Name)
				fmt.Fprintln(w, "Catagory's CreatedAt: ", CatagoryData[i].CreatedAt)
				fmt.Fprintln(w, "Catagory's UpdataAt: ", CatagoryData[i].UpdatedAt)
				fmt.Fprintln(w, "Catagory's Products.. : ")
	
				for j := 0; j < len(CatagoryData[i].Products ); j++ {
					fmt.Fprintln(w, "  -----------------------------")
					fmt.Fprintln(w, "  Product's ID ", CatagoryData[i].Products[j].ID )
					
					fmt.Fprintln(w, "  Product's ProductType ", CatagoryData[i].Products[j].ProdectType)
					fmt.Fprintln(w, "  Product's Name ", CatagoryData[i].Products[j].Name)
					fmt.Fprintln(w, "  Product's Quantity ", CatagoryData[i].Products[j].Quantity  )
					fmt.Fprintln(w, "  Product's Price ", CatagoryData[i].Products[j].Price  )
					fmt.Fprintln(w, "  Product's Price ", CatagoryData[i].Products[j].Available  )
					fmt.Fprintln(w, "  Product's CreatedAt ", CatagoryData[i].Products[j].CreatedAt  )
					fmt.Fprintln(w, "  Product's UptadedAt ", CatagoryData[i].Products[j].UptadedAt )	
					Found= true
					
				}
			}  
			} 

			if Found==false {
				fmt.Fprintln(w ,"ID not Found")
				fmt.Println("ID not found ")
				w.WriteHeader(http.StatusNotFound)

			}
			

			
			
	} else {
		for i := 0; i <len(CatagoryData); i++ { 
			fmt.Fprintln(w, "--------------------------------")
			fmt.Fprintln(w, "Catagory's ID: ", CatagoryData[i].ID)
			fmt.Fprintln(w, "Catagory's Name: ", CatagoryData[i].Name)
			fmt.Fprintln(w, "Catagory's CreatedAt: ", CatagoryData[i].CreatedAt)
			fmt.Fprintln(w, "Catagory's UpdataAt: ", CatagoryData[i].UpdatedAt)
			fmt.Fprintln(w, "Catagory's Products.. : ")

			for j := 0; j < len(CatagoryData[i].Products ); j++ {
				fmt.Fprintln(w, "--------------------------------------")
				fmt.Fprintln(w, "  Product's ID ", CatagoryData[i].Products[j].ID )
				
				fmt.Fprintln(w, "  Product's ProductType ", CatagoryData[i].Products[j].ProdectType)
				fmt.Fprintln(w, "  Product's Name ", CatagoryData[i].Products[j].Name)
				fmt.Fprintln(w, "  Product's Quantity ", CatagoryData[i].Products[j].Quantity  )
				fmt.Fprintln(w, "  Product's Price ", CatagoryData[i].Products[j].Price  )
				fmt.Fprintln(w, "  Product's Price ", CatagoryData[i].Products[j].Available  )
				fmt.Fprintln(w, "  Product's CreatedAt ", CatagoryData[i].Products[j].CreatedAt  )
				fmt.Fprintln(w, "  Product's UptadedAt ", CatagoryData[i].Products[j].UptadedAt )	
			}
		}
		
	}
	 
}
		 

		// Ekranga Chiqarish kerak Hammasini
		// uning uchun for dan Foydalanamiz

		
		// finished time 12:05 20.01.2024

func CreateCatagory (w http.ResponseWriter, r *http.Request) {
	var NewCatagory models.CatagoryModel
	json.NewDecoder(r.Body).Decode(&NewCatagory)

	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData)

	NewCatagory.ID = helper.MaxIdCatagory(CatagoryData)
	
	NewCatagory.CreatedAt = time.Now()

	CatagoryData = append(CatagoryData, NewCatagory)

	res, _ := json.Marshal(CatagoryData)
	os.WriteFile("db/Catagory.json", res, 0)
	// ekranga chiqarish uchun
	fmt.Println("Catagory Created ", NewCatagory.ID)
	fmt.Fprintln(w, "Catagory Created ")
	json.NewEncoder(w).Encode(NewCatagory)
 // finished time 12:23 20:01.2024

}

func UpdateCatagory (w http.ResponseWriter, r *http.Request){
	var UpdateCatagory models.CatagoryModel
	json.NewDecoder(r.Body).Decode(&UpdateCatagory)
	

	
	var NewCatagory []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &NewCatagory)

	var CatagoryFound bool

	for l := 0; l < len(NewCatagory); l++ { 

		if UpdateCatagory.ID==NewCatagory[l].ID {
			if NewCatagory[l].Name !="" {
				NewCatagory[l].Name=UpdateCatagory.Name
			}
			  NewCatagory[l].UpdatedAt = time.Now()

			  

			  CatagoryFound = true
			  break
		}
		
	}
	if !CatagoryFound {
		fmt.Fprintln(w, "Catagory can not found with ID: ", UpdateCatagory.ID)
		fmt.Println("Catagory can not found with ID: ", UpdateCatagory.ID)
		w.WriteHeader(http.StatusNotFound)
		return

	}

	res, _ := json.Marshal(NewCatagory)
	os.WriteFile("db/Catagory.json", res, 0)
	json.NewEncoder(w).Encode(UpdateCatagory)
	// ekranga chiqarish uchun
	fmt.Println("Catagory Created ", UpdateCatagory.ID)
	fmt.Fprintln(w, "Catagory Created ")
// finished 13:04 20.01.2024

}

func DeleteCatagory (w http.ResponseWriter, r *http.Request) {
	var DeleteCatagory models.GetModel

	
	json.NewDecoder(r.Body).Decode(&DeleteCatagory )


	var CatagoryData []models.CatagoryModel
	CatagoryByte, _ := os.ReadFile("db/Catagory.json")
	json.Unmarshal(CatagoryByte, &CatagoryData)

	var CatagoryFound bool

	for i := 0; i < len(CatagoryData); i++ {
		if CatagoryData[i].ID == DeleteCatagory.ID {
			CatagoryData = append(CatagoryData[:i], CatagoryData[i+1:]...)
			CatagoryFound = true
		}

	}

	// endi Berilgan ID topilmagan holarni korib chiqamiz

	if CatagoryFound {
		fmt.Println("Catagorydeleted with ID: ", DeleteCatagory.ID)
		fmt.Fprintln(w, "Catagory deleted with ID: ", DeleteCatagory.ID)
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println("Catagory can not found with ID ", DeleteCatagory.ID)
		fmt.Fprintln(w, "Catagory can nor found with ID: ", DeleteCatagory.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// delete bolgan arrayni endi jsonga o'tkazish kerak
	res, _ := json.Marshal(CatagoryData)
	os.WriteFile("db/Catagory.json", res, 0)

	// ekranga User delete qilinganligi haqidagi malumot chiqarish
	// finished time 13:08 20.01.2024
}
