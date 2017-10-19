package main
import "fmt"

func main(){
	defer done()
	
	returnFunc()

	return
}
func returnFunc(){
	fmt.Println("return")
	return

}
func done(){
	fmt.Println("done")
}
