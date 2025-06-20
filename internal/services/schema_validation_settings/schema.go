// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*SchemaValidationSettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"validation_default_mitigation_action": schema.StringAttribute{
				Description: "The default mitigation action used\nMitigation actions are as follows:\n\n  - `\"log\"` - log request when request does not conform to schema\n  - `\"block\"` - deny access to the site when request does not conform to schema\n  - `\"none\"` - skip running schema validation\nAvailable values: \"none\", \"log\", \"block\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"log",
						"block",
					),
				},
			},
			"validation_override_mitigation_action": schema.StringAttribute{
				Description: "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `\"none\"` - skip running schema validation entirely for the request\n  - `null` - clears any existing override\nAvailable values: \"none\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("none"),
				},
			},
		},
	}
}

func (r *SchemaValidationSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SchemaValidationSettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
