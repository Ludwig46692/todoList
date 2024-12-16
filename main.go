package main

import (
	"fmt"
)

func main(){
    newList := todos{}
    newList.addTodo("Comprar leite", "ir até o mercado comprar leite")
    fmt.Printf("%+v\n", newList[0])
}


