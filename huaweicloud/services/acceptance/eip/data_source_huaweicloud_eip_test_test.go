package eip

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceEipTest_basic(t *testing.T) {
	dataSource := "data.huaweicloud_eip_test.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceDataSourceEipTest_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					// TODO: 请补充你的验证项，越多越好
				),
			},
		},
	})
}

func testDataSourceDataSourceEipTest_basic(name string) string {
	// TODO: 请参考一下代码完善测试脚本
	// 注意以下代码仅供参考，一些复杂类型的字段可能会被遗漏，请与代码详细核对（提交时，请删除本行说明）
	return fmt.Sprintf(`
data "huaweicloud_eip_test" "test" {
  eip_xxx_id = ""
  protocol   = ""
  priority   = ""
}
`, name)
}
