package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadApplyProfilesFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "apply_profile_state.json")

	items, err := loadApplyProfilesFromState(statePath)
	if err != nil {
		t.Fatalf("loadApplyProfilesFromState error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 apply profile, got %d", len(items))
	}

	item := items[0]
	if item.ProfileID != "profile-1" || item.ResourceName != "apply_profile" {
		t.Fatalf("unexpected apply profile: %+v", item)
	}
	if len(item.AccountIDs) != 2 || item.AccountIDs[0] != "account-1" {
		t.Fatalf("unexpected account IDs: %+v", item.AccountIDs)
	}
	if item.Include == nil || item.Include.Exceptions == nil || *item.Include.Exceptions != true {
		t.Fatalf("unexpected include: %+v", item.Include)
	}
}

func TestAppendApplyProfileHCL(t *testing.T) {
	item := applyProfileItem{
		ProfileID:  "profile-1",
		AccountIDs: []string{"account-1", "account-2"},
		Mode:       "overwrite",
		Notes:      "apply profile",
		Include: &applyProfileIncludeItem{
			Exceptions: boolPtr(true),
		},
	}

	var lines []string
	appendApplyProfileHCL(&lines, item, "apply_profile")

	expected := strings.Join([]string{
		"data \"visionone_crm_apply_profile\" \"apply_profile\" {",
		"  profile_id = \"profile-1\"",
		"  account_ids = [\"account-1\", \"account-2\"]",
		"  mode = \"overwrite\"",
		"  notes = \"apply profile\"",
		"  include = {",
		"    exceptions = true",
		"  }",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(lines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}
