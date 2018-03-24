package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/whaangbuu/go-protobuf/customer"
)

func main() {
	firstName := "Richard"
	lastName := "Burk"
	username := "whaangbuu"

	customer := &pb.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}

	data, err := proto.Marshal(customer)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &pb.Customer{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// Now test and newTest contain the same data.
	if customer.GetUsername() != newTest.GetUsername() {
		log.Fatalf("data mismatch %q != %q", customer.GetUsername(), newTest.GetUsername())
	}

}
