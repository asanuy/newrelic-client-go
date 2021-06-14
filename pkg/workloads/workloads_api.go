// Code generated by tutone: DO NOT EDIT
package workloads

import (
	"context"

	"github.com/newrelic/newrelic-client-go/pkg/entities"
)

// Creates a new workload.
func (a *Workloads) WorkloadCreate(
	accountID int,
	workload WorkloadCreateInput,
) (*WorkloadCollection, error) {
	return a.WorkloadCreateWithContext(context.Background(),
		accountID,
		workload,
	)
}

// Creates a new workload.
func (a *Workloads) WorkloadCreateWithContext(
	ctx context.Context,
	accountID int,
	workload WorkloadCreateInput,
) (*WorkloadCollection, error) {

	resp := WorkloadCreateQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"workload":  workload,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, WorkloadCreateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadCollection, nil
}

type WorkloadCreateQueryResponse struct {
	WorkloadCollection WorkloadCollection `json:"WorkloadCreate"`
}

const WorkloadCreateMutation = `mutation(
	$accountId: Int!,
	$workload: WorkloadCreateInput!,
) { workloadCreate(
	accountId: $accountId,
	workload: $workload,
) {
	account {
		id
		name
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entities {
		guid
	}
	entitySearchQueries {
		createdAt
		createdBy {
			email
			gravatar
			id
			name
		}
		id
		query
		updatedAt
	}
	entitySearchQuery
	guid
	id
	name
	permalink
	scopeAccounts {
		accountIds
	}
	status {
		description
		source
		statusDetails {
			__typename
			source
			value
			... on WorkloadRollupRuleStatusResult {
				__typename
			}
			... on WorkloadStaticStatusResult {
				__typename
				description
				summary
			}
		}
		summary
		value
	}
	statusConfig {
		automatic {
			enabled
			rules {
				id
			}
		}
		static {
			description
			enabled
			id
			status
			summary
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Deletes an existing workload.
func (a *Workloads) WorkloadDelete(
	gUID entities.EntityGUID,
) (*WorkloadCollection, error) {
	return a.WorkloadDeleteWithContext(context.Background(),
		gUID,
	)
}

// Deletes an existing workload.
func (a *Workloads) WorkloadDeleteWithContext(
	ctx context.Context,
	gUID entities.EntityGUID,
) (*WorkloadCollection, error) {

	resp := WorkloadDeleteQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, WorkloadDeleteMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadCollection, nil
}

type WorkloadDeleteQueryResponse struct {
	WorkloadCollection WorkloadCollection `json:"WorkloadDelete"`
}

const WorkloadDeleteMutation = `mutation(
	$guid: EntityGuid!,
) { workloadDelete(
	guid: $guid,
) {
	account {
		id
		name
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entities {
		guid
	}
	entitySearchQueries {
		createdAt
		createdBy {
			email
			gravatar
			id
			name
		}
		id
		query
		updatedAt
	}
	entitySearchQuery
	guid
	id
	name
	permalink
	scopeAccounts {
		accountIds
	}
	status {
		description
		source
		statusDetails {
			__typename
			source
			value
			... on WorkloadRollupRuleStatusResult {
				__typename
			}
			... on WorkloadStaticStatusResult {
				__typename
				description
				summary
			}
		}
		summary
		value
	}
	statusConfig {
		automatic {
			enabled
			rules {
				id
			}
		}
		static {
			description
			enabled
			id
			status
			summary
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Duplicates an existing workload.
func (a *Workloads) WorkloadDuplicate(
	accountID int,
	sourceGUID entities.EntityGUID,
	workload WorkloadDuplicateInput,
) (*WorkloadCollection, error) {
	return a.WorkloadDuplicateWithContext(context.Background(),
		accountID,
		sourceGUID,
		workload,
	)
}

// Duplicates an existing workload.
func (a *Workloads) WorkloadDuplicateWithContext(
	ctx context.Context,
	accountID int,
	sourceGUID entities.EntityGUID,
	workload WorkloadDuplicateInput,
) (*WorkloadCollection, error) {

	resp := WorkloadDuplicateQueryResponse{}
	vars := map[string]interface{}{
		"accountId":  accountID,
		"sourceGuid": sourceGUID,
		"workload":   workload,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, WorkloadDuplicateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadCollection, nil
}

type WorkloadDuplicateQueryResponse struct {
	WorkloadCollection WorkloadCollection `json:"WorkloadDuplicate"`
}

const WorkloadDuplicateMutation = `mutation(
	$accountId: Int!,
	$sourceGuid: EntityGuid!,
	$workload: WorkloadDuplicateInput,
) { workloadDuplicate(
	accountId: $accountId,
	sourceGuid: $sourceGuid,
	workload: $workload,
) {
	account {
		id
		name
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entities {
		guid
	}
	entitySearchQueries {
		createdAt
		createdBy {
			email
			gravatar
			id
			name
		}
		id
		query
		updatedAt
	}
	entitySearchQuery
	guid
	id
	name
	permalink
	scopeAccounts {
		accountIds
	}
	status {
		description
		source
		statusDetails {
			__typename
			source
			value
			... on WorkloadRollupRuleStatusResult {
				__typename
			}
			... on WorkloadStaticStatusResult {
				__typename
				description
				summary
			}
		}
		summary
		value
	}
	statusConfig {
		automatic {
			enabled
			rules {
				id
			}
		}
		static {
			description
			enabled
			id
			status
			summary
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Updates an existing workload.
func (a *Workloads) WorkloadUpdate(
	gUID entities.EntityGUID,
	workload WorkloadUpdateInput,
) (*WorkloadCollection, error) {
	return a.WorkloadUpdateWithContext(context.Background(),
		gUID,
		workload,
	)
}

// Updates an existing workload.
func (a *Workloads) WorkloadUpdateWithContext(
	ctx context.Context,
	gUID entities.EntityGUID,
	workload WorkloadUpdateInput,
) (*WorkloadCollection, error) {

	resp := WorkloadUpdateQueryResponse{}
	vars := map[string]interface{}{
		"guid":     gUID,
		"workload": workload,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, WorkloadUpdateMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadCollection, nil
}

type WorkloadUpdateQueryResponse struct {
	WorkloadCollection WorkloadCollection `json:"WorkloadUpdate"`
}

const WorkloadUpdateMutation = `mutation(
	$guid: EntityGuid!,
	$workload: WorkloadUpdateInput!,
) { workloadUpdate(
	guid: $guid,
	workload: $workload,
) {
	account {
		id
		name
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	entities {
		guid
	}
	entitySearchQueries {
		createdAt
		createdBy {
			email
			gravatar
			id
			name
		}
		id
		query
		updatedAt
	}
	entitySearchQuery
	guid
	id
	name
	permalink
	scopeAccounts {
		accountIds
	}
	status {
		description
		source
		statusDetails {
			__typename
			source
			value
			... on WorkloadRollupRuleStatusResult {
				__typename
			}
			... on WorkloadStaticStatusResult {
				__typename
				description
				summary
			}
		}
		summary
		value
	}
	statusConfig {
		automatic {
			enabled
			rules {
				id
			}
		}
		static {
			description
			enabled
			id
			status
			summary
		}
	}
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`
