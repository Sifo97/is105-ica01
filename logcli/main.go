package main
import (
	"fmt"
	"os"
	"math"
	"strconv"
)

func main() {
	// Ta imot argument 1
	ArgumentOne := os.Args[1] //
	// Printe det
	fmt.Println("ARG: " + ArgumentOne)
	// Siden Log2 tar imot Float64 som argument, men command-line argumentet blir parset som en String, må vi bruke strconv for å konvertere Stirng til Float64
	// Så er converterer vi det.
	x, err := strconv.ParseFloat(ArgumentOne, 64)
	// Kan legge til en if argument for å sjekke om err(feilkode) = nill, men whatever.
	fmt.Print("Feil kode: ")
	fmt.Println(err)
	// Kalkulere
	b := math.Log2(x)
	// PRINT
	fmt.Print("Value: ")
	fmt.Println(b)
}

