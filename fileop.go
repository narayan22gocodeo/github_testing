package bank

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"io/ioutil"
)

// Utility Functions
func greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go programming.", name)
}

 func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// File Operations
func writeToFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func readFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Concurrency with Goroutines
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range []rune("ABCDE") {
		fmt.Printf("Letter: %c\n", letter)
		time.Sleep(300 * time.Millisecond)
	}
}

// HTTP Server
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is a simple HTTP server.")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current Time: %s", currentTime)
}

// Main Function
func main() {
	// String Manipulation
	name := "GoCoder"
	fmt.Println(greet(name))
	fmt.Println("Reversed Name:", ReverseString(name))

	// File Operations
	filename := "example.txt"
	content := "This is a sample text written to a file using Go."
	if err := writeToFile(filename, content); err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("File written successfully.")
	}

	fileContent, err := readFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("File Content:", fileContent)
	}

	// Cleanup
	defer os.Remove(filename)

	// Goroutines Example
	go printNumbers()
	go printLetters()

	// HTTP Server
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Starting server at :8080")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait for Goroutines to finish
	time.Sleep(3 * time.Second)
	fmt.Println("Main program exiting.")
}
