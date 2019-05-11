package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB stores the information
type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

type bininfo struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Bin           string        `json:"bin" bson:"bin"`
	Name          string        `json:"name" bson:"name"`
	Constructyear string        `json:"constructyear" bson:"constructyear"`
	LastModDate   string        `json:"lastmoddate" bson:"lastmoddate"`
	LastStateTyp  string        `json:"laststatetyp" bson:"laststatetyp"`
	GeomSource    string        `json:"geomsource" bson:"geomsource"`
	DoittID       string        `json:"doittid" bson:"doittid"`
	Heightroof    string        `json:"heightroof" bson:"heightroof"`
	GroundLev     string        `json:"groundlev" bson:"groundlev"`
	ShapeArea     string        `json:"shapearea" bson:"shapearea"`
	ShapeLen      string        `json:"shapelen" bson:"shapelen"`
	BaseBbl       string        `json:"basebbl" bson:"basebbl"`
	MplutoBbl     string        `json:"mplutobbl" bson:"mplutobbl"`
}

//GetAllBinEndpoint gets all records
func (db *DB) GetAllBinEndpoint(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	var binall []bininfo
	err := db.collection.Find(bson.M{}).All(&binall)

	if err != nil {

		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("content-type", "application/json")
		response, _ := json.Marshal(binall)
		w.Write([]byte("Getting all records..." + "\n"))
		w.Write(response)
	}

}

// GetOneBinEndpoint gets a single record based on ID
func (db *DB) GetOneBinEndpoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	var binone bininfo
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&binone)

	if err != nil {

		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("content-type", "application/json")
		response, _ := json.Marshal(binone)
		w.Write([]byte("Getting one records..." + "\n"))
		w.Write(response)
	}

}

// UpdateOneBinEndpoint gets a single record based on ID
func (db *DB) UpdateOneBinEndpoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var binone bininfo

	putBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(putBody, &binone)

	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &binone})

	if err != nil {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("content-type", "text")

		w.Write([]byte("Updated successfully!"))

	}

}

// CreateOneBinEndpoint gets a single record based on ID
func (db *DB) CreateOneBinEndpoint(w http.ResponseWriter, r *http.Request) {

	var binone bininfo

	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &binone)

	binone.ID = bson.NewObjectId()

	err := db.collection.Insert(binone)

	if err != nil {

		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("content-type", "application/json")
		response, _ := json.Marshal(binone)
		w.Write(response)
		w.Write([]byte("Inserted successfully!"))
	}

}

// DeleteOneBinEndpoint delets a single record based on ID
func (db *DB) DeleteOneBinEndpoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})

	if err != nil {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("content-type", "text")

		w.Write([]byte("Deleted successfully!"))

	}

}

func main() {
	//handleRequest()

	fmt.Println("Starting application...")

	session, err := mgo.Dial("127.0.0.1")
	c := session.DB("test").C("bininfo")
	db := &DB{session: session, collection: c}

	if err != nil {

		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/bin/{id}", db.GetOneBinEndpoint).Methods("GET")
	r.HandleFunc("/binall", db.GetAllBinEndpoint).Methods("GET")
	r.HandleFunc("/binupd/{id}", db.UpdateOneBinEndpoint).Methods("PUT")
	r.HandleFunc("/bindel/{id}", db.DeleteOneBinEndpoint).Methods("DELETE")
	r.HandleFunc("/binadd", db.CreateOneBinEndpoint).Methods("POST")

	srv := &http.Server{

		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
