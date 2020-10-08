package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
)

const (
	LogDir      = "/var/log/sebasshtian"
	KeyPathBase = "/root/keypairs/"
)

var (
	err             error
	SSHMagicVersion string
	GoVersion       string
	Environment     string
	EnvironmentIP   string
	User            string
	Key             string
	KeyPath         string = KeyPathBase + Environment + ".key"
	SSHUser         string = "cloud-user"
)

// check if FILENAME exists and is a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// SetUid to 0
func SetUid() {
	err := syscall.Setuid(0)
	if err != nil {
		os.Exit(1)
	}
}

// check length of username
func CheckUsername() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(user.Username) > 30 {
		fmt.Printf("User name too long: '%s'(%d)", user.Username, len(user.Username))
		os.Exit(23)
	}
}

func CheckSSHAgent() {
	output, err := exec.Command("/usr/bin/ssh-add", KeyPath).Output()
	if err != nil {
		fmt.Println(err)
	}
	if output == 0 {
		_ = exec.Command("/usr/bin/ssh-agent")
	}
}

func LoadSSHKey(KeyPath string) {
	_ = exec.Command("/usr/bin/ssh-add", KeyPath)
}

func main() {
	log.SetFlags(0)
	var (
		show    = flag.Bool("show", false, "Show binary configuration")
		version = flag.Bool("version", false, "Show ssh_magic version")
	)
	flag.Parse()

	if len(strings.TrimSpace(User)) != 0 {
		SSHUser = User
	}

	if len(strings.TrimSpace(Key)) != 0 {
		KeyPath = Key
	}

	if !FileExists(KeyPath) {
		fmt.Printf("'%s' not exists", KeyPath)
		os.Exit(1)
	}

	if *version {
		fmt.Println("SSH_MAGIC version:", SSHMagicVersion)
		fmt.Println("Compiled with GO:", GoVersion)
		os.Exit(0)
	}

	if *show {
		fmt.Println("Environment name:", Environment)
		fmt.Println("Bootstrap IP:", EnvironmentIP)
		fmt.Println("SSH user:", SSHUser)
		fmt.Println("SSH private key:", KeyPath)
		os.Exit(0)
	}

	CheckUsername()
	SetUid()
	CheckSSHAgent()
	SetUid()
	LoadSSHKey(KeyPath)

}
