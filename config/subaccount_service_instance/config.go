package subaccount_trust_configuration

import (
	"context"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("btp_subaccount_service_instance", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "SubaccountServiceInstance"
		// we only use this resource internally to set the client_credentials grant type, so there is no harm in avoiding usage of secrets here
		// it makes the setup a lot easier
		r.TerraformResource.Schema["parameters"].Sensitive = false

		r.ExternalName.OmittedFields = []string{"timeouts"}
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
			// When using "" as ID the API endpoint call will fail, so we need to use anything else that
			// won't yield a result
			if externalName == "" {
				return "NOT_EMPTY_GUID", nil
			}
			return externalName, nil
		}
		// this prevents a callback to the manager, which makes integration of controller calls from within another controller easier
		r.UseAsync = false
	})
}
