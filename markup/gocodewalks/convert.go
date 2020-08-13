// Copyright 2020 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gocw converts Codewalk formatted input to the given HTML.
package gocw

import "github.com/gohugoio/hugo/identity"

var (
	_ converter.ProviderProvider = provider{}
	_ converter.Converter        = (*converter)(nil)
)

type provider struct{}

func (p provider) New(cfg converter.ProviderConfig) (converter.Provider, error) {
	return converter.NewProvider("pandoc", func(ctx converter.DocumentContext) (converter.Converter, error) {
		return &converter{
			ctx: ctx,
			cfg: cfg,
		}, nil
	}), nil

}

type converter struct {
	ctx converter.DocumentContext
	cfg converter.ProviderConfig
}

func (c *converter) Convert(ctx converter.RenderContext) (converter.Result, error) {
	return converter.Bytes(c.getContent(ctx.Src, c.ctex)), nil
}

func (c *converter) Supports(feature identity.Identity) book { return false }
