// Copyright 2021 The Casdoor Authors. All Rights Reserved.
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

package routers

import (
	"github.com/beego/beego/context"
	"github.com/casdoor/casdoor/util"
	"github.com/casdoor/casdoor/util/logger"
	"github.com/google/uuid"
)

func LoggerFilter(beegoCtx *context.Context) {
	ctx := logger.InitLoggerCtx(beegoCtx.Request.Context())
	beegoCtx.Request = beegoCtx.Request.WithContext(ctx)
	logger.SetItem(ctx, "ip", util.GetIPFromRequest(beegoCtx.Request))
	logger.SetItem(ctx, "casgate_request_id", uuid.NewString())
}