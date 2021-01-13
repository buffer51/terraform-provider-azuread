package client

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/manicminer/hamilton/clients"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AadClient *graphrbac.GroupsClient
	MsClient  *clients.GroupsClient
}

func NewClient(o *common.ClientOptions) *Client {
	aadClient := graphrbac.NewGroupsClientWithBaseURI(o.AadGraphEndpoint, o.TenantID)
	msClient := clients.NewGroupsClient(o.TenantID)
	o.ConfigureClient(&msClient.BaseClient, &aadClient.Client)

	return &Client{
		AadClient: &aadClient,
		MsClient:  msClient,
	}
}
