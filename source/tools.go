package source

import "fmt"

const Host, Port string = "localhost", ":8080"
func Open() {
	fmt.Println("server details:")
	fmt.Println("\tstatus: \033[1m\033[92m• Live\033[0m")
	fmt.Println("\t" + Host + Port)

	// // Command to run
	// cmd := exec.Command("bash", "openBrowser.sh") // Example: running 'grep main'

	// // Set up pipes for standard input and output
	// // cmd.Stdin = os.Stdin
	// // cmd.Stdout = os.Stdout
	// // cmd.Stderr = os.Stderr

	// // Run the command
	// cmderr := cmd.Run()
	// if cmderr != nil {
	// 	log.Fatalln(cmderr)
	// }
}
