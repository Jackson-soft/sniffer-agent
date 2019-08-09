package mysql

import (
	"flag"
	"fmt"
	"regexp"
)

var (
	uselessSQLPattern = regexp.MustCompile(`(?i)^\s*(select 1|select @@version_comment limit 1|`+
		`SELECT user, db FROM information_schema.processlist WHERE host=|commit|begin)`)
)

var (
	strictMode bool
	adminUser string
	adminPasswd string
)

func init() {
	flag.BoolVar(&strictMode,"strict_mode", false, "strict mode. Default is false")
	flag.StringVar(&adminUser,"admin_user", "", "admin user name. When set strict mode, must set admin user to query session info")
	flag.StringVar(&adminPasswd,"admin_passwd", "", "admin user passwd. When use strict mode, must set admin user to query session info")
}

func CheckParams()  {
	if !strictMode {
		return
	}

	if len(adminUser) < 1 {
		panic(fmt.Sprintf("In strict mode, admin user name cannot be empty"))
	}

	if len(adminPasswd) < 1 {
		panic(fmt.Sprintf("In strict mode, admin passwd cannot be empty"))
	}
}