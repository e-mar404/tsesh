package picker

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/list"
	"github.com/e-mar404/tsesh/config"
)

type searchTest struct {
	cfg          config.Search
	expectedList []list.Item
}

func TestWorkingDir(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Printf("working dir %v\n", dir)
}

func TestSessionName(t *testing.T) {
	tt := []searchTest{
		{
			cfg: config.Search{
				Paths: []string{
					"testdata/.hiddenDir",
				},
				IgnorePattern: "*",
			},
			expectedList: []list.Item{
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/.hiddenDir",
				},
			},
		},
	}

	for _, tc := range tt {
		got := searchPaths(tc.cfg)

		if !reflect.DeepEqual(tc.expectedList, got) {
			fmt.Printf("expected: %v, got: %v\n", tc.expectedList, got)
			t.Fail()
		}
	}
}

func TestHiddenDirs(t *testing.T) {
	tt := []searchTest{
		{
			cfg: config.Search{
				Paths: []string{
					"testdata",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "testdata",
					Path:        "testdata",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/dir",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"testdata",
					"testdata/.hiddenDir",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "testdata",
					Path:        "testdata",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/dir",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/.hiddenDir",
				},
				Item{
					SessionName: "002-two",
					Path:        "testdata/.hiddenDir/002-two",
				},
				Item{
					SessionName: "003-three",
					Path:        "testdata/.hiddenDir/003-three",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"testdata",
				},
				IgnoreHidden: false,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "testdata",
					Path:        "testdata",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/.hiddenDir",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/dir",
				},
			},
		},
	}

	for _, tc := range tt {
		got := searchPaths(tc.cfg)

		if !reflect.DeepEqual(tc.expectedList, got) {
			fmt.Printf("expected: %v, got: %v\n", tc.expectedList, got)
			t.Fail()
		}
	}
}

func TestNonExistingDir(t *testing.T) {
	tc := searchTest{
		cfg: config.Search{
			Paths: []string{
				"testdata/nonExisting",
			},
		},
		expectedList: []list.Item{},
	}

	got := searchPaths(tc.cfg)

	if !reflect.DeepEqual(tc.expectedList, got) {
		fmt.Printf("expected: %v, got: %v\n", tc.expectedList, got)
		t.Fail()
	}
}

func TestIgnorePattern(t *testing.T) {
	tt := []searchTest{
		{
			cfg: config.Search{
				Paths: []string{
					"testdata",
					"testdata/dir",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "testdata",
					Path:        "testdata",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/dir",
				},
				Item{
					SessionName: "001-hello",
					Path:        "testdata/dir/001-hello",
				},
				Item{
					SessionName: "002-world",
					Path:        "testdata/dir/002-world",
				},
				Item{
					SessionName: "003-go",
					Path:        "testdata/dir/003-go",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"testdata",
					"testdata/dir",
					"testdata/.hiddenDir",
				},
				IgnoreHidden:  true,
				IgnorePattern: "hello|world|002",
			},
			expectedList: []list.Item{
				Item{
					SessionName: "testdata",
					Path:        "testdata",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/dir",
				},
				Item{
					SessionName: "003-go",
					Path:        "testdata/dir/003-go",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/.hiddenDir",
				},
				Item{
					SessionName: "003-three",
					Path:        "testdata/.hiddenDir/003-three",
				},
			},
		},
	}

	for _, tc := range tt {
		got := searchPaths(tc.cfg)

		if !reflect.DeepEqual(tc.expectedList, got) {
			fmt.Printf("expected: %v, got: %v\n", tc.expectedList, got)
			t.Fail()
		}
	}
}
