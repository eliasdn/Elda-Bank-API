package utils

import (
	"log"
)

func CheckErr(err error) error {
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
