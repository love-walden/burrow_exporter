package burrow_exporter

import "testing"

func TestUUIDRegexp(t *testing.T) {
	var uuids []string = []string{
		"3e0dfe45-3c11-458a-8b42-49fb85f8880c",
		"84cd0a41-77b6-49bd-9ea4-2158536c5aaa",
		"72f2c9ed-0a13-49f6-b32b-23fcd28692bc",
	}

	var notAUUID []string = []string{
		"group1",
		"another-group",
		"12group",
	}

	for _, uuid := range uuids {
		if !uuidFilter.MatchString(uuid) {
			t.Error("expected uuid to be caught")
		}
	}

	for _, n := range notAUUID {
		if uuidFilter.MatchString(n) {
			t.Error("unexpected non-uuid caught")
		}
	}
}
