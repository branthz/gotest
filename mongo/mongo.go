package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name    string
	Age     int
	Address string
}

type Filebody struct {
	data string `bson:"data"`
}

var (
	Session    *mgo.Session
	Collection *mgo.Collection
	DB         *mgo.Database
)

func end() {
	Session.Close()
}

func init() {
	var err error
	Session, err = mgo.Dial("localhost")
	if err != nil {
		panic("mongo connection failed")
	}

	Session.SetMode(mgo.Monotonic, true)
	DB = Session.DB("zltest")
	Collection = DB.C("users")
	err = Collection.EnsureIndex(mgo.Index{Key: []string{"name"}, Unique: true})
	if err != nil {
		panic(err)
	}
}

func create(u *User) error {
	return Collection.Insert(u)
}

func del(u *User) error {
	return Collection.Remove(bson.M{"_name": u.Name})
}

func update(id string, buf string) error {
	//str:=fmt.Sprintf("{\"_id\":\"%s\",\"data\":\"%s\"}",id,buf)
	err := Collection.Update(bson.M{"_id": id}, bson.M{"_id": id, "data": buf})
	if err != nil {
		return err
	}
	return nil
}

func getUser(u *User) error {
	err := Collection.Find(bson.M{"name": u.Name}).One(&u)
	return err
}

func main() {
	var u = User{
		Name:    "xiang",
		Age:     10,
		Address: "zhangzhou",
	}
	err := create(&u)
	/* also support map as input
	uu := make(map[string]interface{})
	uu["Name"] = "qian"
	uu["Age"] = 20
	uu["Address"] = "shanghai"
	*/
	fmt.Println("err:", err, u)

	end()
}
