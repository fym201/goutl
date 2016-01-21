package utl

import (
	"regexp"
	"strings"
)

type UserAgent struct {
	OS        string //ios android windows
	OSVersion string //7.1
	Mobile    bool
	Android   bool
	IOS       bool
	IPad 	  bool
	IPod      bool
	IPhone    bool
	WebView   bool
}

func ParseUserAgent(ua string) *UserAgent  {
	ret := &UserAgent{}
	androidReg, _ := regexp.Compile(`(Android);?[\s\/]+([\d.]+)?`)
	ipadReg, _ := regexp.Compile(`(iPad).*OS\s([\d_]+)`)
	ipodReg, _ := regexp.Compile(`(iPod)(.*OS\s([\d_]+))?`)
	iphoneReg, _ := regexp.Compile(`(iPhone\sOS)\s([\d_]+)`)

	if subs := androidReg.FindStringSubmatch(ua); len(subs) > 0 {
		ret.Android = true
		ret.OS = "android"
		ret.OSVersion = strings.Replace(subs[2],"_", ".", -1)

	} else if subs := ipadReg.FindStringSubmatch(ua); len(subs) > 0 {
		ret.IPad = true
		ret.OSVersion = strings.Replace(subs[2],"_", ".", -1)
	} else if subs := ipodReg.FindStringSubmatch(ua); len(subs) > 0 {
		ret.IPod = true
		if len(subs) > 3 {
			ret.OSVersion = strings.Replace(subs[3],"_", ".", -1)
		}
	} else if subs := iphoneReg.FindStringSubmatch(ua); len(subs) > 0 {
		ret.IPhone = true
		ret.OSVersion = strings.Replace(subs[2],"_", ".", -1)
	}
	ret.IOS = ret.IPad || ret.IPod || ret.IPhone
	ret.Mobile = ret.IOS || ret.Android

	if ret.IOS {
		ret.OS = "ios"
	}

	// iOS 8+ changed UA
	if ret.IOS && ret.OSVersion != "" && strings.Index(ua, `Version/`) >= 0 {
		if sub := strings.Split(ret.OSVersion, "."); len(sub) > 0 && sub[0] == "10" {
			tmp := strings.Split(strings.ToLower(ua), `version/`)
			if len(tmp) > 1 {
				ret.OSVersion = strings.Split(tmp[1], " ")[0]
			}
		}
	}

	tmp, _ := regexp.MatchString(`.*AppleWebKit(?!.*Safari)`, ua)
	ret.WebView = ret.IOS && tmp
	return ret
}