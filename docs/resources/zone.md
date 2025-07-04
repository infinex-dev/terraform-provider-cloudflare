---
page_title: "cloudflare_zone Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zone (Resource)



-> If you are attempting to sign up a subdomain of a zone you must first have Subdomain Support entitlement for your account.

## Example Usage

```terraform
resource "cloudflare_zone" "example_zone" {
  account = {
    id = "023e105f4ecef8ad9ca31a8372d0c353"
  }
  name = "example.com"
  type = "full"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account` (Attributes) (see [below for nested schema](#nestedatt--account))
- `name` (String) The domain name.

### Optional

- `paused` (Boolean) Indicates whether the zone is only using Cloudflare DNS services. A
true value means the zone will not receive security or performance
benefits.
- `type` (String) A full zone implies that DNS is hosted with Cloudflare. A partial zone is
typically a partner-hosted zone or a CNAME setup.
Available values: "full", "partial", "secondary", "internal".
- `vanity_name_servers` (List of String) An array of domains used for custom name servers. This is only
available for Business and Enterprise plans.

### Read-Only

- `activated_on` (String) The last time proof of ownership was detected and the zone was made
active.
- `cname_suffix` (String) Allows the customer to use a custom apex.
*Tenants Only Configuration*.
- `created_on` (String) When the zone was created.
- `development_mode` (Number) The interval (in seconds) from when development mode expires
(positive integer) or last expired (negative integer) for the
domain. If development mode has never been enabled, this value is 0.
- `id` (String) Identifier
- `meta` (Attributes) Metadata about the zone. (see [below for nested schema](#nestedatt--meta))
- `modified_on` (String) When the zone was last modified.
- `name_servers` (List of String) The name servers Cloudflare assigns to a zone.
- `original_dnshost` (String) DNS host at the time of switching to Cloudflare.
- `original_name_servers` (List of String) Original name servers before moving to Cloudflare.
- `original_registrar` (String) Registrar for the domain at the time of switching to Cloudflare.
- `owner` (Attributes) The owner of the zone. (see [below for nested schema](#nestedatt--owner))
- `permissions` (List of String, Deprecated) Legacy permissions based on legacy user membership information.
- `plan` (Attributes, Deprecated) A Zones subscription information. (see [below for nested schema](#nestedatt--plan))
- `status` (String) The zone status on Cloudflare.
Available values: "initializing", "pending", "active", "moved".
- `tenant` (Attributes) The root organizational unit that this zone belongs to (such as a tenant or organization). (see [below for nested schema](#nestedatt--tenant))
- `tenant_unit` (Attributes) The immediate parent organizational unit that this zone belongs to (such as under a tenant or sub-organization). (see [below for nested schema](#nestedatt--tenant_unit))
- `verification_key` (String) Verification key for partial zone setup.

<a id="nestedatt--account"></a>
### Nested Schema for `account`

Optional:

- `id` (String) Identifier


<a id="nestedatt--meta"></a>
### Nested Schema for `meta`

Read-Only:

- `cdn_only` (Boolean) The zone is only configured for CDN.
- `custom_certificate_quota` (Number) Number of Custom Certificates the zone can have.
- `dns_only` (Boolean) The zone is only configured for DNS.
- `foundation_dns` (Boolean) The zone is setup with Foundation DNS.
- `page_rule_quota` (Number) Number of Page Rules a zone can have.
- `phishing_detected` (Boolean) The zone has been flagged for phishing.
- `step` (Number)


<a id="nestedatt--owner"></a>
### Nested Schema for `owner`

Read-Only:

- `id` (String) Identifier
- `name` (String) Name of the owner.
- `type` (String) The type of owner.


<a id="nestedatt--plan"></a>
### Nested Schema for `plan`

Read-Only:

- `can_subscribe` (Boolean) States if the subscription can be activated.
- `currency` (String) The denomination of the customer.
- `externally_managed` (Boolean) If this Zone is managed by another company.
- `frequency` (String) How often the customer is billed.
- `id` (String) Identifier
- `is_subscribed` (Boolean) States if the subscription active.
- `legacy_discount` (Boolean) If the legacy discount applies to this Zone.
- `legacy_id` (String) The legacy name of the plan.
- `name` (String) Name of the owner.
- `price` (Number) How much the customer is paying.


<a id="nestedatt--tenant"></a>
### Nested Schema for `tenant`

Read-Only:

- `id` (String) Identifier
- `name` (String) The name of the Tenant account.


<a id="nestedatt--tenant_unit"></a>
### Nested Schema for `tenant_unit`

Read-Only:

- `id` (String) Identifier

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zone.example '<zone_id>'
```
