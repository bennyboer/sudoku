package strategy

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	lcc := &LinearCellChooser{}
	linear, _ := Create(Linear)
	if reflect.TypeOf(linear) != reflect.TypeOf(lcc) {
		t.Errorf("Expected linear cell chooser type")
	}

	rcc := &RandomCellChooser{}
	random, _ := Create(Random)
	if reflect.TypeOf(random) != reflect.TypeOf(rcc) {
		t.Errorf("Expected random cell chooser type")
	}

	_, e := Create(5)
	if e == nil {
		t.Errorf("Expected error because of unknown cell chooser type")
	}
}
