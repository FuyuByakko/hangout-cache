package cache

import (
	"errors"
)

var Cache = make(map[string]string)

func Get(key string) string {
	return Cache[key]
}

// func Add(k string, s structs.EventsAndHotels) {
func Add(k string, info string) error {
	Cache[k] = info
	check := Cache[k]
	if check != "" {
		return nil
	} else {
		return errors.New("Writing of data for key %s failed.\n")
	}
}
