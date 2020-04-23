package main

import "fmt"

type ISiegeTankState interface {
	Move(x, y int)
	Attack()
	State()
}

type SiegeState struct {
}

func (s *SiegeState) Move(x, y int) {
	fmt.Print("Can't move in siege mode.\n")
}

func (s *SiegeState) Attack() {
	fmt.Print("Attacking for 40.\n")
}

func (s *SiegeState) State() {
	fmt.Print("SiegeState\n")
}

type TankState struct {
}

func (t *TankState) Move(x, y int) {
	fmt.Printf("Move to (%d, %d).\n", x, y)
}

func (t *TankState) Attack() {
	fmt.Print("Attacking for 20.\n")
}

func (t *TankState) State() {
	fmt.Print("TankState\n")
}

type Tank struct {
	ISiegeTankState //匿名成员，直接调用ISiegeTankState方法
}

func (t *Tank) SwitchState(s ISiegeTankState) {
	t.ISiegeTankState = s
}

func main() {
	t := Tank{
		&TankState{},
	}
	t.State()
	t.Attack()
	t.Move(1, 2)

	t.SwitchState(&SiegeState{})
	t.State()
	t.Attack()
	t.Move(1, 2)
}
