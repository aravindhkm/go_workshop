package utils

import (
	"WorkShop/api"
	"WorkShop/concurrency"
	contextexample "WorkShop/context_example"
	"WorkShop/ctx"
	errorexample "WorkShop/error_example"
	"WorkShop/generics"
	goroutinework "WorkShop/goroutine_work"
	httphandle "WorkShop/http_handle"
	interfacehandle "WorkShop/interface_handle"
	"WorkShop/interview"
	"WorkShop/json"
	"WorkShop/keyword"
	"WorkShop/leetcode"
	"WorkShop/loop"
	"WorkShop/mapping"
	"WorkShop/pkg"
	"WorkShop/problem"
	ratelimit "WorkShop/rate_limit"
	"WorkShop/routine"
	"WorkShop/sort"
	syncexample "WorkShop/sync_example"
	"WorkShop/task"
	"WorkShop/tool"
	"WorkShop/variadic"
	workerexample "WorkShop/worker_example"
	"WorkShop/workout"
	"WorkShop/zexecute"
	"fmt"
	"os"
	"strings"
)

var runMap = map[string]func(string){
	"generics":         generics.Run,
	"variadic":         variadic.Run,
	"loop":             loop.Run,
	"concurrency":      concurrency.Run,
	"sort":             sort.Run,
	"error_example":    errorexample.Run,
	"leetcode":         leetcode.Run,
	"keyword":          keyword.Run,
	"json":             json.Run,
	"ctx":              ctx.Run,
	"api":              api.Run,
	"mapping":          mapping.Run,
	"http_handle":      httphandle.Run,
	"task":             task.Run,
	"problem":          problem.Run,
	"rate_limit":        ratelimit.Run,
	"routine":          routine.Run,
	"interview":        interview.Run,
	"pkg":              pkg.Run,
	"workout":          workout.Run,
	"tool":             tool.Run,
	"interface_handle": interfacehandle.Run,
	"goroutine_work":   goroutinework.Run,
	"context_example":  contextexample.Run,
	"worker_example":   workerexample.Run,
	"sync_example":     syncexample.Run,
	"zexecute":         zexecute.Run,
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
