//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2019] Last.Backend LLC
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

package types

import "time"

const KindController = "controller"

type System struct {
	AccessToken   string
	AuthServer    string
	CtrlMaster    string
	CtrlUpdated   *time.Time
	CtrlLastEvent *time.Time
	Created       time.Time
	Updated       time.Time
}

// *********************************************
// System distribution options
// *********************************************

type SystemUpdateOptions struct {
	AccessToken *string
	AuthServer  *string
}

type SystemUpdateControllerOptions struct {
	Hostname string
	Pid      int
}

type SystemUpdateControllerLastEventOptions struct {
	LastEvent time.Time
}
