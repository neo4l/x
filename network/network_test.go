package network

import (
	"andui/conf"
	"strconv"
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	param := map[string]string{}
	param["clientId"] = conf.Slock_App_Id
	param["clientSecret"] = conf.Slock_App_Secret
	param["username"] = "sluser0003"
	param["password"] = "e10adc3949ba59abbe56e057f20f883e"
	param["date"] = strconv.FormatInt(time.Now().Unix(), 10)

	reply, err := Post("https://api.sciener.cn/v3/user/register", param)

	t.Fatalf("TestCheckAccountBalance: %s, %s", reply, err)
}
