package commentor

import (
	"fmt"
	. "github.com/countingmars/fb/foundation"
	. "github.com/countingmars/fb/gamesimulator"
)

type Commentor struct {
	GameSimulation *GameSimulation
}

type Comment struct {
	Message string
}

func (this *Commentor) Comment(sim *GameSimulation) Comment {
	homeComment := commentTeam(sim.Game.Home)
	awayComment := commentTeam(sim.Game.Away)
	scoreComment := commentScore(sim)
	highlightComment := commentHighlights(sim.Highlights)
	return Comment {
		Message: fmt.Sprintf("%s vs %s\n%s\n%s",
			homeComment,
			awayComment,
			scoreComment,
			highlightComment),
	}
}
func commentHighlights(highlights []Highlight) string {
	return fmt.Sprintf("%v", highlights)
}
func commentTeam(team *Team) string {
	inspection := team.Ability.Inspect()
	effected := team.Ability.AmendBySelf()
	return fmt.Sprintf("%s (%d[%d] good %s)",
		team.Name,
		team.Ability.Sum(),
		effected.Sum(),
		inspection.Style)
}
func commentScore(sim *GameSimulation) string {
	var result string
	if sim.Draw {
		result = "D"
	} else if sim.Winner == sim.Game.Home {
		result = "W"
	} else {
		result = "L"
	}
	return fmt.Sprintf("%s:%d-%d", result, sim.Score.Home, sim.Score.Away)
}
