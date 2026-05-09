package openrtb

import (
	"os"
	"testing"
)

func TestLexer_ValidJSON(t *testing.T) {
	data, err := os.ReadFile("../testdata/valid_banner.json")
	if err != nil {
		t.Fatalf("could not read testdata/valid_banner.json: %v", err)
	}

	l := newLexer(data)

	tok, err := l.next()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tok.kind != tokLBrace {
		t.Fatalf("expected token tokLBrace, got %v", tok.kind)
	}

	tok, err = l.next()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tok.kind != tokString || string(tok.val) != "id" {
		t.Fatalf("expected token tokString with val \"id\", got kind=%v, val=%q", tok.kind, tok.val)
	}

	tok, err = l.next()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tok.kind != tokColon {
		t.Fatalf("expected token tokColon, got %v", tok.kind)
	}

	tok, err = l.next()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tok.kind != tokString || string(tok.val) != "req-001" {
		t.Fatalf("expected token tokString with val \"req-001\", got kind=%v, val=%q", tok.kind, tok.val)
	}

	for {
		tok, err = l.next()
		if err != nil {
			t.Fatalf("%v", err)
		}
		if tok.kind == tokEOF {
			t.Fatalf("reached EOF before finding \"w\" token")
		}
		if tok.kind == tokString && string(tok.val) == "w" {
			break
		}
	}

	tok, err = l.next()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tok.kind != tokColon {
		t.Fatalf("expected token tokColon, got %v", tok.kind)
	}

	tok, err = l.next()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tok.kind != tokNumber || string(tok.val) != "300" {
		t.Fatalf("expected token tokNumber with val 300, got kind=%v, val=%q", tok.kind, tok.val)
	}
}
