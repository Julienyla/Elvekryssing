package event

var event = "| Kylling---\\________/__________________--- |"

func ViewEvent() string {
	return event
}

func AllGoingOver() {
	event = "| ---\\____Kylling____/__________________--- |"
}

func AllOver() {
	event = "| ---__________________\\________/---Kylling  |"
}

