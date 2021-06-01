package registry

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"net/http"
	"oi.io/apps/zrpc/balance"
	"strings"
	"sync"
	"time"
)

type Discovery struct {
	registry   string
	timeout    time.Duration
	lastUpdate time.Time
	r          *rand.Rand   // generate random number
	mu         sync.RWMutex // protect following
	servers    []string
	index      int // record the selected position for robin algorithm
}

const defaultUpdateTimeout = time.Second * 10

func NewDiscovery(registerAddr string, timeout time.Duration) *Discovery {
	if timeout == 0 {
		timeout = defaultUpdateTimeout
	}

	d := &Discovery{
		servers:  make([]string, 0),
		registry: registerAddr,
		timeout:  timeout,
		r:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	d.index = d.r.Intn(math.MaxInt32 - 1)
	return d
}

func (d *Discovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	d.lastUpdate = time.Now()
	return nil
}

func (d *Discovery) Refresh() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.lastUpdate.Add(d.timeout).After(time.Now()) {
		return nil
	}
	log.Println("rpc registry: refresh servers from registry", d.registry)
	resp, err := http.Get(d.registry)
	if err != nil {
		log.Println("rpc registry refresh err:", err)
		return err
	}
	servers := strings.Split(resp.Header.Get("X-rpc-Servers"), ",")
	d.servers = make([]string, 0, len(servers))
	for _, server := range servers {
		if strings.TrimSpace(server) != "" {
			d.servers = append(d.servers, strings.TrimSpace(server))
		}
	}
	d.lastUpdate = time.Now()
	return nil
}

func (d *Discovery) Get(mode balance.SelectMode) (string, error) {
	if err := d.Refresh(); err != nil {
		return "", err
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	switch mode {
	case balance.RandomSelect:
		return d.servers[d.r.Intn(n)], nil
	case balance.RoundRobinSelect:
		s := d.servers[d.index%n] // servers could be updated, so mode n to ensure safety
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

func (d *Discovery) GetAll() ([]string, error) {
	if err := d.Refresh(); err != nil {
		return nil, err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	// return a copy of d.servers
	servers := make([]string, len(d.servers), len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}
