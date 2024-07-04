package main

import (
	"github.com/zhangyiming748/aac-audio-encoder/aac"
	"github.com/zhangyiming748/aac-audio-encoder/constant"
	"github.com/zhangyiming748/aac-audio-encoder/log"
	"github.com/zhangyiming748/aac-audio-encoder/util"
	"os"
)

func main() {
	p := new(constant.Param)
	p.SetRoot("D:\\.telegram")
	p.SetVbr("3")
	if root := os.Getenv("root"); root != "" {
		p.SetRoot(root)
	}
	if vbr := os.Getenv("vbr"); vbr != "" {
		p.SetVbr(vbr)
	}
	log.SetLog(p)
	files := util.GetAllfiles(p)
	for _, f := range files {
		aac.SingleAAC(f, p)
	}
}
