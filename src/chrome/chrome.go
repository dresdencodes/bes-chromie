package chrome

import (
    "time"
    "bes-chromie/src/chrome/launcher"
)

const nordVPNExtID = "your_nordvpn_extension_id_here"

func Run(target string) {
    
    launcher.Start(&launcher.LaunchOpts{
        UserDataDir:"./ax/chrome/"+target,
    })
time.Sleep(100000 * time.Second)
}
