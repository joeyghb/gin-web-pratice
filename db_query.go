package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"sync"
)

var mu = &sync.Mutex{}

type Configuration struct {
	User string
	Pwd  string
}

type HomeSpent struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Date  string        "bson:`date`"
	Item  string        `bson:"item"`
	Cost  int           `bson:"cost"`
	User  string        `bson:"user"`
	Type  string        `bson:"type"`
	Store string        `bson:"store"`
	Memo  string        `bson:"memo"`
}

func GetConfig(c *gin.Context) *Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": err,
		})
	}

	return &configuration
}

func QueryHomeSpentOne(c *gin.Context) {

	var myConfig = GetConfig(c)
	user := myConfig.User
	Pwd := myConfig.Pwd

	var mydb = getDB()
	err := mydb.Login(user, Pwd)

	if err != nil {
		panic(err)
	}

	defer mydb.Session.Close()
	collections := mydb.C("homespent")

	homespent := HomeSpent{}
	User := "Joey"
	//id := "5c945428409ba03bd6b5ecd6"
	//objectId := bson.ObjectIdHex(id)

	//Query One
	//err = collections.Find(bson.M{"_id": objectId}).One(&homespent)
	err = collections.Find(bson.M{"user": User}).One(&homespent)

	if err != nil {
		println("find data error")
	}

	//ObjectId to Hex value
	strArticleID := bson.ObjectId(homespent.ID).Hex()

	fmt.Printf("ID_Hex : %s\n", homespent.ID)
	fmt.Printf("ID : %s\n", strArticleID)

	fmt.Printf("Date : %s\n", homespent.Date)
	fmt.Printf("Item : %s\n", homespent.Item)
	fmt.Printf("Cost : %d\n", homespent.Cost)
	fmt.Printf("User : %s\n", homespent.User)
	fmt.Printf("Type : %s\n", homespent.Type)
	fmt.Printf("Store : %s\n", homespent.Store)
	fmt.Printf("Memo : %s\n", homespent.Memo)

	c.String(http.StatusOK, fmt.Sprintf("find homespent : %s", homespent.Memo))
}

func QueryHomeSpentAll(c *gin.Context) {

	var myConfig = GetConfig(c)
	user := myConfig.User
	Pwd := myConfig.Pwd

	var mydb = getDB()
	err := mydb.Login(user, Pwd)

	if err != nil {
		panic(err)
	}

	defer mydb.Session.Close()
	collections := mydb.C("homespent")

	var homespents []HomeSpent
	User := "Joey"

	//Query all
	err = collections.Find(bson.M{"user": User}).All(&homespents)

	if err != nil {
		println("find data error")
	}

	//fmt.Printf("All Data\n")
	c.String(http.StatusOK, fmt.Sprintf("find all homespent :\n"))
	c.String(http.StatusOK, fmt.Sprintf("%v\n", homespents))

}

func InsertHomeSpent(c *gin.Context) {

	mu.Lock()
	defer mu.Unlock()

	var myConfig = GetConfig(c)
	user := myConfig.User
	Pwd := myConfig.Pwd

	var mydb = getDB()
	err := mydb.Login(user, Pwd)

	if err != nil {
		panic(err)
	}

	defer mydb.Session.Close()
	collections := mydb.C("homespent")

	homespent := HomeSpent{
		Date:  "2019/03/22",
		Item:  "育兒",
		Cost:  200,
		User:  "Yoyo",
		Type:  "Card",
		Store: "7-11",
		Memo:  "Toy",
	}

	//Insert
	err = collections.Insert(&homespent)

	if err != nil {
		println("insert error")
		c.String(http.StatusOK, fmt.Sprintf("Insert error\n"))
	}

	c.String(http.StatusOK, fmt.Sprintf("Insert Success !!\n"))
}

func getDB() *mgo.Database {

	session, err := mgo.Dial("128.199.143.113:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("home")
	return db
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello docker !")
	})

	r.GET("/getone", func(c *gin.Context) {
		QueryHomeSpentOne(c)
	})

	r.GET("/getall", func(c *gin.Context) {
		QueryHomeSpentAll(c)
	})

	r.GET("/insert", func(c *gin.Context) {
		InsertHomeSpent(c)
	})
}

func main() {
	r := gin.New()
	RegisterRoutes(r)
	r.Run(":18080")
}
