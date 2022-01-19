package main

import (
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/gorilla/mux"
	"github.com/baudhik/departmental-stores/All"
	"github.com/gin-gonic/gin"
)

func main() {
	//All.GetStores()
	//fmt.Println(db)
	r := gin.Default()
	r.GET("/GetStores", All.GetStores)
	r.GET("/GetStoresbyId/:id", All.StoreID)
	r.DELETE("/DeleteStoresbyId/:id",All.DeleteStores)
	r.POST("/AddnewStore", All.AddStore)
	r.PUT("/UpdateStoreName",All.UpdateStore)
	r.Run(":9000")


}
