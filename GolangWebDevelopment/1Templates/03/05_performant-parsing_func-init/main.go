package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() { // use parser before main
	tpl = template.Must(template.ParseGlob("templates/*"))

	/*
		init is called after all the variable declarations
		in the package have evaluated their initializers,
		and those are evaluated only after all the imported
		 packages have been initialized.
	*/
}

func main() {

	fmt.Println("-------------")
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\n-------------")
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
