---
subcategory: "ECS"
---

# huaweicloud_ecs_yyy

Use this data source to get the traffic mirror sessions.

## Example Usage

```hcl
data "huaweicloud_ecs_yyy" "test1" {
  name = "mirror-session-a6b5"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `ecs_yyy` - (Optional, String) Specifies the traffic mirror session ID used to query.
