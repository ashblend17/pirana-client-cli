package ter

import "fmt"

func Help() {
	fmt.Println("Welcome to The Alfred Project")
	fmt.Println("Usage:")
	fmt.Println("  alfred [command] {regex}")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  help       Show this help message")
	fmt.Println("  find       Find what you seek, for you have alzheimers")
	fmt.Println("  s          For Samwise the great")
	fmt.Println("  c          For the race-course of life")
	fmt.Println("  p          For the pass of Caradhras")
	fmt.Println()
	fmt.Println("Note:")
	fmt.Println("  You can use multiple getch commands together at once in a single command: eg. alfred scp RNO123456")
}
