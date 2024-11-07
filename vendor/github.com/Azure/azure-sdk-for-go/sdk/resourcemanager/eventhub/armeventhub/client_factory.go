//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewApplicationGroupClient creates a new instance of ApplicationGroupClient.
func (c *ClientFactory) NewApplicationGroupClient() *ApplicationGroupClient {
	return &ApplicationGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClustersClient creates a new instance of ClustersClient.
func (c *ClientFactory) NewClustersClient() *ClustersClient {
	return &ClustersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationClient creates a new instance of ConfigurationClient.
func (c *ClientFactory) NewConfigurationClient() *ConfigurationClient {
	return &ConfigurationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConsumerGroupsClient creates a new instance of ConsumerGroupsClient.
func (c *ClientFactory) NewConsumerGroupsClient() *ConsumerGroupsClient {
	return &ConsumerGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDisasterRecoveryConfigsClient creates a new instance of DisasterRecoveryConfigsClient.
func (c *ClientFactory) NewDisasterRecoveryConfigsClient() *DisasterRecoveryConfigsClient {
	return &DisasterRecoveryConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEventHubsClient creates a new instance of EventHubsClient.
func (c *ClientFactory) NewEventHubsClient() *EventHubsClient {
	return &EventHubsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNamespacesClient creates a new instance of NamespacesClient.
func (c *ClientFactory) NewNamespacesClient() *NamespacesClient {
	return &NamespacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkSecurityPerimeterConfigurationClient creates a new instance of NetworkSecurityPerimeterConfigurationClient.
func (c *ClientFactory) NewNetworkSecurityPerimeterConfigurationClient() *NetworkSecurityPerimeterConfigurationClient {
	return &NetworkSecurityPerimeterConfigurationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient.
func (c *ClientFactory) NewNetworkSecurityPerimeterConfigurationsClient() *NetworkSecurityPerimeterConfigurationsClient {
	return &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSchemaRegistryClient creates a new instance of SchemaRegistryClient.
func (c *ClientFactory) NewSchemaRegistryClient() *SchemaRegistryClient {
	return &SchemaRegistryClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
