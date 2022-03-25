package presentation

import "github.com/Jeffthedoor/lazygit/pkg/gui/style"

func OpensMenuStyle(str string) string {
	return style.FgMagenta.Sprintf("%s...", str)
}
