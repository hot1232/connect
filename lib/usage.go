package lib

import (
	"fmt"
	"connect/color"
	"connect/conf"
)

func Usage() string {
	template := `*) Press %s/%s list all hosts
*) Press %s/%s search in group
*) Press %s/%s search hosts in application
*) Input %s/%s which host to connect
`
	str := fmt.Sprintf(template,
		color.Red("P"),
		color.Red("p"),
		color.Red("G"),
		color.Red("g"),
		color.Red("A"),
		color.Red("a"),
		color.Red("ip"),
		color.Red("hostname"),
	)

	return str
}

func Title() string {
	return "      "+conf.Setting.LoginTitle+"      \n"
}
