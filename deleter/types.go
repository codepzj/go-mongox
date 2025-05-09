// Generated by [optioner] command-line tool; DO NOT EDIT
// If you have any questions, please create issues and submit contributions at:
// https://github.com/chenmingyong0423/go-optioner

// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deleter

import (
	"context"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

//go:generate optioner -type OpContext -output types.go -mode append
type OpContext struct {
	Col          *mongo.Collection `opt:"-"`
	Filter       any               `opt:"-"`
	MongoOptions any
	ModelHook    any
	StartTime    time.Time

	Fields []*field.Filed

	// result of the collection operation
	Result any
}

type (
	BeforeHookFn func(ctx context.Context, opContext *OpContext, opts ...any) error
	AfterHookFn  func(ctx context.Context, opContext *OpContext, opts ...any) error
)

type OpContextOption func(*OpContext)

func NewOpContext(col *mongo.Collection, filter any, opts ...OpContextOption) *OpContext {
	opContext := &OpContext{
		Col:    col,
		Filter: filter,
	}

	for _, opt := range opts {
		opt(opContext)
	}

	return opContext
}

func WithMongoOptions(mongoOptions any) OpContextOption {
	return func(opContext *OpContext) {
		opContext.MongoOptions = mongoOptions
	}
}

func WithModelHook(modelHook any) OpContextOption {
	return func(opContext *OpContext) {
		opContext.ModelHook = modelHook
	}
}

func WithStartTime(startTime time.Time) OpContextOption {
	return func(opContext *OpContext) {
		opContext.StartTime = startTime
	}
}

func WithFields(fields []*field.Filed) OpContextOption {
	return func(opContext *OpContext) {
		opContext.Fields = fields
	}
}

func WithResult(result any) OpContextOption {
	return func(opContext *OpContext) {
		opContext.Result = result
	}
}
