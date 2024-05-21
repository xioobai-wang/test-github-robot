---
subcategory: "wafxxx"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_waf_test"
description: |-
  wafxxx
---

# huaweicloud_waf_test

wafxxx

## Example Usage

```hcl
wafxxx
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `idxxxxxx` - (Optional, List) VPC ID, which can be used to filter VPCs.

* `name` - (Optional, List) VPC name, which can be used to filter VPCs.

* `description` - (Optional, List) Supplementary information about the VPC, which can be used to filter VPCs.

* `cidr` - (Optional, List) VPC CIDR block, which can be used to filter VPCs.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `vpcs` - Response body of VPCs

  The [vpcs](#vpcs_struct) structure is documented below.

* `page_info` - Pagination information

  The [page_info](#page_info_struct) structure is documented below.

<a name="vpcs_struct"></a>
The `vpcs` block supports:

* `project_id` - ID of the project to which the VPC belongs

* `created_at` - Time when the VPC is created.
  UTC time in the format of yyyy-MM-ddTHH:mmss

* `updated_at` - Time when the VPC is updated.
  UTC time in the format of yyyy-MM-ddTHH:mmss

* `cloud_resources` - Type and number of resources associated with the VPC.
  Currently, only route tables and subnets of the VPC are returned.

  The [cloud_resources](#vpcs_cloud_resources_struct) structure is documented below.

* `extend_cidrs` - Secondary CIDR blocks of VPCs.
  The value can be:
  Currently, only IPv4 CIDR blocks are supported.

* `status` - VPC status.
  The value can be PENDING (indicates that the VPC is being created) or ACTIVE (indicates that the VPC is created successfully).

* `description` - Provides supplementary information about the VPC.
  The value can contain no more than 255 characters and cannot contain angle brackets (< or >).

* `cidr` - Available VPC CIDR blocks.
  The value can be:
  - 10.0.0.0/8-10.255.255.240/28
  - 172.16.0.0/12-172.31.255.240/28
  - 192.168.0.0/16-192.168.255.240/28
  If cidr is not specified, the default value is "".
  The value must be in IPv4 CIDR format, for example, 192.168.0.0/16.

* `enterprise_project_id` - ID of the enterprise project to which the VPC belongs.
  The value is 0 or a string that contains a maximum of 36 characters in UUID format with hyphens (-). Value 0 indicates the default enterprise project.

* `tags` - VPC tags. For details, see the Tag objects.
  Value range: 0 to 10 tag key-value pairs.

  The [tags](#vpcs_tags_struct) structure is documented below.

* `id` - VPC ID, which uniquely identifies the VPC.
  The value is in UUID format with hyphens (-).

* `name` - VPC name.
  The value can contain 1 to 64 characters, including letters, digits, underscores (_), hyphens (-), and periods (.).

<a name="vpcs_cloud_resources_struct"></a>
The `cloud_resources` block supports:

* `resource_type` - Resource type

* `resource_count` - Number of resources

<a name="vpcs_tags_struct"></a>
The `tags` block supports:

* `value` - 1. Tag value
  2. A tag value contains a maximum of 43 Unicode characters and can be left blank. It cannot contain non-printable ASCII characters (0–31) or the following special characters: asterisks (*), left angle brackets (<), right angle brackets (>), backslashes (\), or equal signs (=).

* `key` - 1. Tag key.
  2. A tag key contains a maximum of 36 Unicode characters. A tag key cannot be left blank. It cannot contain non-printable ASCII characters (0–31) or the following special characters: asterisks (*), left angle brackets (<), right angle brackets (>), backslashes (\), or equal signs (=).

<a name="page_info_struct"></a>
The `page_info` block supports:

* `previous_marker` - First record on the current page

* `current_count` - Total number of records on the current page

* `next_marker` - Last record on the current page. This parameter does not exist if the page is the last one.
