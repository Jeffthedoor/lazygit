package git_commands

import (
	"sync"

	gogit "github.com/Jeffthedoor/go-git/v5"
	"github.com/Jeffthedoor/lazygit/pkg/commands/oscommands"
	"github.com/Jeffthedoor/lazygit/pkg/common"
)

type GitCommon struct {
	*common.Common
	cmd       oscommands.ICmdObjBuilder
	os        *oscommands.OSCommand
	dotGitDir string
	repo      *gogit.Repository
	config    *ConfigCommands
	// mutex for doing things like push/pull/fetch
	syncMutex *sync.Mutex
}

func NewGitCommon(
	cmn *common.Common,
	cmd oscommands.ICmdObjBuilder,
	osCommand *oscommands.OSCommand,
	dotGitDir string,
	repo *gogit.Repository,
	config *ConfigCommands,
	syncMutex *sync.Mutex,
) *GitCommon {
	return &GitCommon{
		Common:    cmn,
		cmd:       cmd,
		os:        osCommand,
		dotGitDir: dotGitDir,
		repo:      repo,
		config:    config,
		syncMutex: syncMutex,
	}
}
