package osx

import (
	"os/exec"
	"runtime"
)

// OpenDefault opens the specified file or URL with the default associated application.
//
// You may use it to open a web site:
//   OpenDefault("https://google.com")
//
// Or open a file:
//   OpenDefault("/home/bob/story.txt")
//
// Or open a folder in your default file manager:
//  OpenDefault("/home/bob")
//
// For details, see https://stackoverflow.com/a/39324149/1705598
func OpenDefault(fileOrURL string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, fileOrURL)
	return exec.Command(cmd, args...).Start()
}
