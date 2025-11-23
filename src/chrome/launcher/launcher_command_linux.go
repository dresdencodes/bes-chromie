//go:build linux || darwin
package launcher

import (
    "fmt"
    "os/exec"
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

    fmt.Println("Launching Chrome (Linux/macOS detached)...")

    cmd := exec.Command("google-chrome", args...)
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Setsid: true,
    }

    cmd.Stdout = nil
    cmd.Stderr = nil

    return cmd.Start()
}
