// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package pprof is a middleware that provides pprof service for Macaron.
package pprof

import (
	"net/http/pprof"
	"strings"

	"github.com/Unknwon/macaron"
)

// Pprofer is a middleware provides pprof service for your application.
func Pprofer() macaron.Handler {
	return func(ctx *macaron.Context) {
		if !strings.HasPrefix(ctx.Req.URL.Path, "/debug/pprof/") {
			return
		}

		switch ctx.Req.URL.Path[13:] {
		case "cmdline":
			pprof.Cmdline(ctx.Resp, ctx.Req)
		case "profile":
			pprof.Profile(ctx.Resp, ctx.Req)
		case "symbol":
			pprof.Symbol(ctx.Resp, ctx.Req)
		default:
			pprof.Index(ctx.Resp, ctx.Req)
		}
	}
}
