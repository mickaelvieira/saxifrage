package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSection(t *testing.T) {
	s := make(Sections, 2)
	s[0] = &Section{Type: HostType, Matching: "localhost"}
	s[1] = &Section{Type: MatchType, Matching: "localhost"}

	got := s.Find(func(s *Section) bool {
		return s.Matching == "localhost"
	})

	assert.Equal(t, HostType, got.Type)
	assert.Equal(t, "localhost", got.Matching)

	got = s.Find(func(s *Section) bool {
		return s.Matching == "localhost" && s.Type == MatchType
	})

	assert.Equal(t, MatchType, got.Type)
	assert.Equal(t, "localhost", got.Matching)

	got = s.Find(func(s *Section) bool {
		return s.Matching == "foo"
	})
	assert.Nil(t, got)
}

func TestFindSectionIndex(t *testing.T) {
	s := make(Sections, 2)
	s[0] = &Section{Type: HostType, Matching: "localhost"}
	s[1] = &Section{Type: MatchType, Matching: "localhost"}

	got := s.FindIndex(func(s *Section) bool {
		return s.Matching == "localhost"
	})

	assert.Equal(t, 0, got)

	got = s.FindIndex(func(s *Section) bool {
		return s.Matching == "localhost" && s.Type == MatchType
	})

	assert.Equal(t, 1, got)

	got = s.FindIndex(func(s *Section) bool {
		return s.Matching == "foo"
	})

	assert.Equal(t, -1, got)
}

func TestFindSectionAt(t *testing.T) {
	s1 := &Section{Type: HostType, Matching: "localhost"}
	s2 := &Section{Type: MatchType, Matching: "localhost"}

	s := make(Sections, 2)
	s[0] = s1
	s[1] = s2

	got := s.FindAt(0)
	assert.Equal(t, s1, got)

	got = s.FindAt(1)
	assert.Equal(t, s2, got)

	got = s.FindAt(-1)
	assert.Nil(t, got)

	got = s.FindAt(2)
	assert.Nil(t, got)
}

func TestFilterSections(t *testing.T) {
	s1 := &Section{Type: HostType, Matching: "localhost"}
	s2 := &Section{Type: MatchType, Matching: "localhost"}

	s := make(Sections, 2)
	s[0] = s1
	s[1] = s2

	got := s.Filter(func(s *Section) bool {
		return s.Matching == "localhost"
	})
	assert.Equal(t, s1, got[0])
	assert.Equal(t, s2, got[1])
	assert.Equal(t, 2, len(got))

	got = s.Filter(func(s *Section) bool {
		return s.Type == MatchType
	})
	assert.Equal(t, s2, got[0])
	assert.Equal(t, 1, len(got))

	got = s.Filter(func(s *Section) bool {
		return s.Type == HostType
	})
	assert.Equal(t, s1, got[0])
	assert.Equal(t, 1, len(got))

	got = s.Filter(func(s *Section) bool {
		return s.Matching == "foo"
	})
	assert.Equal(t, 0, len(got))
}

func TestSectionKeyFiles(t *testing.T) {

	home := os.Getenv("HOME")

	cases := []struct {
		input string
		want  []string
	}{
		{"~/.ssh/id_rsa", []string{
			home + "/.ssh/id_rsa",
			home + "/.ssh/id_rsa.pub",
		}},
		{"~/.ssh/foo/id_rsa", []string{
			home + "/.ssh/foo/id_rsa",
			home + "/.ssh/foo/id_rsa.pub",
			// home + "/.ssh/foo",
		}},
		{"~/.ssh/baz/bar", []string{
			home + "/.ssh/baz/bar",
			home + "/.ssh/baz/bar.pub",
			// home + "/.ssh/baz",
		}},
		// {"~/foo", []string{}},
	}

	for i, tc := range cases {
		s := Section{}
		s.Options = append(s.Options, NewOption("IdentityFile", " ", tc.input))

		f := s.GetKeyFiles()

		assert.Equal(t, len(tc.want), len(f), "Test Case %d [length] %v", i, tc)

		for j, p := range tc.want {
			assert.Equal(t, p, f[j], "Test Case %d [paths] %v", i, tc)
		}
	}
}

func TestFindOptions(t *testing.T) {
	o1 := &Option{Name: "host"}
	o2 := &Option{Name: "IdentityFile"}

	o := make(Options, 2)
	o[0] = o1
	o[1] = o2

	got := o.Find("Host")
	assert.Equal(t, o1, got)

	got = o.Find("identityfile")
	assert.Equal(t, o2, got)

	got = o.Find("foo")
	assert.Nil(t, got)
}
