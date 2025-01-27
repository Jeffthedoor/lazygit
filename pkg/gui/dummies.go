package gui

import (
	"github.com/Jeffthedoor/lazygit/pkg/commands/git_config"
	"github.com/Jeffthedoor/lazygit/pkg/commands/oscommands"
	"github.com/Jeffthedoor/lazygit/pkg/config"
	"github.com/Jeffthedoor/lazygit/pkg/updates"
	"github.com/Jeffthedoor/lazygit/pkg/utils"
)

// NewDummyGui creates a new dummy GUI for testing
func NewDummyUpdater() *updates.Updater {
	newAppConfig := config.NewDummyAppConfig()
	dummyUpdater, _ := updates.NewUpdater(utils.NewDummyCommon(), newAppConfig, oscommands.NewDummyOSCommand())
	return dummyUpdater
}

func NewDummyGui() *Gui {
	newAppConfig := config.NewDummyAppConfig()
	dummyGui, _ := NewGui(utils.NewDummyCommon(), newAppConfig, git_config.NewFakeGitConfig(nil), NewDummyUpdater(), false, "")
	return dummyGui
}
