package main

import (
	"fmt"
	"time"
    "errors"
)

type Todo struct{
    title string
    desc string
    completed bool
    createdAt time.Time
    completedAt* time.Time
}

type todos[]Todo

func (todos *todos) addTodo(title string, desc string) {
    newTodo := Todo{
        title: title,
        desc: desc,
        completed: false,
        createdAt: time.Now(),
        completedAt: nil,
    }
    *todos = append(*todos, newTodo)
}

func (todos *todos) validateIndex(index int) error{
    if index < 0 || index >= len(*todos){
        err := errors.New("Invalid index")
        fmt.Println(err.Error())
        return err
    }
    return nil
}

func (todos *todos) deleteTodo(index int) error{
    t := *todos
    if err := t.validateIndex(index); err != nil{
        return err
    }
    *todos = append(t[:index], t[index+1:]...)
    return nil
}
