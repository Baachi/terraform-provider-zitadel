---
page_title: "zitadel_project Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing the project, which can then be granted to different organizations or users directly, containing different applications.
---

# zitadel_project (Resource)

Resource representing the project, which can then be granted to different organizations or users directly, containing different applications.

## Example Usage

```terraform
resource "zitadel_project" "default" {
  name                     = "projectname"
  org_id                   = data.zitadel_org.default.id
  project_role_assertion   = true
  project_role_check       = true
  has_project_check        = true
  private_labeling_setting = "PRIVATE_LABELING_SETTING_ENFORCE_PROJECT_RESOURCE_OWNER_POLICY"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the project

### Optional

- `has_project_check` (Boolean) ZITADEL checks if the org of the user has permission to this project
- `org_id` (String) ID of the organization
- `private_labeling_setting` (String) Defines from where the private labeling should be triggered, supported values: PRIVATE_LABELING_SETTING_UNSPECIFIED, PRIVATE_LABELING_SETTING_ENFORCE_PROJECT_RESOURCE_OWNER_POLICY, PRIVATE_LABELING_SETTING_ALLOW_LOGIN_USER_RESOURCE_OWNER_POLICY
- `project_role_assertion` (Boolean) describes if roles of user should be added in token
- `project_role_check` (Boolean) ZITADEL checks if the user has at least one on this project

### Read-Only

- `id` (String) The ID of this resource.
- `state` (String) State of the project

## Import

```terraform
# The resource can be imported using the ID format `<id[:org_id]>`, e.g.
terraform import project.imported '123456789012345678:123456789012345678'
```