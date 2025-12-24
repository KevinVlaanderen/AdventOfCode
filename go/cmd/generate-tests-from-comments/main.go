package main

import (
	"aoc/framework/tasks/tests"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

type Task struct {
	Name  string
	Cases []Case
}

type Case struct {
	DataType tests.DataType
	Filename string
	Param    string
	Expected string
	Skip     bool
}

func main() {
	in, _ := os.LookupEnv("GOFILE")
	pkg, _ := os.LookupEnv("GOPACKAGE")
	out := "generated_test.go"

	tasks, err := parseTestCases(in)
	if err != nil {
		log.Fatal(err)
	} else if len(tasks) == 0 {
		return
	}

	f := jen.NewFile(pkg)

	for _, task := range tasks {
		f.Func().Id(fmt.Sprintf("Test%v", task.Name)).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
			lo.Map(task.Cases, func(testCase Case, index int) jen.Code {
				return jen.Id("t").Dot("Run").Call(
					jen.Lit(fmt.Sprintf("%v(%v)", testCase.DataType, testCase.Filename)),
					jen.Qual("aoc/framework/tasks/tests", "RunTest").Call(
						jen.Id(task.Name),
						parseData(testCase),
						parseParam(testCase),
						parseExpected(testCase),
						jen.Lit(testCase.Skip),
					),
				)
			})...,
		)
		f.Func().Id(fmt.Sprintf("Benchmark%v", task.Name)).Params(jen.Id("b").Op("*").Qual("testing", "B")).Block(
			lo.Map(task.Cases, func(testCase Case, index int) jen.Code {
				return jen.Id("b").Dot("Run").Call(
					jen.Lit(fmt.Sprintf("%v(%v)", testCase.DataType, testCase.Filename)),
					jen.Qual("aoc/framework/tasks/tests", "RunBenchmark").Call(
						jen.Id(task.Name),
						parseData(testCase),
						parseParam(testCase),
						jen.Lit(testCase.Skip),
					),
				)
			})...,
		)
	}

	if err = f.Save(out); err != nil {
		log.Fatal(err)
	}
}

func parseTestCases(in string) ([]Task, error) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, in, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)

	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.FuncDecl)
		if !ok || gd.Doc == nil {
			continue
		}

		task := Task{
			Name:  gd.Name.String(),
			Cases: make([]Case, 0),
		}

		for _, line := range gd.Doc.List {
			def, ok := strings.CutPrefix(line.Text, fmt.Sprintf("// %v ", gd.Name.String()))
			if !ok {
				continue
			}

			testCase := Case{}

			for _, field := range strings.Fields(def) {
				kv := strings.Split(field, ":")
				if len(kv) != 2 {
					return nil, fmt.Errorf("invalid key/value pair: %v", field)
				}

				switch kv[0] {
				case "type":
					switch kv[1] {
					case "real":
						testCase.DataType = tests.Real
					case "mock":
						testCase.DataType = tests.Mock
					default:
						return nil, fmt.Errorf("invalid data type: %v", kv[1])
					}
				case "file":
					testCase.Filename = kv[1]
				case "param":
					testCase.Param = kv[1]
				case "expected":
					testCase.Expected = kv[1]
				case "skip":
					switch kv[1] {
					case "true":
						testCase.Skip = true
					case "false":
						testCase.Skip = false
					default:
						return nil, fmt.Errorf("invalid value: %v", kv[1])
					}
				default:
					return nil, fmt.Errorf("invalid field: %v", kv[0])
				}
			}

			task.Cases = append(task.Cases, testCase)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
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
	var paramStmt jen.Code
	if testCase.Param == "" {
		paramStmt = jen.Qual("github.com/samber/lo", "Empty").Types(jen.Qual("go/types", "Nil")).Call()
	} else {
		paramStmt = jen.Id(testCase.Param)
	}
	return paramStmt
}

func parseExpected(testCase Case) jen.Code {
	return jen.Id(testCase.Expected)
}
