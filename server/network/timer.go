package network

import (
	"log"
	"os"
	"time"
)

var timerLog = log.New(os.Stdout, "[timer] ", 0)

// TIMEPERTURN Temps pour jouer, en secondes
const TIMEPERTURN = 8

// Timer Objet pour gérer les timers
type Timer struct {
	curentID      int
	autoIncrement int
}

func createTimer() *Timer {
	t := &Timer{}
	return t
}


// LaunchNewTimer Démarre un nouveau timer (et annule le précédent)
func (t *Timer) LaunchNewTimer() {
	t.autoIncrement++
	t.curentID = t.autoIncrement
	timerLog.Println("LuanchNewTimer, curentID = ", t.curentID)
	go t.endTimer(t.curentID)
}

// Cancel Annule le timer en cours
func (t *Timer) Cancel() {
	timerLog.Println("Canceled")
	t.curentID = 0
}

// Fonction callback s'executant à la fin d'un timer
func (t *Timer) endTimer(id int) {
	timerLog.Println("wait end timer ", id)
	time.Sleep(TIMEPERTURN * time.Second)
	timerLog.Println("endTimer ? ", id, " vs ", t.curentID)
	if t.curentID == id {
		AI()
	}
}
