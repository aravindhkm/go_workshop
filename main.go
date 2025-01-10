package main

import (
	"WorkShop/generics"
	"WorkShop/variadic"
	"WorkShop/loop"
	"WorkShop/concurrency"
	"WorkShop/sort"
	"WorkShop/error"
	"WorkShop/leetcode"
	"WorkShop/keyword"
	"WorkShop/json"
	"WorkShop/ctx"
	"WorkShop/api"
	"WorkShop/mapping"
	"WorkShop/http_handle"
	"WorkShop/task"
	"WorkShop/problem"
	"WorkShop/routine"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <your_string>")
		return
	}

	inputString := os.Args[1]

	if inputString == "routine" || inputString ==  "routine/" {
		routine.RoutineMain()
		return
	}

	if inputString == "leetcode" || inputString ==  "leetcode/" {
		leetcode.LeetcodeExecute()
		return
	}
	
	if inputString == "concurrency" || inputString ==  "concurrency/" {
		concurrency.ConcurrencyExecute()
		return
	}

	if inputString == "sort" || inputString == "sort/"{
		sort.SortExecute()
		return
	}

	if inputString == "loop" || inputString == "loop/"{
		loop.LoopExecute()
		return
	}
	
	if inputString == "variadic" || inputString == "variadic/"{
		variadic.VariadicExecute()
		return
	}

	if inputString == "generics" || inputString == "generics/"{
		generics.GenericsExecute()
		return
	}

	if inputString == "error" || inputString == "error/"{
		error.ErrorExecute()
		return
	}

	if inputString == "keyword" || inputString == "keyword/"{
		keyword.KeywordExecute()
		return
	}

	if inputString == "json" || inputString == "json/"{
		json.JsonExecute()
		return
	}

	if inputString == "ctx" || inputString == "ctx/"{
		ctx.ContextExecute()
		return
	}

	if inputString == "api" || inputString == "api/"{
		api.ApiExecute()
		return
	}

	if inputString == "mapping" || inputString == "mapping/"{
		mapping.MappingExecute()
		return
	}

	if inputString == "http_handle" || inputString == "http_handle/"{
		httphandle.HttpMain()
		return
	}

	if inputString == "task" || inputString == "task/"{
		task.TaskExecute()
		return
	}

	if inputString == "problem" || inputString == "problem/"{
		problem.ProblemExecute()
		return
	}

	fmt.Println("running cmd not found")

}
