package loaders

import (
	"testing"

	"github.com/Jeffthedoor/lazygit/pkg/commands/models"
	"github.com/Jeffthedoor/lazygit/pkg/commands/oscommands"
	"github.com/Jeffthedoor/lazygit/pkg/utils"
	"github.com/stretchr/testify/assert"
)

const tagsOutput = `v0.34
v0.33
v0.32.2
v0.32.1
v0.32
testtag
`

func TestGetTags(t *testing.T) {
	type scenario struct {
		testName      string
		runner        *oscommands.FakeCmdObjRunner
		expectedTags  []*models.Tag
		expectedError error
	}

	scenarios := []scenario{
		{
			testName: "should return no tags if there are none",
			runner: oscommands.NewFakeRunner(t).
				Expect(`git tag --list --sort=-creatordate`, "", nil),
			expectedTags:  []*models.Tag{},
			expectedError: nil,
		},
		{
			testName: "should return tags if present",
			runner: oscommands.NewFakeRunner(t).
				Expect(`git tag --list --sort=-creatordate`, tagsOutput, nil),
			expectedTags: []*models.Tag{
				{Name: "v0.34"},
				{Name: "v0.33"},
				{Name: "v0.32.2"},
				{Name: "v0.32.1"},
				{Name: "v0.32"},
				{Name: "testtag"},
			},
			expectedError: nil,
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.testName, func(t *testing.T) {
			loader := &TagLoader{
				Common: utils.NewDummyCommon(),
				cmd:    oscommands.NewDummyCmdObjBuilder(scenario.runner),
			}

			tags, err := loader.GetTags()

			assert.Equal(t, scenario.expectedTags, tags)
			assert.Equal(t, scenario.expectedError, err)

			scenario.runner.CheckForMissingCalls()
		})
	}
}
