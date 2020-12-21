package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("****************")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n")
	fmt.Println("****************")
	tpl, err = tpl.ParseFiles("two.gmao", "three.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n")
	fmt.Println("****************")
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n")
	fmt.Println("****************")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n")
	fmt.Println("****************")
}