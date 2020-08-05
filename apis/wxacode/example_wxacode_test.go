// Copyright 2020 FastWeGo
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

package wxacode_test

import (
	"fmt"

	"github.com/fastwego/miniprogram"
	"github.com/fastwego/miniprogram/apis/wxacode"
)

func ExampleCreateQRCode() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := wxacode.CreateQRCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := wxacode.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUnlimited() {
	var ctx *miniprogram.Miniprogram

	payload := []byte("{}")
	resp, err := wxacode.GetUnlimited(ctx, payload)

	fmt.Println(resp, err)
}
