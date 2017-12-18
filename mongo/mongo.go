package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "fmt"
)

type Collection struct {
	_id   string
	Data  string
}

type Filebody struct {
	data string `bson:"data"`
}

func mgoSaveIn(id string, buf string) error {

	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("sensor").C("tasklist")
    //str:=fmt.Sprintf("{\"_id\":\"%s\",\"data\":\"%s\"}",id,buf)
    err = c.Update(bson.M{"_id":id},bson.M{"_id":id,"data":buf})
	if err != nil {
		return err
	}
	return nil
}
/*
func getData(md5 string) ([]byte, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("sensor").C("tasklist")
	result := Filebody{}
	c.Find(bson.M{"md5": md5}).One(&result)
	buf, err := packdata([]byte(result.data))
	if err != nil {
		return nil, err
	}
	return buf, nil
}*/
func main(){
    err:= mgoSaveIn("11","{\"word\":\"Behind ervery beautiful thing,there is some kind of pain\"}")
    fmt.Println("err:",err)
}
