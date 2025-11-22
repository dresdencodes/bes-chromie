package launcher

import (
	"io"
	"fmt"
	"os/exec"
)

func launchChrome(opts *LaunchOpts) (*exec.Cmd, error) {

	// define args
	args := []string{
		"--remote-debugging-port="+remoteDebugPort,
		"--user-data-dir="+opts.UserDataDir,
		"--disable-popup-blocking",
		"--disable-infobars",
		"--no-first-run",
		"--no-default-browser-check",
	}

	
	if opts.NordVPNExtensionID != "" {
		args = append(args, "--load-extension=" + opts.NordVPNExtensionID)
	}

	// if windows
	if runtime.GOOS == "windows" {

		fmt.Println("Launching Chrome (Windows)...")

		// exec command
		cmd := exec.Command(
			"chrome.exe",
			...args,
		)

		// hide console window
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Start(), cmd
	}

	// LINUX DEFAULT
	fmt.Println("Launching Chrome (Linux)...")

	// define cmd
	cmd := exec.Command(
		"google-chrome",
		...args,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start(), cmd
}
