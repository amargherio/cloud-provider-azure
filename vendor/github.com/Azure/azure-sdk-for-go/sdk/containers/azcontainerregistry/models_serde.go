//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azcontainerregistry

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// UnmarshalJSON implements the json.Unmarshaller interface for type acrAccessToken.
func (a *acrAccessToken) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "access_token":
			err = unpopulate(val, "AccessToken", &a.AccessToken)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type acrRefreshToken.
func (a *acrRefreshToken) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "refresh_token":
			err = unpopulate(val, "RefreshToken", &a.RefreshToken)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactManifestPlatform.
func (a *ArtifactManifestPlatform) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "architecture":
			err = unpopulate(val, "Architecture", &a.Architecture)
			delete(rawMsg, key)
		case "digest":
			err = unpopulate(val, "Digest", &a.Digest)
			delete(rawMsg, key)
		case "os":
			err = unpopulate(val, "OperatingSystem", &a.OperatingSystem)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactManifestProperties.
func (a *ArtifactManifestProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "manifest":
			err = unpopulate(val, "Manifest", &a.Manifest)
			delete(rawMsg, key)
		case "registry":
			err = unpopulate(val, "RegistryLoginServer", &a.RegistryLoginServer)
			delete(rawMsg, key)
		case "imageName":
			err = unpopulate(val, "RepositoryName", &a.RepositoryName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactTagProperties.
func (a *ArtifactTagProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "registry":
			err = unpopulate(val, "RegistryLoginServer", &a.RegistryLoginServer)
			delete(rawMsg, key)
		case "imageName":
			err = unpopulate(val, "RepositoryName", &a.RepositoryName)
			delete(rawMsg, key)
		case "tag":
			err = unpopulate(val, "Tag", &a.Tag)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContainerRepositoryProperties.
func (c *ContainerRepositoryProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "changeableAttributes":
			err = unpopulate(val, "ChangeableAttributes", &c.ChangeableAttributes)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateDateTimeRFC3339(val, "CreatedOn", &c.CreatedOn)
			delete(rawMsg, key)
		case "lastUpdateTime":
			err = unpopulateDateTimeRFC3339(val, "LastUpdatedOn", &c.LastUpdatedOn)
			delete(rawMsg, key)
		case "manifestCount":
			err = unpopulate(val, "ManifestCount", &c.ManifestCount)
			delete(rawMsg, key)
		case "imageName":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "registry":
			err = unpopulate(val, "RegistryLoginServer", &c.RegistryLoginServer)
			delete(rawMsg, key)
		case "tagCount":
			err = unpopulate(val, "TagCount", &c.TagCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManifestAttributes.
func (m *ManifestAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "architecture":
			err = unpopulate(val, "Architecture", &m.Architecture)
			delete(rawMsg, key)
		case "changeableAttributes":
			err = unpopulate(val, "ChangeableAttributes", &m.ChangeableAttributes)
			delete(rawMsg, key)
		case "configMediaType":
			err = unpopulate(val, "ConfigMediaType", &m.ConfigMediaType)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateDateTimeRFC3339(val, "CreatedOn", &m.CreatedOn)
			delete(rawMsg, key)
		case "digest":
			err = unpopulate(val, "Digest", &m.Digest)
			delete(rawMsg, key)
		case "lastUpdateTime":
			err = unpopulateDateTimeRFC3339(val, "LastUpdatedOn", &m.LastUpdatedOn)
			delete(rawMsg, key)
		case "mediaType":
			err = unpopulate(val, "MediaType", &m.MediaType)
			delete(rawMsg, key)
		case "os":
			err = unpopulate(val, "OperatingSystem", &m.OperatingSystem)
			delete(rawMsg, key)
		case "references":
			err = unpopulate(val, "RelatedArtifacts", &m.RelatedArtifacts)
			delete(rawMsg, key)
		case "imageSize":
			err = unpopulate(val, "Size", &m.Size)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &m.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ManifestWriteableProperties.
func (m ManifestWriteableProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "deleteEnabled", m.CanDelete)
	populate(objectMap, "listEnabled", m.CanList)
	populate(objectMap, "readEnabled", m.CanRead)
	populate(objectMap, "writeEnabled", m.CanWrite)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ManifestWriteableProperties.
func (m *ManifestWriteableProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deleteEnabled":
			err = unpopulate(val, "CanDelete", &m.CanDelete)
			delete(rawMsg, key)
		case "listEnabled":
			err = unpopulate(val, "CanList", &m.CanList)
			delete(rawMsg, key)
		case "readEnabled":
			err = unpopulate(val, "CanRead", &m.CanRead)
			delete(rawMsg, key)
		case "writeEnabled":
			err = unpopulate(val, "CanWrite", &m.CanWrite)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Manifests.
func (m *Manifests) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "manifests":
			err = unpopulate(val, "Attributes", &m.Attributes)
			delete(rawMsg, key)
		case "link":
			err = unpopulate(val, "Link", &m.Link)
			delete(rawMsg, key)
		case "registry":
			err = unpopulate(val, "RegistryLoginServer", &m.RegistryLoginServer)
			delete(rawMsg, key)
		case "imageName":
			err = unpopulate(val, "Repository", &m.Repository)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Repositories.
func (r *Repositories) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "link":
			err = unpopulate(val, "Link", &r.Link)
			delete(rawMsg, key)
		case "repositories":
			err = unpopulate(val, "Names", &r.Names)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RepositoryWriteableProperties.
func (r RepositoryWriteableProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "deleteEnabled", r.CanDelete)
	populate(objectMap, "listEnabled", r.CanList)
	populate(objectMap, "readEnabled", r.CanRead)
	populate(objectMap, "writeEnabled", r.CanWrite)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RepositoryWriteableProperties.
func (r *RepositoryWriteableProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deleteEnabled":
			err = unpopulate(val, "CanDelete", &r.CanDelete)
			delete(rawMsg, key)
		case "listEnabled":
			err = unpopulate(val, "CanList", &r.CanList)
			delete(rawMsg, key)
		case "readEnabled":
			err = unpopulate(val, "CanRead", &r.CanRead)
			delete(rawMsg, key)
		case "writeEnabled":
			err = unpopulate(val, "CanWrite", &r.CanWrite)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TagAttributes.
func (t *TagAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "changeableAttributes":
			err = unpopulate(val, "ChangeableAttributes", &t.ChangeableAttributes)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateDateTimeRFC3339(val, "CreatedOn", &t.CreatedOn)
			delete(rawMsg, key)
		case "digest":
			err = unpopulate(val, "Digest", &t.Digest)
			delete(rawMsg, key)
		case "lastUpdateTime":
			err = unpopulateDateTimeRFC3339(val, "LastUpdatedOn", &t.LastUpdatedOn)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TagList.
func (t *TagList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "link":
			err = unpopulate(val, "Link", &t.Link)
			delete(rawMsg, key)
		case "registry":
			err = unpopulate(val, "RegistryLoginServer", &t.RegistryLoginServer)
			delete(rawMsg, key)
		case "imageName":
			err = unpopulate(val, "Repository", &t.Repository)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &t.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagWriteableProperties.
func (t TagWriteableProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "deleteEnabled", t.CanDelete)
	populate(objectMap, "listEnabled", t.CanList)
	populate(objectMap, "readEnabled", t.CanRead)
	populate(objectMap, "writeEnabled", t.CanWrite)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TagWriteableProperties.
func (t *TagWriteableProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deleteEnabled":
			err = unpopulate(val, "CanDelete", &t.CanDelete)
			delete(rawMsg, key)
		case "listEnabled":
			err = unpopulate(val, "CanList", &t.CanList)
			delete(rawMsg, key)
		case "readEnabled":
			err = unpopulate(val, "CanRead", &t.CanRead)
			delete(rawMsg, key)
		case "writeEnabled":
			err = unpopulate(val, "CanWrite", &t.CanWrite)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}