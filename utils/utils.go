package utils

import (
	"WorkShop/api"
	"WorkShop/concurrency"
	"WorkShop/ctx"
	"WorkShop/error"
	"WorkShop/generics"
	httphandle "WorkShop/http_handle"
	"WorkShop/interview"
	"WorkShop/json"
	"WorkShop/keyword"
	"WorkShop/leetcode"
	"WorkShop/loop"
	"WorkShop/mapping"
	"WorkShop/pkg"
	"WorkShop/problem"
	"WorkShop/routine"
	"WorkShop/sort"
	"WorkShop/task"
	"WorkShop/variadic"
	"WorkShop/workout"
	"fmt"
	"os"
	"strings"
)

var runMap = map[string]func(string){
	"generics":    generics.Run,
	"variadic":    variadic.Run,
	"loop":        loop.Run,
	"concurrency": concurrency.Run,
	"sort":        sort.Run,
	"error":       error.Run,
	"leetcode":    leetcode.Run,
	"keyword":     keyword.Run,
	"json":        json.Run,
	"ctx":         ctx.Run,
	"api":         api.Run,
	"mapping":     mapping.Run,
	"http_handle": httphandle.Run,
	"task":        task.Run,
	"problem":     problem.Run,
	"routine":     routine.Run,
	"interview":   interview.Run,
	"pkg":         pkg.Run,
	"workout": workout.Run,
}

func FindRunMap(pkgName string) func(string) {
	pkgNameTrim := strings.TrimSuffix(pkgName, "/")

	if runFunc, ok := runMap[pkgNameTrim]; ok {
		return runFunc
	} else {
		fmt.Printf("Unknown module: %s\n", pkgName)
	}

	os.Exit(1)
	return nil
}
