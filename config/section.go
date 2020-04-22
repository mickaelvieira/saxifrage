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

// Section --
type Section struct {
	Type     SectionType
	Matching string
	Configs  map[string]string
}

// NewSection --
func NewSection(t SectionType, hosts string) *Section {
	return &Section{
		Type:     t,
		Matching: hosts,
		Configs:  make(map[string]string),
	}
}
