package commentor

import (
	"fmt"
	. "github.com/countingmars/fb/base"
	. "github.com/countingmars/fb/simulater"
)

func Comment(sim *Simulation) string {
	return fmt.Sprintf("%s vs %s, %s\n" +
		"possession: %d vs %d\n" +
		"first half: %s\n" +
		"second half: %s\n",
		commentTeam(sim.Home), commentTeam(sim.Away), commentResult(sim),
		sim.Possession(Left), sim.Possession(Right),
		commentHighlights(sim.First),
		commentHighlights(sim.Second),
	)
}
func commentHighlights(highlights Highlights) string {
	var comments = []string{}
	for _, hl := range highlights {
		if hl.Type != "" {
			comments = append(comments, string(hl.Side)+": "+hl.Message)
		}
	}
	return fmt.Sprintf(
		"possessions: %d vs %d\n"+
		"shoots: %d/%d, %d/%d\n" +
		"highlights: %v\n",
		highlights.Possession(Left), highlights.Possession(Right),
		highlights.Goals(Left), highlights.Shoots(Left) + highlights.Goals(Left),
		highlights.Goals(Right), highlights.Shoots(Right) + highlights.Goals(Right),
		comments,
	)
}
func commentTeam(team *Team) string {
	return fmt.Sprintf("%s (%d good %s)",
		team.Name,
		team.Ability.Sum(),
		team.Ability.Inspect().Style,
	)
}
func commentResult(sim *Simulation) string {
	var result string
	if sim.Draw() {
		result = "D"
	} else if sim.Win() {
		result = "W"
	} else {
		result = "L"
	}
	return fmt.Sprintf("%s:%d-%d",
		result,
		sim.Goals().Home,
		sim.Goals().Away,
	)
}
