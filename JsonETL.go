package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

//Udata exported data
type jf struct {
	Bin, Name string
}

// type Udata struct {
// 	BIN           string `json:"BIN"`
// 	Name          string `json:"NAME"`
// 	Constructyear string `json:"CNSTRCT_YR"`
// 	LastModDate   string `json:"LSTMODDATE"`
// 	LastStateTyp  string `json:"LSTSTATYPE"`
// 	GeomSource    string `json:"GEOMSOURCE"`
// 	DoittID       string `json:"DOITT_ID"`
// 	Heightroof    string `json:"HEIGHTROOF"`
// 	Featcode      string `json:"FEAT_CODE"`
// 	GroundLev     string `json:"GROUNDELEV"`
// 	ShapeArea     string `json:"SHAPE_AREA"`
// 	ShapeLen      string `json:"SHAPE_LEN"`
// 	BaseBbl       string `json:"BASE_BBL"`
// 	MplutoBbl     string `json:"MPLUTO_BBL"`
// }

func main() {

	connectionURL := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(connectionURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	jsonFle, err := os.Open("building_data.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFle.Close()

	r, err := ioutil.ReadAll(jsonFle)

	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		return
	}

	var v []interface{}

	collection := client.Database("test").Collection("jsonbin")

	json.Unmarshal(r, &v)

	if collection.InsertMany(context.TODO(), v); err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("I am here!")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")

}
