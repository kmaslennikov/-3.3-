package main

import "fmt"

func main() {
  fmt.Println("Программа, которая меняет значения переменных местами.")
 fmt.Println("Введите значение a:")
 var a int
 fmt.Scan(&a)

 fmt.Println("Введите значение b:")
 var b int
 fmt.Scan(&b)

 var c int

 c = a
 a = b
 b = c

fmt.Println("Значения переменных поменялись местами.")

 
 

 

 fmt.Println("a:", a)

 fmt.Println("b:", b)

}