package gook

import "testing"

func TestUrlize_Val(t *testing.T) {
	tests := []struct {
		name string
		u    Urlize
		want string
	}{
		{name: "simple", u: Urlize("Title of The Article"), want: "title-of-the-article"},
		{name: "normal", u: Urlize("Toto's hat is yellow"), want: "totos-hat-is-yellow"},
		{name: "medium", u: Urlize("We have 5 apples"), want: "we-have-5-apples"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Val(); got != tt.want {
				t.Errorf("Urlize.Val() = %v, want %v", got, tt.want)
			}
		})
	}
}
