package superlog

import "log"

// SuperLog logs super log
func SuperLog(text string) {
	log.Printf("This is super log: %s", text)
}
