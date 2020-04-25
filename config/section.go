package config

// SectionType is the type of the section
type SectionType string

// Sections types
const (
	HostType  SectionType = "Host"
	MatchType SectionType = "Match"
)

// Sections contains the section found in ssh_config files
type Sections []*Section

// Filter the sections
func (s Sections) Filter() []*Section {
	panic("not implemented")
}

// Section represents a section in a configuration file
// It can be a Host or Match section
type Section struct {
	Type      SectionType
	Matching  string
	Separator string
	Options   []*Option
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
