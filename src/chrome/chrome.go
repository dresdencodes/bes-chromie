package chrome

import (
    "context"
    "fmt"
    "log"
    "time"

    "bes-chromie/src/chrome/launcher"

    "github.com/chromedp/cdproto/target"
    "github.com/chromedp/chromedp"
)

const remoteDebuggerURL = "http://localhost:9222/json"
const nordVPNExtID = "your_nordvpn_extension_id_here"

func Run(target string) {
    
    launch.Start(&LaunchOpts{
        TargetFolder:"./ax/chrome/"+target,
    })

}
