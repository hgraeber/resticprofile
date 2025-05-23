//go:build !windows

package user

import (
	"os"
	"os/user"
	"strconv"
)

func Current() User {
	var username, homedir string
	sudoed := false
	uid := os.Getuid()
	gid := os.Getgid()
	if uid == 0 {
		// after a sudo, both os.Getuid() and os.Geteuid() return 0 (the root user)
		// to detect the logged on user after a sudo, we need to use the environment variable
		if userid, sudo := os.LookupEnv("SUDO_UID"); sudo {
			if temp, err := strconv.Atoi(userid); err == nil {
				uid = temp
				sudoed = true
			}
		}
	}
	current, err := user.LookupId(strconv.Itoa(uid))
	if err == nil {
		gid, _ = strconv.Atoi(current.Gid)
		username = current.Username
		homedir = current.HomeDir
	}
	return User{
		Uid:      uid,
		Gid:      gid,
		Username: username,
		HomeDir:  homedir,
		SudoRoot: sudoed,
	}
}
