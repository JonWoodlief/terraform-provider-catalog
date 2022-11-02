// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmPresetDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmPresetDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_preset.cm_preset", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_preset.cm_preset", "preset"),
				),
			},
		},
	})
}

func testAccCheckIBMCmPresetDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_cm_preset" "cm_preset" {
			id = <PUT ID HERE>
		}
	`)
}
