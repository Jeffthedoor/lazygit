package gui

import (
	"fmt"

	"github.com/Jeffthedoor/gocui"
	"github.com/Jeffthedoor/lazygit/pkg/gui/types"
)

func (gui *Gui) showUpdatePrompt(newVersion string) error {
	return gui.c.Ask(types.AskOpts{
		Title:  "New version available!",
		Prompt: fmt.Sprintf("Download version %s? (enter/esc)", newVersion),
		HandleConfirm: func() error {
			gui.startUpdating(newVersion)
			return nil
		},
	})
}

func (gui *Gui) onUserUpdateCheckFinish(newVersion string, err error) error {
	if err != nil {
		return gui.c.Error(err)
	}
	if newVersion == "" {
		return gui.c.ErrorMsg("New version not found")
	}
	return gui.showUpdatePrompt(newVersion)
}

func (gui *Gui) onBackgroundUpdateCheckFinish(newVersion string, err error) error {
	if err != nil {
		// ignoring the error for now so that I'm not annoying users
		gui.c.Log.Error(err.Error())
		return nil
	}
	if newVersion == "" {
		return nil
	}
	if gui.c.UserConfig.Update.Method == "background" {
		gui.startUpdating(newVersion)
		return nil
	}
	return gui.showUpdatePrompt(newVersion)
}

func (gui *Gui) startUpdating(newVersion string) {
	gui.State.Updating = true
	statusId := gui.statusManager.addWaitingStatus("updating")
	gui.Updater.Update(newVersion, func(err error) error { return gui.onUpdateFinish(statusId, err) })
}

func (gui *Gui) onUpdateFinish(statusId int, err error) error {
	gui.State.Updating = false
	gui.statusManager.removeStatus(statusId)
	gui.OnUIThread(func() error {
		_ = gui.renderString(gui.Views.AppStatus, "")
		if err != nil {
			return gui.c.ErrorMsg("Update failed: " + err.Error())
		}
		return nil
	})

	return nil
}

func (gui *Gui) createUpdateQuitConfirmation() error {
	return gui.c.Ask(types.AskOpts{
		Title:  "Currently Updating",
		Prompt: "An update is in progress. Are you sure you want to quit?",
		HandleConfirm: func() error {
			return gocui.ErrQuit
		},
	})
}
