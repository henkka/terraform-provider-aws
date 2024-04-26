// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshift

// Exports for use in tests only.
var (
	ResourceAuthenticationProfile        = resourceAuthenticationProfile
	ResourceCluster                      = resourceCluster
	ResourceClusterIAMRoles              = resourceClusterIAMRoles
	ResourceClusterSnapshot              = resourceClusterSnapshot
	ResourceDataShareAuthorization       = newResourceDataShareAuthorization
	ResourceDataShareConsumerAssociation = newResourceDataShareConsumerAssociation
	ResourceEndpointAccess               = resourceEndpointAccess
	ResourceLogging                      = newResourceLogging
	ResourceSnapshotCopy                 = newResourceSnapshotCopy

	FindAuthenticationProfileByID        = findAuthenticationProfileByID
	FindClusterByID                      = findClusterByID
	FindClusterSnapshotByID              = findClusterSnapshotByID
	FindDataShareAuthorizationByID       = findDataShareAuthorizationByID
	FindDataShareConsumerAssociationByID = findDataShareConsumerAssociationByID
	FindEndpointAccessByName             = findEndpointAccessByName
	FindLoggingByID                      = findLoggingByID
	FindSnapshotCopyByID                 = findSnapshotCopyByID
)
