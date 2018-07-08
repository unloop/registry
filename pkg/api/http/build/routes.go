//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package build

import (
	"github.com/lastbackend/registry/pkg/util/http"
	"github.com/lastbackend/registry/pkg/util/http/middleware"
)

var Routes = []http.Route{
	// Build handlers
	{Path: "/build", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Context}, Handler: BuildCreateH},
	{Path: "/build/{build}/logs", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Context}, Handler: BuildLogsH},
	{Path: "/build/{build}/cancel", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Context}, Handler: BuildCancelH},

	{Path: "/build/task/{task}/info", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Context}, Handler: BuildTaskInfoUpdateH},
	{Path: "/build/task/{task}/status", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Context}, Handler: BuildTaskStatusUpdateH},
}
