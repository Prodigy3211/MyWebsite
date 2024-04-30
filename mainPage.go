//This is For the main page
package main
import("fmt")
import ("text/template")

type page struct {
	Header string
	Paragraph string
}

func main(){
	p:= page{Header:"Hello World", Paragraph:"How's it going"}
	fmt.Println(p)
}

func hello(){
	fmt.Println("Hello")
}

func main(){
	hello()
}