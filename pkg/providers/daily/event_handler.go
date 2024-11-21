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

package daily

import (
	"log/slog"

	"github.com/notedownorg/notedown/pkg/fileserver/reader"
	"github.com/notedownorg/notedown/pkg/providers/pkg/traits"
)

func onLoad(c *DailyClient) traits.EventHandler {
	return func(event reader.Event) {
		if c.handleChanges(event) {
			c.publisher.Events <- Event{Op: Load}
		}
	}
}

func onChange(c *DailyClient) traits.EventHandler {
	return func(event reader.Event) {
		if c.handleChanges(event) {
			c.publisher.Events <- Event{Op: Change}
		}
	}
}

func onDelete(c *DailyClient) traits.EventHandler {
	return func(event reader.Event) {
		c.notesMutex.Lock()
		delete(c.notes, event.Key)
		c.notesMutex.Unlock()
		c.publisher.Events <- Event{Op: Delete}
		slog.Debug("removed daily note", "path", event.Key)
	}
}

func (c *DailyClient) handleChanges(event reader.Event) bool {
	if event.Document.Metadata.Type() != MetadataKey {
		return false
	}
	c.notesMutex.Lock()
	c.notes[event.Key] = NewDaily(NewIdentifier(event.Key, event.Document.Checksum))
	c.notesMutex.Unlock()
	slog.Debug("updated daily note in cache", "path", event.Key)
	return true
}
