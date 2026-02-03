package tmux

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

/*
Taking inspiration from bubbletea/exec_test.go of making a small tea program to see if the error gets returned properly
*/

type attachModel struct {
	sessionName string
	Err error
}

func (am attachModel) Init() tea.Cmd {
	return Attach(am.sessionName)
}

func (am attachModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TmuxMsg:
		am.Err = msg.Err
		return am, tea.Quit
	}
	return am, nil
}

func (am attachModel) View() string {
	return "\n"
}

func TestAttachArgs(t *testing.T) {
	t.Setenv("TMUX", "")
	cmdRunner = mockCommand()
	Attach("test-attach")
	expectedArgs := []string {"attach-session", "-t", "test-attach"}
	if !reflect.DeepEqual(expectedArgs, capturedArgs) {
		failArgsDoNotMatch(t, expectedArgs, capturedArgs)
	}
}

func TestAttachInsideTmux(t *testing.T) {
	t.Setenv("TMUX", "inside")
	cmdRunner = mockCommand()

	var in, out bytes.Buffer 

	initModel := attachModel {
		sessionName: "attach inside tmux",
	}

	p := tea.NewProgram(initModel, tea.WithInput(&in), tea.WithOutput(&out))
	outModel, _ := p.Run()
	finalModel := outModel.(attachModel)

	if finalModel.Err != ErrNestedSession {
		fmt.Printf("tea.Cmd returned something unexpected\n")
		fmt.Printf("expected: %v, got: %v\n", ErrNestedSession, finalModel.Err)
		t.FailNow()
	}
}

func TestAttachOutsideTmux(t *testing.T) {
	t.Setenv("TMUX", "")
	
	tt := []struct {
		sessionExists bool	
		expectedErr error
	}{
		{
			sessionExists: true,
			expectedErr: nil,
		},
		{
			sessionExists: false,
			expectedErr: fmt.Errorf("can't find session: non-existing"),
		},
	}

	var in, out bytes.Buffer
	for _, tc := range tt {
		var sessionName string
		if tc.sessionExists {
			sessionName = "existing"
			cmdRunner = mockCommand()
		} else {
			sessionName = "non-existing"
			cmdRunner = mockCommand(
				withNonExistingSession,
			)
		}

		initModel := attachModel {
			sessionName: sessionName,
		}
		p := tea.NewProgram(initModel, tea.WithInput(&in), tea.WithOutput(&out))
		outModel, err := p.Run()
		if err != nil {
			t.Fatalf("something went wrong with the tea program: %v\n", err)
		}
		finalModel := outModel.(attachModel)
		
		// TODO: see if I should simplify this error check
		if tc.sessionExists {
			if finalModel.Err != nil {
				fmt.Printf("expected: %v, got: %v\n", tc.expectedErr, finalModel.Err)
				t.FailNow()
			}

			return
		} 

		if errors.Is(finalModel.Err, tc.expectedErr) {
			fmt.Printf("expected: %v, got: %v\n", tc.expectedErr, finalModel.Err)
			t.FailNow()
		}
	}
}
