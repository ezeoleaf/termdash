// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package testdraw provides helpers for tests that use the draw package.
package testdraw

import (
	"fmt"
	"image"

	"github.com/mum4k/termdash/canvas"
	"github.com/mum4k/termdash/draw"
)

// MustBorder draws border on the canvas or panics.
func MustBorder(c *canvas.Canvas, border image.Rectangle, opts ...draw.BorderOption) {
	if err := draw.Border(c, border, opts...); err != nil {
		panic(fmt.Sprintf("draw.Border => unexpected error: %v", err))
	}
}

// MustText draws the text on the canvas or panics.
func MustText(c *canvas.Canvas, text string, start image.Point, opts ...draw.TextOption) {
	if err := draw.Text(c, text, start, opts...); err != nil {
		panic(fmt.Sprintf("draw.Text => unexpected error: %v", err))
	}
}

// MustRectangle draws the rectangle on the canvas or panics.
func MustRectangle(c *canvas.Canvas, r image.Rectangle, opts ...draw.RectangleOption) {
	if err := draw.Rectangle(c, r, opts...); err != nil {
		panic(fmt.Sprintf("draw.Rectangle => unexpected error: %v", err))
	}
}
