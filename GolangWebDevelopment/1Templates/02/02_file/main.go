package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func TemplateWithFile(s string) {
	tpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` +
		s +
		`</h1>
</body>
</html>
	`
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(tpl))
}

func main() {
	name := "Erdem Ozgen"
	//you can use sprint here
	str := fmt.Sprint(` 
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` +
		name +
		`</h1>
</body>
</html>
	`)
	fmt.Println(str)
	TemplateWithFile(name)

}
