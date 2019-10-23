package log

import "fmt"



func Black(item string)  {
	fmt.Print(format(item,black))
}

//print red
func Red(item string)  {
	fmt.Print(format(item,red))
}


func Green(item string)  {
	fmt.Print(format(item,green))
}

func Yellow(item string)  {
	fmt.Print(format(item,yellow))
}

func Blue(item string)  {
	fmt.Print(format(item,blue))
}


func Fuchsia(item string)  {
	fmt.Print(format(item,fuchsia))
}

func Ultramarine(item string)  {
	fmt.Print(format(item,ultramarine))
}
func White(item string)  {
	fmt.Print(format(item,white))
}



