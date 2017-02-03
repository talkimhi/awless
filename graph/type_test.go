package graph

import (
	"testing"

	"github.com/google/badwolf/triple/node"
)

func TestResourceTypeToRdfType(t *testing.T) {
	if got, want := Region.ToRDFString(), "/region"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := Region.String(), "region"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}

	resourceTypes := []ResourceType{Region, Vpc, Subnet, Instance, User, Role, Group, Policy}
	for _, r := range resourceTypes {
		_, err := node.NewType(r.ToRDFString())
		if err != nil {
			t.Fatal(err)
		}
		if got, want := "/"+r.String(), r.ToRDFString(); got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	}
}