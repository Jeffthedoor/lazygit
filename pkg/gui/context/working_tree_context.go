package context

import (
	"github.com/Jeffthedoor/gocui"
	"github.com/Jeffthedoor/lazygit/pkg/commands/models"
	"github.com/Jeffthedoor/lazygit/pkg/gui/filetree"
	"github.com/Jeffthedoor/lazygit/pkg/gui/types"
)

type WorkingTreeContext struct {
	*filetree.FileTreeViewModel
	*ListContextTrait
}

var _ types.IListContext = (*WorkingTreeContext)(nil)

func NewWorkingTreeContext(
	getModel func() []*models.File,
	view *gocui.View,
	getDisplayStrings func(startIdx int, length int) [][]string,

	onFocus func(...types.OnFocusOpts) error,
	onRenderToMain func(...types.OnFocusOpts) error,
	onFocusLost func() error,

	c *types.HelperCommon,
) *WorkingTreeContext {
	viewModel := filetree.NewFileTreeViewModel(getModel, c.Log, c.UserConfig.Gui.ShowFileTree)

	return &WorkingTreeContext{
		FileTreeViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				ViewName:   "files",
				WindowName: "files",
				Key:        FILES_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
			}), ContextCallbackOpts{
				OnFocus:        onFocus,
				OnFocusLost:    onFocusLost,
				OnRenderToMain: onRenderToMain,
			}),
			list:              viewModel,
			viewTrait:         NewViewTrait(view),
			getDisplayStrings: getDisplayStrings,
			c:                 c,
		},
	}
}

func (self *WorkingTreeContext) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.ID()
}
