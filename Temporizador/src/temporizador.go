package main

import (
	"fmt"
	"time"
)

type generic struct {}

func wait(canal chan<- string, tiempo time.Duration) {
	time.Sleep(tiempo)
	close (canal)
}


/*func closed (c chan generic) bool {
	x,y :=<-ch
}*/
func estaListo(what string, sec time.Duration, ch chan<- string) {

	go wait(ch, sec)

	time.Sleep(sec)
	envio:=what+" esta listo"
	ch <- envio 
	
}

func main() {

	ch := make(chan string)

	go estaListo("trabajo1", 10, ch)

	
	status,x := <-ch
	
	if x {
		fmt.Println(status)		
	} else {
		fmt.Println("Se cerro el canal antes de tiempo")
		
	}
}
