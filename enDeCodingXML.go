package main

import (
	"encoding/xml"
	"fmt"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

func multiply(number int64) int64 {
	return number * number
}

func Encode() {
	m := Message{"Alice",
		"Hello",
		1294706395881547000}
	fmt.Println("\nEncode:\n", m)

	j, err := xml.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded JSON-Object:\n%s\n\n", j)
	fmt.Println("Alternative:")
	pretty, err := xml.MarshalIndent(m, "", "   ")
	fmt.Println(string(pretty))
}

func Decode() {
	b := []byte(`<name>Alice</name><body>Hello</body><time>1294706395881547000</time>`)
	fmt.Println("\nDecode:\n", string(b))

	var m Message
	err := xml.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded string:\n%+v\n\n", m)
}

func PartialDecode() {
	b := []byte(`<name>Bob</name><food>Pickle</food>`)

	var m Message
	err := xml.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	expected := Message{
		Name: "Bob",
	}
	fmt.Printf("Partial Decode:\n %s\n", string(b))
	fmt.Printf("expected: %+v\n", expected)
	fmt.Printf("m:        %+v\n", m)
}

func main() {
	Encode()
	Decode()
	PartialDecode()
}
