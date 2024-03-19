package main

import "fmt"

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't do it kid. You just can't")
	}
}

func main() {
	var myEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine, 50)

}
