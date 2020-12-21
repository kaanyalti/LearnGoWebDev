package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl*template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("****************")
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n")
	fmt.Println("****************")

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