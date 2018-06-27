package main

import (
	pb "github.com/siman-man/golang-sandbox/protobuf_app/addressbook"
	"github.com/golang/protobuf/proto"
	"log"
	"io/ioutil"
	"fmt"
)

const FNAME = "person.pb"

func main() {
	encodePerson()
	person := decodePerson()

	fmt.Printf("person id = %d\n", person.Id)
}

func encodePerson() {
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(FNAME, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func decodePerson() *pb.Person {
	in, err := ioutil.ReadFile(FNAME)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.Person{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	return book
}
