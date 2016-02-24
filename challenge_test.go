package main

import (
	"testing"
)

func TestSingleLevelPermissionCheck(t *testing.T) {
	Load_rule_csv()
	if !HasAuthorized("D1", "IN") {
		t.Errorf("D1 has authorization to distribute in IN")
	}

	if !HasAuthorized("D1", "US") {
		t.Errorf("D1 has authorization to distribute in US")
	}

	if HasAuthorized("D1", "UK") {
		t.Errorf("D1 doesn't have authoriUKzation to distribute in UK")
	}

	if !HasAuthorized("D2", "IN") {
		t.Errorf("D2 authorization to distribute in IN")
	}

	if !HasAuthorized("D2", "PUNE-MH-IN") {
		t.Errorf("D2 has authorization to distribute in MH-IN")
	}

	// As when checking only we are identifying invalid authorization mapping.
	// Could be added when creating the data structure itself and remove the
	// conflicting distributer and throw error.
	if HasAuthorized("D3", "HUBLI-KA-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in HUBLI-KA-IN as D2 Doesn't hve it.")
	}
}

func TestMultipleLevelPermissionCheck(t *testing.T) {
	Load_rule_csv()

	if HasAuthorized("D2", "TN-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in TN-IN")
	}

	if HasAuthorized("D2", "KA-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D2", "CENAI-TN-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D3", "TN-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in TN-IN")
	}

	if HasAuthorized("D3", "KA-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in KA-IN")
	}

	if !HasAuthorized("D3", "MH-IN") {
		t.Errorf("D3 has authorization to distribute in MH-IN")
	}

	if !HasAuthorized("D3", "PUNE-MH-IN") {
		t.Errorf("D3 has authorization to distribute in PUNE-MH-IN")
	}
	if HasAuthorized("D3", "CN") {
		t.Errorf("D3 has doesn't have authorization to distribute in CN")
	}

	if HasAuthorized("D4", "KA-IN") {
		t.Errorf("D4 has doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D4", "KL-IN") {
		t.Errorf("D4 doesn't have authorization to distribute in KL-IN")
	}

	if !HasAuthorized("D4", "PUNE-MH-IN") {
		t.Errorf("D4 has authorization to distribute in PUNE-MH-IN")
	}

	if !HasAuthorized("D5", "IN") {
		t.Errorf("D5 has authorization to distribute in IN")
	}

	if HasAuthorized("D5", "PY-IN") {
		t.Errorf("D5 doesn't have authorization to distribute in PY-IN")
	}

	if HasAuthorized("D5", "KA-IN") {
		t.Errorf("D5 doesn't have authorization to distribute in KA-IN")
	}
	if HasAuthorized("D5", "KL-IN") {
		t.Errorf("D5 doesn't have authorization to distribute in KL-IN")
	}

	if HasAuthorized("D6", "KL-IN") {
		t.Errorf("D6 doesn't have authorization to distribute in KL-IN")
	}
	if HasAuthorized("D6", "PY-IN") {
		t.Errorf("D6 doesn't have authorization to distribute in PY-IN")
	}
	if !HasAuthorized("D6", "TN-IN") {
		t.Errorf("D6 has authorization to distribute in TN-IN")
	}
}
