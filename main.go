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
	fmt.Println("done")

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	//dat, err := ioutil.ReadFile("P05D150701.TXT")
	//check(err)
	//fmt.Print(string(dat))
	//fmt.Print("done")

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	// f, err := os.Open("P05D150701.TXT")
	// check(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	// b1 := make([]byte, 50)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// You can also `Seek` to a known location in the file
	// and `Read` from there.
	// o2, err := f.Seek(6, 0)
	// check(err)
	// b2 := make([]byte, 2)
	// n2, err := f.Read(b2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.
	// o3, err := f.Seek(6, 0)
	// check(err)
	// b3 := make([]byte, 2)
	// n3, err := io.ReadAtLeast(f, b3, 2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but `Seek(0, 0)`
	// accomplishes this.
	// _, err = f.Seek(0, 0)
	// check(err)

	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.
	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(50)
	// check(err)
	// fmt.Printf("50 bytes: %s\n", string(b4))

	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).
	//f.Close()

}
