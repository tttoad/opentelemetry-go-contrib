// Copyright The OpenTelemetry Authors
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

//nolint:all // Linter is executed at the same time as tests which leads to race conditions and failures.
package app

import (
	"fmt"
	__atel_context "context"
	__atel_otel "go.opentelemetry.io/otel"
)

type BasicSerializer struct {
}

func (b BasicSerializer) Serialize(__atel_tracing_ctx __atel_context.Context,) {
	__atel_child_tracing_ctx, __atel_span := __atel_otel.Tracer("Serialize").Start(__atel_tracing_ctx, "Serialize")
	_ = __atel_child_tracing_ctx
	defer __atel_span.End()

	fmt.Println("Serialize")
}
