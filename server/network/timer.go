package network

import (
	"log"
	"os"
	"time"
)

var timerLog = log.New(os.Stdout, "[timer] ", 0)

// TIMEPERTURN Définition de la durée d'un tour pour un joueur, en secondes
const TIMEPERTURN = 9

// Timer Objet pour gérer les timers
type Timer struct {
	currentID     int
	autoIncrement int
}

func createTimer() *Timer {
	t := &Timer{}
	return t
}

// LaunchNewTimer Démarre un nouveau timer (et annule le précédent)
func (t *Timer) LaunchNewTimer() {
	t.autoIncrement++
	t.currentID = t.autoIncrement
	timerLog.Println("LaunchNewTimer, currentID = ", t.currentID)
	go t.endTimer(t.currentID)
}

// Cancel Annule le timer en cours
func (t *Timer) Cancel() {
	timerLog.Println("Canceled")
	t.currentID = 0
}

// Fonction callback s'executant à la fin d'un timer
func (t *Timer) endTimer(id int) {
	timerLog.Println("wait end timer ", id)
	time.Sleep(TIMEPERTURN * time.Second)
	timerLog.Println("endTimer ? ", id, " vs ", t.currentID)
	if t.currentID == id {
		AI()
	}
}
