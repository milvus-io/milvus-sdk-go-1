// Copyright (C) 2019-2021 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

// Package client provides milvus client functions
package client

import (
	"context"

	"github.com/milvus-io/milvus-sdk-go/entity"
	"google.golang.org/grpc"
)

// Client is the interface used to communicate with Milvus
type Client interface {
	// Connect connect to the address provided
	Connect(ctx context.Context, addr string) error
	// Close close the remaining connection
	Close() error

	// -- collection --

	// ListCollections list collections from connection
	ListCollections(ctx context.Context) ([]*entity.Collection, error)
	// CreateCollection create collection using provided schema
	CreateCollection(ctx context.Context, schema entity.Schema, shardsNum int32) error
	// DescribeCollection describe collection meta
	DescribeCollection(ctx context.Context, collName string) (*entity.Collection, error)
	// DropCollection drop the specified collection
	DropCollection(ctx context.Context, collName string) error
	// GetCollectionStatistics get collection statistics
	GetCollectionStatistics(ctx context.Context, collName string) (map[string]string, error)
	// LoadCollection load collection into memory
	LoadCollection(ctx context.Context, collName string, async bool) error
	// ReleaseCollection release loaded collection
	ReleaseCollection(ctx context.Context, collName string) error
	// HasCollection check whether collection exists
	HasCollection(ctx context.Context, collName string) (bool, error)

	// -- partition --

	// CreatePartition create partition for collection
	CreatePartition(ctx context.Context, collName string, partitionName string) error
	// DropPartition drop partition from collection
	DropPartition(ctx context.Context, collName string, partitionName string) error
	// ShowParitions list all partitions from collection
	ShowPartitions(ctx context.Context, collName string) ([]*entity.Partition, error)
	// HasPartition check whether partition exists in collection
	HasPartition(ctx context.Context, collName string, partitionName string) (bool, error)

	// -- index --

	// CreateIndex create index for field of specified collection
	// currently index naming is not supported, so only one index on vector field is supported
	CreateIndex(ctx context.Context, collName string, fieldName string, idx entity.Index) error
	// DescribeIndex describe index on collection
	// currently index naming is not supported, so only one index on vector field is supported
	DescribeIndex(ctx context.Context, collName string) ([]entity.Index, error)
	// DropINdex drop index from collection with specified field name
	DropIndex(ctx context.Context, collName string, fieldName string) error
	// GetIndexState get index state with specified collection and field name
	// index naming is not supported yet
	GetIndexState(ctx context.Context, collName string, fieldName string) (entity.IndexState, error)

	// -- basic operation --

	// Insert column-based data into collection, returns id column values
	Insert(ctx context.Context, collName string, partitionName string, columns ...entity.Column) (entity.Column, error)
	// Flush flush collection, specified
	Flush(ctx context.Context, collName string, async bool) error
	// Search search with bool expression
	Search(ctx context.Context, collName string, partitions []string,
		expr string, outputFields []string, vectors []entity.Vector, vectorField string, metricType entity.MetricType, topK int, params map[string]string) ([]SearchResult, error)
}

// SearchResult search result
type SearchResult struct {
	ResultCount int
	IDs         entity.Column
	Fields      []entity.Column
	Scores      []float32
	Err         error
}

// alias type for context field key
type grpcKey int

const (
	dialOption grpcKey = 1
)

// NewGrpcClient create client with grpc addr
func NewGrpcClient(ctx context.Context, addr string, dialOptions ...grpc.DialOption) (Client, error) {
	c := &grpcClient{}
	// since different client may have different type of connect option(s), it's hard to put concrete type in Connect method def
	// interface{} is ugly, so use context value may be a solution
	ctx = context.WithValue(ctx, dialOption, dialOptions)
	err := c.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}
	return c, nil
}
