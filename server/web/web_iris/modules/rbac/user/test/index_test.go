package test

import (
	"fmt"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/tests"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

var (
	loginUrl  = "/api/v1/auth/login"
	logoutUrl = "/api/v1/users/logout"
	url       = "/api/v1/users"
)

func TestList(t *testing.T) {
	if TestServer == nil {
		t.Errorf("TestServer is nil")
	}
	client := TestServer.GetTestLogin(t, loginUrl, nil)
	if client != nil {
		defer client.Logout(logoutUrl, nil)
	} else {
		return
	}
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "items", Value: []tests.Responses{
				{
					{Key: "id", Value: 1, Type: "ge"},
					{Key: "name", Value: "超级管理员"},
					{Key: "username", Value: "admin"},
					{Key: "intro", Value: "超级管理员"},
					{Key: "avatar", Value: str.Join("http://", web_iris.CONFIG.System.Addr, web_iris.CONFIG.System.StaticPrefix, "/images/avatar.jpg")},
					{Key: "roles", Value: []string{"超级管理员"}},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	client.GET(url, pageKeys, tests.RequestParams)
}

func TestCreate(t *testing.T) {
	client := TestServer.GetTestLogin(t, loginUrl, nil)
	if client != nil {
		defer client.Logout(logoutUrl, nil)
	} else {
		return
	}
	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "create_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)
}

func TestUpdate(t *testing.T) {
	client := TestServer.GetTestLogin(t, loginUrl, nil)
	if client != nil {
		defer client.Logout(logoutUrl, nil)
	} else {
		return
	}
	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "update_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)

	update := map[string]interface{}{
		"name":     "更新测试名称",
		"username": "update_test_username",
		"intro":    "更新测试描述信息",
		"avatar":   "",
		"password": "123456",
	}

	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
	}
	client.POST(fmt.Sprintf("%s/%d", url, id), pageKeys, update)
}

func TestGetById(t *testing.T) {
	client := TestServer.GetTestLogin(t, loginUrl, nil)
	if client != nil {
		defer client.Logout(logoutUrl, nil)
	} else {
		return
	}
	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "getbyid_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)

	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "id", Value: 1, Type: "ge"},
			{Key: "name", Value: data["name"].(string)},
			{Key: "username", Value: data["username"].(string)},
			{Key: "intro", Value: data["intro"].(string)},
			{Key: "avatar", Value: data["avatar"].(string)},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "roles", Value: []string{}, Type: "null"},
		},
		},
	}
	client.GET(fmt.Sprintf("%s/%d", url, id), pageKeys)
}

func Create(client *tests.Client, data map[string]interface{}) uint {
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "id", Value: 1, Type: "ge"},
		},
		},
	}
	return client.POST(url, pageKeys, data).GetId()
}

func Delete(client *tests.Client, id uint) {
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
	}
	client.DELETE(fmt.Sprintf("%s/%d", url, id), pageKeys)
}