package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	nv "./navigation"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	menu := `
	===================================
	=== WebPage Pagination Exercise ===
	===================================
	==[1] - Insert input
	==[0] - Quit
	===================================
	>`

	choice := ""

	for true {
		fmt.Print(menu)
		choice, _ = reader.ReadString('\n')
		choice = strings.Trim(choice, " \n\t")
		if choice == "1" {
			fmt.Print("==> Insert Current Page: ")
			currentPage, _ := reader.ReadString('\n')
			cp, _ := strconv.Atoi(strings.Trim(currentPage, " \n\t"))
			fmt.Print("\n==> Insert Total Number of Pages: ")
			totalPages, _ := reader.ReadString('\n')
			tp, _ := strconv.Atoi(strings.Trim(totalPages, " \n\t"))
			fmt.Print("\n==> Insert Boundaries: ")
			boundaries, _ := reader.ReadString('\n')
			b, _ := strconv.Atoi(strings.Trim(boundaries, " \n\t"))
			fmt.Print("\n==> Insert Around: ")
			around, _ := reader.ReadString('\n')
			a, _ := strconv.Atoi(strings.Trim(around, " \n\t"))

			res, err := nv.GetPagination(cp, tp, b, a)
			if err != nil {
				fmt.Println(err, "\n\n")
			} else {
				fmt.Println("===Result:\n", res, "\n\n")
			}
		} else if choice == "0" {
			break
		} else {
			fmt.Printf(" PLEASE PICK ONE OF THE OPTIONS\n\n")
		}
	}
}
