package main
import "fmt"
    
type Animal interface {
    Speak() string
}

type Dog struct {}
func (d Dog) Speak() string {
    return "Woof!"
}

type JavaProgrammer struct {}
func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

func main() {
    animals := []Animal{Dog{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}