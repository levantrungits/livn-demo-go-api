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

const (
	host     = "localhost"
	port     = "5432"
	user     = "levantrungits"
	password = "thanhtruc@011188"
	dbname   = "demoapidb"
)

type DataPost struct {
	ID    int32  `json:"id"`
	Name  string `json:name`
	Email string `json:email`
}

func main() {
	route := gin.Default()

	route.GET("/users", getUsers)
	route.POST("/users", func(c *gin.Context) {
		var dataPost DataPost
		c.BindJSON(&dataPost)
		// Do something...
		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"data":   dataPost,
		})
	})

	route.Run(":8081")
}

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
		ID:     1,
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

func logInfo() {
	// insert to LOG table - db
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("===>> LOG DONE")
}
