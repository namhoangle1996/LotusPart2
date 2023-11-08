package main

import (
	"LotusPart2/conf"
	"LotusPart2/pkg/route"
	"context"
	"fmt"
	"os"

	"gitlab.com/goxp/cloud0/logger"
)

const (
	APPNAME = "eLotus SVC"
)

// @title eLotus SVC API
// @version 1.0
// @description This is eLotus SVC api docs.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost
// @BasePath  /
func main() {
	conf.SetEnv()
	logger.Init(APPNAME)

	app := route.StartNewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}

// renew all elements that last e plus 1
// [0,2,3] => [0,2,4]
// [0,9] => [0,1,0]
func solution1(input []int64) []int64 {

	var lastSliceValue = input[len(input)-1]
	if lastSliceValue < 9 {
		input[len(input)-1] = lastSliceValue + 1
		return input
	}

	var tmpSlice = make([]int64, len(input)+1)
	for i, v := range input {
		if i == len(input)-1 {
			tmpSlice[i] = 1
			tmpSlice[i+1] = 0
		} else {
			tmpSlice[i] = v
		}
	}

	return tmpSlice
}

func solution2(input [][]int64) int64 {
	var resp int64

	var hmap = map[int]int64{}

	for i, v1 := range input {
		for _, v2 := range v1 {
			hmap[i] += v2
		}
	}

	for v := range hmap {
		if hmap[v] > resp {
			resp = hmap[v]
		}
	}

	return resp
}

// jewels-and-stones
func solution3(jewels, stone string) int64 {
	var resp int64

	var hmap = make(map[string]bool, len(jewels))

	for _, v := range jewels {
		hmap[fmt.Sprintf("%c", v)] = true
	}

	for _, v := range stone {
		if _, ok := hmap[fmt.Sprintf("%c", v)]; ok {
			resp++
		}
	}

	return resp
}

// find the index of elements that total = k
func solution4(input []int64, k int64) (index []int64) {
	var hmap = make(map[int64]int64)

	for i, v := range input {
		if indexHmap, ok := hmap[k-v]; !ok {
			hmap[v] = int64(i)
		} else {
			index = append(index, indexHmap)
			index = append(index, int64(i))
			break
		}
	}

	return index
}
