package helper

import (
	"Dokon/models"
	
)
// bring MaxID For User
 

func MaxIdCatagory(PostArray []models.CatagoryModel) int {
	var maxID =  0
	 for i := 0; i <len(PostArray); i++ {
		if maxID < PostArray[i].ID {
			maxID = PostArray[i].ID
		}
	 }
	 return maxID+1

	 // Bring MaxID for Comment
}
	 // Bring MaxID for Post
	 
func MaxIdProduct(PostArray []models.ProdectModel) int {
	var maxID =  0
	 for i := 0; i <len(PostArray); i++ {
		if maxID < PostArray[i].ID {
			maxID = PostArray[i].ID
		}
	 }
	 return maxID+1

	 // Bring MaxID for Comment
}
	 // Bring MaxID for Post