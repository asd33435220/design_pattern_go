package main

import "fmt"

type client struct {

}

func (c *client)insertLightingConnectorIntoComputer(com computer){
	fmt.Println("Client inserts Lighting connector into computer.")
	com.insertIntoLightingPort()
}