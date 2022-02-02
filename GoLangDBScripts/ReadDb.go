package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Entry defines both the CSV layout and database schema
type Customer_companies struct {
	gorm.Model

	Company_id   string `csv:"company_id"`
	Company_name string `csv:"company_name"`
}

type Customers struct {
	gorm.Model

	User_id      string `csv:"user_id"`
	Password     string `csv:"password"`
	Name         string `csv:"name"`
	Company_id   string `csv:"company_id"`
	Credit_cards string `csv:"credit_cards"`
}

type Deliveries struct {
	gorm.Model

	Id                 string `csv:"id"`
	Order_item_id      string `csv:"order_item_id"`
	Delivered_quantity string `csv:"delivered_quantity"`
}
type Order_items struct {
	gorm.Model

	Id             string `csv:"id"`
	Order_id       string `csv:"order_id"`
	Price_per_unit string `csv:"price_per_unit"`
	Quantity       string `csv:"quantity"`
	Product        string `csv:"product"`
}

type Orders struct {
	gorm.Model

	Id         string `csv:"id"`
	Created_at string `csv:"created_at"`

	Order_name  string `csv:"order_name"`
	Customer_id string `csv:"customer_id"`
}

func ReadFileName() {

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// fmt.Println(f.Name())
		// fmt.Println(filepath.Ext(f.Name()))

		if filepath.Ext(f.Name()) == ".csv" {
			path, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"))
			if err != nil {
				panic(err)
			}
			file, err := os.Open(filepath.Join(path, f.Name()))

			switch strings.Title(strings.ToLower(strings.Split(f.Name(), ".")[0])) {

			case "Customer_companies":

				if err != nil {
					panic(err)
				}
				entries := []*Customer_companies{}

				// err = gocsv.Unmarshal(file, &entries)

				if err := gocsv.UnmarshalFile(file, &entries); err != nil { // Load clients from file
					panic(err)
				}

				defer file.Close()

				// Create `entries` table if not exists
				err = db.AutoMigrate(&Customer_companies{})
				if err != nil {
					panic(err)
				}

				// Save all the records at once in the database
				result := db.Create(entries)
				if result.Error != nil {
					panic(result.Error)
				}

			case "Customers":
				entries := []*Customers{}

				// err = gocsv.Unmarshal(file, &entries)

				if err := gocsv.UnmarshalFile(file, &entries); err != nil { // Load clients from file
					panic(err)
				}

				defer file.Close()

				// Create `entries` table if not exists
				err = db.AutoMigrate(&Customers{})
				if err != nil {
					panic(err)
				}

				// Save all the records at once in the database
				result := db.Create(entries)
				if result.Error != nil {
					panic(result.Error)
				}
			case "Deliveries":
				entries := []*Deliveries{}

				// err = gocsv.Unmarshal(file, &entries)

				if err := gocsv.UnmarshalFile(file, &entries); err != nil { // Load clients from file
					panic(err)
				}

				defer file.Close()

				// Create `entries` table if not exists
				err = db.AutoMigrate(&Deliveries{})
				if err != nil {
					panic(err)
				}

				// Save all the records at once in the database
				result := db.Create(entries)
				if result.Error != nil {
					panic(result.Error)
				}

			case "Order_items":
				entries := []*Order_items{}

				// err = gocsv.Unmarshal(file, &entries)

				if err := gocsv.UnmarshalFile(file, &entries); err != nil { // Load clients from file
					panic(err)
				}

				defer file.Close()

				// Open a postgres database connection using GORM

				// Create `entries` table if not exists
				err = db.AutoMigrate(&Order_items{})
				if err != nil {
					panic(err)
				}

				// Save all the records at once in the database
				result := db.Create(entries)
				if result.Error != nil {
					panic(result.Error)
				}
			case "Orders":
				entries := []*Orders{}

				// err = gocsv.Unmarshal(file, &entries)

				if err := gocsv.UnmarshalFile(file, &entries); err != nil { // Load clients from file
					panic(err)
				}

				defer file.Close()

				// Open a postgres database connection using GORM

				// Create `entries` table if not exists
				err = db.AutoMigrate(&Orders{})
				if err != nil {
					panic(err)
				}

				// Save all the records at once in the database
				result := db.Create(entries)
				if result.Error != nil {
					panic(result.Error)
				}

			}

		}

	}

}

func InitializeDb(class interface{}, file *os.File) {

}

func main() {
	// Open the CSV file for reading

	ReadFileName()

	// handle error
	// ...

	// defer file.Close()

	// Parse CSV into a slice of typed data `[]Entry` (just like json.Unmarshal() does)
	// The builtin package `encoding/csv` does not support unmarshaling into a struct
	// so you need to use an external library to avoid writing for-loops.
	// var entries []customer_companies

}
