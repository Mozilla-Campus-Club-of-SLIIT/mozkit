package engine

import (
	"testing"
)

func TestScriptList(t *testing.T) {
	items := ScriptList()

	if len(items) == 0 {
		t.Fatalf("expected at least one script, got 0")
	}

	t.Logf("Found %d scripts:", len(items))
	for i, item := range items {
		t.Logf("%d. %s (%s)", i+1, item.TitleStr, item.Target)
		if item.TitleStr == "" {
			t.Errorf("item %d has empty title", i)
		}
		if item.DescStr == "" {
			t.Errorf("item %d has empty description", i)
		}
		if item.Target == "" {
			t.Errorf("item %d has empty target", i)
		}
	}

	// Verify sorting by title
	for i := 1; i < len(items); i++ {
		prev := items[i-1].TitleStr
		curr := items[i].TitleStr
		if prev > curr {
			t.Errorf("items are not sorted by title: '%s' comes before '%s'", prev, curr)
		}
	}
}
