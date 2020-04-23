package template

import "testing"

func TestStyler(t *testing.T) {
	t.Run("renders a single code", func(t *testing.T) {
		got := Styler(FGRed)("hi")
		want := "\033[31mhi\033[0m"
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("combines multiple codes", func(t *testing.T) {
		got := Styler(FGRed, FGBold)("hi")
		want := "\033[31;1mhi\033[0m"
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("should not repeat reset codes for nested styles", func(t *testing.T) {
		red := Styler(FGRed)("hi")
		got := Styler(FGBold)(red)
		want := "\033[1m\033[31mhi\033[0m"
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
