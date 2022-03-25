package types

import (
	"github.com/Jeffthedoor/lazygit/pkg/gui/modes/cherrypicking"
	"github.com/Jeffthedoor/lazygit/pkg/gui/modes/diffing"
	"github.com/Jeffthedoor/lazygit/pkg/gui/modes/filtering"
)

type Modes struct {
	Filtering     filtering.Filtering
	CherryPicking *cherrypicking.CherryPicking
	Diffing       diffing.Diffing
}
