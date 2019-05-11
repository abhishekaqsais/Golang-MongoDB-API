package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type csvFile struct {
	Bin, Name, Constructyear, LastModDate, LastStateTyp, GeomSource                   string
	DoittID, Heightroof, featcode, GroundLev, ShapeArea, ShapeLen, BaseBbl, MplutoBbl string
}

// this is a comment

func main() {

	//connectionURL := "mongodb+srv://Abhishek:Stevens765326!@go678-4ke5r.mongodb.net/test?retryWrites=true"
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

	filename := "sample.csv"
	f, err := os.Open(filename)

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		panic(err)
	}

	var sum int

	collection := client.Database("test").Collection("bininfo")

	for _, line := range lines {

		sum++

		data := csvFile{
			Bin:           line[2],
			Name:          line[3],
			Constructyear: line[4],
			LastModDate:   line[5],
			LastStateTyp:  line[6],
			GeomSource:    line[7],
			DoittID:       line[8],
			Heightroof:    line[9],
			featcode:      line[10],
			GroundLev:     line[11],
			ShapeArea:     line[12],
			ShapeLen:      line[13],
			BaseBbl:       line[14],
			MplutoBbl:     line[15],
		}

		bininfo := csvFile{line[2], line[3], line[4], line[5], line[6], line[7], line[8], line[9], line[10], line[11], line[12], line[13], line[14], line[15]}

		insertResult, err := collection.InsertOne(context.TODO(), bininfo)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data.Bin + "\t " + data.Name + "\t " + data.Constructyear)
		fmt.Println(insertResult.InsertedID)
	}

	sum--

	fmt.Printf("%d", sum)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")

}
