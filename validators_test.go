package do_validator

import "testing"

func TestLuhn(t *testing.T) {
	t.Run("Only with string parameter returns boolean", func(t *testing.T) {
		type test struct {
			input string
			want  bool
		}

		tests := []test{
			{input: "79927398713", want: true},
			{input: "79927398711", want: false},
			{input: "22300878232", want: true},
			{input: "something", want: false},
		}
		for _, te := range tests {
			t.Run("Lunhs -> with value -> "+te.input, func(t *testing.T) {
				if got := Lunh(te.input); got != te.want {
					t.Errorf("Lunh(%s) = %v, didn't return %v", te.input, got, te.want)
				}
			})
		}
	})
}

func TestIsCedula(t *testing.T) {
	type args struct {
		number string
	}
	type test struct {
		name  string
		input string
		want  bool
	}
	tests := []test{
		{name: "valid with slash", input: "223-0087823-2", want: true},
		{name: "invalid with slash", input: "223-0087823-1", want: false},
		{name: "less than 11 characters", input: "223-0087823", want: false},
		{name: "valid without slash", input: "22300878232", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCedula(tt.input); got != tt.want {
				t.Errorf("IsCedula() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRncLunh(t *testing.T) {
	type test struct {
		name  string
		input string
		want  bool
	}
	tests := []test{
		{name: "valid", input: "130508844", want: true},
		{name: "invalid", input: "131497911", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RncLunh(tt.input); got != tt.want {
				t.Errorf("RncLunh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRnc(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "valid", input: "130-50884-4", want: true},
		{name: "invalid", input: "131-49791-1", want: false},
		{name: "invalid with length 8", input: "131-49791", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRnc(tt.input); got != tt.want {
				t.Errorf("IsRnc() = %v, want %v", got, tt.want)
			}
		})
	}
}
