package log

import (
	"fmt"
)

const blackBg=40
const redBg  =41
const greenBg  =42
const yellowBg  =43
const  blueBg =44
const fuchsiaBg  =45
const  ultramarineBg=46
const whiteBg  =47

const black  =30
const red  = 31
const green  =32
const yellow  =33
const blue  =34
const fuchsia  =35
const  ultramarine=36
const white=37


// background special variables
const terminalDefault=0
const highlight=1
const underline=4
const reverseVideo=7
const invisible=8


//default background of terminal
func format(item string,color int) string {
	return fmt.Sprintf("%c[0;%d;%dm%s%c[0m\n", 0x1B,terminalDefault,color,item, 0x1B)
}

//transform
func formatWithBackground(item string,color string,bgColor string) string {
	return formatWithBackground2(item,transform(color),transformBg(bgColor))

}
//format
func formatWithBackground2(item string,color int,bgColor int) string {
	return fmt.Sprintf("%c[0;%d;%dm%s%c[0m\n", 0x1B,bgColor,color,item, 0x1B)
}

func formatWithUnderLine(item string,color int)  string{
	return fmt.Sprintf("%c[0;%d;%dm%s%c[0m\n", 0x1B,underline,color,item, 0x1B)
}

func formatWithHighLight(item string,color int) string {
	return fmt.Sprintf("%c[0;%d;%dm%s%c[0m\n", 0x1B,highlight,color,item, 0x1B)
}

// print with font color and background
func Combine(item string,color string,bgColor string)  {
	fmt.Print(formatWithBackground(item,color,bgColor))
}

func UnderLine(item string,color string)  {
	fmt.Print(formatWithUnderLine(item,transform(color)))
}
func HighLight(item string,color string)  {
	fmt.Print(formatWithHighLight(item,transform(color)))
}
//transform color string to its corresponding int format
func transform(color string) int  {
	fontColor:=0
	switch color {
	case "red":
		fontColor=red
	case "black":
		fontColor=black
	case "green":
		fontColor=green
	case "yellow":
		fontColor=yellow
	case "blue":
		fontColor=blue
	case "ultramarine":
		fontColor=ultramarine
	case "white":
		fontColor=white
	default:
		fontColor=white
	}
	return fontColor
}

func transformBg(bgColor string) int {
	background:=0

	switch bgColor {
	case "red":
		background=redBg
	case "black":
		background=blackBg
	case "green":
		background=greenBg
	case "yellow":
		background=yellowBg
	case "blue":
		background=blueBg
	case "ultramarine":
		background=ultramarineBg
	case "white":
		background=whiteBg
	default:
		background=blackBg
	}
	return background
}