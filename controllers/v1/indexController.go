package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/gorms"
	"test/models"
)

var (
	con = gorms.DB
)

func init() {
	con.AutoMigrate(&models.User{}, &models.Inventory{})
}

func Index(r *gin.Context) {
	r.HTML(http.StatusOK, "index.html", nil)
}

func Add(r *gin.Context) {

	var data models.Inventory
	var user models.User

	err1 := r.BindJSON(&data)
	if err1 != nil {
		return
	}
	con.Take(&user, 1)
	err := con.Model(&user).Association("Inventorys").Append(&data).Error()
	if err != "" {
		r.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		r.JSON(http.StatusOK, data)
	}

}

func Del(r *gin.Context) {
	id := r.Param("id")
	var inventory models.Inventory
	con.Take(&inventory, id)
	err := con.Delete(&inventory).Error
	if err != nil {
		return
	} else {
		r.JSON(http.StatusOK, inventory)
	}

}

func Update(r *gin.Context) {
	id := r.Param("id")
	var inventory models.Inventory
	con.Take(&inventory, id)
	switch inventory.Status {
	case 1:
		inventory.Status = 0
		con.Save(&inventory)
	default:
		inventory.Status = 1
		con.Save(&inventory)
	}

}

func Find(r *gin.Context) {

	var inventoryslist []models.Inventory

	if err := con.Debug().Preload("User").Where("user_id = ?", 1).Find(&inventoryslist).Error; err != nil {
		r.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		r.JSON(http.StatusOK, inventoryslist)
	}
}
