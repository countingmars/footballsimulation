package commentor

import (
	"fmt"
	"github.com/countingmars/fb/simulater"
	"github.com/countingmars/fb/base/team"
	"github.com/countingmars/fb/base/zone"
)

func Comment(sim *simulater.Simulation) string {
	return fmt.Sprintf("%s vs %s, %s\n" +
		"possession: %d vs %d\n" +
		"first half: %s\n" +
		"second half: %s\n",
		commentTeam(sim.Home), commentTeam(sim.Away), commentResult(sim),
		sim.Possession(simulater.Left), sim.Possession(simulater.Right),
		commentHighlights(sim.First),
		commentHighlights(sim.Second),
	)
}
func commentHighlights(highlights simulater.Highlights) string {
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
		highlights.Possession(simulater.Left), highlights.Possession(simulater.Right),
		highlights.Goals(simulater.Left), highlights.Shoots(simulater.Left) + highlights.Goals(simulater.Left),
		highlights.Goals(simulater.Right), highlights.Shoots(simulater.Right) + highlights.Goals(simulater.Right),
		comments,
	)
}
func commentTeam(aTeam *team.Team) string {
	return fmt.Sprintf("%s (%d)",
		aTeam.Name,
		zone.ZonesFrom(aTeam.Formation).Sum(),
	)
}
func commentResult(sim *simulater.Simulation) string {
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
