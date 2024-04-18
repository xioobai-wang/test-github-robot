package huaweicloud

import "fmt"

func Provider() {
	DataSourcesMap := map[string]string{
		"huaweicloud_ecs_xxx": "DataSourceEcsXxx",

		"huaweicloudd_vpc_xxx": "DataSourceVpcXxx()",
		"huaweicloudd_vpc_yyy": "DataSourceVpcYyy",
	}
	ResourcesMap := map[string]string{
		"huaweicloud_ecs_yyy": "ResourceEcsYyy",
	}
	fmt.Println(DataSourcesMap, ResourcesMap)
}
