package zone

import (
	"testing"
	"sort"
)

func TestEntries_Sort(t *testing.T) {
	entries := Entries{
		&Entry{Stats:Stats{LastPoint: 1}},
		&Entry{Stats:Stats{LastPoint: 2}},
		&Entry{Stats:Stats{LastPoint: 3}},
	}

	sort.Sort(entries)

	if entries[0].Stats.LastPoint != 3 {
		t.Error("expected 3, but ", entries[0].Stats.LastPoint)
	}
}
