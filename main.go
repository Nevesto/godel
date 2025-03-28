// read this: follow my friend <3: https://www.instagram.com/_jarxdd/

package main

import (
	"fmt"

	"github.com/Nevesto/godel/cmd"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.GreenString(`
                           /$$           /$$
                          | $$          | $$
  /$$$$$$   /$$$$$$   /$$$$$$$  /$$$$$$ | $$
 /$$__  $$ /$$__  $$ /$$__  $$ /$$__  $$| $$
| $$  \ $$| $$  \ $$| $$  | $$| $$$$$$$$| $$
| $$  | $$| $$  | $$| $$  | $$| $$_____/| $$
|  $$$$$$$|  $$$$$$/|  $$$$$$$|  $$$$$$$| $$
 \____  $$ \______/  \_______/ \_______/|__/
 /$$  \ $$
|  $$$$$$/
 \______/ 
 `))
	cmd.Execute()
}
