package vision

import "testing"

func TestParseString(t *testing.T) {
	text := "Hello World"
	str, err := ParseString(text)
	if err != nil {
		t.Fatalf("Unable to parse text result. Error: %v", err.Error())
	}
	if str.String() != text {
		t.Fatalf("Unexpected string result. Found %v, expected %v", str, text)
	}
}

func TestStringValue(t *testing.T) {
	expected := "Hello World"
	text := String(expected)
	if text.String() != expected {
		t.Fatalf("Unexpected string result. Found %v, expected %v", text, expected)
	}
}

func TestStringToInteger(t *testing.T) {
	num := 1
	str, err := ParseString(num)
	if err != nil {
		t.Fatalf("Unable to parse integer value. Error: %v", err.Error())
	}
	if int(str.Integer()) != num {
		t.Fatalf("Unexpected string result. Found %d, expected %d", str.Integer(), num)
	}
}

func TestStringToFloat(t *testing.T) {
	num := 1.2
	str, err := ParseString(num)
	if err != nil {
		t.Fatalf("Unable to parse float value. Error: %v", err.Error())
	}
	if float64(str.Float()) != num {
		t.Fatalf("Unexpected string result. Found %f, expected %f", str.Float(), num)
	}
}

func TestStringToBoolTrue(t *testing.T) {
	res := true
	str, err := ParseString(res)
	if err != nil {
		t.Fatalf("Unable to parse bool value. Error: %v", err.Error())
	}
	if bool(str.Bool()) != res {
		t.Fatalf("Unexpected string result. Found %v, expected %v", str.Bool(), res)
	}
}

func TestStringToBoolFalse(t *testing.T) {
	res := false
	str, err := ParseString(res)
	if err != nil {
		t.Fatalf("Unable to parse bool value. Error: %v", err.Error())
	}
	if bool(str.Bool()) != res {
		t.Fatalf("Unexpected string result. Found %v, expected %v", str.Bool(), res)
	}
}

func TestStringIsEmpty(t *testing.T) {
	text := String("")
	if text.IsEmpty() != true {
		t.Fatalf("Unexpected result to be empty. Found %v, expected %v", text.IsEmpty(), true)
	}
}

func TestWhitespaceStringIsEmpty(t *testing.T) {
	text := String(" ")
	if text.IsEmpty() {
		t.Fatalf("Unexpected result not to be empty. Found %v, expected %v", text.IsEmpty(), false)
	}
}

func TestStringMustEmpty(t *testing.T) {
	text := String(" ")
	if !text.MustEmpty() {
		t.Fatalf("Unexpected result to be empty. Found %v, expected %v", text.MustEmpty(), true)
	}
}

func TestStringIsLength(t *testing.T) {
	text := String("")
	if !text.IsLength(0, 1) {
		t.Fatalf("Unexpected result to be true. Found %v, expected %v", text.IsLength(0, 1), true)
	}
}

func TestWhitespaceStringIsLength(t *testing.T) {
	text := String(" ")
	if !text.IsLength(0, 1) {
		t.Fatalf("Unexpected result to be false. Found %v, expected %v", text.IsLength(0, 1), false)
	}
}

func TestWhitespaceStringMustLength(t *testing.T) {
	text := String(" /.")
	if !text.MustLength(0, 0) {
		t.Fatalf("Unexpected result to be true. Found %v, expected %v", text.MustLength(0, 0), true)
	}
}

func TestStringIsEmail(t *testing.T) {
	tests := map[string]bool{
		"localhost":           false,
		"127.0.0.1":           false,
		"test@localhost":      false,
		"awesome@example.com": true,
		"test@gmail.com":      true,
	}
	for v := range tests {
		expected := tests[v]
		if valid := String(v).IsEmail(); valid != expected {
			t.Fatalf("Unexpected result of checking valid email address. Found %v, expected %v", valid, expected)
		}
	}
}

func TestStringIsURL(t *testing.T) {
	tests := map[string]bool{
		"http://google.ca":    true,
		"127.0.0.1":           true,
		"gmail.com":           true,
		"https://example.com": true,
		"ftp://google.ca":     false,
		"localhost":           true,
	}
	for v := range tests {
		expected := tests[v]
		if valid := String(v).IsURL(); valid != expected {
			t.Fatalf("Unexpected result of checking valid URL %v. Found %v, expected %v", v, valid, expected)
		}
	}
}
