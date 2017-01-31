package main
import "fmt"
import "os"
import "math"
import "strconv"
func main() {
// Ta imot argument 1
logMethod := os.Args[1]
ArgumentTwo := os.Args[2]
// Printe det
fmt.Println("logMethod: " + logMethod)
x, err := strconv.ParseFloat(ArgumentTwo, 64)

fmt.Print("Feil kode: ")
fmt.Println(err)

if logMethod == "10" {
b := math.Log10(x)
fmt.Print("Value: ")
fmt.Println(b)
}
if logMethod == "2" {
b := math.Log2(x)
fmt.Print("Value: ")
fmt.Println(b)
}
}

