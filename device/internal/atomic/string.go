// "A little copying is better than a little dependency." - Rob Pike
// Original source:
// https://github.com/uber-go/atomic/blob/78a3b8ec6cb2e156b91140c7d53b16e12b7b3f91/string.go
// This file contains several modifications to the original source. All code modifications
// to this file are marked with a comment containing the string "MODIFICATION" and are
// licensed under the terms of this project's license. See project LICENSE file for details
// on the license and copyright of modifications made to this file.

// Original source license and copyright:
//
// Copyright (c) 2020-2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import "sync/atomic" // MODIFICATION: import atomic.Value

// String is an atomic type-safe wrapper for string values.
type String struct {
	_ nocmp // disallow non-atomic comparison

	v atomic.Value // MODIFICATION: use atomic.Value instead - Uber's version is a wrapper around it.
}

var _zeroString string

// MODIFICATION: remove NewString function since constructor functions
// are not in stdlib atomic types. (modified for consistency)

// Load atomically loads the wrapped string.
func (x *String) Load() string {
	if v := x.v.Load(); v != nil {
		return v.(string)
	}
	return _zeroString
}

// Store atomically stores the passed string.
func (x *String) Store(val string) {
	x.v.Store(val)
}

// CompareAndSwap is an atomic compare-and-swap for string values.
func (x *String) CompareAndSwap(old, new string) (swapped bool) {
	return x.v.CompareAndSwap(old, new)
}

// Swap atomically stores the given string and returns the old
// value.
func (x *String) Swap(val string) (old string) {
	return x.v.Swap(val).(string)
}
