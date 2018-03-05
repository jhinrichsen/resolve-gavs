package main

import "testing"

func TestEmptyGav(t *testing.T) {
	want := Gav{"", "", "", "", ""}
	got := Concise("")
	if want != got {
		t.Fatalf("want %s but got %s\n", want, got)
	}
}

func TestGav(t *testing.T) {
	want := Gav{"g", "a", "v", "", ""}
	got := Concise("g:a:v")
	if want != got {
		t.Fatalf("want %s but got %s\n", want, got)
	}
}

func TestGroupOnlyGav(t *testing.T) {
	want := Gav{"g", "", "", "", ""}
	got := Concise("g")
	if want != got {
		t.Fatalf("want %s but got %s\n", want, got)
	}
}

func TestArtifactOnlyGav(t *testing.T) {
	want := Gav{"", "a", "", "", ""}
	got := Concise(":a")
	if want != got {
		t.Fatalf("want %s but got %s\n", want, got)
	}
}

func TestPackagingOnlyGav(t *testing.T) {
	want := Gav{"", "", "", "", "jar"}
	got := Concise("@jar")
	if want != got {
		t.Fatalf("want %s but got %s\n", want, got)
	}
}

func TestMatch(t *testing.T) {
	want := true
	got := Gav{"g", "", "", "", ""}.Includes(Gav{"g", "a", "1.0", "", ""})
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestConciseGav(t *testing.T) {
	want := "g:a:v"
	got := Gav{"g", "a", "v", "", ""}.Concise()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestConciseGroup(t *testing.T) {
	want := "g"
	got := Gav{"g", "", "", "", ""}.Concise()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestConciseArtifact(t *testing.T) {
	want := ":a"
	got := Gav{"", "a", "", "", ""}.Concise()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestConciseGroupVersion(t *testing.T) {
	want := "g::v"
	got := Gav{"g", "", "v", "", ""}.Concise()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestConciseClassifier(t *testing.T) {
	want := ":::d"
	got := Gav{"", "", "", "d", ""}.Concise()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestConcisePackaging(t *testing.T) {
	want := "@ear"
	got := Gav{"", "", "", "", "ear"}.Concise()
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
