package middleware

import "github.com/jellycheng/gosupport"

var (
	globalCfg *gosupport.DataManage
	globalEnv *gosupport.DataManage //存env文件配置内容
)

func init()  {
	globalEnv = gosupport.NewGlobalEnvSingleton()
	globalCfg = gosupport.NewGlobalCfgSingleton()

}
