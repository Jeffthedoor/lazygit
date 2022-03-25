package helpers

import (
	"github.com/Jeffthedoor/lazygit/pkg/commands"
	"github.com/Jeffthedoor/lazygit/pkg/commands/types/enums"
	"github.com/Jeffthedoor/lazygit/pkg/gui/types"
)

type IPatchBuildingHelper interface {
	ValidateNormalWorkingTreeState() (bool, error)
}

type PatchBuildingHelper struct {
	c   *types.HelperCommon
	git *commands.GitCommand
}

func NewPatchBuildingHelper(
	c *types.HelperCommon,
	git *commands.GitCommand,
) *PatchBuildingHelper {
	return &PatchBuildingHelper{
		c:   c,
		git: git,
	}
}

func (self *PatchBuildingHelper) ValidateNormalWorkingTreeState() (bool, error) {
	if self.git.Status.WorkingTreeState() != enums.REBASE_MODE_NONE {
		return false, self.c.ErrorMsg(self.c.Tr.CantPatchWhileRebasingError)
	}
	return true, nil
}
