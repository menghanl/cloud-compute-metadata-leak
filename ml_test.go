package ml

import (
	"testing"

	"github.com/codyoss/cloud-compute-metadata-leak/leakcheck"

	"cloud.google.com/go/compute/metadata"
)

func TestML(t *testing.T) {

	defer leakcheck.Check(t)

	if metadata.OnGCE() {
		t.Logf("on gce")
	}
}
