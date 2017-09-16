package commentor

import (
	"fmt"
	. "github.com/countingmars/fb/core"
	. "github.com/countingmars/fb/simulate"
)

func Comment(sim *Simulation) string {
	homeComment := commentTeam(sim.Home)
	awayComment := commentTeam(sim.Away)
	scoreComment := commentScore(sim)
	firstHalfHighlightComment := commentHighlights(sim.First)
	secondHalfHighlightComment := commentHighlights(sim.Second)
	return fmt.Sprintf("%s vs %s\n%s\nfirst half\n%s\nsecond half\n%s",
			homeComment,
			awayComment,
			scoreComment,
			firstHalfHighlightComment,
			secondHalfHighlightComment,
		)
}
func commentHighlights(highlights []Highlight) string {
	return fmt.Sprintf("%v", highlights)
}
func commentTeam(team *Team) string {
	inspection := team.Ability.Inspect()
	return fmt.Sprintf("%s (%d good %s)",
		team.Name,
		team.Ability.Sum(),
		inspection.Style)
}
func commentScore(sim *Simulation) string {
	var result string
	if sim.Draw() {
		result = "D"
	} else if sim.Win() {
		result = "W"
	} else {
		result = "L"
	}
	return fmt.Sprintf("%s:%d-%d", result, sim.Goals().Home, sim.Goals().Away)
}
