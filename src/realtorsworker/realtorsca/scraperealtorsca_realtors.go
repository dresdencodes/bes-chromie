package realtorsca

import (
	"time"
	"net/url"
	"strconv"

	"bes-chromie/src/realtorsworker/realtorsca/chrome"
)

type ScrapeRealtors struct {

	GeoName 		string 	
	Page 			int
	ActionTimeout 	time.Duration
	Chrome 			*chrome.Chrome
}


func (sr *ScrapeRealtors) BuildURL(page int, base ...bool)  string {
	
	// Base URL
	u := &url.URL{
		Scheme: "https",
		Host:   "www.realtor.ca",
		Path:   "/map",
	}

	if len(base) > 0 {return u.String()}

	// Build the fragment parameters (everything after #)
	frag := url.Values{}
	
	//frag.Set("Center", "51.028188,-114.086920")
	//frag.Set("LatitudeMax", "51.31062")
	//frag.Set("LongitudeMax", "-113.46276")
	//frag.Set("LatitudeMin", "50.74402")
	//frag.Set("LongitudeMin", "-114.71108")
	frag.Set("view", "list")
	frag.Set("CurrentPage", strconv.Itoa(page))
	//frag.Set("Sort", "6-D")
	//frag.Set("PGeoIds", "g30_c3nfkdtg")
	frag.Set("GeoName", sr.GeoName)
	frag.Set("PropertyTypeGroupID", "1")
	frag.Set("PropertySearchTypeId", "0")
	//frag.Set("Currency", "CAD")

	// Assign encoded fragment
	u.Fragment = frag.Encode()

	// return string
	return u.String()
}


func Run(geoName string) error {

	//
	// defs
	//
	sr := &ScrapeRealtors{
		GeoName:geoName,
		ActionTimeout:time.Second * time.Duration(30),
	}
	var deferFn func()
	var err error


	//**
	//	stage 1 start chrome
	//**
	sr.Chrome, deferFn, err = chrome.New()
	defer deferFn()
	if err!=nil {
		return err
	}

	//**
	//	navigate 
	//**
	err = sr.NavAndWaitForListings(sr.BuildURL(1), "NAVIGATED AND LISTING CARDS FOUND")
	if err!=nil {
		return err
	}

	//
	// get nodes
	//
	nodes, err := sr.GetCardNodes("CARD NODES SCRAPED")
	if err!=nil {
		return err
	}

	//**
	//  find acceptiable listing
	//**
	err = sr.ScrapeListing(nodes)
	if err!=nil {
		return err
	}

	return nil
}