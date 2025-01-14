---
page_title: "zitadel_smtp_config Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing the SMTP configuration of an instance.
---

# zitadel_smtp_config (Resource)

Resource representing the SMTP configuration of an instance.

## Example Usage

```terraform
resource "zitadel_smtp_config" "default" {
  sender_address   = "sender@example.com"
  sender_name      = "no-reply"
  tls              = true
  host             = "localhost:25"
  user             = "user"
  password         = "secret_password"
  reply_to_address = "replyto@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) Host and port address to your SMTP server.
- `sender_address` (String) Address used to send emails.
- `sender_name` (String) Sender name used to send emails.

### Optional

- `password` (String, Sensitive) Password used to communicate with your SMTP server.
- `reply_to_address` (String) Address to reply to.
- `tls` (Boolean) TLS used to communicate with your SMTP server.
- `user` (String) User used to communicate with your SMTP server.

### Read-Only

- `id` (String) The ID of this resource.

## Import

```terraform
# The resource can be imported using the ID format `<[password]>`, e.g.
terraform import smtp_config.imported 'p4ssw0rd'
```
