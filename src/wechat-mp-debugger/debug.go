/* debug for wechat mp server */

package main

import (
	"encoding/xml"
	"fmt"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"
)

func debug() {
	var costTick int64
	var finishedCnt = 0

	wg := sync.WaitGroup{}
	wg.Add(times)

	costTick = time.Now().UnixNano() / 1e6
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()

			timestamp := strconv.FormatInt(time.Now().Unix(), 10)
			nonce := randStringRunes(8)
			sign := signature(timestamp, nonce, getArg(1))

			var message = makeMessage(timestamp)

			if message == nil {
				return
			}

			postURL := fmt.Sprintf("%s?signature=%s&timestamp=%s&nonce=%s", getArg(0), sign, timestamp, nonce)

			xml, err := xml.Marshal(message)

			if err != nil {
				fmt.Println(err)
				return
			}

			resp, err := send(postURL, string(xml))

			if err != nil {
				fmt.Println(err)
				return
			}

			xs, errs := prettyXML(xml)
			if errs == nil && resp != nil {
				xr, errr := prettyXML(resp)
				if errr == nil {
					tc <- struct{}{}
					finishedCnt++
					<-tc
					fmt.Println("URL:", getArg(0))
					fmt.Println("--------------------")
					fmt.Println("Send Message:")
					fmt.Println(string(xs))
					fmt.Println("--------------------")
					fmt.Println("Response:")
					fmt.Println(string(xr))
					fmt.Println("--------------------")
				} else {
					fmt.Println("URL:", getArg(0))
					fmt.Println(errr)
				}
			} else {
				fmt.Println("URL:", getArg(0))
				fmt.Println(errs)
			}
			fmt.Println()
			return
		}()
		if sleep != 0 {
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}
	wg.Wait()
	costTick = time.Now().UnixNano()/1e6 - costTick

	sec := costTick / 1e3
	costTick -= sec * 1e3
	fmt.Printf("Finished! successed count: %d, fialed count: %d, cost time: %ds%dms\n", finishedCnt, times-finishedCnt, sec, costTick)
}
