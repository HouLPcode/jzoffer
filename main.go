package main

import (
	"fmt"
	"github.com/HouLPcode/jzoffer/alg"
)

func main() {
	fmt.Println("剑指offer运行结果")

	// A B T G
	// C F C S
	// J D E H
	//path: BFCE

	//ABCEHJIG
	//SFCSLOPQ
	//ADEEMNOE
	//ADIDEJFM
	//VCEIFGGS
	b := alg.HasPath("ABCEHJIGSFCSLOPQADEEMNOEADIDEJFMVCEIFGGS", 5, 8, "SGGFIECVAASABCEHJIGQEM")
	if b == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
