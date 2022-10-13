package utils

import (
	"log"
)

func Terminate(err error) {
	log.Fatalf("Terminating: %s", err.Error())
}
