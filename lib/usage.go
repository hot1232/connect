package lib

import (
	"fmt"
	"connect/color"
	"connect/conf"
)

func Usage() string {
	template := `
      0x0) 输入 %s 直接登陆 或 输入 %s 进行搜索登录(如果唯一)
      0x1) 输入 %s + %s 搜索. 如: /ip
      0x1) 输入 %s/%s 列出所有主机
      0x2) 输入 %s/%s 按主机组搜索主机
      0x3) 输入 %s/%s 按应用搜索主机
      0x4) 输入 %s/%s 选择主机
      0x5) 输入 %s/%s 帮助
      0x6) 输入 %s/%s 退出

`
	str := fmt.Sprintf(template,
		color.Green("ID"),
		color.Green("部分IP,主机名,备注"),
		color.Green("／"),
		color.Green("IP, 主机名 or 备注 or 主机组 or 项目名"),
		color.Green("P"),
		color.Green("p"),
		color.Green("G"),
		color.Green("g"),
		color.Green("A"),
		color.Green("a"),
		color.Green("ip"),
		color.Green("hostname"),
		color.Green("H"),
		color.Green("h"),
		color.Green("q"),
		color.Green("exit"),
	)

	return str
}

func Title() string {
	return "###      "+color.Title(conf.Setting.LoginTitle)+"      ###\n"
}

func Prompt(str string) string {
	//\x1b[?25h 是显示光标
	//\x1b[?25h 是显示光标
	switch str {
		case "option":
			return color.Title(fmt.Sprintf("%s >:\x1b[?25h",str))
		default:
			return color.Title("Opt or ID >:\x1b[?25h")
	}
}
