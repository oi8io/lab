package consistent

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestHashMap_Add(t *testing.T) {
	hashMap := NewHashMap(2, nil)
	hashMap.Add("n1", "n2", "n3", "n4", "n5")
	fmt.Println(hashMap.hashMap)
	fmt.Println(hashMap.keys)
}

func TestHashMap_Get(t *testing.T) {
	hashMap := NewHashMap(2, nil)
	hashMap.Add("n1", "n2", "n3", "n4", "n5")
	get, err := hashMap.Get("9")
	fmt.Println(get, err)

}

func TestHashing(t *testing.T) {
	hash := NewHashMap(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if get, _ := hash.Get(k); get != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// Adds 8, 18, 28
	hash.Add("8")

	// 27 should now map to 8.
	testCases["27"] = "8"

	for k, v := range testCases {
		if get, _ := hash.Get(k); get != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}

func TestHashMap_Add1(t *testing.T) {
	type fields struct {
		hashFunc hashFunc
		replicas int
		keys     []int
		hashMap  map[int]string
	}
	type args struct {
		nodeName []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &HashMap{
				hashFunc: tt.fields.hashFunc,
				replicas: tt.fields.replicas,
				keys:     tt.fields.keys,
				hashMap:  tt.fields.hashMap,
			}
			fmt.Println(m)
		})
	}
}

func TestHashMap_Get1(t *testing.T) {
	type fields struct {
		hashFunc hashFunc
		replicas int
		keys     []int
		hashMap  map[int]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantNodeName string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &HashMap{
				hashFunc: tt.fields.hashFunc,
				replicas: tt.fields.replicas,
				keys:     tt.fields.keys,
				hashMap:  tt.fields.hashMap,
			}
			gotNodeName, err := m.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNodeName != tt.wantNodeName {
				t.Errorf("Get() gotNodeName = %v, want %v", gotNodeName, tt.wantNodeName)
			}
		})
	}
}

func TestNewHashMap(t *testing.T) {
	type args struct {
		replicas int
		hashFunc hashFunc
	}
	var hashfunc = func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	}
	hash := NewHashMap(3, hashfunc)
	tests := []struct {
		name string
		args args
		want *HashMap
	}{
		{name: "add1", args: args{
			replicas: 3,
			hashFunc: hashfunc,
		}, want: hash},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMap(tt.args.replicas, tt.args.hashFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
