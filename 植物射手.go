package main

import (
	"fmt"
)

type Snowpea interface {
	shoot()
	retard()
}


type Repeatshoot interface {
	repeat()
}

type Repeater struct {
	Snowpea
}

func  (use Repeater)shoot(){
	fmt.Println("this is shoot")
}

func (user Repeater) repeat() {
	fmt.Println("this is repeat")
}
func main() {
	sf := &Repeater{nil}
	var ss Snowpea = sf
	var rr Repeatshoot =sf
	ss.shoot()
	rr.repeat()


}
