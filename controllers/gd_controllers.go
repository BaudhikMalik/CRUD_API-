package controllers

import (
	//"encoding/json"
	"fmt"
	//"net/http"

	"github.com/baudhik/departmental-stores/models"
	//"github.com/baudhik/departmental-stores/models"
	// "github.com/gin-gonic/gin"
)

func GetStores() {
	Stores := models.GetAllStores()
	for _, gd := range Stores {
		fmt.Println(gd.Id, gd.Name, gd.Address)
	}
	//res, _ := json.Marshal(Stores)
	//w.Header().Set("Content-Type", "pkglication/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(res)
}
