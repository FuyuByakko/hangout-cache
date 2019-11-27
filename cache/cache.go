package cache

// var Cache = make(map[string] EventsAndHotels)
var Cache = make(map[string]string)

// var Cache = make(map[string]string)

func Get(key string) string {
	return Cache[key]
}

func Add(k string, s string) {
	Cache[k] = s
}

// func (ev *EventsAndHotels) Add(event Event, hotel Hotel) {
// 	ev.Events = append(ev.Events, event);
// 	ev.Hotels = append(ev.Hotels, hotel);
// }
