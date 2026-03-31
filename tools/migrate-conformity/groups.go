package main

import "fmt"

func appendGroupsHCL(lines *[]string, importLines *[]string, mappingLines *[]string, groups []groupItem, dryRun bool) {
	if len(groups) == 0 {
		return
	}

	*lines = append(*lines, "# Groups")
	nameCounter := map[string]int{}
	for _, item := range groups {
		resourceName := item.ResourceName
		if resourceName == "" {
			resourceName = toTerraformName(item.Name, nameCounter, "group")
		} else {
			resourceName = uniqueResourceName(resourceName, nameCounter)
		}

		*lines = append(*lines, fmt.Sprintf("resource \"visionone_crm_group\" \"%s\" {", resourceName))
		*lines = append(*lines, fmt.Sprintf("  name = \"%s\"", escapeHCL(item.Name)))
		if len(item.Tags) > 0 {
			*lines = append(*lines, fmt.Sprintf("  tags = [%s]", formatTags(item.Tags)))
		} else if item.TagsSet {
			*lines = append(*lines, "  tags = []")
		}
		*lines = append(*lines, "}")
		*lines = append(*lines, "")

		if mappingLines != nil {
			sourceName := item.ResourceName
			if sourceName == "" {
				sourceName = resourceName
			}
			*mappingLines = append(*mappingLines, formatMappingLine("conformity_group", sourceName, "visionone_crm_group", resourceName))
			*mappingLines = append(*mappingLines, formatAttributeMappingLine("name", "name"))
			if len(item.Tags) > 0 || item.TagsSet {
				*mappingLines = append(*mappingLines, formatAttributeMappingLine("tags", "tags"))
			}
		}

		*importLines = append(*importLines, formatImportLine("visionone_crm_group", resourceName, item.ID, dryRun))
	}
}
