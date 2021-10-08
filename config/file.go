package config

import (
	"errors"
)

// Files errors
var (
	ErrMissingUserConfig        = errors.New("unable to find user configuration file")
	ErrMissingSection           = errors.New("a Host or Match keyword was expected")
	ErrMissingIdentityFileValue = errors.New("section does not have IdentityFile value")
	ErrIsSSHBasedDirection      = errors.New("the key directory is the based directory")
	ErrDirectoryIsNotEmpty      = errors.New("the directory is not empty")
)

func contains(n []int, e int) bool {
	for _, a := range n {
		if a == e {
			return true
		}
	}
	return false
}

// Files list of ssh_config files
type Files []*File

// GetUserConfig retrieves the user configuration file
func (f Files) GetUserConfig() *File {
	for _, f := range f {
		if f.IsUserConfig() {
			return f
		}
	}
	return nil
}

// File a SSH configuration file
type File struct {
	Path  string
	Lines Lines
}

// RemoveLineNumbers removes the lines from file with the provided numbers
func (f *File) RemoveLineNumbers(n []int) {
	var lines Lines
	for _, l := range f.Lines {
		if !contains(n, l.Number) {
			lines = append(lines, l)
		}
	}
	f.Lines = lines
}

// CommentLineNumbers comments the lines from file with the provided numbers
func (f *File) CommentLineNumbers(n []int) {
	for _, l := range f.Lines {
		if contains(n, l.Number) {
			l.Comment()
		}
	}
}

// FindSectionLines finds the section line and its related configuration
func (f *File) FindSectionLines(s string) (lines Lines) {
	var found bool

	for _, l := range f.Lines {
		if l.IsSection() && l.HasValue(s) {
			found = true
		}
		if found {
			if l.IsSection() && !l.HasValue(s) {
				found = false
			} else if !l.IsComment() && !l.IsEmpty() {
				lines = append(lines, l)
			}
		}
	}

	return lines
}

func (f *File) buildKeywords() ([]*keyword, error) {
	var keywords []*keyword

	kw := &keyword{}
	for _, l := range f.Lines {
		for _, t := range l.tokens {
			if t.IsSection() ||
				t.IsKeyword() ||
				t.IsSeparator() ||
				t.IsValue() {
				if err := kw.add(t); err != nil {
					return keywords, err
				}
				if kw.isComplete() {
					keywords = append(keywords, kw)
					kw = &keyword{}
				}
			}
		}
	}
	return keywords, nil
}

// BuildSections builds the sections from the file's lines
func (f *File) BuildSections() (Sections, error) {
	var sections Sections
	var section *Section

	keywords, err := f.buildKeywords()
	if err != nil {
		return sections, err
	}

	for _, kw := range keywords {
		name := kw.name()
		separator := kw.separator()
		value := kw.value()

		if name.IsSection() {
			t := HostType
			if name.IsMatchSection() {
				t = MatchType
			}
			if section != nil {
				sections = append(sections, section)
			}
			section = NewSection(t, separator.Value, value.Value)
		} else {
			if section == nil {
				return sections, ErrMissingSection
			}
			section.Options = append(section.Options, NewOption(
				name.Value,
				separator.Value,
				value.Value,
			))
		}
	}

	if section != nil {
		sections = append(sections, section)
	}

	return sections, nil
}

// String returns the file content as a string
func (f *File) String() (s string) {
	for _, l := range f.Lines {
		s += l.String()
	}
	return s
}

// Bytes returns the file content as a slice of bytes
func (f *File) Bytes() (b []byte) {
	for _, l := range f.Lines {
		b = append(b, l.Bytes()...)
	}
	return b
}

// IsUserConfig identifies whether the file is the user local configuration
func (f *File) IsUserConfig() bool {
	return f.Path == GetUserConfigPath()
}
