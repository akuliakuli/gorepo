package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

var secretWord string = "Pss, It's me mario"
const age int = 18


func main(){

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())

}