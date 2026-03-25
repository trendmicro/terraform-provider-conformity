# Migration Steps

## 0. Prerequisites

Ensure the following dependencies are installed before running the migration:

- Go 1.21 or later (required for `go run`).
- Terraform CLI available in `PATH`.
- Access to the `migrate-conformity` tool in this repository.

## 1. Prepare the workspace

```bash
cd /path/to/your/terraform/workspace

# Export state first when using a remote backend
terraform state pull > state.json

# IMPORTANT: run terraform refresh to ensure state is up to date before migration
terraform refresh

# Backup state and configuration
cp terraform.tfstate terraform.tfstate.backup
cp main.tf main.tf.backup
```

**Do not skip `terraform refresh`.** The migration tool reads the Terraform state as-is, and stale state can map rule settings to the wrong rules.

Confirm the state file exists locally before proceeding. If the workspace uses a remote backend, use the exported `state.json` file as input.

## 2. Choose an output location

Run the tool from the directory where you want `main.tf` written.

## 3. Generate migrated resources

```bash
go run ./tools/migrate-conformity -state terraform.tfstate
```

## 4. Understand tool arguments

- `-state` (required): Path to a Terraform state JSON file (for example, local `.tfstate` or exported `state.json`).
- The tool always writes `main.tf` to the current working directory.

## 5. Review generated output

The tool generates `main.tf` containing:

- Vision One provider configuration and variables.
- Migrated Vision One resource blocks.
- Mapping comments that describe Conformity-to-Vision One resource translation.
- `terraform import` commands in comments, placed after the mapping section.

Validate resource names, arguments, and comments before proceeding to further steps and merging them into your target workspace.

For schema and attribute parity checks, review the corresponding [Vision One provider Cloud Risk Management resource and data source documentation](https://registry.terraform.io/providers/trendmicro/vision-one/latest/docs).

If account IDs are required during manual adjustments, use [`visionone_crm_account` data source](https://registry.terraform.io/providers/trendmicro/vision-one/latest/docs/data-sources/crm_account) lookups.

## 6. Remove Conformity resources from state and code

```bash
cd /path/to/your/terraform/workspace

# Inspect resources before removal
terraform state list | grep conformity_

# Remove Conformity resources from state only
terraform state rm $(terraform state list | grep conformity_)
```

Remove `conformity_*` resource blocks from `.tf` files so they are not recreated on subsequent plans.

> ⚠️ `terraform state rm` detaches resources from Terraform state only; it does not delete cloud resources.

## 7. Merge generated configuration

Merge generated Vision One resources from `main.tf` into your Terraform configuration, including provider configuration and required variables.

Ensure the target project contains the Vision One provider configuration in `terraform.required_providers` and `provider "visionone"`. See [the provider documentation](https://registry.terraform.io/providers/trendmicro/vision-one/latest/docs) for more details.

## 8. Initialize provider and import resources

```bash
terraform init
```

Run the generated `terraform import` commands from `main.tf` to bind Vision One resources to existing cloud objects.

## 9. Validate migration

```bash
terraform validate
terraform plan
```

Expected result: `No changes. Your infrastructure matches the configuration.`

If changes are detected, review the diffs and adjust fields that are not directly translatable between providers.
Diffs may also result from computed fields, sensitive fields, or list ordering.

## 10. Apply changes (if required)

```bash
terraform plan
terraform apply
```

Apply only after validating that planned changes are expected and safe.

## 11. Post-migration cleanup

After all resources are consolidated into the original project and validation is complete:

- Remove temporary migration files/directories created only for generation.
- Keep backups until at least one successful plan/apply cycle is completed.

## Rollback procedure

```bash
cp terraform.tfstate.backup terraform.tfstate
cp main.tf.backup main.tf
terraform plan
```

## Troubleshooting

### Import command fails

- Error: `Cannot import non-existent remote object`
- Cause: Source state resource ID does not match the Vision One import ID format.
- Resolution: Verify the target resource ID from the Vision One API and rerun the import command.

### Plan shows extensive drift

- Cause: Some provider attributes do not have a one-to-one mapping.
- Resolution:
	1. Review warnings and mapping comments in generated `main.tf`.
	2. Compare unmapped or transformed attributes.
	3. Add or adjust fields manually in Terraform configuration.

### Rule settings mapped to the wrong rules

- Cause: State was not refreshed before running the migration.
- Resolution: Run `terraform refresh`, then rerun the migration tool.

### Provider not initialized

- Error: `Provider "registry.terraform.io/trendmicro/vision-one" not available`
- Resolution: Run `terraform init` in the workspace.

## Best practices

1. Validate in a non-production workspace first.
2. Keep state and configuration backups until migration is complete.
3. Review generated warnings before running imports.
4. Run `terraform plan` immediately after imports to validate drift.

## Limitations

- Complex state dependencies can require manual import ordering.
- Provider-specific features may require manual adaptation.

## Support

- Review generated warnings and mapping comments in `main.tf`.
- Consult Vision One provider documentation for target schema details.