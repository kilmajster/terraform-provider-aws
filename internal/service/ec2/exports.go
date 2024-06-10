// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

// Exports for use in other modules.
var (
	CustomFiltersBlock                                               = customFiltersBlock
	DeleteNetworkInterface                                           = deleteNetworkInterface
	DetachNetworkInterface                                           = detachNetworkInterface
	FindImageByID                                                    = findImageByID
	FindInstanceByID                                                 = findInstanceByID
	FindNetworkInterfacesByAttachmentInstanceOwnerIDAndDescriptionV2 = findNetworkInterfacesByAttachmentInstanceOwnerIDAndDescriptionV2
	FindNetworkInterfacesV2                                          = findNetworkInterfacesV2
	FindVPCByIDV2                                                    = findVPCByIDV2
	FindVPCEndpointByIDV2                                            = findVPCEndpointByIDV2
	NewCustomFilterListFrameworkV2                                   = newCustomFilterListFrameworkV2
	NewFilter                                                        = newFilter
	NewFilterV2                                                      = newFilterV2
	ResourceAMI                                                      = resourceAMI
	ResourceTransitGateway                                           = resourceTransitGateway
	ResourceTransitGatewayConnectPeer                                = resourceTransitGatewayConnectPeer
	VPCEndpointCreationTimeout                                       = vpcEndpointCreationTimeout
	WaitVPCEndpointAvailableV2                                       = waitVPCEndpointAvailable
)
