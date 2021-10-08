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

// GetIdentityFile returns the path the private key
func (s *Section) GetIdentityFile() (p string) {
	o := s.Options.FindByName("IdentityFile")
	if o != nil {
		p = o.GetUnquotedValue()
	}
	return p
}

// Options is a list of options
type Options []*Option

// FindByName returns the first section matching the predicate
func (o Options) FindByName(n string) *Option {
	for _, option := range o {
		if strings.EqualFold(n, option.Name) {
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

// GetUnquotedValue returns the value without leading and trailing quotes
func (o *Option) GetUnquotedValue() string {
	return strings.Trim(o.Value, "\"")
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
