// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning_expression

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ContentScanningExpressionResultEnvelope struct {
	Result *[]*ContentScanningExpressionBodyModel `json:"result"`
}

type ContentScanningExpressionModel struct {
	ID     types.String                           `tfsdk:"id" json:"id,computed"`
	ZoneID types.String                           `tfsdk:"zone_id" path:"zone_id,required"`
	Body   *[]*ContentScanningExpressionBodyModel `tfsdk:"body" json:"body,required"`
}

func (m ContentScanningExpressionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m ContentScanningExpressionModel) MarshalJSONForUpdate(state ContentScanningExpressionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}

type ContentScanningExpressionBodyModel struct {
	Payload types.String `tfsdk:"payload" json:"payload,required"`
}
