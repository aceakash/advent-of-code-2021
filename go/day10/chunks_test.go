package day10

import "testing"
import "github.com/stretchr/testify/assert"

func TestFindFirstWrongCloser(t *testing.T) {
	testData := []struct{
		name string
		input string
		want string
	}{
		{`empty input is valid`, ``, ``},
		{`() is valid`, `()`, ``},
		{`[] is valid`, `[]`, ``},
		{`<> is valid`, `<>`, ``},
		{`{} is valid`, `{}`, ``},
		{`(] is invalid coz of ]`, `(]`, `]`},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			assert.Equal(t, td.want, FindFirstWrongCloser(td.input))
		})
	}
}
