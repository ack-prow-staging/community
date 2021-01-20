// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package snapshot

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/elasticache/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ElastiCache{}
	_ = &svcapitypes.Snapshot{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeSnapshotsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeSnapshots", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "CacheClusterNotFound" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.Snapshots {
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.AutoMinorVersionUpgrade != nil {
			ko.Status.AutoMinorVersionUpgrade = elem.AutoMinorVersionUpgrade
		}
		if elem.AutomaticFailover != nil {
			ko.Status.AutomaticFailover = elem.AutomaticFailover
		}
		if elem.CacheClusterCreateTime != nil {
			ko.Status.CacheClusterCreateTime = &metav1.Time{*elem.CacheClusterCreateTime}
		}
		if elem.CacheClusterId != nil {
			ko.Spec.CacheClusterID = elem.CacheClusterId
		}
		if elem.CacheNodeType != nil {
			ko.Status.CacheNodeType = elem.CacheNodeType
		}
		if elem.CacheParameterGroupName != nil {
			ko.Status.CacheParameterGroupName = elem.CacheParameterGroupName
		}
		if elem.CacheSubnetGroupName != nil {
			ko.Status.CacheSubnetGroupName = elem.CacheSubnetGroupName
		}
		if elem.Engine != nil {
			ko.Status.Engine = elem.Engine
		}
		if elem.EngineVersion != nil {
			ko.Status.EngineVersion = elem.EngineVersion
		}
		if elem.KmsKeyId != nil {
			ko.Spec.KMSKeyID = elem.KmsKeyId
		}
		if elem.NodeSnapshots != nil {
			f11 := []*svcapitypes.NodeSnapshot{}
			for _, f11iter := range elem.NodeSnapshots {
				f11elem := &svcapitypes.NodeSnapshot{}
				if f11iter.CacheClusterId != nil {
					f11elem.CacheClusterID = f11iter.CacheClusterId
				}
				if f11iter.CacheNodeCreateTime != nil {
					f11elem.CacheNodeCreateTime = &metav1.Time{*f11iter.CacheNodeCreateTime}
				}
				if f11iter.CacheNodeId != nil {
					f11elem.CacheNodeID = f11iter.CacheNodeId
				}
				if f11iter.CacheSize != nil {
					f11elem.CacheSize = f11iter.CacheSize
				}
				if f11iter.NodeGroupConfiguration != nil {
					f11elemf4 := &svcapitypes.NodeGroupConfiguration{}
					if f11iter.NodeGroupConfiguration.NodeGroupId != nil {
						f11elemf4.NodeGroupID = f11iter.NodeGroupConfiguration.NodeGroupId
					}
					if f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone != nil {
						f11elemf4.PrimaryAvailabilityZone = f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone
					}
					if f11iter.NodeGroupConfiguration.PrimaryOutpostArn != nil {
						f11elemf4.PrimaryOutpostARN = f11iter.NodeGroupConfiguration.PrimaryOutpostArn
					}
					if f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones != nil {
						f11elemf4f3 := []*string{}
						for _, f11elemf4f3iter := range f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones {
							var f11elemf4f3elem string
							f11elemf4f3elem = *f11elemf4f3iter
							f11elemf4f3 = append(f11elemf4f3, &f11elemf4f3elem)
						}
						f11elemf4.ReplicaAvailabilityZones = f11elemf4f3
					}
					if f11iter.NodeGroupConfiguration.ReplicaCount != nil {
						f11elemf4.ReplicaCount = f11iter.NodeGroupConfiguration.ReplicaCount
					}
					if f11iter.NodeGroupConfiguration.ReplicaOutpostArns != nil {
						f11elemf4f5 := []*string{}
						for _, f11elemf4f5iter := range f11iter.NodeGroupConfiguration.ReplicaOutpostArns {
							var f11elemf4f5elem string
							f11elemf4f5elem = *f11elemf4f5iter
							f11elemf4f5 = append(f11elemf4f5, &f11elemf4f5elem)
						}
						f11elemf4.ReplicaOutpostARNs = f11elemf4f5
					}
					if f11iter.NodeGroupConfiguration.Slots != nil {
						f11elemf4.Slots = f11iter.NodeGroupConfiguration.Slots
					}
					f11elem.NodeGroupConfiguration = f11elemf4
				}
				if f11iter.NodeGroupId != nil {
					f11elem.NodeGroupID = f11iter.NodeGroupId
				}
				if f11iter.SnapshotCreateTime != nil {
					f11elem.SnapshotCreateTime = &metav1.Time{*f11iter.SnapshotCreateTime}
				}
				f11 = append(f11, f11elem)
			}
			ko.Status.NodeSnapshots = f11
		}
		if elem.NumCacheNodes != nil {
			ko.Status.NumCacheNodes = elem.NumCacheNodes
		}
		if elem.NumNodeGroups != nil {
			ko.Status.NumNodeGroups = elem.NumNodeGroups
		}
		if elem.Port != nil {
			ko.Status.Port = elem.Port
		}
		if elem.PreferredAvailabilityZone != nil {
			ko.Status.PreferredAvailabilityZone = elem.PreferredAvailabilityZone
		}
		if elem.PreferredMaintenanceWindow != nil {
			ko.Status.PreferredMaintenanceWindow = elem.PreferredMaintenanceWindow
		}
		if elem.PreferredOutpostArn != nil {
			ko.Status.PreferredOutpostARN = elem.PreferredOutpostArn
		}
		if elem.ReplicationGroupDescription != nil {
			ko.Status.ReplicationGroupDescription = elem.ReplicationGroupDescription
		}
		if elem.ReplicationGroupId != nil {
			ko.Spec.ReplicationGroupID = elem.ReplicationGroupId
		}
		if elem.SnapshotName != nil {
			ko.Spec.SnapshotName = elem.SnapshotName
		}
		if elem.SnapshotRetentionLimit != nil {
			ko.Status.SnapshotRetentionLimit = elem.SnapshotRetentionLimit
		}
		if elem.SnapshotSource != nil {
			ko.Status.SnapshotSource = elem.SnapshotSource
		}
		if elem.SnapshotStatus != nil {
			ko.Status.SnapshotStatus = elem.SnapshotStatus
		}
		if elem.SnapshotWindow != nil {
			ko.Status.SnapshotWindow = elem.SnapshotWindow
		}
		if elem.TopicArn != nil {
			ko.Status.TopicARN = elem.TopicArn
		}
		if elem.VpcId != nil {
			ko.Status.VPCID = elem.VpcId
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)

	// custom set output from response
	ko, err = rm.CustomDescribeSnapshotSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}

	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeSnapshotsInput, error) {
	res := &svcsdk.DescribeSnapshotsInput{}

	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	customResp, customRespErr := rm.CustomCreateSnapshot(ctx, r)
	if customResp != nil || customRespErr != nil {
		return customResp, customRespErr
	}
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateSnapshotWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateSnapshot", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Snapshot.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Snapshot.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Snapshot.AutoMinorVersionUpgrade != nil {
		ko.Status.AutoMinorVersionUpgrade = resp.Snapshot.AutoMinorVersionUpgrade
	}
	if resp.Snapshot.AutomaticFailover != nil {
		ko.Status.AutomaticFailover = resp.Snapshot.AutomaticFailover
	}
	if resp.Snapshot.CacheClusterCreateTime != nil {
		ko.Status.CacheClusterCreateTime = &metav1.Time{*resp.Snapshot.CacheClusterCreateTime}
	}
	if resp.Snapshot.CacheNodeType != nil {
		ko.Status.CacheNodeType = resp.Snapshot.CacheNodeType
	}
	if resp.Snapshot.CacheParameterGroupName != nil {
		ko.Status.CacheParameterGroupName = resp.Snapshot.CacheParameterGroupName
	}
	if resp.Snapshot.CacheSubnetGroupName != nil {
		ko.Status.CacheSubnetGroupName = resp.Snapshot.CacheSubnetGroupName
	}
	if resp.Snapshot.Engine != nil {
		ko.Status.Engine = resp.Snapshot.Engine
	}
	if resp.Snapshot.EngineVersion != nil {
		ko.Status.EngineVersion = resp.Snapshot.EngineVersion
	}
	if resp.Snapshot.NodeSnapshots != nil {
		f11 := []*svcapitypes.NodeSnapshot{}
		for _, f11iter := range resp.Snapshot.NodeSnapshots {
			f11elem := &svcapitypes.NodeSnapshot{}
			if f11iter.CacheClusterId != nil {
				f11elem.CacheClusterID = f11iter.CacheClusterId
			}
			if f11iter.CacheNodeCreateTime != nil {
				f11elem.CacheNodeCreateTime = &metav1.Time{*f11iter.CacheNodeCreateTime}
			}
			if f11iter.CacheNodeId != nil {
				f11elem.CacheNodeID = f11iter.CacheNodeId
			}
			if f11iter.CacheSize != nil {
				f11elem.CacheSize = f11iter.CacheSize
			}
			if f11iter.NodeGroupConfiguration != nil {
				f11elemf4 := &svcapitypes.NodeGroupConfiguration{}
				if f11iter.NodeGroupConfiguration.NodeGroupId != nil {
					f11elemf4.NodeGroupID = f11iter.NodeGroupConfiguration.NodeGroupId
				}
				if f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone != nil {
					f11elemf4.PrimaryAvailabilityZone = f11iter.NodeGroupConfiguration.PrimaryAvailabilityZone
				}
				if f11iter.NodeGroupConfiguration.PrimaryOutpostArn != nil {
					f11elemf4.PrimaryOutpostARN = f11iter.NodeGroupConfiguration.PrimaryOutpostArn
				}
				if f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones != nil {
					f11elemf4f3 := []*string{}
					for _, f11elemf4f3iter := range f11iter.NodeGroupConfiguration.ReplicaAvailabilityZones {
						var f11elemf4f3elem string
						f11elemf4f3elem = *f11elemf4f3iter
						f11elemf4f3 = append(f11elemf4f3, &f11elemf4f3elem)
					}
					f11elemf4.ReplicaAvailabilityZones = f11elemf4f3
				}
				if f11iter.NodeGroupConfiguration.ReplicaCount != nil {
					f11elemf4.ReplicaCount = f11iter.NodeGroupConfiguration.ReplicaCount
				}
				if f11iter.NodeGroupConfiguration.ReplicaOutpostArns != nil {
					f11elemf4f5 := []*string{}
					for _, f11elemf4f5iter := range f11iter.NodeGroupConfiguration.ReplicaOutpostArns {
						var f11elemf4f5elem string
						f11elemf4f5elem = *f11elemf4f5iter
						f11elemf4f5 = append(f11elemf4f5, &f11elemf4f5elem)
					}
					f11elemf4.ReplicaOutpostARNs = f11elemf4f5
				}
				if f11iter.NodeGroupConfiguration.Slots != nil {
					f11elemf4.Slots = f11iter.NodeGroupConfiguration.Slots
				}
				f11elem.NodeGroupConfiguration = f11elemf4
			}
			if f11iter.NodeGroupId != nil {
				f11elem.NodeGroupID = f11iter.NodeGroupId
			}
			if f11iter.SnapshotCreateTime != nil {
				f11elem.SnapshotCreateTime = &metav1.Time{*f11iter.SnapshotCreateTime}
			}
			f11 = append(f11, f11elem)
		}
		ko.Status.NodeSnapshots = f11
	}
	if resp.Snapshot.NumCacheNodes != nil {
		ko.Status.NumCacheNodes = resp.Snapshot.NumCacheNodes
	}
	if resp.Snapshot.NumNodeGroups != nil {
		ko.Status.NumNodeGroups = resp.Snapshot.NumNodeGroups
	}
	if resp.Snapshot.Port != nil {
		ko.Status.Port = resp.Snapshot.Port
	}
	if resp.Snapshot.PreferredAvailabilityZone != nil {
		ko.Status.PreferredAvailabilityZone = resp.Snapshot.PreferredAvailabilityZone
	}
	if resp.Snapshot.PreferredMaintenanceWindow != nil {
		ko.Status.PreferredMaintenanceWindow = resp.Snapshot.PreferredMaintenanceWindow
	}
	if resp.Snapshot.PreferredOutpostArn != nil {
		ko.Status.PreferredOutpostARN = resp.Snapshot.PreferredOutpostArn
	}
	if resp.Snapshot.ReplicationGroupDescription != nil {
		ko.Status.ReplicationGroupDescription = resp.Snapshot.ReplicationGroupDescription
	}
	if resp.Snapshot.SnapshotRetentionLimit != nil {
		ko.Status.SnapshotRetentionLimit = resp.Snapshot.SnapshotRetentionLimit
	}
	if resp.Snapshot.SnapshotSource != nil {
		ko.Status.SnapshotSource = resp.Snapshot.SnapshotSource
	}
	if resp.Snapshot.SnapshotStatus != nil {
		ko.Status.SnapshotStatus = resp.Snapshot.SnapshotStatus
	}
	if resp.Snapshot.SnapshotWindow != nil {
		ko.Status.SnapshotWindow = resp.Snapshot.SnapshotWindow
	}
	if resp.Snapshot.TopicArn != nil {
		ko.Status.TopicARN = resp.Snapshot.TopicArn
	}
	if resp.Snapshot.VpcId != nil {
		ko.Status.VPCID = resp.Snapshot.VpcId
	}

	rm.setStatusDefaults(ko)

	// custom set output from response
	ko, err = rm.CustomCreateSnapshotSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateSnapshotInput, error) {
	res := &svcsdk.CreateSnapshotInput{}

	if r.ko.Spec.CacheClusterID != nil {
		res.SetCacheClusterId(*r.ko.Spec.CacheClusterID)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.ReplicationGroupID != nil {
		res.SetReplicationGroupId(*r.ko.Spec.ReplicationGroupID)
	}
	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	return rm.customUpdateSnapshot(ctx, desired, latest, diffReporter)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteSnapshotWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteSnapshot", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteSnapshotInput, error) {
	res := &svcsdk.DeleteSnapshotInput{}

	if r.ko.Spec.SnapshotName != nil {
		res.SetSnapshotName(*r.ko.Spec.SnapshotName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Snapshot,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	// custom update conditions
	customUpdate := rm.CustomUpdateConditions(ko, r, err)
	if terminalCondition != nil || customUpdate {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidParameter",
		"InvalidParameterValue",
		"InvalidParameterCombination",
		"SnapshotAlreadyExistsFault",
		"CacheClusterNotFound",
		"ReplicationGroupNotFoundFault",
		"SnapshotQuotaExceededFault",
		"SnapshotFeatureNotSupportedFault":
		return true
	default:
		return false
	}
}
