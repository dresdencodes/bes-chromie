package realtors

import (
	"bes-chromie/src/realtorsworker/realtorsca/chrome"
)

type ScrapeRealtors struct {

	City 			string 	
	Page 			int
	Chrome 			*chrome.Chrome

}


func (sr *ScrapeRealtors) BuildURL()  string {

	// create url
	u := &url.URL{
		Scheme: "https",
		Host:   "www.realtor.ca",
		Path:   "/realtor-search-results",
	}

	// Build fragment parameters
	frag := url.Values{}
	frag.Set("city", sr.City)
	frag.Set("page", strconv.Itoa(sr.Page))
	frag.Set("sort", "11-A")

	// Assign encoded fragment
	u.Fragment = frag.Encode()

	// return string
	return u.String()
}


func Run(city string) {

	//
	// defs
	//
	sr := &ScrapeRealtors{
		City:city,
	}
	var deferFn func()

	//**
	//	stage 1 start chrome
	//**
	sr.Chrome, deferFn = chrome.New()
	defer deferFn()

}