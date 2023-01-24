// Tell compiler that we are creating an executable package here
package main

// import any methods we need in the application, these will get automatically imported as you write
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//  open a file using os
	file, err := os.Open("myfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if  err != nil{
			break
		}
		if strings.Contains(str, "ERROR") {
			fmt.Println(str)
		}
	}
}