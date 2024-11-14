// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server is the main server
package server

import (
	"context"

	pbv3 "github.com/datacommonsorg/mixer/internal/proto/v3"
)

// V3Node implements API for mixer.V3Node.
func (s *Server) V3Node(ctx context.Context, in *pbv3.NodeRequest) (
	*pbv3.NodeResponse, error,
) {
	return s.dataSources.Node(ctx, in)
}