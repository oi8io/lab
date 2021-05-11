package lru

import (
	"fmt"
	"reflect"
	"testing"
)

type String string

func (s String) Len() int64 {
	return int64(len(s))
}

func TestCache_Get(t *testing.T) {
	c := NewCache(0, nil)
	c.Add("one", String("<--------value-------->"))
	get, ok := c.Get("one")
	if !ok {
		t.Fatal("not ok")
	} else {
		fmt.Println(get.(String))
	}
}

func TestCache_add(t *testing.T) {
	c := NewCache(0, nil)
	c.Add("one", String("<--------1-------->"))
	c.Add("two", String("<--------2-------->"))
	c.Add("three", String("<--------3-------->"))
	get, ok := c.Get("one")
	if !ok {
		t.Fatal("not ok")
	} else {
		fmt.Println(get.(String))
	}
}

func TestCache_Del(t *testing.T) {
	c := NewCache(0, nil)
	c.Add("one", String("<--------1-------->"))
	get, ok := c.Get("one")
	if !ok {
		t.Fatal("not ok")
	} else {
		fmt.Println(get.(String))
	}
	c.Add("one", String("<--------x-------->"))
	c.Add("two", String("<--------2-------->"))
	c.Add("three", String("<--------3-------->"))
	get, ok = c.Get("one")
	if !ok {
		t.Fatal("not ok")
	} else {
		fmt.Println(get.(String))
	}

	c.Del("one")
	get, ok = c.Get("one")
	if !ok {
		t.Fatal("one is not ok")
	} else {
		fmt.Println(get.(String))
	}

}
func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		fmt.Println(key, "removed")
		keys = append(keys, key)
	}
	c := NewCache(int64(10), callback)
	c.Add("key1", String("123456"))
	c.Add("k2", String("k2"))
	c.Add("k3", String("k3"))
	c.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	c := NewCache(int64(cap), nil)
	c.Add(k1, String(v1))
	c.Add(k2, String(v2))
	c.Add(k3, String(v3))

	if _, ok := c.Get("key1"); ok || c.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}
