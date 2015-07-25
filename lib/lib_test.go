package lib

import (
	"os"
	"testing"
)

var apps Apps = Apps{
	"Pkg1": {"Name1"},
	"Pkg2": {"Name2"},
}
var yamlApps string = `---
Pkg1:
  label: Name1
Pkg2:
  label: Name2
`

func TestDumpYaml(t *testing.T) {
	yaml, err := apps.toYaml()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if yaml != yamlApps {
		t.Fatalf("error: yaml doesn't match\nGot:\n%v\nExpected:\n%v\n", yaml, yamlApps)
	}
}

func TestFromYaml(t *testing.T) {
	yaml, err := apps.toYaml()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if yaml != yamlApps {
		t.Fatalf("error: yaml doesn't match\nGot:\n%v\nExpected:\n%v\n", yaml, yamlApps)
	}
}

func TestYamlRoundtrip(t *testing.T) {
	tmpfile := ".tmpfile"
	defer os.Remove(tmpfile)

	err := apps.DumpYaml(tmpfile)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	newApps, err := LoadYaml(tmpfile)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	newYaml, err := newApps.toYaml()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if newYaml != yamlApps {
		t.Fatalf("error: yaml doesn't match\nGot:\n%v\nExpected:\n%v\n", newYaml, yamlApps)
	}

}

func TestGetUserApps(t *testing.T) {
	// TODO
	t.Skip()
}

func TestBackupUserApps(t *testing.T) {
	// TODO
	t.Skip()
}
