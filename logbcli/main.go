package main
import "fmt"
import "os"
import "strconv"
func main() {
// Ta imot argument 1
logMethod := os.Args[1]
ArgumentTwo := os.Args[2]
// Printe det
fmt.Println("logMethod: " + logMethod)
fmt.Println("Number: " + ArgumentTwo)
x, err := strconv.ParseFloat(ArgumentTwo, 64)
fmt.Print("Feil kode: ")
fmt.Println(err)
fmt.Print("Base" + logMethod + ": ")
fmt.Println(CalcLogArg(logMethod, x))
}

