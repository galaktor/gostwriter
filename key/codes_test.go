/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package key

import(
	"testing"
)

func TestGetAllCodes_ReturnsCODE_CNT_Entries(t *testing.T) {
	expected := int(CODE_CNT)
	
	keys := getAllCodes()

	actual := len(keys)
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}

func TestAllCodes_ReturnsCODE_CNT_Entries(t *testing.T) {
	expected := int(CODE_CNT)
	
	keys := ALL_CODES

	actual := len(keys)
	if actual != expected {
		t.Errorf("expected '%v' but found '%v'", expected, actual)
	}
}
