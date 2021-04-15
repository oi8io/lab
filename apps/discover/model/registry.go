package model

import (
	"oi.io/apps/discover/configs"
	"oi.io/apps/discover/errcode"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Registry struct {
	apps map[string]*Application
	lock sync.RWMutex
	gd   *Guard
}

func NewRegistry() *Registry {
	registry := &Registry{
		apps: make(map[string]*Application),
		gd:   new(Guard),
	}
	return registry
}

func (r *Registry) Register(instance *Instance, latestTimestamp int64) (*Application, *errcode.Error) {
	key := getKey(instance.AppId, instance.Env)
	r.lock.RLock()
	app, ok := r.apps[key]
	r.lock.RUnlock()
	if !ok { //new app
		app = NewApplication(instance.AppId)
	}
	//add instance
	_, isNew := app.AddInstance(instance, latestTimestamp)
	if isNew { //todo }
		//add into registry apps
		r.lock.Lock()
		r.apps[key] = app
		r.gd.incrNeed()
		r.lock.Unlock()
		return app, nil
	}
	return app, nil
}

func (r *Registry) Fetch(env, appid string, status uint32, latestTime int64) (*FetchData, *errcode.Error) {
	app, ok := r.getApplication(appid, env)
	if !ok {
		return nil, errcode.NotFound
	}
	return app.GetInstance(status, latestTime)
}
func (r *Registry) getApplication(appid, env string) (*Application, bool) {
	key := getKey(appid, env)
	r.lock.RLock()
	app, ok := r.apps[key]
	r.lock.RUnlock()
	return app, ok
}

func (r *Registry) Cancel(env, appid, hostname string, latestTimestamp int64) (*Instance, *errcode.Error) {
	log.Println("action cancel...")
	//find app
	app, ok := r.getApplication(appid, env)
	if !ok {
		return nil, errcode.NotFound
	}
	instance, ok, insLen := app.Cancel(hostname, latestTimestamp)
	if !ok {
		return nil, errcode.NotFound
	}
	//if instances is empty, delete app from apps
	if insLen == 0 {
		r.lock.Lock()
		delete(r.apps, getKey(appid, env))
		r.lock.Unlock()
	}
	return instance, nil
}

func (r *Registry) Renew(env, appid, hostname string) (*Instance, *errcode.Error) {
	app, ok := r.getApplication(appid, env)
	if !ok {
		return nil, errcode.NotFound
	}
	in, ok := app.Renew(hostname)
	if !ok {
		return nil, errcode.NotFound
	}
	return in, nil
}

func (r *Registry) evictTask() {
	ticker := time.Tick(configs.CheckEvictInterval)
	resetTicker := time.Tick(configs.ResetGuardNeedCountInterval)

	for {
		select {
		case <-ticker:
			r.gd.storeLastCount()
			r.evict()
		case <-resetTicker:
			var count int64
			for _, app := range r.getAllApplications() {
				count += int64(app.GetInstanceLen())
			}
			r.gd.setNeed(count)
		}
	}
}

func (r *Registry) evict() {
	now := time.Now().UnixNano()
	var expiredInstances []*Instance
	apps := r.getAllApplications()
	var registryLen int
	for _, app := range apps {
		registryLen += app.GetInstanceLen()
		allInstances := app.GetAllInstances()
		for _, instance := range allInstances {
			delta := now - instance.RenewTimestamp
			if !r.gd.selfProtectStatus() && delta > int64(configs.InstanceExpireDuration) || now-instance.RenewTimestamp > int64(configs.InstanceExpireDuration) {
				expiredInstances = append(expiredInstances, instance)
			}
		}
	}
	evictionLimit := registryLen - int(float64(registryLen)*configs.SelfProtectThreshold)
	expiredLen := len(expiredInstances)
	if expiredLen > evictionLimit {
		expiredLen = evictionLimit
	}

	if expiredLen == 0 {
		return
	}
	for i := 0; i < expiredLen; i++ {
		j := i + rand.Intn(len(expiredInstances)-i)
		expiredInstances[i], expiredInstances[j] = expiredInstances[j], expiredInstances[i]
		expiredInstance := expiredInstances[i]
		r.Cancel(expiredInstance.Env, expiredInstance.AppId, expiredInstance.Hostname, now)
	}
}

func (r *Registry) getAllApplications() map[string]*Application {
	return r.apps
}

//
//r := NewRegistry()
//instance := NewInstance(&req)
//r.Register(instance, req.LatestTimestamp)
