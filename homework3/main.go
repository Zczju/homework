package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	// maxFloor 楼层总数
	maxFloor int
	// nowFloor 当前所处楼层
	nowFloor int
	// direction 方向
	direction string
	// upRequests 上行请求
	upRequests []int
	// downRequests 下行请求
	downRequests []int
}

func (e *Elevator) request(requestFloor int) error {
	if requestFloor >= 1 && requestFloor <= e.maxFloor {
		if requestFloor == e.nowFloor {
			// 当请求为当前楼层时，直接返回，不进行任何操作
			return nil
		}
		if len(e.upRequests)+len(e.downRequests) == 0 {
			e.updateDirection(requestFloor)
		}

		if requestFloor > e.nowFloor {
			e.upRequests = append(e.upRequests, requestFloor)
			sort.Ints(e.upRequests)
		} else {
			e.downRequests = append(e.downRequests, requestFloor)
			sort.Slice(e.downRequests, func(i, j int) bool {
				return e.downRequests[j] < e.downRequests[i]
			})
		}

		return nil
	} else {
		return errors.New("invalid floor request")
	}
}

func (e *Elevator) updateDirection(requestFloor int) {
	if requestFloor > e.nowFloor {
		e.direction = "向上"
	} else if requestFloor < e.nowFloor {
		e.direction = "向下"
	} else {
		e.direction = ""
	}
}

func (e *Elevator) elevatorMove() {
	if e.direction == "向上" {
		if len(e.upRequests) == 0 {
			if len(e.downRequests) != 0 {
				e.updateDirection(e.downRequests[0])
				e.elevatorMove()
			} else {
				e.direction = ""
				return
			}
		} else {
			e.move(e.upRequests[0])
			e.upRequests = e.upRequests[1:]
		}
	} else if e.direction == "向下" {
		if len(e.downRequests) == 0 {
			if len(e.upRequests) != 0 {
				e.updateDirection(e.upRequests[0])
				e.elevatorMove()
			} else {
				e.direction = ""
				return
			}
		} else {
			e.move(e.downRequests[0])
			e.downRequests = e.downRequests[1:]
		}
	}
}
func (e *Elevator) Open() {
	fmt.Printf("open the door at floor %d\n", e.nowFloor)
}
func (e *Elevator) Close() {
	fmt.Printf("close the door at floor %d\n", e.nowFloor)
}
func (e *Elevator) move(to int) {
	for {
		if e.nowFloor == to {
			e.Open()
			time.Sleep(1 * time.Second)
			e.Close()
			return
		} else {
			if e.nowFloor > to {
				time.Sleep(1 * time.Second)
				e.nowFloor -= 1
			} else {
				time.Sleep(1 * time.Second)
				e.nowFloor += 1
			}
		}
	}
}
