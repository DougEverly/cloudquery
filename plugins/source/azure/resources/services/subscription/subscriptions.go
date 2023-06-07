package subscription

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_subscription_subscriptions",
		Resolver:             fetchSubscriptions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list?tabs=HTTP#subscription",
		Transform:            transformers.TransformWithStruct(&armsubscription.Subscription{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
		Relations: []*schema.Table{
			locations(),
		},
	}
}
