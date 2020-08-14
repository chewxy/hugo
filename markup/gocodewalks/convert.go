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

import (
	"github.com/gohugoio/hugo/identity"
	"github.com/gohugoio/hugo/markup/converter"
)

var (
	Provider converter.ProviderProvider = provider{}
	_        converter.Converter        = (*cv)(nil)
)

type provider struct{}

func (p provider) New(cfg converter.ProviderConfig) (converter.Provider, error) {
	return converter.NewProvider("gocodewalk", func(ctx converter.DocumentContext) (converter.Converter, error) {
		return &cv{
			ctx: ctx,
			cfg: cfg,
		}, nil
	}), nil

}

type cv struct {
	ctx converter.DocumentContext
	cfg converter.ProviderConfig
}

func (c *cv) Convert(ctx converter.RenderContext) (converter.Result, error) {
	return converter.Bytes(c.convert(ctx.Src)), nil
}

func (c *cv) Supports(feature identity.Identity) bool { return false } // TODO

func (c *cv) convert(bs []byte) []byte {

	switch c.ctx.Document.(type) {
	//case page.Page
	// cw := loadCodeWalk1()
	// return applyTemplate(codewalkHTML, "codewalk", cw)
	// case page bundles?
	case nil:
	}

	return nil // TODO
}
