package structs

//define the test cases and structs
type Hotel struct {
	Address   string `json:"address"`
	Available string `json:"available"`
	City      string `json:"city"`
	Image     string `json:"image"`
	Price     string `json:"price"`
	Review    string `json:"review"`
	Url       string `json:"url"`
	Zip       string `json:"zip"`
}

type Hotels []Hotel

type Event struct {
	Image     string `json:"image"`
	StartTime string `json:"startTime"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Venue     string `json:"venue"`
}

type Events []Event

type EventsAndHotels struct {
	Events []Event `json:"Events"`
	Hotels []Hotel `json:"Hotel"`
}

func New() *EventsAndHotels {
	return &EventsAndHotels{}
}

func TestResponse() EventsAndHotels {
	return EventsAndHotels{
		Events{
			Event{
				Image:     "testEvent1Image",
				StartTime: "testEvent1Start",
				Title:     "testEvent1Title",
				Url:       "testEvent1Url",
				Venue:     "testEvent1Venue",
			},
			Event{
				Image:     "testEvent2Image",
				StartTime: "testEvent2Start",
				Title:     "testEvent2Title",
				Url:       "testEvent2Url",
				Venue:     "testEvent2Venue",
			},
		},
		Hotels{
			Hotel{
				Address:   "test1address",
				Available: "test1available",
				City:      "test1test",
				Image:     "test1img",
				Price:     "test1price",
				Review:    "test1Review",
				Url:       "test1url",
				Zip:       "test1zip"},
			Hotel{
				Address:   "test2address",
				Available: "test2available",
				City:      "test2test",
				Image:     "test2img",
				Price:     "test2price",
				Review:    "test2Review",
				Url:       "test2url",
				Zip:       "test2zip"},
		},
	}
}
