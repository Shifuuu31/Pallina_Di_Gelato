package backend

import (
	"fmt"
	"os"
)

const Host string = "http://localhost"

var Port string = ":" + os.Getenv("PORT")

func Open() {
	if Port == ":" {
		Port = ":8080"
	}
	fmt.Println("server details:")
	fmt.Println("\tstatus: \033[1m\033[92m• Live\033[0m")
	fmt.Println("\t" + Host + Port)
	// fmt.Println("\t" + Port)

	// Command to run
	// cmd := exec.Command("bash", "openBrowser.sh") // Example: running 'grep main'

	// // Set up pipes for standard input and output
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// // Run the command
	// cmderr := cmd.Run()
	// if cmderr != nil {
	// 	log.Fatalln(cmderr)
	// }
}
