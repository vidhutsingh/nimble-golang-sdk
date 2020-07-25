// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Group is a collection of arrays operating together organized into storage pools.
const (
	groupPath = "groups"
)

// GroupObjectSet
type GroupObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Group object
func (objectSet *GroupObjectSet) CreateObject(payload *nimbleos.Group) (*nimbleos.Group, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Group")
}

// UpdateObject Modify existing Group object
func (objectSet *GroupObjectSet) UpdateObject(id string, payload *nimbleos.Group) (*nimbleos.Group, error) {
	resp, err := objectSet.Client.Put(groupPath, id, payload, &nimbleos.Group{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.Group), err
}

// DeleteObject deletes the Group object with the specified ID
func (objectSet *GroupObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Group")
}

// GetObject returns a Group object with the given ID
func (objectSet *GroupObjectSet) GetObject(id string) (*nimbleos.Group, error) {
	resp, err := objectSet.Client.Get(groupPath, id, nimbleos.Group{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Group), err
}

// GetObjectList returns the list of Group objects
func (objectSet *GroupObjectSet) GetObjectList() ([]*nimbleos.Group, error) {
	resp, err := objectSet.Client.List(groupPath)
	if err != nil {
		return nil, err
	}
	return buildGroupObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Group objects using the given params query info
func (objectSet *GroupObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Group, error) {
	groupObjectSetResp, err := objectSet.Client.ListFromParams(groupPath, params)
	if err != nil {
		return nil, err
	}
	return buildGroupObjectSet(groupObjectSetResp), err
}

// generated function to build the appropriate response types
func buildGroupObjectSet(response interface{}) []*nimbleos.Group {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Group, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Group{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}