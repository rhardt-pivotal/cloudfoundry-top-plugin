// Copyright (c) 2017 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.Space/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package space

import "github.com/ecsteam/cloudfoundry-top-plugin/metadata/common"

type SpaceMetadata struct {
	*common.Metadata
	*Space
}

func NewSpaceMetadata(Space Space) *SpaceMetadata {
	return &SpaceMetadata{Metadata: &common.Metadata{}, Space: &Space}
}

func NewSpaceMetadataById(id string) *SpaceMetadata {
	return NewSpaceMetadata(Space{EntityCommon: common.EntityCommon{Guid: id}, Name: id})
}

func (metadataItem *SpaceMetadata) GetName() string {
	return metadataItem.Name
}
