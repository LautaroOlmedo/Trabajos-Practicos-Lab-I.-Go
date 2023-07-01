package exercises

import "testing"

func TestValidationOfInputValues(t *testing.T) {

	if validationOfInputValues(13, 3) {
		t.Logf("PASS. Expected: %v, got: %v", true, true)
	} else {
		t.Fatalf("ERROR. Expected: %v, got: %v", true, false)
	}
}

func TestValidateOutApplication(t *testing.T) {

}
