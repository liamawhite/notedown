// Copyright 2024 Notedown Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package projects_test

import (
	"testing"

	"github.com/notedownorg/notedown/pkg/fileserver/reader"
	"github.com/notedownorg/notedown/pkg/fileserver/writer"
	"github.com/notedownorg/notedown/pkg/providers/projects"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	client, _ := buildClient(loadEvents(),
		// Create
		func(doc writer.Document, metadata reader.Metadata, content []byte, feed chan reader.Event) error {
			assert.Equal(t, writer.Document{Path: "projects/project.md"}, doc)
			assert.Equal(t, reader.Metadata{reader.MetadataTypeKey: projects.MetadataKey, projects.StatusKey: projects.Backlog}, metadata)
			assert.Equal(t, []byte("# project\n\n"), content)
			return nil
		},
	)

	assert.NoError(t, client.CreateProject("projects/project.md", "project", projects.Backlog))
}