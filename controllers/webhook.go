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

package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/casdoor/casdoor/object"
	"github.com/casdoor/casdoor/util"
)

// GetWebhooks
// @Title GetWebhooks
// @Tag Webhook API
// @Description get webhooks
// @Param   owner     query    string  built-in/admin	true        "The owner of webhooks"
// @Success 200 {array} object.Webhook The Response object
// @router /get-webhooks [get]
// @Security test_apiKey
func (c *ApiController) GetWebhooks() {
	request := c.ReadRequestFromQueryParams()
	c.ContinueIfHasRightsOrDenyRequest(request)

	paginator, err := object.GetPaginator(c.Ctx, request.Owner, request.Field, request.Value, request.Limit, object.Webhook{Organization: request.Organization})
	if err != nil {
		c.ResponseDBError(err)
		return
	}

	webhooks, err := object.GetPaginationWebhooks(request.Owner, request.Organization, paginator.Offset(),
		request.Limit, request.Field, request.Value, request.SortField, request.SortOrder)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(webhooks, paginator.Nums())
}

// GetWebhook
// @Title GetWebhook
// @Tag Webhook API
// @Description get webhook
// @Param   id     query    string  built-in/admin	true        "The id ( owner/name ) of the webhook"
// @Success 200 {object} object.Webhook The Response object
// @router /get-webhook [get]
func (c *ApiController) GetWebhook() {
	request := c.ReadRequestFromQueryParams()
	c.ContinueIfHasRightsOrDenyRequest(request)

	webhook, err := object.GetWebhook(request.Id)
	c.ValidateOrganization(webhook.Organization)

	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(webhook)
}

// UpdateWebhook
// @Title UpdateWebhook
// @Tag Webhook API
// @Description update webhook
// @Param   id     query    string  built-in/admin true        "The id ( owner/name ) of the webhook"
// @Param   body    body   object.Webhook  true        "The details of the webhook"
// @Success 200 {object} controllers.Response The Response object
// @router /update-webhook [post]
func (c *ApiController) UpdateWebhook() {
	request := c.ReadRequestFromQueryParams()
	c.ContinueIfHasRightsOrDenyRequest(request)

	var webhook object.Webhook
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &webhook)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}
	c.ValidateOrganization(webhook.Organization)

	webhookFromDb, _ := object.GetWebhook(request.Id)
	if webhookFromDb == nil {
		c.Data["json"] = wrapActionResponse(false)
		c.ServeJSON()
		return
	}
	c.ValidateOrganization(webhookFromDb.Organization)
	c.validateWebhookURLs(webhook)

	c.Data["json"] = wrapActionResponse(object.UpdateWebhook(request.Id, &webhook))
	c.ServeJSON()
}

// AddWebhook
// @Title AddWebhook
// @Tag Webhook API
// @Description add webhook
// @Param   body    body   object.Webhook  true        "The details of the webhook"
// @Success 200 {object} controllers.Response The Response object
// @router /add-webhook [post]
func (c *ApiController) AddWebhook() {
	request := c.ReadRequestFromQueryParams()
	c.ContinueIfHasRightsOrDenyRequest(request)

	var webhook object.Webhook
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &webhook)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}
	c.ValidateOrganization(webhook.Organization)
	c.validateWebhookURLs(webhook)

	c.Data["json"] = wrapActionResponse(object.AddWebhook(&webhook))
	c.ServeJSON()
}

// DeleteWebhook
// @Title DeleteWebhook
// @Tag Webhook API
// @Description delete webhook
// @Param   body    body   object.Webhook  true        "The details of the webhook"
// @Success 200 {object} controllers.Response The Response object
// @router /delete-webhook [post]
func (c *ApiController) DeleteWebhook() {
	request := c.ReadRequestFromQueryParams()
	c.ContinueIfHasRightsOrDenyRequest(request)
	var webhook object.Webhook
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &webhook)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	webhookFromDb, _ := object.GetWebhook(webhook.GetId())
	if webhookFromDb == nil {
		c.ResponseBadRequest("webhook does't exist")
		return
	}
	c.ValidateOrganization(webhookFromDb.Organization)

	c.Data["json"] = wrapActionResponse(object.DeleteWebhook(&webhook))
	c.ServeJSON()
}

func (c *ApiController) validateWebhookURLs(webhook object.Webhook) {
	if webhook.Url != "" && !util.IsURLValid(webhook.Url) {
		c.ResponseError(fmt.Sprintf(c.T("general:%s field is not valid URL"), c.T("webhook:Url")))
		return
	}
}
