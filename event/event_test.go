package event

import "testing"

func TestEvent(t *testing.T) {
	wanted := "| Kylling---\\________/__________________--- |"
	event := Event()
	if event != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", event, wanted)
	}
}

