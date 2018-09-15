package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func RunGet() {
	resp, err := http.Get("http://dict.cn/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}

func RunGetNoZip() {
	trs := &http.Transport{DisableCompression: true}
	client := http.Client{Transport: trs}
	resp, err := client.Get("http://dict.cn/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n%v\n", bytes, err)
}

func GetRunPooled() {
	// use keep alive (connect pool)
	trs := &http.Transport{
		DisableCompression:  false,
		MaxConnsPerHost:     24,
		MaxIdleConnsPerHost: 12,
		MaxIdleConns:        32,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 1800 * time.Second,
		}).DialContext,
	}
	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: trs,
	}
	var wg sync.WaitGroup
	wg.Add(10)
	log.Printf("%s\n", "Pooled get running...")
	for i := 0; i < 10; i++ {
		go runPooled(client, &wg)
	}
	wg.Wait()
	log.Printf("%s\n", "All done.")
}

func runPooled(client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := client.Get("http://dict.cn")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
	log.Println("Done.")
}

func PostPooled() {
	trs := &http.Transport{
		DisableCompression:  false,
		MaxConnsPerHost:     24,
		MaxIdleConnsPerHost: 12,
		MaxIdleConns:        32,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 1800 * time.Second,
		}).DialContext,
	}
	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: trs,
	}
	req := map[string][]string{}
	req["20180915"] = []string{"id1"}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", req)
	resp, err := client.Post("http://192.168.157.36:7600/query/bid", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
}
