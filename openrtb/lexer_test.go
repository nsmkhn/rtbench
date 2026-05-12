package openrtb

import (
	"testing"
)

func TestLexer_Helpers(t *testing.T) {
	input := `{"id":"req-001","at":2,"w":300}`
	l := newLexer([]byte(input))

	if err := l.skipOpen(); err != nil {
		t.Fatalf("skipOpen: %v", err)
	}

	// "id": "req-001"
	key, err := l.readKey()
	if err != nil {
		t.Fatalf("readKey: %v", err)
	}
	if string(key) != "id" {
		t.Fatalf(`expected key "id", got %q`, string(key))
	}
	if err = l.skipColon(); err != nil {
		t.Fatalf("skipColon: %v", err)
	}
	val, err := l.readStringVal()
	if err != nil {
		t.Fatalf("readStringVal: %v", err)
	}
	if val != "req-001" {
		t.Fatalf(`expected "req-001", got %q`, val)
	}
	done, err := l.readSep('}')
	if err != nil {
		t.Fatalf("readSep: %v", err)
	}
	if done {
		t.Fatalf("expected more fields, got closing brace")
	}

	// "at": 2
	key, err = l.readKey()
	if err != nil {
		t.Fatalf("readKey: %v", err)
	}
	if string(key) != "at" {
		t.Fatalf(`expected key "at", got %q`, string(key))
	}
	if err = l.skipColon(); err != nil {
		t.Fatalf("skipColon: %v", err)
	}
	num, err := l.readNumberBytes()
	if err != nil {
		t.Fatalf("readNumberBytes: %v", err)
	}
	if string(num) != "2" {
		t.Fatalf(`expected "2", got %q`, string(num))
	}
	done, err = l.readSep('}')
	if err != nil {
		t.Fatalf("readSep: %v", err)
	}
	if done {
		t.Fatalf("expected more fields, got closing brace")
	}

	// "w": 300
	key, err = l.readKey()
	if err != nil {
		t.Fatalf("readKey: %v", err)
	}
	if string(key) != "w" {
		t.Fatalf(`expected key "w", got %q`, string(key))
	}
	if err = l.skipColon(); err != nil {
		t.Fatalf("skipColon: %v", err)
	}
	num, err = l.readNumberBytes()
	if err != nil {
		t.Fatalf("readNumberBytes: %v", err)
	}
	if string(num) != "300" {
		t.Fatalf(`expected "300", got %q`, string(num))
	}
	done, err = l.readSep('}')
	if err != nil {
		t.Fatalf("readSep: %v", err)
	}
	if !done {
		t.Fatalf("expected closing brace")
	}
}

func TestLexer_ScanRaw(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"string", `"hello"`, `"hello"`},
		{"number", `42`, `42`},
		{"float", `3.14`, `3.14`},
		{"true", `true`, `true`},
		{"false", `false`, `false`},
		{"null", `null`, `null`},
		{"object", `{"a":1}`, `{"a":1}`},
		{"array", `[1,2,3]`, `[1,2,3]`},
		{"nested", `{"a":{"b":2}}`, `{"a":{"b":2}}`},
		{"string with escape", `"he\"llo"`, `"he\"llo"`},
		{"object with escaped string val", `{"a":"b\"c"}`, `{"a":"b\"c"}`},
		{"object with brace in string", `{"a":"x}y"}`, `{"a":"x}y"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLexer([]byte(tt.input))
			got, err := l.scanRaw()
			if err != nil {
				t.Fatalf("scanRaw: %v", err)
			}
			if string(got) != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, string(got))
			}
		})
	}
}

func TestLexer_ReadSep_Array(t *testing.T) {
	// Test readSep with ']' close
	input := `[1,2]`
	l := newLexer([]byte(input))

	if err := l.skipOpen(); err != nil {
		t.Fatalf("skipOpen: %v", err)
	}

	for i, want := range []string{"1", "2"} {
		num, err := l.readNumberBytes()
		if err != nil {
			t.Fatalf("[%d] readNumberBytes: %v", i, err)
		}
		if string(num) != want {
			t.Fatalf("[%d] expected %q, got %q", i, want, string(num))
		}
		done, err := l.readSep(']')
		if err != nil {
			t.Fatalf("[%d] readSep: %v", i, err)
		}
		if i < 1 && done {
			t.Fatalf("[%d] unexpected closing bracket", i)
		}
		if i == 1 && !done {
			t.Fatalf("[%d] expected closing bracket", i)
		}
	}
}
