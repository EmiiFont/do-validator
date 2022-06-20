package do_validator

import "testing"

func TestLuhn(t *testing.T) {
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
}
