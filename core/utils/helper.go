package utils

import "log"

func ErrorLog(message string, err error) {
	if err != nil {
		log.Printf("ERROR !!!: %s => %v", message, err)
	}
}
