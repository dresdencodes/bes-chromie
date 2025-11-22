package launcher

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
)

func launchChrome(opts *LaunchOpts) error {

	args := []string{
		"--remote-debugging-port=" + remoteDebugPort,
		"--user-data-dir=" + opts.UserDataDir,
		"--disable-popup-blocking",
		"--disable-infobars",
		"--no-first-run",
		"--no-default-browser-check",
	}

	if opts.NordVPNExtensionID != "" {
		args = append(args, "--load-extension="+opts.NordVPNExtensionID)
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		fmt.Println("Launching Chrome (Windows detached)...")
		cmd = exec.Command("chrome.exe", args...)

		// Detach from parent process
		cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP | syscall.DETACHED_PROCESS,
		}
	} else {
		// Linux/macOS
		fmt.Println("Launching Chrome (Linux/macOS detached)...")
		cmd = exec.Command("google-chrome", args...)

		// Detach process
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}
	}

	// Optional: redirect stdout/stderr to null to prevent blocking
	cmd.Stdout = nil
	cmd.Stderr = nil

	// Start the process and immediately return
	return cmd.Start()
}
