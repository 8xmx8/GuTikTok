package checks

import "testing"

/*
=== RUN   TestIsValidString
=== RUN   TestIsValidString/qwertyuuio
=== RUN   TestIsValidString/123456789
=== RUN   TestIsValidString/asdqwe5451232
=== RUN   TestIsValidString/1
=== RUN   TestIsValidString/#00
=== RUN   TestIsValidString/1564165....
=== RUN   TestIsValidString/!!!!....
=== RUN   TestIsValidString/@@@@@@aaa...
=== RUN   TestIsValidString/###%%%%...
=== RUN   TestIsValidString/+-()==
=== RUN   TestIsValidString/]][[
=== RUN   TestIsValidString/}}{{
=== RUN   TestIsValidString/....
=== RUN   TestIsValidString/++++
=== RUN   TestIsValidString/----
=== RUN   TestIsValidString/~~~~
=== RUN   TestIsValidString/~~~
=== RUN   TestIsValidString/:::”'""|||\\///
--- PASS: TestIsValidString (0.00s)

	--- PASS: TestIsValidString/qwertyuuio (0.00s)
	--- PASS: TestIsValidString/123456789 (0.00s)
	--- PASS: TestIsValidString/asdqwe5451232 (0.00s)
	--- PASS: TestIsValidString/1 (0.00s)
	--- PASS: TestIsValidString/#00 (0.00s)
	--- PASS: TestIsValidString/1564165.... (0.00s)
	--- PASS: TestIsValidString/!!!!.... (0.00s)

	--- PASS: TestIsValidString/@@@@@@aaa... (0.00s)
	--- PASS: TestIsValidString/###%%%%... (0.00s)
	--- PASS: TestIsValidString/+-()== (0.00s)
	--- PASS: TestIsValidString/]][[ (0.00s)
	--- PASS: TestIsValidString/}}{{ (0.00s)
	--- PASS: TestIsValidString/.... (0.00s)
	--- PASS: TestIsValidString/++++ (0.00s)
	--- PASS: TestIsValidString/---- (0.00s)
	--- PASS: TestIsValidString/~~~~ (0.00s)
	--- PASS: TestIsValidString/~~~ (0.00s)
	--- PASS: TestIsValidString/:::'''""|||\\/// (0.00s)

# PASS

Process finished with the exit code 0
*/
func TestIsValidString(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"qwertyuuio", true},
		{"123456789", true},
		{"asdqwe5451232", true},
		{"1", true},
		{"", false},
		{"1564165....", true},
		{"!!!!....", true},
		{"@@@@@@aaa...", true},
		{"###%%%%...", true},
		{"+-()==", false},
		{"]][[", false},
		{"}}{{", false},
		{"....", true},
		{"++++", false},
		{"----", false},
		{"~~~~", false},
		{"~~~", false},
		{":::'''\"\"|||\\\\///", false},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := isValidString(tt.args); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
