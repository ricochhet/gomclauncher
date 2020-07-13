package download

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/xmdhs/gomclauncher/launcher"
)

func Getversionlist(atype string) (*Version, error) {
	f := auto(atype)
	var rep *http.Response
	var err error
	for i := 0; i < 4; i++ {
		if i == 3 {
			return nil, err
		}
		rep, _, err = Aget(source(`https://launchermeta.mojang.com/mc/game/version_manifest.json`, f))
		if err != nil {
			if rep != nil {
				rep.Body.Close()
			}
			f = fail(f)
			continue
		}
		err = nil
		break
	}
	if rep != nil {
		defer rep.Body.Close()
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}
	v := Version{}
	json.Unmarshal(b, &v)
	v.atype = atype
	return &v, nil
}

type Version struct {
	Latest   VersionLatest    `json:"latest"`
	Versions []VersionVersion `json:"versions"`
	atype    string
}

type VersionLatest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type VersionVersion struct {
	ID          string `json:"id"`
	ReleaseTime string `json:"releaseTime"`
	Time        string `json:"time"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

func (v Version) Downjson(version string) error {
	f := auto(v.atype)
	for _, vv := range v.Versions {
		if vv.ID == version {
			s := strings.Split(vv.URL, "/")
			path := launcher.Minecraft + `/versions/` + vv.ID + `/` + vv.ID + `.json`
			if ver(path, s[len(s)-2]) {
				return nil
			}
			for i := 0; i < 4; i++ {
				if i == 3 {
					return errors.New("file download fail")
				}
				err := get(source(vv.URL, f), path)
				if err != nil {
					fmt.Println("似乎是网络问题，重试", source(vv.URL, f), err)
					f = fail(f)
					continue
				}
				if !ver(path, s[len(s)-2]) {
					fmt.Println("文件效验失败，重新下载", source(vv.URL, f))
					f = fail(f)
					continue
				}
				break
			}
			return nil
		}
	}
	return errors.New("no such")
}
