package examples

import "testing"

func TestHelloLogger(t *testing.T) {
	tests := []struct {
		testName string
		num int
	}{
		{"info", 10},
		{"error", 20},
	}

	for _, test := range tests {
		t.Logf("testing: %s", test.testName)
		printLogMessage(test.num)
	}

	tests2 := []struct{
		testName string
		f func()
	}{
		{"withFields", printLogMessageWithFields},
		{"withHooks", printLogMessageWithHooks},
	}
	for _, test := range tests2 {
		name := test.testName
		t.Run(name, func(t *testing.T) {
			test.f()
		})
	}

	printLogMessageToFile("hello.log")
}

