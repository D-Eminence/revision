package main

import "fmt"

// PrintMemory takes a fixed-size array of 10 bytes and displays its contents
// in a structured memory dump style, showing 5 bytes per line.
// 1. Hexadecimal representation (two digits per byte).
// 2. ASCII representation (printable characters or a dot).
func PrintMemory(arr [10]byte) {
	hexOutput := ""
	asciiOutput := ""

	const bytesPerLine = 5 // We choose 5 bytes per line for a 10-byte array

	for i, b := range arr {
		// A. When the line is complete (after 5 bytes), print the buffer and reset.
		if i > 0 && i%bytesPerLine == 0 {
			// Print the accumulated hex and ASCII buffer for the completed line.
			// %-15s ensures the hex column has 15 characters width for alignment.
			fmt.Printf("%-15s %s\n", hexOutput, asciiOutput)
			// Reset buffers for the next line
			hexOutput = ""
			asciiOutput = ""
		}

		// B. Hexadecimal Conversion: Append two-digit hex representation with a space.
		hexOutput += fmt.Sprintf("%02x ", b)

		// C. ASCII Character Check: 
		// Printable ASCII range is 32 (space) to 126 (tilde '~').
		if b >= 32 && b <= 126 {
			asciiOutput += string(b)
		} else {
			asciiOutput += "."
		}
	}

	// 2. Print the final remaining buffer (the last line).
	fmt.Printf("%-15s %s\n", hexOutput, asciiOutput)
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}