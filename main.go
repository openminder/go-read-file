// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
	"fmt"
	// "io"
	//"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Accessory struct {
	ItemNumber string
	Price      float64
}

func main() {
	fmt.Println("App started")
	db, err := sql.Open("mysql", "root:sdij2015,gt@tcp(:3306)/pdb_sie_car_20140625")
	check(err)
	defer db.Close()
	rows, err := db.Query("Select accessories.id, accessories.item_number, accessories.price, accessorieslinks.price FROM pdb_sie_car_20140625.accessories INNER JOIN accessorieslinks ON accessories.id = accessorieslinks.accessory_id WHERE item_number IS NOT NULL")
	check(err)
	var dbAccessories []Accessory
  for rows.Next() {
    var aid int64
		var itemNumber string
		var pr sql.NullFloat64
		if !pr.Valid {
			err = rows.Scan(&aid, &itemNumber, &pr)
			check(err)
			dbAccessories = append(dbAccessories, Accessory{itemNumber, pr.Float64})
		}

  }
  rows.Close()

	// slow version from internet, but maybe correct one?
	file, err := os.Open("P05D150701.TXT")
	check(err)
	//
	defer file.Close()
	//
	var lines []Accessory
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		itemNumber := line[1 : len(line)-64]
		price, err := strconv.ParseFloat(strings.Replace(line[47:len(line)-23], ",", ".", -1), 2)
		check(err)
		l := Accessory{itemNumber, price}
		lines = append(lines, l)
	}
	//fmt.Print(lines)

	var wg sync.WaitGroup

	wg.Add(len(lines))

	for _, v := range dbAccessories {
		go func(a1 Accessory, file []Accessory) {
			for _, l := range file {
				defer wg.Done()
				if a1.ItemNumber == l.ItemNumber {
					fmt.Println(a1.ItemNumber)
					//f, err := os.Create("output.txt")
					//check(err)
					//defer f.Close()
					//d1 := []byte(a1.ItemNumber)
					//_, err = f.WriteString(string(d1))
					//err := ioutil.WriteFile("output.txt", d1, 0644)
					//check(err)
				}
			}
		}(v, lines)
	}

	wg.Wait()
	//fmt.Print(otherAccessories)
	fmt.Println("done")

}
