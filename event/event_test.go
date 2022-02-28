package event

import "testing"

func TestViewEvent(t *testing.T) {
	wanted := "| Kylling---\\________/__________________--- |"
	event := ViewEvent()
	if event != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", event, wanted)
	}
}

