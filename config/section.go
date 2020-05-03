package config

import (
	"strings"
)

// SectionType is the type of the section
type SectionType string

// Sections types
const (
	HostType  SectionType = "Host"
	MatchType SectionType = "Match"
)

// Sections contains the section found in ssh_config files
type Sections []*Section

// GetMatchingValues ...
func (s Sections) GetMatchingValues() []string {
	values := make([]string, len(s))
	for k, v := range s {
		values[k] = v.Matching
	}
	return values
}

// FindAt returns section at the given index
func (s Sections) FindAt(i int) *Section {
	if i >= 0 && i < len(s) {
		return s[i]
	}
	return nil
}

// FindSectionByMatchingValue ...
func (s Sections) FindSectionByMatchingValue(value string) *Section {
	return s.Find(func(s *Section) bool {
		return s.Matching == value
	})
}

// Find returns the first section matching the predicate
func (s Sections) Find(predicate func(s *Section) bool) *Section {
	for _, section := range s {
		if predicate(section) {
			return section
		}
	}
	return nil
}

// FindIndex returns the index of the first section matching the predicate
func (s Sections) FindIndex(predicate func(s *Section) bool) int {
	for i, section := range s {
		if predicate(section) {
			return i
		}
	}
	return -1
}

// Filter the sections
func (s Sections) Filter(predicate func(s *Section) bool) Sections {
	sections := make(Sections, 0)
	for _, section := range s {
		if predicate(section) {
			sections = append(sections, section)
		}
	}
	return sections
}

// Section represents a section in a configuration file
// It can be a Host or Match section
type Section struct {
	Type      SectionType
	Matching  string
	Separator string
	Options   Options
}

// GetKeyFiles returns the private and public keys paths
// as well as their parent directory when the key is stored in a subdirectory
// @TODO move this func out of the struct and test it properly
func (s *Section) GetKeyFiles() []string {
	files := make([]string, 0)
	option := s.Options.Find("IdentityFile")

	if option == nil {
		return files
	}

	// @TODO there is a bug here
	// The value might be between quotes

	privateKey := ToAbsolutePath(option.Value)
	publickey := privateKey + ".pub"
	// directory := filepath.Dir(p)

	// if _, err := os.Stat(privateKey); err == nil {
	files = append(files, privateKey)
	// }
	// if _, err := os.Stat(privateKey); err == nil {
	files = append(files, publickey)
	// }

	// https://stackoverflow.com/questions/30697324/how-to-check-if-directory-on-path-is-empty
	// if !IsBaseSSHDirectory(directory) {
	// 	f = append(f, directory)
	// }

	return files
}

// Options is a list of options
type Options []*Option

// Find returns the first section matching the predicate
func (o Options) Find(n string) *Option {
	for _, option := range o {
		if strings.ToLower(n) == strings.ToLower(option.Name) {
			return option
		}
	}
	return nil
}

// Option represents a section's option
type Option struct {
	Name      string
	Value     string
	Separator string
}

// NewSection creates a new section
func NewSection(t SectionType, s string, v string) *Section {
	return &Section{
		Type:      t,
		Matching:  v,
		Separator: s,
	}
}

// NewOption creates a new option
func NewOption(n string, s string, v string) *Option {
	return &Option{
		Name:      n,
		Separator: s,
		Value:     v,
	}
}
