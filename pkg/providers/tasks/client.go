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

package tasks

import (
	"sync"
	"time"

	"github.com/notedownorg/notedown/pkg/fileserver/reader"
	"github.com/notedownorg/notedown/pkg/fileserver/writer"
	"github.com/notedownorg/notedown/pkg/providers/pkg/collections"
	"github.com/notedownorg/notedown/pkg/providers/pkg/traits"
)

// Use a type alias to hide the implementation details of the traits
type watcher = traits.Watcher
type publisher = traits.Publisher[Event]

type DocumentUpdater interface {
	UpdateContent(doc writer.Document, mutations ...writer.LineMutation) error
}

type Client struct {
	*watcher
	*publisher
	writer DocumentUpdater

	// tasks maps between file paths and line numbers to tasks it should ONLY be updated in response
	// to events from the docuuments client and should otherwise be read-only.
	tasks      map[string]map[int]Task
	tasksMutex sync.RWMutex
}

type clientOptions func(*Client)

// Inform NewClient to wait for the initial load to complete before returning
func WithInitialLoadWaiter(tick time.Duration) clientOptions {
	return func(client *Client) {
		traits.WithInitialLoadWaiter(client.watcher)(tick)
	}
}

func NewClient(writer DocumentUpdater, feed <-chan reader.Event, opts ...clientOptions) *Client {
	client := &Client{
		tasks:  make(map[string]map[int]Task),
		writer: writer,
	}

	client.publisher = traits.NewPublisher[Event]()
	client.watcher = traits.NewWatcher(feed, onLoad(client), onChange(client), onDelete(client))

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *Client) Summary() int {
	tasks := 0
	c.tasksMutex.RLock()
	defer c.tasksMutex.RUnlock()
	for _, doc := range c.tasks {
		tasks += len(doc)
	}
	return tasks
}

// Opts are applied in order so filters should be applied before sorters
func (c *Client) ListTasks(fetcher collections.Fetcher[Client, Task], opts ...collections.ListOption[Task]) []Task {
	return collections.List(c, fetcher, opts...)
}
