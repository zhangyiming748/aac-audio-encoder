package aac

import (
	"github.com/zhangyiming748/aac-audio-encoder/constant"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

/*
* 将mp3文件转为aac文件
接收源文件的绝对路径
*/
func SingleAAC(fp string, p *constant.Param) {
	ext := path.Ext(fp)
	out := strings.Replace(fp, ext, ".aac", 1)
	vbr := p.GetVbr()
	if vbr == "" {
		vbr = "2"
	}
	cmd := exec.Command("ffmpeg", "-i", fp, "-vbr", vbr, "-ac", "1", "-map_chapters", "-1", out)
	log.Printf("开始执行命令%v\n", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("命令%v执行失败%v\n", cmd.String(), err)
		return
	} else {
		log.Printf("命令成功执行%v\n", string(output))
		os.Remove(fp)
	}
}
