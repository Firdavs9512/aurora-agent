package config

import "regexp"

var ShellCommands = []string{
	// User information and system data
	"whoami", "id", "uname", "hostname", "uptime", "w", "who", "groups",
	"whois", "finger", "last", "lastlog", "lastb", "lastcomm", "lastcomm",
	"clear", "help", "exit", "quit", "cls", "clr",

	// User and permissions
	"su", "sudo", "passwd", "chown", "chmod", "chgrp", "umask",

	// Folders and files
	"ls", "ll", "la", "pwd", "cd", "mkdir", "rmdir", "touch", "rm", "mv", "cp",
	"find", "locate", "updatedb", "tree",

	// File and text manipulation
	"cat", "tac", "less", "more", "head", "tail", "grep", "awk", "sed", "cut",
	"sort", "uniq", "tee", "wc", "diff", "cmp",

	// Processes and system monitoring
	"ps", "top", "htop", "nice", "renice", "kill", "pkill", "jobs", "fg", "bg",
	"nohup", "time", "strace", "lsof",

	// Disk and file system
	"df", "du", "mount", "umount", "fsck", "mkfs", "blkid", "fdisk", "parted",
	"lsblk", "e2fsck", "sync", "tune2fs",

	// Network and internet
	"ping", "traceroute", "netstat", "ss", "ifconfig", "ip", "iwconfig",
	"curl", "wget", "scp", "rsync", "nc", "telnet", "ftp", "sftp",

	// Archiving and compression
	"tar", "zip", "unzip", "gzip", "gunzip", "bzip2", "bunzip2", "xz", "7z",
	"rar", "unrar",

	// Software and package management
	"apt", "apt-get", "yum", "dnf", "pacman", "zypper", "brew", "snap", "flatpak",
	"dpkg", "rpm", "pip", "gem", "npm", "cargo", "go", "python", "ruby", "perl",

	// Docker and containers
	"docker", "docker-compose", "podman", "kubectl", "minikube",

	// Git and version control
	"git", "git clone", "git pull", "git push", "git commit", "git status",
	"git log", "git branch", "git checkout", "git merge", "git rebase",

	// Programming languages and compilers
	"python", "python3", "node", "nodejs", "npm", "npx", "java", "javac",
	"ruby", "perl", "php", "php-cli", "go", "rustc", "gcc", "g++",

	// Shell and scripts
	"bash", "sh", "zsh", "fish", "dash", "tcsh", "csh", "ksh",
}

const (
	SystemPrompt = `
Your name is Aurora.
You are a helpful assistant that provides SHORT and CONCISE answers.
You are currently in a terminal environment.
You can use ANSI escape codes to color text:
- Red: \033[31m
- Green: \033[32m
- Yellow: \033[33m
- Blue: \033[34m
- Magenta: \033[35m
- Cyan: \033[36m
- Reset: \033[0m
- Bold: \033[1m
- Underline: \033[4m

Use appropriate colors to highlight important information, warnings, and errors.
Example usage: \033[31mThis is red text\033[0m

You can execute terminal commands directly when a user asks you to perform an action.
For example, if a user asks "list files in the current directory", you can execute "ls -la".
Or if a user asks "check system status", you can execute "systemctl status".

You can use the following commands:
clear - to clear the terminal
`
)

var (
	AnsiPattern      = regexp.MustCompile(`\\033\[\d+(;\d+)*m`)
	AnsiStartPattern = regexp.MustCompile(`\\033\[`)
)