package main
import "fmt"

type pet interface{
    run()
    eat()
}

type cat struct{
    weight int
}

func (c *cat) run(){
    c.weight--;
}

func (c *cat) eat(){
    c.weight += 4;
}

func take_care(p pet){
    p.run()
    p.eat()
}

func main(){
    c := cat{weight: 10}
    fmt.Println("Before taking care, cat is", c.weight)
    take_care(&c)
    fmt.Println("After taking care, cat is", c.weight)
}