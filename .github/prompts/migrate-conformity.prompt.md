# migrate-conformity LLM prompt

You are an assistant that generates migration scripts for this repository.
Your goal is to add new migrations to the `tools/migrate-conformity` tool (like the custom rule migration).

## Project context
- Repo: `terraform-provider-conformity`
- Migration tool: `tools/migrate-conformity`
- The tool reads Terraform state (`.tfstate` or exported `state.json`) and generates Vision One HCL + import commands.

## What you should build
When asked to add a new migration (e.g., `conformity_custom_rule`), implement the following:

1. **State parsing**
	- Add a new file in `tools/migrate-conformity/` for the resource (example: `custom_rule.go`).
	- Define a model struct that represents the Conformity resource fields you need.
	- Add a `load<Resource>sFromState()` function that:
	  - Reads the state JSON.
	  - Filters `state.Resources` by `Type`.
	  - Extracts data from `Instances[].Attributes`.
	  - Returns a sorted slice by name or title for stable output.

2. **HCL generation**
	- Add an `append<Resource>HCL()` function that emits valid Vision One HCL blocks.
	- Ensure all strings are escaped (`escapeHCL`, `escapeHCLMultiline`).
	- Preserve deterministic ordering for nested blocks.

3. **Import lines**
	- Use `formatImportLine()` with IDs from state.
	- In dry-run, prefix import lines with `#`.

4. **Wire into the main flow**
	- Update `tools/migrate-conformity/main.go` to:
	  - Call the new `load*FromState()`.
	  - Include it in the "no resources" check.
	  - Append HCL and import lines to output.

5. **Tests + fixtures**
	- Add a fixture under `tools/migrate-conformity/testdata/`.
	- Add unit tests in `tools/migrate-conformity/*_test.go` for:
	  - State parsing.
	  - HCL generation.
	  - Any value formatting helpers.

6. **Docs**
	- Update `docs/tools/migrate-conformity.md` to list the new resource and mapping rules.

## Mapping guidance
When mapping Conformity -> Vision One:
- Prefer direct field mapping when semantics match.
- If there is no Vision One equivalent, omit and mention it in docs.
- Keep transformations explicit and deterministic.
- For JSON-like string fields, preserve booleans, numbers, nulls, and JSON objects where possible.

## Output expectations
- Use ASCII only.
- HCL must be valid and ready to paste into `main.tf`.
- Sorting and spacing should match existing style.
