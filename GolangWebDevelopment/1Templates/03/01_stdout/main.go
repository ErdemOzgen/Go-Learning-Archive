package main

import (
	"log"
	"os"
	"text/template"
)

func checkError(err *error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//https://pkg.go.dev/text/template
func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") //templete is struct impelemented parse tree ==> https://youtu.be/a1OLDU1QAfw
	checkError(&err)
	err = tpl.Execute(os.Stdout, nil)
	checkError(&err)
	//fmt.Println(tpl.Root)
	//fmt.Println(tpl.Tree)
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
