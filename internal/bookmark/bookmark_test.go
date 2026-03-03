package bookmark

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"testing"
)

type action struct {
	name   string
	rawUrl string
}

type AddRemoveTest struct {
	actions     []action
	expctedUrls []Bookmark
}

func TestValidateUrl(t *testing.T) {
	tt := []struct {
		rawUrl   string
		expected error
	}{
		{
			rawUrl:   "https://google.com",
			expected: nil,
		},
	}

	for _, tc := range tt {
		_, got := convertToUrl(tc.rawUrl)

		if !errors.Is(got, tc.expected) {
			fmt.Printf("expected: %v, got: %v\n", tc.expected, got)
			t.Fail()
		}
	}
}

func TestAddRemoveBookmarks(t *testing.T) {
	tt := []AddRemoveTest{
		{
			actions: []action{
				{
					name:   "add",
					rawUrl: "https://google.com",
				},
			},
			expctedUrls: []Bookmark{
				{
					Name: "google.com",
					Url:  "https://google.com",
				},
			},
		},
		{
			actions: []action{
				{
					name:   "remove",
					rawUrl: "https://google.com",
				},
			},
			expctedUrls: nil,
		},
		{
			actions: []action{
				{
					name:   "add",
					rawUrl: "https://google.com/one",
				},
				{
					name:   "add",
					rawUrl: "https://google.com/two",
				},
				{
					name:   "remove",
					rawUrl: "google",
				},
			},
			expctedUrls: nil,
		},
		{
			actions: []action{
				{
					name:   "add",
					rawUrl: "https://google.com",
				},
				{
					name:   "add",
					rawUrl: "https://github.com",
				},
				{
					name:   "remove",
					rawUrl: "google",
				},
			},
			expctedUrls: []Bookmark{
				{
					Name: "github.com",
					Url:  "https://github.com",
				},
			},
		},
	}

	for _, tc := range tt {
		got := Data{}
		for _, action := range tc.actions {
			switch action.name {
			case "add":
				got.Add(action.rawUrl)
			case "remove":
				got.Remove(action.rawUrl)
			}
		}

		cwd, _ := os.Getwd()
		if !slices.Equal(got[cwd], tc.expctedUrls) {
			fmt.Printf("expected: %v, got: %v\n:", tc.expctedUrls, got)
			t.Fail()
		}
	}
}
