package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadGroupsFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "groups_state.json")

	groups, err := loadGroupsFromState(statePath)
	if err != nil {
		t.Fatalf("loadGroupsFromState error: %v", err)
	}
	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}

	if groups[0].Name != "Alpha" || groups[0].ID != "group-1" || groups[0].ResourceName != "alpha" {
		t.Fatalf("unexpected first group: %+v", groups[0])
	}
	if groups[1].Name != "Beta" || groups[1].ID != "group-2" || groups[1].ResourceName != "beta" {
		t.Fatalf("unexpected second group: %+v", groups[1])
	}
}

func TestAppendGroupsHCL(t *testing.T) {
	groups := []groupItem{
		{ID: "group-1", Name: "Alpha", ResourceName: "alpha", Tags: []string{"prod", "core"}},
		{ID: "group-2", Name: "Beta", ResourceName: "beta", TagsSet: true},
	}

	var hclLines []string
	var importLines []string
	appendGroupsHCL(&hclLines, &importLines, nil, groups, true)

	expectedHCL := strings.Join([]string{
		"# Groups",
		"resource \"visionone_crm_group\" \"alpha\" {",
		"  name = \"Alpha\"",
		"  tags = [\"prod\", \"core\"]",
		"}",
		"",
		"resource \"visionone_crm_group\" \"beta\" {",
		"  name = \"Beta\"",
		"  tags = []",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(hclLines, "\n") + "\n"; got != expectedHCL {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}

	expectedImports := []string{
		"terraform import visionone_crm_group.alpha group-1",
		"terraform import visionone_crm_group.beta group-2",
	}
	if len(importLines) != len(expectedImports) {
		t.Fatalf("expected %d import lines, got %d", len(expectedImports), len(importLines))
	}
	for i, line := range expectedImports {
		if importLines[i] != line {
			t.Fatalf("unexpected import line %d: %s", i, importLines[i])
		}
	}
}

func TestFormatImportLineMissingID(t *testing.T) {
	line := formatImportLine("visionone_crm_group", "alpha", "", true)
	if line != "terraform import visionone_crm_group.alpha <resource_id>" {
		t.Fatalf("unexpected import line: %s", line)
	}
}

func TestWriteMainTFOverwritesExistingFile(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "main.tf")
	if err := os.WriteFile(path, []byte("# existing\n"), 0o644); err != nil {
		t.Fatalf("write main.tf: %v", err)
	}

	gotPath, err := writeMainTF(tempDir, []string{"# new"})
	if err != nil {
		t.Fatalf("writeMainTF error: %v", err)
	}
	if gotPath != path {
		t.Fatalf("unexpected output path: %s", gotPath)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read main.tf: %v", err)
	}
	if string(content) != "# new\n" {
		t.Fatalf("main.tf content mismatch: %q", string(content))
	}
}
