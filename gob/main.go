package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

//var xxx [1024]P

func main(){
	var buf bytes.Buffer
	var mmap = make(map[string]P)
	p1 := P{1,1,1,"ZHANG"}
	p2 := P{2,2,2,"lei"}
	mmap["hello"]=p1
	mmap["world"]=p2
	enc:=gob.NewEncoder(&buf)
	//for k,v:=range mmap{
	//	gob.Register(k)
	//	gob.Register(v.Name)
	//}
	err := enc.Encode(&mmap)
	if err!=nil{
		fmt.Printf("====%v\n",err)
		return
	}

	var bb = buf.Bytes()
	var xx = make(map[string]interface{})
	gob.NewDecoder(bytes.NewReader(bb)).Decode(&xx)
	fmt.Printf("%+v\n",xx)
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func xxx() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// Encode (send) some values.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Decode (receive) and print the values.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

	// Output:
	// "Pythagoras": {3, 4}
	// "Treehouse": {1782, 1841}
}
