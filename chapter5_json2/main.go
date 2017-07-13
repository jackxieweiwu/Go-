// chapter5_json2 project main.go
package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var decoderMap map[string]interface{}

		if err := dec.Decode(&decoderMap); err != nil {
			log.Println(err)
			return
		}
		//		for key := range decoderMap {
		//			if key != "Title" {
		//				delete(decoderMap, key)
		//			}
		//		}

		if err := enc.Encode(&decoderMap); err != nil {
			log.Println(err)
		}
	}
}

//测试：
//{"Title":      "go语言编程","Authors":     ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan"],"Publisher":   "ituring.com.cn","IsPublished": true,"Price":       9.99}
