package tmux

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

/*

SwitchClient:
	- TestSwitchClientArgs
	- TestSwitchClientOutsideTmux
	- TestSwitchClientInsideTmux
	- TestSwitchClientExist
	- TestSwitchClientDoesNotExist
*/

/*
Taking inspiration from bubbletea/exec_test.go of making a small tea program to see if the error gets returned properly
*/

type mockOption func(*exec.Cmd) *exec.Cmd
type execCommand func(string, ...string) *exec.Cmd
type model struct {
	sessionName string
	testCmd func(args...string) tea.Cmd
	Err error
}
type tmuxTest struct {
	cmdRunner execCommand
	sessionName string
	testCmd func(args...string) tea.Cmd
	expectedErr error
}

var capturedArgs []string

func (m model) Init() tea.Cmd {
	return m.testCmd(m.sessionName) 
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TmuxMsg:
		m.Err = msg.Err
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	return "\n"
}

func TestMain(m *testing.M) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		var exitCode int

		codeStr := os.Getenv("MOCK_CMD_EXIT_CODE")
		if codeStr == "" {
			exitCode = 0
		} else {
			exitCode, _ = strconv.Atoi(codeStr)
		}
		
		os.Exit(exitCode)
	}

	os.Exit(m.Run())
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

func TestInsideTmux(t *testing.T) {
	t.Setenv("TMUX", "inside")
	var in, out bytes.Buffer 

	tt := []tmuxTest {
		{
			cmdRunner: mockCommand(),
			sessionName: "attach inside tmux",
			testCmd: func(args...string) tea.Cmd {
				return Attach(args[0])
			},
			expectedErr: ErrNestedSession,
		},
	}

	for _, tc := range tt {
		cmdRunner = tc.cmdRunner 
		initModel := model {
			sessionName: tc.sessionName,
			testCmd: tc.testCmd,
		}

		p := tea.NewProgram(initModel, tea.WithInput(&in), tea.WithOutput(&out))
		outModel, _ := p.Run()
		finalModel := outModel.(model)

		if !errors.Is(finalModel.Err, tc.expectedErr) {
			fmt.Printf("tea.Cmd returned something unexpected\n")
			fmt.Printf("expected: %v, got: %v\n", ErrNestedSession, finalModel.Err)
			t.FailNow()
		}
	}
}

func TestOutsideTmux(t *testing.T) {
	t.Setenv("TMUX", "")

	tt := []tmuxTest {
		{
			cmdRunner: mockCommand(),
			testCmd: func(args...string) tea.Cmd {
				return Attach(args[0])
			},
			sessionName: "attach-existing-session",
			expectedErr: nil,
		},
		{
			cmdRunner: mockCommand(
				withNonExistingSession,
			),
			testCmd: func(args...string) tea.Cmd {
				return Attach(args[0])
			},
			sessionName: "attach-non-existing-session",
			expectedErr: ErrSessionNotFound, 
		},
	}

	var in, out bytes.Buffer
	for _, tc := range tt {
		cmdRunner = tc.cmdRunner
		initModel := model {
			sessionName: tc.sessionName,
			testCmd: tc.testCmd,
		}

		p := tea.NewProgram(initModel, tea.WithInput(&in), tea.WithOutput(&out))
		outModel, err := p.Run()
		if err != nil {
			t.Fatalf("something went wrong with the tea program: %v\n", err)
		}
		finalModel := outModel.(model)

		if !errors.Is(finalModel.Err, tc.expectedErr) {
			fmt.Printf("expected: %v, got: %v\n", tc.expectedErr, finalModel.Err)
			t.FailNow()
		}
	}
}

// Unless another option that modifies the exit code is passed it will default to exiting with code 0
func mockCommand(mockOpts...mockOption) execCommand {
	return func(command string, args...string) *exec.Cmd {
		capturedArgs = args
		testBinary := os.Args[0]

		cs := []string{"--", command}
		cs = append(cs, args...) 

		cmd := exec.Command(testBinary, cs...)
		cmd.Env = []string{
			"GO_WANT_HELPER_PROCESS=1",
		}
		
		for _, f := range mockOpts {
			cmd = f(cmd)
		}

		return cmd
	}
}

func withExitCodeOne(cmd *exec.Cmd) *exec.Cmd {
	cmd.Env = append(cmd.Env, "MOCK_CMD_EXIT_CODE=1")
	return cmd
}

func withNonExistingSession(cmd *exec.Cmd) *exec.Cmd {
	cmd.Err = fmt.Errorf("can't find session") 
	return cmd
}

func failArgsDoNotMatch(t *testing.T, expectedArgs, capturedArgs []string) {
	fmt.Println("arguments passed to tmux do not match")
	fmt.Printf("expected: %v, got:%v\n", expectedArgs, capturedArgs)
	t.FailNow()
}
