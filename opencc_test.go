package opencc

import (
	"context"
	"testing"
)

func TestConvert(t *testing.T) {
	var testcase = []string{
		"说起来你可能不信,我是考试考进来的",
		"说起来你可能不信,我是花钱找关系进来的",
		"在中国,资讯类移动应用的人均阅读时长是 5 分钟,而在知乎日报,这个数字是 21",
		"我开挖掘机拆屋的时候听特别带感",
		"1990年真实记录，当时的秋名山的日常",
		"1990年藤原豆腐店成了连锁店 没错这些车都是送豆腐的",
		"Go语言,从底层到应用，视Golang的环境搭建、基础知识、进阶知识、项目实践、Redis基础及其项目实践（海量用户通讯系统）、算法与数据结构基础知识的golang实现。",
	}
	Init()
	for _, c := range testcase {
		str, err := Convert(context.Background(), c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("\n%s:\t%s\n", c, str)
	}
}

func BenchmarkConvert(b *testing.B) {
	var testcase = []string{
		"说起来你可能不信,我是考试考进来的",
		"说起来你可能不信,我是花钱找关系进来的",
		"在中国,资讯类移动应用的人均阅读时长是 5 分钟,而在知乎日报,这个数字是 21",
		"我开挖掘机拆屋的时候听特别带感",
		"1990年真实记录，当时的秋名山的日常",
		"1990年藤原豆腐店成了连锁店 没错这些车都是送豆腐的",
		"Go语言,从底层到应用，视Golang的环境搭建、基础知识、进阶知识、项目实践、Redis基础及其项目实践（海量用户通讯系统）、算法与数据结构基础知识的golang实现。",
	}
	Init()
	for i := 0; i < b.N; i++ {
		out, err := Convert(context.Background(), testcase[i%len(testcase)])
		if err != nil {
			b.Fatal(err)
		}
		b.Logf("in:%s,out:%s", testcase[i%len(testcase)], out)
	}
}
