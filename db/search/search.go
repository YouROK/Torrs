package search

import (
	"github.com/agnivade/levenshtein"
	"strings"
	"torrsru/db/sync"
	"torrsru/db/utils"
	"torrsru/models/fdb"
	"torrsru/web/global"
)

var (
	index map[string]struct{}
)

func FindTitle(query string) []*fdb.Torrent {
	list := sync.ListTitles()
	queryarr := strings.Split(utils.ClearStrSpace(query), " ")
	find := map[string]struct{}{}

	for _, name := range list {
		isFound := true
		for _, q := range queryarr {
			if !strings.Contains(name, q) {
				isFound = false
				break
			}
		}

		if isFound {
			lastColon := strings.LastIndex(name, ":")
			if lastColon != -1 {
				name = name[:lastColon]
			}
			find[name] = struct{}{}
		}
	}

	ret := []*fdb.Torrent{}
	for ind := range find {
		torrs := sync.GetTorrentsByName(ind)
		ret = append(ret, torrs...)
	}
	return ret
}

func FindName(query string, accurate bool) []*fdb.Torrent {
	if global.IsUpdateIndex {
		UpdateIndex()
	}

	listRes := map[string]bool{}
	query = utils.ClearStr(query)
	for s, _ := range index {
		if strings.Contains(s, query) {
			listRes[s] = true
		}
	}

	if accurate {
		passDst := 1
		for s := range listRes {
			arr := strings.Split(s, ":")
			if len(arr) == 0 {
				continue
			}
			if len(arr) < 2 {
				arr = append(arr, arr[0])
			}
			if len(arr[1]) == 0 {
				arr[1] = arr[0]
			}
			dst1 := levenshtein.ComputeDistance(query, arr[0])
			dst2 := levenshtein.ComputeDistance(query, arr[1])
			if dst2 < dst1 {
				dst1 = dst2
			}

			if dst1 > passDst {
				listRes[s] = false
			}
		}
	}

	var listTorr []*fdb.Torrent
	for s, b := range listRes {
		if b {
			trs := sync.GetTorrentsByName(s)
			listTorr = append(listTorr, trs...)
		}
	}
	return listTorr
}

func UpdateIndex() {
	list := sync.ListNames()
	index = make(map[string]struct{})
	for _, s := range list {
		index[s] = struct{}{}
	}
	global.IsUpdateIndex = false
}
