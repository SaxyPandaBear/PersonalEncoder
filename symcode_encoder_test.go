package main

import "testing"

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func assertTrue(t *testing.T, value bool) {
	if !value {
		t.Fatal()
	}
}

func assertFalse(t *testing.T, value bool) {
	if value {
		t.Fatal()
	}
}

func TestCharacterEncoding(t *testing.T) {
	cipher := "symantec"
	for i := 0; i < len(cipher); i++ {
		expected := string(cipher[i])
		actual := EncodeChar(i)
		assertEquals(t, expected, actual)
	}
}

func TestNoExtraPadding(t *testing.T) {
	s := "111111111111111111111111" // should be len 24
	length := len(s)
	s = PadBinaryString(s)
	assertTrue(t, length == len(s))
}

func TestPaddingRequired(t *testing.T) {
	s := "1"
	actual := PadBinaryString(s)
	assertTrue(t, len(actual) > len(s))

	expected := "100000000000000000000000"
	assertEquals(t, expected, actual)
}

func TestGetPadNoPadding(t *testing.T) {
	s := "111111111111111111111111" // should be len 24
	length := len(s)
	assertTrue(t, length == GetPad(s))
}

func TestGetPadWithPadding(t *testing.T) {
	s := "1111111111111111111111" // len(s) = 22
	assertEquals(t, 24, GetPad(s))
}

func TestEncodeSimpleString(t *testing.T) {
	s := "gimble"
	expected := "ayeentttasneeynt"

	assertEquals(t, expected, EncodeString(s))
}

func TestEncodeStringWithWhiteSpace(t *testing.T) {
	s := "\t\n   wabe \n\n  \t"
	expected := "ateestnmaym$$$$$"

	assertEquals(t, expected, EncodeString(s))
}
