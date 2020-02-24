package errorz

import "log"

func Fs(err error) error {
	log.Print("fs: unknown filesystem error: "+err.Error())
	return err
}