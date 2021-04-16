package test

import (
	"encoding/json"
	"fmt"
	"oi.io/apps/discover/model"
	"testing"
)

var req = &model.RequestRegister{AppId: "com.xx.testapp", Hostname: "myhost", Addrs: []string{"http://testapp.xx.com/myhost"}, Status: 1}

func TestRegister(t *testing.T) {
	r := model.NewRegistry()
	instance := model.NewInstance(req)
	app, _ := r.Register(instance, req.LatestTimestamp)
	t.Log(app)
}

func TestAaaB(t *testing.T) {
	r := model.NewRegistry()
	r.Cancel(req.Env, req.AppId, req.Hostname, 0)
}

func TestRe(t *testing.T) {
	x:=new(model.RequestRegister)
	marshal, err := json.Marshal(x)
	if err!=nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshal))
}
