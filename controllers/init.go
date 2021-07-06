package controllers

import (
	"github.com/jellycheng/gosupport"
)

var (
	globalCfg *gosupport.DataManage
	globalEnv *gosupport.DataManage
)


func init()  {
	globalCfg = gosupport.NewGlobalCfgSingleton()
	globalEnv = gosupport.NewGlobalEnvSingleton()

}
