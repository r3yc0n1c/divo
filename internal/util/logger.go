package util

import "fmt"

const (
	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Reset  = "\033[0m"
)

func PrettyInfo(msg string) {
	fmt.Println(Cyan + msg + Reset)
}

func PrettyWarn(msg string) {
	fmt.Println(Yellow + msg + Reset)
}

func PrettyError(msg string) {
	fmt.Println(Red + msg + Reset)
}

func PrettyPrint(color, msg string) {
	fmt.Println(color + msg + Reset)
}

func PrettyPrintf(color, format string, a ...any) {
	fmt.Printf(color+format+Reset, a...)
}