package middlestring

import "testing"

func TestGet(t *testing.T) {
	r := Get("(hello)(world)", "(", ")")
	if r != "hello" {
		t.Fatal(r)
	}

	t.Log(r)
}

func TestGetMore(t *testing.T) {
	r := Get("(hello)(world)", "(", ")", true)
	if r != "hello)(world" {
		t.Fatal(r)
	}

	t.Log(r)
}

func TestGetEmpty(t *testing.T) {
	r := Get("(hello)(world)", "[", "]")
	if r != "" {
		t.Fatal(r)
	}

	t.Log(r)
}
