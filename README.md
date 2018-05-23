# truing-bot

使用[tuling123](http://www.tuling123.com)的接口实现的图灵聊天机器人

## how to ues

先去tuling123注册账号，然后在个人页面找到API KEY，添加到`turing/config.go`的APIKEY里

go run main.go

## example

Input

    message := turing.NewTuringMessage("北京天气")
	messageResp := turing.GetTuringBotResp(message)
	fmt.Println(messageResp.Result[0].Values["text"])

Output

    北京:周六 05月19日,多云转小雨 南风微风,最低气温16度，最高气温27度


