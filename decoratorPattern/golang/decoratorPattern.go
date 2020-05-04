package main

import (
	"fmt"
	"time"
)

type Executor interface {
	Execute()
}

type Thread struct {
}

func (t *Thread) Execute() {
	fmt.Println("start a thread.")
	time.Sleep(time.Microsecond * 200)
}

type LogDecorator struct {
	e Executor
}

func (l *LogDecorator) Execute() {
	fmt.Printf("start execute at:%v\n", time.Now())
	l.e.Execute()
	fmt.Printf("finish execute at:%v\n", time.Now())
}

func main() {
	t := Thread{}
	t.Execute()

	l := LogDecorator{e: &t}
	l.Execute()
}
