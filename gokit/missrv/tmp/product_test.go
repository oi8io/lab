package tmp

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestGetProduct(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	config := hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 0,
		SleepWindow:            0,
		ErrorPercentThreshold:  0,
	}
	hystrix.ConfigureCommand("getProd", config)
	for {
		err := hystrix.Do("getProd", func() error {
			p, _ := GetProduct()
			fmt.Println(p)
			time.Sleep(time.Second * 1)
			return nil
		}, func(err error) error {
			fmt.Println(RecProduct())
			return errors.New("服务不可用")
		})
		if err != nil {
			fmt.Println("hystrix.Do", err)
		}
	}
}

func TestGetProductWithGo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	config := hystrix.CommandConfig{
		Timeout:                3000,
		MaxConcurrentRequests:  20,
		RequestVolumeThreshold: 5, //熔断器请求阀值，意思是20个请求才进行错误百分比计算
		SleepWindow:            10,
		ErrorPercentThreshold:  50, // 错误百分比。默认50%，超过指定值直接执行降级方法
	}
	hystrix.ConfigureCommand("getProd", config)
	circuit, _, _ := hystrix.GetCircuit("getProd")
	var prodChan = make(chan Product)
	var wg = sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		//go func() {
		wg.Add(1)
		defer wg.Done()
		err := hystrix.Go("getProd", func() error {
			p, err := GetProduct()
			if err != nil {
				return err
			}
			prodChan <- p
			time.Sleep(time.Second)
			return nil
		}, func(err error) error {
			prodChan <- RecProduct()
			return errors.New("服务不可用")
		})
		select {
		case prod := <-prodChan:
			fmt.Println(prod)
		case errs := <-err:
			fmt.Println(errs)
		}
		if circuit.IsOpen() {
			fmt.Println("触发熔断")
		}else{
			fmt.Println("没有触发熔断")
		}
		//}()
	}
	wg.Wait()
}

func TestCreateToken(t *testing.T) {
	CreateToken()
}