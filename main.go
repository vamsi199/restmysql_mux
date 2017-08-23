package main



func main ()  {
	a := App{}
	a.Initialize("root","pass","107.21.2.246","vamsigo" )
	a.Run(":8080")

}
//create the main.go file which will contain the entry point for the application