package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TDDD96-G11-BLHub/dbman/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func graphHandler(client *mongo.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		fmt.Println("FETCH DATAPOINTS ")

		//s := db.GetAllDatabases(client)
		//collections := db.GetAllCollections(client, s[0])
		bytedata := db.FetchManyDocuments(client, "Sensordata", "deepoidsensor", bson.D{})
		//fmt.Println(bayson)
		// jayson := EJSON.serialize(bytedata)
		// fmt.Print(jayson)
		//fmt.Println(bytedata)

		var data []interface{}
		err := json.Unmarshal(bytedata, &data)

		if err != nil {
			fmt.Println("Error unmarshalling data: ", err)
		}
		//¡LLLETTTSSS FUCCKIINGG GGOOOOOOOOOOOOOOOOOOOOOOO!
		//Solid kommentar, 7/10
		res := data[:5]
		fmt.Println(res)

		// Include a custom message in the response
		c.JSON(http.StatusOK, gin.H{
			"datapoints": res,
		})
	}

	return gin.HandlerFunc(fn)
}
