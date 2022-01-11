package main

import (
	"fmt"
	"sort"
	"time"
)

var requestFloors []int

type Elevator struct {
	floor     int
	nowFloor  int
	direction string
}

func (e *Elevator) getDirection(requestFloors []int) string {
	if len(requestFloors) == 0 {
		e.direction = "不动"
		return e.direction
	} else {
		if requestFloors[0] > e.nowFloor {
			e.direction = "向上"
			return e.direction
		} else if requestFloors[0] < e.nowFloor {
			e.direction = "向下"
			return e.direction
		}

	}
	return ""
}
func (e *Elevator) elevatorMove(direction string, nowFloor int) {
	sort.Ints(requestFloors)
	if direction == "向上" {
		e.upMove(nowFloor)
	}
	if direction == "向下" {
		e.downMove(nowFloor)
	}
}
func (e *Elevator) Open() {
	fmt.Println("open the door")
}
func (e *Elevator) Close() {
	fmt.Println("close the door")
}
func (e *Elevator) upMove(nowFloor int) {
	for {
		time.Sleep(1 * time.Second)
		nowFloor++
		for _, requestFloor := range requestFloors {
			if nowFloor == requestFloor {
				e.Open()
				time.Sleep(1 * time.Second)
				e.Close()
				e.nowFloor = nowFloor
			}
			if nowFloor == e.floor {
				return
			}
		}
	}
}
func (e *Elevator) downMove(nowFloor int) {
	for {
		time.Sleep(1 * time.Second)
		nowFloor--
		for _, requestFloor := range requestFloors {
			if nowFloor == requestFloor {
				e.Open()
				time.Sleep(1 * time.Second)
				e.Close()
				e.nowFloor = nowFloor
			}
			if nowFloor == 1 {
				return
			}
		}
	}
}
