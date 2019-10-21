// Package demo-api with Golang
// Designed by TRUNGLV
package main

import (
	"demo-api/driver"
	repoimpl "demo-api/repositories/repoimpl"
	"fmt"
	"net/http"
	"time"

	models "demo-api/model"

	"github.com/gin-gonic/gin"
)

// Const host, port, user, password, dbname
const (
	host     = "localhost"
	port     = "5432"
	user     = "levantrungits"
	password = "thanhtruc@011188"
	dbname   = "demoapidb"
)

// Struct DataPost
type DataPost struct {
	ID    int32  `json:"id"`
	Name  string `json:name`
	Email string `json:email`
}

// Main Function
func main() {
	route := gin.Default()
	route.GET("/users", GetMockDataUsers)
	route.POST("/users", PostDataUser)
	route.Run(":8080")
}

// GetMockDataUsers mock data return json
func GetMockDataUsers(c *gin.Context) {
	users := []models.User {
		{ID:1, Name:"Name 1", Gender:"Male", Email:"name1@gmail.com"},
		{ID:1, Name:"Name 1", Gender:"Male", Email:"name1@gmail.com"},
		{ID:1, Name:"Name 1", Gender:"Male", Email:"name1@gmail.com"},
	}
	// Do something...
	c.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"message":   "get list Users success.",
		"error-msg": nil,
		"data":      users,
	})
}

// PostDataUser return this data
func PostDataUser(c *gin.Context) {
	var dataPost DataPost
	c.BindJSON(&dataPost)
	// Do something...
	c.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"message":   "data is posted.",
		"error-msg": nil,
		"data":      dataPost,
	})
}

// Private getUsers Function
func getUsers(c *gin.Context) {
	start := time.Now()

	// CONNECT DB
	db, err := driver.Connect(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	defer db.SQL.Close()

	// NEW user repo
	userRepo := repoimpl.NewUserRepo(db.SQL)

	// GET ALL users -> show
	users, _ := userRepo.Select()
	for idx := range users {
		fmt.Println(users[idx])
	}

	// sleep 5s to LOG db
	//logInfo()
	//go logInfo()

	t := time.Now()
	elapsed := t.Sub(start)

	c.JSON(http.StatusOK, gin.H{
		"code":            http.StatusOK,
		"message":         "Success",
		"elapsed(second)": elapsed,
		"data":            users,
	})
}

// Private initData Function
func initData() {
	// TEST CONNECT DB
	db, err := driver.Connect(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	db.SQL.Ping()
	fmt.Println("connection postgres is OK.")

	// NEW user repo
	userRepo := repoimpl.NewUserRepo(db.SQL)

	// INIT 2 models user
	userF := models.User{
		ID:     1,
		Name:   "User Name 1",
		Gender: "Male",
		Email:  "us1@gmail.com",
	}
	userS := models.User{
		ID:     2,
		Name:   "User Name 2",
		Gender: "Male",
		Email:  "us2@gmail.com",
	}

	// INSERT 2 users
	userRepo.Insert(userF)
	userRepo.Insert(userS)

	// GET ALL users -> show
	users, _ := userRepo.Select()
	for idx := range users {
		fmt.Println(users[idx])
	}
}

// Private logInfo Function
func logInfo() {
	// insert to LOG table - db
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("===>> LOG DONE")
}
