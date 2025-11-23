package chrome

import (
	"log"
	"context"

    "bes-chromie/src/chrome/launcher"
)


type ChromeInstance struct {
	Ctx 				context.Context
	Launch				*launcher.Launch
	LaunchOpts			*launcher.LaunchOpts
}


func Run(target string) (*ChromeInstance, error) {
    
	instance := &ChromeInstance{}
	var err error

	instance.LaunchOpts = &launcher.LaunchOpts{
        UserDataDir:"./ax/chrome/"+target,
    }

    instance.Launch, instance.Ctx, err = launcher.Start(instance.LaunchOpts)
	if err!=nil {
		return instance, err
	}

	return instance, err
}
