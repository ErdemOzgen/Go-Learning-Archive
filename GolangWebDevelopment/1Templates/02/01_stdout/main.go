package main

import (
	"fmt"
)

// TemplateWithStdout function takes strting , stdout template and returns it
func TemplateWithStdout(name string) string {
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
	return tpl
}
func main() {
	name := "Erdem Ozgen"
	TemplateWithStdout(name)

}
