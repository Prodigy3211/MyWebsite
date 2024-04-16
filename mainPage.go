//This is For the main page
package main
import("fmt")
import ("text/template")

type page struct {
	Header string
	Paragraph string
}

func main(){
	p:= page{Header:"Fuck you very much", Paragraph:"You're a Bitch"}
	fmt.Println(p)
}

func fuckYou(){
	fmt.Println("Fuck You!")
}

func main(){
	fuckYou()
}