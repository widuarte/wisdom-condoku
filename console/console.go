package console

import (
	"fmt"
	"os"
	"os/exec"
)

type Color func(msg string)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

func Clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Red(msg string) {
	print(fmt.Sprintf("%s%s%s", red, msg, reset))
}

func Green(msg string) {
	print(fmt.Sprintf("%s%s%s", green, msg, reset))
}

func Yellow(msg string) {
	print(fmt.Sprintf("%s%s%s", yellow, msg, reset))
}

func Blue(msg string) {
	print(fmt.Sprintf("%s%s%s", blue, msg, reset))
}

func Purple(msg string) {
	print(fmt.Sprintf("%s%s%s", purple, msg, reset))
}

func Cyan(msg string) {
	print(fmt.Sprintf("%s%s%s", cyan, msg, reset))
}

func Gray(msg string) {
	print(fmt.Sprintf("%s%s%s", gray, msg, reset))
}

func White(msg string) {
	print(fmt.Sprintf("%s%s%s", white, msg, reset))
}
