package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"github.com/spf13/cobra"
	"markdowndoc/routers"
	"markdowndoc/middleware"
	"os"
)

var envFile string

func init() {
	webCmd.Flags().StringVarP(&envFile, "config","c", ".env", "指定配置文件")
	rootCmd.AddCommand(webCmd) //添加命令
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "启动web服务",
	Long:  `启动web服务`,
	Run: func(cmd *cobra.Command, args []string) {
		const banner = `
#    # #####     #####  ####     #    # ##### #    # #      
##  ## #    #      #   #    #    #    #   #   ##  ## #      
# ## # #    #      #   #    #    ######   #   # ## # #      
#    # #    #      #   #    #    #    #   #   #    # #      
#    # #    #      #   #    #    #    #   #   #    # #      
#    # #####       #    ####     #    #   #   #    # ###### 

 Welcome to go-markdowndoc, starting application ...
`
		fmt.Println(gosupport.ToYellow(banner))

		if envFile == "" {
			envFile = ".env"
		}
		if envFile!="" {
			if gosupport.IsFile(envFile) {
				globalCfg.Set("ENV_FILE", envFile) //项目加载的env文件
			} else {
				fmt.Println(gosupport.ToRed(fmt.Sprintf("配置文件%s不存在", envFile)))
				os.Exit(0)
			}

		} else {
			fmt.Println(gosupport.ToYellow("请输入指定配置文件"))
			os.Exit(0)
		}

		commonActionInit(envFile)

		r := gin.Default()
		r.Use(middleware.LogMiddle)
		r.Use(middleware.RecoveryMiddle)
		r.Use(middleware.CommonMiddle)

		routers.RegisterRouters(r)

		r.Run(":" + globalEnv.GetString("APP_PORT"))

	},
}

