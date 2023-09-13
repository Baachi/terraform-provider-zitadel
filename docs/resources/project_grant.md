---
page_title: "zitadel_project_grant Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.
---

# zitadel_project_grant (Resource)

Resource representing the grant of a project to a different organization, also containing the available roles which can be given to the members of the projectgrant.

## Example Usage

```terraform
resource "zitadel_project_grant" "default" {
  org_id         = data.zitadel_org.default.id
  project_id     = data.zitadel_project.default.id
  granted_org_id = data.zitadel_org.granted_org.id
  role_keys      = ["super-user"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `granted_org_id` (String) ID of the organization granted the project
- `project_id` (String) ID of the project

### Optional

- `org_id` (String) ID of the organization
- `role_keys` (Set of String) List of roles granted

### Read-Only

- `id` (String) The ID of this resource.

## Import

```terraform
# The resource can be imported using the ID format `<id:project_id[:org_id]>`, e.g.
terraform import project_grant.imported '123456789012345678:123456789012345678:123456789012345678'
```