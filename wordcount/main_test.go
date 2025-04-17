package main

import (
	"testing"
)

func TestCountLines(t *testing.T) {
	text := `The colossal squid shares features common to all squids: a mantle for locomotion, one pair of gills, a beak or tooth, and certain external characteristics like eight arms and two tentacles, a head, and two fins.
In general, the morphology and anatomy of the colossal squid are the same as any other squid.However, there are certain morphological characteristics that separate the colossal squid from other squids in its family: the colossal squid is the only squid in its family whose arms and tentacles are equipped with hooks, either swiveling or three-pointed.
There are squids in other families that also have hooks, but no other squid in the family Cranchiidae.
`
	expected := 3
	got := countLines(text)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCountWords(t *testing.T) {
	text := `The colossal squid shares features common to all squids: a mantle for locomotion, one pair of gills, a beak or tooth, and certain external characteristics like eight arms and two tentacles, a head, and two fins.
In general, the morphology and anatomy of the colossal squid are the same as any other squid.However, there are certain morphological characteristics that separate the colossal squid from other squids in its family: the colossal squid is the only squid in its family whose arms and tentacles are equipped with hooks, either swiveling or three-pointed.
There are squids in other families that also have hooks, but no other squid in the family Cranchiidae.
`
	expected := 109
	got := countWords(text)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCountCharacters(t *testing.T) {
	text := `The colossal squid shares features common to all squids: a mantle for locomotion, one pair of gills, a beak or tooth, and certain external characteristics like eight arms and two tentacles, a head, and two fins.
In general, the morphology and anatomy of the colossal squid are the same as any other squid.However, there are certain morphological characteristics that separate the colossal squid from other squids in its family: the colossal squid is the only squid in its family whose arms and tentacles are equipped with hooks, either swiveling or three-pointed.
There are squids in other families that also have hooks, but no other squid in the family Cranchiidae.
`
	expected := 667
	got := countCharacters(text)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCheckIfFileExists(t *testing.T) {
	filename := "text"
	expected := true
	got := checkIfFileExists(filename)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCheckIfFileNotExists(t *testing.T) {
	filename := "test-file.txt"
	expected := false
	got := checkIfFileExists(filename)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCheckIfDirOrFile(t *testing.T) {
	filename := "test-dir"
	expected := false
	got := checkIfFileOrDir(filename)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCheckIfFileOrDir(t *testing.T) {
	filename := "test-file"
	expected := true
	got := checkIfFileOrDir(filename)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestCheckFilePermissions(t *testing.T) {
	filename := "protected-file"
	expected := false
	got := checkFilePermissions(filename)

	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
