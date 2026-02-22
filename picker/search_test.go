package picker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbles/list"
	"github.com/e-mar404/tsesh/config"
)

type searchTest struct {
	cfg          config.Search
	expectedList []list.Item
}

func TestSessionName(t *testing.T) {
	tt := []searchTest{
		{
			cfg: config.Search{
				Paths: []string{
					"../testData/search/.emptyHiddenDir",
				},
				IgnoreHidden: false,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "_emptyHiddenDir",
					Path:        "../testData/search/.emptyHiddenDir",
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
					"../testData/search",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "../testData/search",
				},
				Item{
					SessionName: "dir",
					Path:        "../testData/search/dir",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"../testData/search",
					"../testData/search/.hiddenDir",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "../testData/search",
				},
				Item{
					SessionName: "dir",
					Path:        "../testData/search/dir",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "../testData/search/.hiddenDir",
				},
				Item{
					SessionName: "002-two",
					Path:        "../testData/search/.hiddenDir/002-two",
				},
				Item{
					SessionName: "003-three",
					Path:        "../testData/search/.hiddenDir/003-three",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"../testData/search",
				},
				IgnoreHidden: false,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "../testData/search",
				},
				Item{
					SessionName: "_emptyHiddenDir",
					Path:        "../testData/search/.emptyHiddenDir",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "../testData/search/.hiddenDir",
				},
				Item{
					SessionName: "dir",
					Path:        "../testData/search/dir",
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
				"../testData/search/nonExisting",
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
					"../testData/search",
					"../testData/search/dir",
				},
				IgnoreHidden: true,
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "../testData/search",
				},
				Item{
					SessionName: "dir",
					Path:        "../testData/search/dir",
				},
				Item{
					SessionName: "001-hello",
					Path:        "../testData/search/dir/001-hello",
				},
				Item{
					SessionName: "002-world",
					Path:        "../testData/search/dir/002-world",
				},
				Item{
					SessionName: "003-go",
					Path:        "../testData/search/dir/003-go",
				},
			},
		},
		{
			cfg: config.Search{
				Paths: []string{
					"../testData/search",
					"../testData/search/dir",
					"../testData/search/.hiddenDir",
				},
				IgnoreHidden:  true,
				IgnorePattern: "hello|world|002",
			},
			expectedList: []list.Item{
				Item{
					SessionName: "search",
					Path:        "../testData/search",
				},
				Item{
					SessionName: "dir",
					Path:        "../testData/search/dir",
				},
				Item{
					SessionName: "003-go",
					Path:        "../testData/search/dir/003-go",
				},
				Item{
					SessionName: "_hiddenDir",
					Path:        "../testData/search/.hiddenDir",
				},
				Item{
					SessionName: "003-three",
					Path:        "../testData/search/.hiddenDir/003-three",
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
