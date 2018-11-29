# opencc - Golang version OpenCC

## Introduction 介紹

Open Chinese Convert (OpenCC, 開放中文轉換) is an opensource project for conversion between Traditional Chinese and Simplified Chinese, supporting character-level conversion, phrase-level conversion, variant conversion and regional idioms among Mainland China, Taiwan and Hong kong.

中文簡繁轉換開源項目，支持詞彙級別的轉換、異體字轉換和地區習慣用詞轉換（中國大陸、臺灣、香港）。


### Links 相關鏈接

* Introduction 詳細介紹 https://github.com/BYVoid/OpenCC/wiki/%E7%B7%A3%E7%94%B1
* OpenCC Online (在線轉換) http://opencc.byvoid.com/
* 現代漢語常用簡繁一對多字義辨析表 http://ytenx.org/byohlyuk/KienxPyan


## Usage 使用
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/meission/opencc"
)

func main() {
    s2t, err := gocc.New("s2t")
    if err != nil {
        log.Fatal(err)
    }
    in := "说起来你可能不信,我是考试考进来的"
    out, err := s2t.Convert(in)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
}
```


#### 預設配置文件

* `s2t.json` Simplified Chinese to Traditional Chinese 簡體到繁體
* `t2s.json` Traditional Chinese to Simplified Chinese 繁體到簡體
* `s2tw.json` Simplified Chinese to Traditional Chinese (Taiwan Standard) 簡體到臺灣正體
* `tw2s.json` Traditional Chinese (Taiwan Standard) to Simplified Chinese 臺灣正體到簡體
* `s2hk.json` Simplified Chinese to Traditional Chinese (Hong Kong Standard) 簡體到香港繁體（香港小學學習字詞表標準）
* `hk2s.json` Traditional Chinese (Hong Kong Standard) to Simplified Chinese 香港繁體（香港小學學習字詞表標準）到簡體
* `s2twp.json` Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom 簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙
* `tw2sp.json` Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom 繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙
* `t2tw.json` Traditional Chinese (OpenCC Standard) to Taiwan Standard 繁體（OpenCC 標準）到臺灣正體
* `t2hk.json` Traditional Chinese (OpenCC Standard) to Hong Kong Standard 繁體（OpenCC 標準）到香港繁體（香港小學學習字詞表標準）