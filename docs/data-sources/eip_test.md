---
subcategory: "EIP"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_eip_test"
description: |-
  xxx
---

# huaweicloud_eip_test

xxx

## Example Usage

```hcl
uuu
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `eip_xxx_id` - (Optional, String) Filter or sort by ID.

* `protocol` - (Optional, String) Filter by protocol.

* `priority` - (Optional, String) Filter by priority.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `traffic_mirror_filter_rules` - List of traffic mirror filter rules.

  The [traffic_mirror_filter_rules](#traffic_mirror_filter_rules_struct) structure is documented below.

<a name="traffic_mirror_filter_rules_struct"></a>
The `traffic_mirror_filter_rules` block supports:

* `direction` - Traffic direction
  The value can be ingress (inbound direction) or egress (outbound direction).

* `protocol` - Protocol of the mirrored traffic
  The value can be TCP, UDP, ICMP, ICMPV6, or ALL.
