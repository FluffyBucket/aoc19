package helpers

import "io/ioutil"

func LoadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(dat)
}