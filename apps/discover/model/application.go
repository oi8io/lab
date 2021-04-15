package model

import (
	"oi.io/apps/discover/errcode"
	"log"
	"sync"
	"time"
)

type Application struct {
	appID           string
	instances       map[string]*Instance
	latestTimestamp int64
	lock            sync.RWMutex
}

func NewApplication(appid string) *Application {
	return &Application{
		appID:     appid,
		instances: make(map[string]*Instance),
	}
}

func (app *Application) AddInstance(in *Instance, latestTimestamp int64) (*Instance, bool) {
	app.lock.Lock()
	defer app.lock.Unlock()
	appIns, ok := app.instances[in.Hostname]
	if ok { //exist
		in.UpTimestamp = appIns.UpTimestamp
		//dirtytimestamp
		if in.DirtyTimestamp < appIns.DirtyTimestamp {
			log.Println("register exist dirty timestamp")
			in = appIns
		}
	}
	//add or update instances
	app.instances[in.Hostname] = in
	app.upLatestTimestamp(latestTimestamp)
	returnIns := new(Instance)
	*returnIns = *in
	return returnIns, !ok
}

func (app *Application) GetInstance(status uint32, latestTime int64) (*FetchData, *errcode.Error) {
	app.lock.RLock()
	defer app.lock.RUnlock()
	if latestTime >= app.latestTimestamp {
		return nil, errcode.NotModified
	}
	fetchData := FetchData{
		Instances:       make([]*Instance, 0),
		LatestTimestamp: app.latestTimestamp,
	}
	var exists bool
	for _, instance := range app.instances {
		if status&instance.Status > 0 {
			exists = true
			newInstance := copyInstance(instance)
			fetchData.Instances = append(fetchData.Instances, newInstance)
		}
	}
	if !exists {
		return nil, errcode.NotFound
	}
	return &fetchData, nil
}

//deep copy
func copyInstance(src *Instance) *Instance {
	dst := new(Instance)
	*dst = *src
	//copy addrs
	dst.Addrs = make([]string, len(src.Addrs))
	for i, addr := range src.Addrs {
		dst.Addrs[i] = addr
	}
	return dst
}

func (app *Application) Cancel(hostname string, latestTimestamp int64) (*Instance, bool, int) {
	newInstance := new(Instance)
	app.lock.Lock()
	defer app.lock.Unlock()
	appIn, ok := app.instances[hostname]
	if !ok {
		return nil, ok, 0
	}
	//delete hostname
	delete(app.instances, hostname)
	appIn.LatestTimestamp = latestTimestamp
	app.upLatestTimestamp(latestTimestamp)
	*newInstance = *appIn
	return newInstance, true, len(app.instances)
}

func (app *Application) Renew(hostname string) (*Instance, bool) {
	app.lock.Lock()
	defer app.lock.Unlock()
	appIn, ok := app.instances[hostname]
	if !ok {
		return nil, ok
	}
	appIn.RenewTimestamp = time.Now().UnixNano()
	return copyInstance(appIn), true
}

func (app *Application) upLatestTimestamp(latestTimestamp int64) {
	app.latestTimestamp = latestTimestamp
}

func (app *Application) GetInstanceLen() int {
	return len(app.instances)
}

func (app *Application) GetAllInstances() map[string]*Instance {
	return app.instances
}
