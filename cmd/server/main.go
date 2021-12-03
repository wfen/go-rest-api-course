package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")
}
func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}
