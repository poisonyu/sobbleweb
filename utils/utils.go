package utils

import (
	"github.com/cyansobble/global"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

func GetRandonPicture() []byte {
	client := resty.New()
	// Referer:https://weibo.com/
	//resp, err := client.R().EnableTrace().Get("https://sese.iw233.top/iapi.php?sort=cdnrandom")
	resp, err := client.R().SetQueryParams(map[string]string{
		"sort": "random",
		"type": "json",
		"num":  "1",
	}).SetHeaders(map[string]string{
		"Referer": "https://weibo.com/",
		"Accept":  "application/json",
	}).Get("https://iw233.cn/api.php")
	if err != nil {
		global.LOGGER.Error("[request]:", zap.Error(err))
		return nil
	}
	global.LOGGER.Info(string(resp.Body()))
	return resp.Body()
	// if resp.StatusCode() == 200 {
	// 	return string(resp.Body())
	// }
	// return ""
}
