package ml

import (
	"testing"

	"ml.test.com/leakcheck"

	"cloud.google.com/go/compute/metadata"

	_ "google.golang.org/grpc"
)

func TestML(t *testing.T) {

	defer leakcheck.Check(t)

	if metadata.OnGCE() {
		t.Logf("on gce")
	}
}
