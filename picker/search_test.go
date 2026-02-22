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
					"testdata/search/.emptyHiddenDir",
				},
				IgnoreHidden: false,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "_emptyHiddenDir",
					Path:        "testdata/search/.emptyHiddenDir",
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
					"testdata/search",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "testdata/search",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/search/dir",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"testdata/search",
					"testdata/search/.hiddenDir",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "testdata/search",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/search/dir",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/search/.hiddenDir",
				},
				Item{
					SessionName: "002-two",
					Path:        "testdata/search/.hiddenDir/002-two",
				},
				Item{
					SessionName: "003-three",
					Path:        "testdata/search/.hiddenDir/003-three",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"testdata/search",
				},
				IgnoreHidden: false,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "testdata/search",
				},
				Item{
					SessionName: "_emptyHiddenDir",
					Path:        "testdata/search/.emptyHiddenDir",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/search/.hiddenDir",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/search/dir",
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
				"testdata/search/nonExisting",
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
					"testdata/search",
					"testdata/search/dir",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "testdata/search",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/search/dir",
				},
				Item{
					SessionName: "001-hello",
					Path:        "testdata/search/dir/001-hello",
				},
				Item{
					SessionName: "002-world",
					Path:        "testdata/search/dir/002-world",
				},
				Item{
					SessionName: "003-go",
					Path:        "testdata/search/dir/003-go",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"testdata/search",
					"testdata/search/dir",
					"testdata/search/.hiddenDir",
				},
				IgnoreHidden:  true,
				IgnorePattern: "hello|world|002",
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "testdata/search",
				},
				Item{
					SessionName: "dir",
					Path:        "testdata/search/dir",
				},
				Item{
					SessionName: "003-go",
					Path:        "testdata/search/dir/003-go",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "testdata/search/.hiddenDir",
				},
				Item{
					SessionName: "003-three",
					Path:        "testdata/search/.hiddenDir/003-three",
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
