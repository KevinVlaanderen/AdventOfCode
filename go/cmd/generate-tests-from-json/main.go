package main

import (
	"aoc/framework/tasks/tests"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

type Tasks map[int][]Case

type Case struct {
	DataType tests.DataType
	Filename string
	Param    any `json:",omitempty"`
	Expected any
	Skip     bool `json:",omitempty"`
}

func main() {
	if err := filepath.Walk("../data/data", func(dataPath string, info fs.FileInfo, err error) error {
		if info.Name() != "cases.json" {
			return nil
		}

		var tasks Tasks

		if f, err := os.OpenFile(dataPath, os.O_RDONLY, os.ModePerm); err != nil {
			return err
		} else {
			d := json.NewDecoder(f)
			d.UseNumber()
			if err = d.Decode(&tasks); err != nil {
				return err
			}
		}

		parts := strings.Split(dataPath, "/")
		year, err := strconv.Atoi(parts[len(parts)-3])
		if err != nil {
			return err
		}
		dayString, ok := strings.CutPrefix(parts[len(parts)-2], "day")
		if !ok {
			return fmt.Errorf("failed to parse day number")
		}
		day, err := strconv.Atoi(dayString)
		if err != nil {
			return err
		}

		out := path.Join(strconv.Itoa(year), "tasks", fmt.Sprintf("day%v", dayString), "generated_test.go")
		f := jen.NewFile(fmt.Sprintf("day%v", day))

		for task, cases := range tasks {
			taskName := fmt.Sprintf("Task%v", strconv.Itoa(task))

			f.Func().Id(fmt.Sprintf("Test%v", taskName)).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
				lo.Map(cases, func(testCase Case, index int) jen.Code {
					return jen.Id("t").Dot("Run").Call(
						jen.Lit(fmt.Sprintf("%v(%v)", testCase.DataType, testCase.Filename)),
						jen.Qual("aoc/framework/tasks/tests", "RunTest").Call(
							jen.Id(taskName),
							parseData(testCase),
							parseParam(testCase),
							parseExpected(testCase),
							jen.Lit(testCase.Skip),
						),
					)
				})...,
			)
			f.Func().Id(fmt.Sprintf("Benchmark%v", taskName)).Params(jen.Id("b").Op("*").Qual("testing", "B")).Block(
				lo.Map(cases, func(testCase Case, index int) jen.Code {
					return jen.Id("b").Dot("Run").Call(
						jen.Lit(fmt.Sprintf("%v(%v)", testCase.DataType, testCase.Filename)),
						jen.Qual("aoc/framework/tasks/tests", "RunBenchmark").Call(
							jen.Id(taskName),
							parseData(testCase),
							parseParam(testCase),
							jen.Lit(testCase.Skip),
						),
					)
				})...,
			)
		}

		return f.Save(out)
	}); err != nil {
		panic(err)
	}
}

func parseData(testCase Case) jen.Code {
	switch testCase.DataType {
	case tests.Mock:
		return jen.Qual("aoc/framework/tasks/tests", "MockData").Call(jen.Lit(testCase.Filename))
	case tests.Real:
		return jen.Qual("aoc/framework/tasks/tests", "RealData").Call(jen.Lit(testCase.Filename))
	default:
		panic(fmt.Errorf("invalid data type: %v", testCase.DataType))
	}
}

func parseParam(testCase Case) jen.Code {
	if testCase.Param == nil {
		return jen.Qual("github.com/samber/lo", "Empty").Types(jen.Qual("go/types", "Nil")).Call()
	}

	switch testCase.Param.(type) {
	case json.Number:
		if iv, err := testCase.Param.(json.Number).Int64(); err == nil {
			return jen.Lit(int(iv))
		} else if fv, err := testCase.Param.(json.Number).Float64(); err == nil {
			return jen.Lit(fv)
		}
		panic(fmt.Errorf("invalid expected value: %v", testCase.Param))
	default:
		return jen.Lit(testCase.Param)
	}
}

func parseExpected(testCase Case) jen.Code {
	switch testCase.Expected.(type) {
	case json.Number:
		if iv, err := testCase.Expected.(json.Number).Int64(); err == nil {
			return jen.Lit(int(iv))
		} else if fv, err := testCase.Expected.(json.Number).Float64(); err == nil {
			return jen.Lit(fv)
		}
		panic(fmt.Errorf("invalid expected value: %v", testCase.Expected))
	default:
		return jen.Lit(testCase.Expected)
	}
}
