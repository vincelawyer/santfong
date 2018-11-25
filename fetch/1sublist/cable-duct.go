package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadCableDuctJSON() productList {
	b, err := ioutil.ReadFile("cable-duct.json")
	if err != nil {
		panic(err)
	}

	list := productList{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		panic(err)
	}

	return list
}

func SetCableDuctChinese(enlist, zhlist productList) productList {
	for i, item := range enlist.Items {
		zhlist.Items[i].EnTitle = item.Title
		zhlist.Items[i].ImageSrcs = item.ImageSrcs
	}

	return zhlist
}
