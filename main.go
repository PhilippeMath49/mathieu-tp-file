package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	tableau := ReadFile()
	fmt.Println("Le fraguement 1 est : ", tableau[0])            //prend la première valeur du tableau
	fmt.Println("Le fraguement 2 est:", tableau[len(tableau)-1]) // prend la dernière valeur du tableau
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == "before" {
			fmt.Println("b")
		}
	}
	fmt.Println("Le fraguement 3 est : ")

}

func ReadFile() []string {
	var word string
	var tab []string

	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("File reading error", err)

	} else {

		for _, c := range data {
			if c != '\n' {
				word += string(c)

			} else {
				tab = append(tab, word)
				word = ""
			}

		}

	}
	tab = append(tab, word) //rajoute "a" au tableau
	return tab

}
