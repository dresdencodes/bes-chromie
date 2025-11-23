package launcher

import (
    "os/exec"
)
func launchWindows(opts *LaunchOpts) error {
	cmd := exec.Command(
		`C:\Program Files\Google\Chrome\Application\chrome.exe`,
		"--remote-debugging-port="+opts.RemoteDebugPort,
		"--user-data-dir="+opts.UserDataDir,
		"--disable-popup-blocking",
		"--disable-infobars",
		"--no-first-run",
		"--no-default-browser-check",
	)
	if opts.NordVPNExtensionID != "" {
		cmd.Args = append(cmd.Args, "--load-extension="+opts.NordVPNExtensionID)
	}

	return cmd.Start()  // async, Chrome runs independently
}
