package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var p int

	tableau := ReadFile()
	fmt.Println("Le fraguement 1 est :", tableau[0])              //prend la première valeur du tableau
	fmt.Println("Le fraguement 2 est :", tableau[len(tableau)-1]) // prend la dernière valeur du tableau

	for i := 0; i < len(tableau); i++ {
		if strings.Contains(tableau[i], "before") {
			p, _ = strconv.Atoi(tableau[i+1])
		}
	}

	fmt.Println("Le fraguement 3 est :", tableau[p-1])

	for j := 0; j < len(tableau)-1; j++ {
		if strings.Contains(tableau[j+1], "now") {
			p = j
		}
	}
	for i, c := range tableau[p] {
		if i == 1 {
			p = int(c)/len(tableau) - 1
		}
	}
	fmt.Println("Le fraguement 4 est :", tableau[p])
}

func ReadFile() []string {
	var word string
	var tab []string

	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("File reading error", err)

	} else {

		for _, c := range data {
			if c == '\n' {
				tab = append(tab, word)
				word = ""

			} else if c != '\r' {
				word += string(c)
			}
		}
	}
	tab = append(tab, word) //rajoute "a" au tableau
	return tab
}
