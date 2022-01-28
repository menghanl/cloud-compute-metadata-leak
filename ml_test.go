package ml

import (
	"testing"
	"time"

	"cloud.google.com/go/compute/metadata"
	"go.uber.org/goleak"

)

func TestML(t *testing.T) {
	defer func() {
		time.Sleep(time.Second * 5)
		goleak.VerifyNone(t)
	}()

	if metadata.OnGCE() {
		t.Logf("on gce")
	}
}
