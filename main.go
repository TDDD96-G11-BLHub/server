package main

//Test

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	fmt.Fprintln(c.Writer,
		`<!DOCTYPE html>
	<html>
	<form method="post" action="/name">
    <label for="firstname">First name:</label>
    <input type="text" name="firstname" /><br />
    <label for="lastname">Last name:</label>
    <input type="text" name="lastname" /><br />
    <input type="submit" />
	</form></html>`)
}

func postHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	name := c.Request.Form.Get("firstname")
	lastname := c.Request.Form.Get("lastname")
	fmt.Println(name, lastname)

	fmt.Fprintln(c.Writer, "Form submitted fine")
}

func main() {
	fmt.Println("Hello from BLHub server!")

	engine := gin.Default()

	engine.GET("/index.html", helloHandler)
	engine.POST("/name", postHandler)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
