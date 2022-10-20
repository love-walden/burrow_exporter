package burrow_exporter

import "testing"

func TestUUIDRegexp(t *testing.T) {
	var badGroups []string = []string{
		"3e0dfe45-3c11-458a-8b42-49fb85f8880c",
		"84cd0a41-77b6-49bd-9ea4-2158536c5aaa",
		"72f2c9ed-0a13-49f6-b32b-23fcd28692bc",
		"main-dfbfdb8f7-6wc8w",
		"main-84669d549c-btrcn",
	}

	var okGroups []string = []string{
		"group1",
		"another-group",
		"12group",
		"valid-consumer-group",
	}

	for _, badGroup := range badGroups {
		if !filterGroup(badGroup) {
			t.Errorf("expected %s to be caught", badGroup)
		}
	}

	for _, n := range okGroups {
		if filterGroup(n) {
			t.Errorf("unexpected group %s caught", n)
		}
	}
}
