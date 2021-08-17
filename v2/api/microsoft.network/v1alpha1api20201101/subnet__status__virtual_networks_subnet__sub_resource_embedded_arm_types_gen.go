// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

//Generated from:
type Subnet_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the subnet.
	Properties *SubnetPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type SubnetPropertiesFormat_StatusARM struct {
	//AddressPrefix: The address prefix for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	//AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	//ApplicationGatewayIpConfigurations: Application gateway IP configurations of
	//virtual network resource.
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_StatusARM `json:"applicationGatewayIpConfigurations,omitempty"`

	//Delegations: An array of references to the delegations on the subnet.
	Delegations []Delegation_StatusARM `json:"delegations,omitempty"`

	//IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResource_StatusARM `json:"ipAllocations,omitempty"`

	//IpConfigurationProfiles: Array of IP configuration profiles which reference this
	//subnet.
	IpConfigurationProfiles []IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"ipConfigurationProfiles,omitempty"`

	//IpConfigurations: An array of references to the network interface IP
	//configurations using subnet.
	IpConfigurations []IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"ipConfigurations,omitempty"`

	//NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResource_StatusARM `json:"natGateway,omitempty"`

	//NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"networkSecurityGroup,omitempty"`

	//PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on
	//private end point in the subnet.
	PrivateEndpointNetworkPolicies *SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies `json:"privateEndpointNetworkPolicies,omitempty"`

	//PrivateEndpoints: An array of references to private endpoints.
	PrivateEndpoints []PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"privateEndpoints,omitempty"`

	//PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on
	//private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies `json:"privateLinkServiceNetworkPolicies,omitempty"`

	//ProvisioningState: The provisioning state of the subnet resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//Purpose: A read-only string identifying the intention of use for this subnet
	//based on delegations and other user-defined properties.
	Purpose *string `json:"purpose,omitempty"`

	//ResourceNavigationLinks: An array of references to the external resources using
	//subnet.
	ResourceNavigationLinks []ResourceNavigationLink_StatusARM `json:"resourceNavigationLinks,omitempty"`

	//RouteTable: The reference to the RouteTable resource.
	RouteTable *RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"routeTable,omitempty"`

	//ServiceAssociationLinks: An array of references to services injecting into this
	//subnet.
	ServiceAssociationLinks []ServiceAssociationLink_StatusARM `json:"serviceAssociationLinks,omitempty"`

	//ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"serviceEndpointPolicies,omitempty"`

	//ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat_StatusARM `json:"serviceEndpoints,omitempty"`
}

//Generated from:
type ApplicationGatewayIPConfiguration_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: Name of the IP configuration that is unique within an Application Gateway.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the application gateway IP configuration.
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type Delegation_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a subnet. This name can be
	//used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the IP configuration profile.
	Properties *IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Type: Sub Resource type.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the IP configuration.
	Properties *IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

//Generated from:
type NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

//Generated from:
type PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

//Generated from:
type ResourceNavigationLink_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: Name of the resource that is unique within a resource group. This name can
	//be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Resource navigation link properties format.
	Properties *ResourceNavigationLinkFormat_StatusARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

//Generated from:
type ServiceAssociationLink_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: Name of the resource that is unique within a resource group. This name can
	//be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Resource navigation link properties format.
	Properties *ServiceAssociationLinkPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

//Generated from:
type ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Kind: Kind of service endpoint policy. This is metadata used for the Azure
	//portal experience.
	Kind *string `json:"kind,omitempty"`
}

//Generated from:
type ServiceEndpointPropertiesFormat_StatusARM struct {
	//Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	//ProvisioningState: The provisioning state of the service endpoint resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//Service: The type of the endpoint service.
	Service *string `json:"service,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies string

const (
	SubnetPropertiesFormatStatusPrivateEndpointNetworkPoliciesDisabled = SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies("Disabled")
	SubnetPropertiesFormatStatusPrivateEndpointNetworkPoliciesEnabled  = SubnetPropertiesFormatStatusPrivateEndpointNetworkPolicies("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies string

const (
	SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPoliciesDisabled = SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies("Disabled")
	SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPoliciesEnabled  = SubnetPropertiesFormatStatusPrivateLinkServiceNetworkPolicies("Enabled")
)

//Generated from:
type ApplicationGatewayIPConfigurationPropertiesFormat_StatusARM struct {
	//ProvisioningState: The provisioning state of the application gateway IP
	//configuration resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//Subnet: Reference to the subnet resource. A subnet from where application
	//gateway gets its private address.
	Subnet *SubResource_StatusARM `json:"subnet,omitempty"`
}

//Generated from:
type IPConfigurationProfilePropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//ProvisioningState: The provisioning state of the IP configuration profile
	//resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`
}

//Generated from:
type IPConfigurationPropertiesFormat_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	//PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_Status `json:"privateIPAllocationMethod,omitempty"`

	//ProvisioningState: The provisioning state of the IP configuration resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//PublicIPAddress: The reference to the public IP resource.
	PublicIPAddress *PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`
}

//Generated from:
type ResourceNavigationLinkFormat_StatusARM struct {
	//Link: Link to the external resource.
	Link *string `json:"link,omitempty"`

	//LinkedResourceType: Resource type of the linked resource.
	LinkedResourceType *string `json:"linkedResourceType,omitempty"`

	//ProvisioningState: The provisioning state of the resource navigation link
	//resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`
}

//Generated from:
type ServiceAssociationLinkPropertiesFormat_StatusARM struct {
	//AllowDelete: If true, the resource can be deleted.
	AllowDelete *bool `json:"allowDelete,omitempty"`

	//Link: Link to the external resource.
	Link *string `json:"link,omitempty"`

	//LinkedResourceType: Resource type of the linked resource.
	LinkedResourceType *string `json:"linkedResourceType,omitempty"`

	//Locations: A list of locations.
	Locations []string `json:"locations,omitempty"`

	//ProvisioningState: The provisioning state of the service association link
	//resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`
}

//Generated from:
type ServiceDelegationPropertiesFormat_StatusARM struct {
	//Actions: The actions permitted to the service upon delegation.
	Actions []string `json:"actions,omitempty"`

	//ProvisioningState: The provisioning state of the service delegation resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//ServiceName: The name of the service to whom the subnet should be delegated
	//(e.g. Microsoft.Sql/servers).
	ServiceName *string `json:"serviceName,omitempty"`
}

//Generated from:
type PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbeddedARM struct {
	//ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Sku: The public IP address SKU.
	Sku *PublicIPAddressSku_StatusARM `json:"sku,omitempty"`

	//Zones: A list of availability zones denoting the IP allocated for the resource
	//needs to come from.
	Zones []string `json:"zones,omitempty"`
}
