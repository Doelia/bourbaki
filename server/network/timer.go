package network

import (
	"log"
	"os"
	"time"
)

var timerLog = log.New(os.Stdout, "[timer] ", 0)

// TIMEPERTURN Temps pour jouer, en secondes
const TIMEPERTURN = 1

// Timer TODO
type Timer struct {
	curentID      int
	autoIncrement int
}

func createTimer() *Timer {
	t := &Timer{}
	return t
}

// LuanchNewTimer TODO
func (t *Timer) LuanchNewTimer() {
	t.autoIncrement++
	t.curentID = t.autoIncrement
	timerLog.Println("LuanchNewTimer, curentID = ", t.curentID)
	go t.endTimer(t.curentID)
}

// Cancel TODO
func (t *Timer) Cancel() {
	timerLog.Println("Canceled")
	t.curentID = 0
}

func (t *Timer) endTimer(id int) {
	timerLog.Println("wait end timer ", id)
	time.Sleep(300 * time.Millisecond)
	timerLog.Println("endTimer ? ", id, " vs ", t.curentID)
	if t.curentID == id {
		AI()
	}
}
