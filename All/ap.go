package All

import (
	//"context"
	"encoding/json"
	"fmt"
	"strconv"

	//	"net/http"
	"github.com/baudhik/departmental-stores/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10/translations/id"
)

type Stores struct {
	Id   int    `json:"Id"`
	Name string `json:"storeName"`
	Add  string `json:"address"`
}

func GetStores(c *gin.Context) {
	db := config.Connect()
	selDB, err := db.Query("SELECT * FROM GD.Stores ORDER BY Id DESC")
	if err != nil {
		panic(err.Error())
	}
	str := Stores{}
	res := []Stores{}
	for selDB.Next() {
		var id int
		var name, add string
		err = selDB.Scan(&id, &name, &add)
		if err != nil {
			fmt.Println("No stores or error loading them")
			panic(err.Error())
		}
		str.Id=id
		str.Name = name
		str.Add = add
		res = append(res, str)
	}
	_, e := json.Marshal(res)
	if e != nil {
		fmt.Println("DAta is not converted in Marshall")
	}

	for _, emp := range res {
		fmt.Println(emp.Id, emp.Name, emp.Add)
	}
	c.JSON(200, gin.H{"all fields": res})
}

func StoreID(c *gin.Context) {
	db := config.Connect()
	//I:= c.GetInt("id")
	clusterID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("no table with that id")
	}
	// It,err :=json.Unmarshal()
	//i,err:=strconv.Atoi(I)
	fmt.Println(clusterID)
	selDB, err := db.Query("SELECT Id,Name,Address From GD.Stores Where ID=?", clusterID)
	if err != nil {

		fmt.Println("No stores or error loading them")
	}
	str := Stores{}
	res := []Stores{}
	for selDB.Next() {
		var id int
		var name, add string
		err = selDB.Scan(&id, &name, &add)
		if err != nil {
			fmt.Println("No stores or error loading them")
			panic(err.Error())
		}
		str.Id = id
		str.Name = name
		str.Add = add
		res = append(res, str)
	}
	// ress, e := json.Marshal(str)
	// if e != nil {
	//	fmt.Println("DAta is not converted in Marshall")
	// }

	/*for _, emp := range res {

	}*/
	fmt.Println(res)
	c.JSON(200, gin.H{"all fields": res})
}

func DeleteStores(c *gin.Context) {
	db := config.Connect()
	cid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Problem in converting id")
	}
	_, e := db.Query("Delete From GD.Stores Where ID=?", cid)
	if e != nil {
		fmt.Println("unable to Delete record")
	}else{
	c.JSON(200, gin.H{"Successful": " "})
	}
}
func AddStore(c *gin.Context) {

	db := config.Connect()
	//ictx, _ := c.Get("context")
	//var s Stores
	// c.BindJSON(&s)
	//ctx := ictx.(context.Context)
	var jsonData map[string]interface{}

	x, _ := c.GetRawData()
	er := json.Unmarshal(x, &jsonData)
	if er != nil {
		fmt.Println("error in Converting", er.Error())
	}
	//fmt.Println(ctx)
	/*for k, v := range jsonData {
		switch c := v.(type) {
		case string:
		  fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
		  fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
		  fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	  }*/
	fmt.Println(jsonData["ID"])

	//fmt.Println(s)
	//fmt.Println(x)
	_, e := db.Query("Insert into GD.Stores(Id,Name,Address) values (?,?,?)", jsonData["ID"], jsonData["NAME"], jsonData["ADDRESS"])
	if e != nil {
		// 	err = json.Unmarshal(x, &jsonData)
		fmt.Println("NOt  able to inserting statement", e.Error())
	}else{
		c.JSON(200, gin.H{"Successful in adding new values":" ?",})}
}

// func CreateUser(c *gin.Context) {
// 	var user Models.User
// 	c.BindJSON(&user)
// 	err := Models.CreateUser(&user)
// 	if err != nil {
// 	 fmt.Println(err.Error())
// 	 c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 	 c.JSON(http.StatusOK, user)
// 	}
//    }
func UpdateStore(c *gin.Context){
	db :=config.Connect()
	var jsonData map[string]interface{}
	x, _ := c.GetRawData()
	err:=json.Unmarshal(x,&jsonData)
	if err!=nil{
		fmt.Println("Error in Marshalling",err.Error())
	}
	_,e:=db.Query("Update GD.Stores Set NAME=? WHERE ID =?",jsonData["NAME"],jsonData["ID"])
	if e != nil {
		// 	err = json.Unmarshal(x, &jsonData)
		fmt.Println("CAn't update", e.Error())
	}
	c.JSON(200, gin.H{"Updated new values":jsonData["ID"]})
}