package biu

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"oi.io/apps/biu/biu/pb"
)

// PeerPicker() 的 PickPeer() 方法用于根据传入的 key 选择相应节点 PeerGetter。
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 接口 PeerGetter 的 \Get() 方法用于从对应 group 查找缓存值。PeerGetter 就对应于上述流程中的 HTTP 客户端。
type PeerGetter interface {
	Get(request *pb.Request, response *pb.Response) error
	Name() string
}

type httpGetter struct {
	name    string
	baseURL string
}

func (h *httpGetter) Name() string {
	return h.name
}

func (h *httpGetter) Get(request *pb.Request, response *pb.Response) error {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(request.Group),
		url.QueryEscape(request.Key),
	)
	log.Printf("start request [%s]", u)
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body error: %v", err)
	}
	if err := proto.Unmarshal(bytes, response); err != nil {
		return fmt.Errorf("proto Unmarshal body error: %v", err)
	}
	return nil
}

var _ PeerGetter = (*httpGetter)(nil)
