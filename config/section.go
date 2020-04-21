package config

// SectionType is the type of the section
type SectionType string

// Sections types
const (
	HostType  SectionType = "Host"
	MatchType SectionType = "Match"
)

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
