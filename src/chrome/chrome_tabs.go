package chrome


func ListTabs(ctx context.Context) ([]*target.TargetInfo, error) {
    infos, err := target.GetTargets().Do(ctx)
    if err != nil {
        return nil, err
    }

    var tabs []*target.TargetInfo
    for _, t := range infos {
        if t.Type == "page" {
            tabs = append(tabs, t)
        }
    }
    return tabs, nil
}

func NewTab(ctx context.Context, url string) (context.Context, error) {
    tid, err := target.CreateTarget(url).Do(ctx)
    if err != nil {
        return nil, err
    }

    tabCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(target.ID(tid)))
    return tabCtx, nil
}
