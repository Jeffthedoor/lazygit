package presentation

import (
	"github.com/Jeffthedoor/generics/slices"
	"github.com/Jeffthedoor/lazygit/pkg/commands/models"
	"github.com/Jeffthedoor/lazygit/pkg/theme"
)

func GetTagListDisplayStrings(tags []*models.Tag, diffName string) [][]string {
	return slices.Map(tags, func(tag *models.Tag) []string {
		diffed := tag.Name == diffName
		return getTagDisplayStrings(tag, diffed)
	})
}

// getTagDisplayStrings returns the display string of branch
func getTagDisplayStrings(t *models.Tag, diffed bool) []string {
	textStyle := theme.DefaultTextColor
	if diffed {
		textStyle = theme.DiffTerminalColor
	}
	return []string{textStyle.Sprint(t.Name)}
}
