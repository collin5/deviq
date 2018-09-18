package test

import (
	"testing"
	"util"
)

func TestNewContainer(t *testing.T) {
	cont := util.NewContainer("foo", "bar")
	if cont == nil {
		t.Errorf("Failed to create container")
	}
}

func TestNewContainerProps(t *testing.T) {
	expected := map[string]string{
		"name":  "foo",
		"lang ": "bar",
	}
	cont := util.NewContainer(expected["name"], expected["lang"])
	if cont.Name != expected["name"] {
		t.Error("Invalid container name")
	}
	if cont.Lang != expected["lang"] {
		t.Error("Invalid lang for container")
	}
}

func TestContainerRun(t *testing.T) {

}
