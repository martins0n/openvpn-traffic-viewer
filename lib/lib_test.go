package lib

import (
	"testing"
	"time"
)

func TestParseClientListRecord(t *testing.T) {
	pattern := "ma,178.205.243.63:27437,1229770,20585938,Thu Apr 21 20:25:19 2022"
	got := ParseClientListRecord(pattern)
	want := OpenVpnStatus{
		CommonName:     "ma",
		BytesReceived:  1229770,
		BytesSent:      20585938,
		RealAddress:    "178.205.243.63:27437",
		ConnectedSince: time.Date(2022, time.April, 21, 20, 25, 19, 0, time.UTC),
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}