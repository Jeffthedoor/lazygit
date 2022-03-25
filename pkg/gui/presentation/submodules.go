package presentation

import (
	"github.com/Jeffthedoor/generics/slices"
	"github.com/Jeffthedoor/lazygit/pkg/commands/models"
	"github.com/Jeffthedoor/lazygit/pkg/theme"
)

func GetSubmoduleListDisplayStrings(submodules []*models.SubmoduleConfig) [][]string {
	return slices.Map(submodules, func(submodule *models.SubmoduleConfig) []string {
		return getSubmoduleDisplayStrings(submodule)
	})
}

func getSubmoduleDisplayStrings(s *models.SubmoduleConfig) []string {
	return []string{theme.DefaultTextColor.Sprint(s.Name)}
}
