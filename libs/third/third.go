// 对接第三方系统
package third

import (
	"transformer_mz/libs/config"
)

var AI_SYSTEM_URL string

type ConfigWrap struct {
	URLink string `yaml:"url_link"`
}

type Config struct {
	ConfigWrap `yaml:"third_party"`
}

// Init Third Party 初始化第三方URL ，本系统里面对应AI 算法连接
func InitThirdParty() error {
	thirdPartyConfig := &Config{}
	err := config.Load(thirdPartyConfig)
	if err != nil {
		return err
	}
	AI_SYSTEM_URL = thirdPartyConfig.URLink
	return nil
}
