package presentation

import (
	"github.com/Jeffthedoor/generics/slices"
	"github.com/Jeffthedoor/lazygit/pkg/gui/types"
)

func GetSuggestionListDisplayStrings(suggestions []*types.Suggestion) [][]string {
	return slices.Map(suggestions, func(suggestion *types.Suggestion) []string {
		return getSuggestionDisplayStrings(suggestion)
	})
}

func getSuggestionDisplayStrings(suggestion *types.Suggestion) []string {
	return []string{suggestion.Label}
}
