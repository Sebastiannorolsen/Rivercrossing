package main


import (
	"bufio"
	"fmt"

	"github.com/Sebastiannorolsen/Rivercrossing/event"
	"os"

	//"rc1-verden/event"
)

/*var items = [] string{
"kylling",
"rev",
"korn",
}*/


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("For å løse gåten må du ta med riktig ting først, hva vil du ta med først? kylling, rev eller korn")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(event.FirstPut(text))
		fmt.Println(event.GetIn(text))
		fmt.Println(event.CrossRiver(text))
		fmt.Println(event.TakeOut(text))
		fmt.Println(event.GetOut(text))
		if(text == "kylling"){
			break
		}
	}
}
