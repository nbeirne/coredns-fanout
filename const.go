// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fanout

import "time"

const (
	maxIPCount           = 100
	maxLoadFactor        = 100
	minLoadFactor        = 1
	policyWeightedRandom = "weighted-random"
	policySequential     = "sequential"
	maxWorkerCount       = 32
	minWorkerCount       = 2
	maxTimeout           = 2 * time.Second
	defaultTimeout       = 30 * time.Second
	readTimeout          = 2 * time.Second
	attemptDelay         = time.Millisecond * 100
	TCPTLS               = "tcp-tls" // TCPTLS Net type for a Client.
	TCP                  = "tcp"     // TCP Net type for a Client.
	UDP                  = "udp"     // UDP Net type for a Client.
)
