package params

import (
	"fmt"
	"github.com/editor/lib/utils"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"strings"
	"sync"
)

var (
	BaseDirect string
	ExecPath   = ""
	AppInfo    *Application
	Io         = utils.NewSystemOut().SetInfoType(utils.LogLevelInfo)
)

type Application struct {
	AppId      string `json:"app_id" yaml:"appid"`        // 当前运行环境
	AppEnv     string `json:"app_env" yaml:"env"`         // 当前运行环境
	AppName    string `json:"app_name" yaml:"name"`       // 应用名称
	AppVersion string `json:"app_version" yaml:"version"` // 应用版本以前缀v 开头
}

//加载配置文件
func LoadConfig(file string) (err error) {

	var syncLock sync.Mutex
	syncLock.Lock()
	defer syncLock.Unlock()
	AppInfo = &Application{}

	var yamlFile []byte
	filePath := GetConfigFilePath(file, true)
	if yamlFile, err = os.ReadFile(filePath); err != nil {
	}
	if err = yaml.Unmarshal(yamlFile, AppInfo); err != nil {
		Io.SystemOutFatalf("load app config err(%#v) \n", err)
	}
	return
}

// GetConfigFilePath 获取配置文件的路径
func GetConfigFilePath(fileName string, notEnv ...bool) (res string) {
	dir := GetConfigFileDirectory(notEnv...)
	res = getConfigFilePathContent(dir, fileName, notEnv...)
	return
}

// GetConfigFileDirectory 获得配置文件所在目录
func GetConfigFileDirectory(notEnv ...bool) (res string) {
	var (
		dir = ExecPath
		err error
	)
	env := getEnvPath()

	if BaseDirect == "" {
		if ExecPath == "" {
			if dir, err = os.Getwd(); err != nil {
				Io.SystemOutPrintf("Template GetConfigFileDirectory is :'%s'", res)
			}
		}

		if len(notEnv) > 0 && notEnv[0] {
			res = fmt.Sprintf("%s/config/", dir)
			return
		}
		res = fmt.Sprintf("%s/config/apps/%s/%s", dir, AppInfo.AppName, env)
		return

	}

	if len(notEnv) > 0 && notEnv[0] {
		res = fmt.Sprintf("%s/config/", BaseDirect)
		return
	}

	res = fmt.Sprintf("%s/config/%s", BaseDirect, env)
	return
}

func getEnvPath() (env string) {
	if AppInfo != nil && AppInfo.AppEnv != "" {
		env = AppInfo.AppEnv + "/"
	}
	return
}

func getConfigFilePathContent(dir, fileName string, notEnv ...bool) (res string) {
	res = fmt.Sprintf("%s%s", dir, fileName)
	extString := path.Ext(fileName)
	var ext string
	if extString != "" {
		ext = strings.TrimLeft(extString, ".")
	}
	switch ext {
	case "yaml":
		if ok, _ := PathExists(res); ok {
			return
		}
		res = fmt.Sprintf("%s%s.yml", dir, strings.TrimSuffix(path.Base(fileName), extString))
		return
	case "yml":
		if ok, _ := PathExists(res); ok {
			return
		}
		res = fmt.Sprintf("%s%s.yaml", dir, strings.TrimSuffix(path.Base(fileName), extString))
		return
	}

	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
