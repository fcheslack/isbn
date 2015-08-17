package isbn

import (
	"testing"
)

type ChecksumTest struct {
	Input    string
	Expected bool
}

type InstringOutstringTest struct {
	Input  string
	Output string
}

type InstringOutintTest struct {
	Input  string
	Output int
}

func TestChecksum13digit(t *testing.T) {
	testResults := []InstringOutintTest{
		InstringOutintTest{"", -1},
		InstringOutintTest{"123", -1},
		InstringOutintTest{"9781937522414", 4},
		InstringOutintTest{"978193752241", 4},
	}

	for _, test := range testResults {
		if checksum13digit(test.Input) != test.Output {
			t.Errorf("Unexpected checksum13digit result for %s. Should be %d", test.Input, test.Output)
		}
	}
}

func TestChecksum10digit(t *testing.T) {
	testResults := []InstringOutintTest{
		InstringOutintTest{"", -1},
		InstringOutintTest{"123", -1},
		InstringOutintTest{"089791988-2", 2},
		InstringOutintTest{"089791988", 2},
		InstringOutintTest{"0306406157", 2},
	}

	for _, test := range testResults {
		if checksum10digit(test.Input) != test.Output {
			t.Errorf("Unexpected checksum result for %s. Should be %d but got %d", test.Input, test.Output, checksum10digit(test.Input))
		}
	}
}

func TestChecksum13(t *testing.T) {
	testResults := []ChecksumTest{
		ChecksumTest{"", false},
		ChecksumTest{"", false},
		ChecksumTest{"123", false},
		ChecksumTest{"9780306406157", true},
		ChecksumTest{"9781937522414", true},
		ChecksumTest{"978-1-937522-41", false},
		ChecksumTest{"9780306406151", false},
		ChecksumTest{"978-1-937522-41-2", false},
	}

	for _, test := range testResults {
		if checksum13(test.Input) != test.Expected {
			t.Errorf("Unexpected checksum result for %s. Should be %t", test.Input, test.Expected)
		}
	}
}

type NormalizeTest struct {
	Input  string
	Output string
	Error  bool
}

func TestNormalizeChecksum13(t *testing.T) {
	//for these inputs we expect Normalize to return an error and empty string
	expectNormalizeErrors := []NormalizeTest{
		NormalizeTest{"", "", true},
		NormalizeTest{"123", "", true},
		NormalizeTest{"1234567890123", "", true},
		NormalizeTest{"9780306406157", "9780306406157", false},
		NormalizeTest{"0306406157", "", true},
		NormalizeTest{"0-306-40-6157", "", true},
		NormalizeTest{"0-306-40615-2", "9780306406157", false},
		NormalizeTest{"0-89791-988-2", "9780897919883", false},
		NormalizeTest{"978-1-937522-41-4", "9781937522414", false},
		NormalizeTest{"0-201-83595-9", "9780201835953", false},
		NormalizeTest{"0-321-48681-1", "9780321486813", false},
		NormalizeTest{"0-465-03914-6", "9780465039142", false},
		NormalizeTest{"0-201-89683-4", "9780201896831", false},
	}

	for _, test := range expectNormalizeErrors {
		s, err := Normalize(test.Input)
		if s != test.Output {
			t.Errorf("Unexpected output string for input %s. Expected %s but got %s", test.Input, test.Output, s)
		}
		if err != nil && !test.Error {
			t.Errorf("Got error when expecting none: %s", err.Error())
		}
		if err == nil && test.Error {
			t.Error("Got nil for error when we expected an error")
		}
	}
}
