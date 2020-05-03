package config

import (
	"errors"
	"fmt"
)

// Files errors
var (
	ErrMissingUserConfig = errors.New("Unable to find user configuration file")
)

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

// RemoveLinesWithNumbers removes the lines from file with the provided numbers
func (f *File) RemoveLinesWithNumbers(n []int) {
	var lines Lines

	contains := func(e int) bool {
		for _, a := range n {
			if a == e {
				return true
			}
		}
		return false
	}

	for _, l := range f.Lines {
		if !contains(l.Number) {
			lines = append(lines, l)
		}
	}

	f.Lines = lines
}

// FindSectionLines ...
func (f *File) FindSectionLines(s string) Lines {
	var lines Lines

	var found bool
	for _, l := range f.Lines {
		if l.IsSectionMatching(s) {
			found = true
		}
		if found {
			if l.IsSection() && !l.IsSectionMatching(s) {
				found = false
			} else if !l.IsComment() && !l.IsEmpty() {
				lines = append(lines, l)
			}
		}
	}

	return lines
}

// BuildSections builds the sections from the file's lines
func (f *File) BuildSections() (Sections, error) {

	var sections Sections

	var a []*keyValue

	kv := &keyValue{}
	for _, l := range f.Lines {
		for _, t := range l.tokens {
			if t.IsSection() ||
				t.IsKeyword() ||
				t.IsSeparator() ||
				t.IsValue() {
				if err := kv.add(t); err != nil {
					return sections, err
				}
				if kv.isComplete() {
					a = append(a, kv)
					kv = &keyValue{}
				}
			}
		}
	}

	var se *Section
	for _, kv := range a {
		k := kv.tokens[0] // keyword
		s := kv.tokens[1] // separator
		v := kv.tokens[2] // value

		if k.IsSection() {
			if se != nil {
				sections = append(sections, se)
			}
			t := HostType
			if k.IsMatchSection() {
				t = MatchType
			}
			se = NewSection(t, s.Value, v.Value)
		} else {
			if se == nil {
				return sections, fmt.Errorf("Section is null")
			}
			se.Options = append(se.Options, NewOption(k.Value, s.Value, v.Value))
		}
	}

	if se != nil {
		sections = append(sections, se)
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
