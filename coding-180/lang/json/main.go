package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct{
	Id int				`json:"id"`
	Name string             `json:"name"`
	Price float64			`json:"price"`
}

type Order struct{
	//tag
	Id  string    				`json:"id"`
	Item   OrderItem  			`json:"item,omitempty"`  //omitempty   : 空字段不保留
	Quantity  int               `json:"quantity"`
	TotalPrice  float64			`json:"total_price"`
}
func main(){
	unmarshall()
	marshall()
}
func marshall(){
	o :=Order{
		Id: "1234",
		//Name: "learn go",
		Item: OrderItem{
			Id: 12,
			Name: "beer",
			Price: 23.32,

		},
		Quantity: 3,
		TotalPrice: 30.00,
	}
	//fmt.Printf("%+v\n", o)	
	b,err := json.Marshal(o)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func unmarshall(){
	s := `{"id":"1234","item":{"id":12,"name":"beer","price":23.32},"quantity":3,"total_price":30}`
	var  o  Order
	err :=json.Unmarshal([]byte(s),&o)
	if err!=nil {
		panic(err)
	}

	fmt.Printf("%+v\n", o)

}
