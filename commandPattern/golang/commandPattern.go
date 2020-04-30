package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Command interface {
	Execute()
}

type Light struct {
	Name string
}

func (l *Light) On() {
	fmt.Printf("light:%s is on\n", l.Name)
}

func (l *Light) Off() {
	fmt.Printf("light:%s is off\n", l.Name)
}

type LightOnCommand struct {
	*Light
}

func (l *LightOnCommand) Execute() {
	l.On()
}

type LightOffCommand struct {
	*Light
}

func (l *LightOffCommand) Execute() {
	l.Off()
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light,
	}
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		light,
	}
}

type RemoteControl struct {
	onCommand  [10]Command
	offCommand [10]Command
}

func (o *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	if slot >= 10 {
		panic("out of range")
	}
	o.onCommand[slot] = onCommand
	o.offCommand[slot] = offCommand
}

func (o *RemoteControl) OnButtonPushed(slot int) {
	if slot >= 10 {
		panic("out of range")
	}
	o.onCommand[slot].Execute()
}

func (o *RemoteControl) OffButtonPushed(slot int) {
	o.offCommand[slot].Execute()
}

func main() {
	rControl := RemoteControl{}
	for i := 0; i < 10; i++ {
		light := &Light{
			Name: "light" + fmt.Sprint(i),
		}
		on := NewLightOnCommand(light)
		off := NewLightOffCommand(light)
		rControl.SetCommand(i, on, off)
	}

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Int31() % 10
		if index%2 == 0 {
			rControl.OnButtonPushed(i)
			continue
		}
		rControl.OffButtonPushed(i)
	}
}
