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

package spanner

import (
	"context"
	"fmt"

	v2 "github.com/datacommonsorg/mixer/internal/proto/v2"
	"github.com/datacommonsorg/mixer/internal/server/datasource"
)

// SpannerDataSource represents a data source that interacts with Spanner.
type SpannerDataSource struct {
	client *SpannerClient
}

func NewSpannerDataSource(client *SpannerClient) *SpannerDataSource {
	return &SpannerDataSource{client: client}
}

// Type returns the type of the data source.
func (sds *SpannerDataSource) Type() datasource.DataSourceType {
	return datasource.TypeSpanner
}

// Node retrieves node data from Spanner.
func (sds *SpannerDataSource) Node(ctx context.Context, req *v2.NodeRequest) (*v2.NodeResponse, error) {
	edges, err := sds.client.GetNodeEdgesByID(ctx, req.Nodes)
	if err != nil {
		return nil, fmt.Errorf("error getting node edges: %v", err)
	}
	return nodeEdgesToNodeResponse(edges), nil
}