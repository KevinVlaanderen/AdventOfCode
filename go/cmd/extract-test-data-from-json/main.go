package main

import (
	"aoc/framework/tasks/tests"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/samber/lo"
)

type Years map[int]Days

type Days map[int]Tasks

type Tasks map[int][]Case

type Case struct {
	DataType tests.DataType
	Filename string
	Param    any `json:",omitempty"`
	Expected any
	Skip     bool `json:",omitempty"`
}

type TestCase struct {
	Year, Day, Task int
	Case
}

func main() {
	var paths []string

	if err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.Name() == "tasks.go" {
			paths = append(paths, path)
		}
		return nil
	}); err != nil {
		panic(err)
	}

	var (
		testCases []TestCase
		errs      error
	)

	for _, p := range paths {
		if currentTestCases, err := extractTestCases(p); err != nil {
			errs = multierror.Append(errs, err)
		} else {
			testCases = append(testCases, currentTestCases...)
		}
	}
	if errs != nil {
		panic(errs)
	}

	years := make(Years)
	for _, testCase := range testCases {
		if _, ok := years[testCase.Year]; !ok {
			years[testCase.Year] = make(Days)
		}
		if _, ok := years[testCase.Year][testCase.Day]; !ok {
			years[testCase.Year][testCase.Day] = make(Tasks)
		}
		if _, ok := years[testCase.Year][testCase.Day][testCase.Task]; !ok {
			years[testCase.Year][testCase.Day][testCase.Task] = make([]Case, 0)
		}

		years[testCase.Year][testCase.Day][testCase.Task] = append(years[testCase.Year][testCase.Day][testCase.Task], testCase.Case)
	}

	for year, days := range years {
		for day, tasks := range days {
			jsonOutput, err := json.MarshalIndent(tasks, "", "\t")

			dirPath := path.Join("../data/data", strconv.Itoa(year), fmt.Sprintf("day%v", strconv.Itoa(day)))
			err = os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}

			filePath := path.Join(dirPath, "cases.json")
			f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}

			if _, err = f.WriteString(string(jsonOutput)); err != nil {
				errs = multierror.Append(errs, err)
			}

			if err = f.Close(); err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
		}
	}
	if errs != nil {
		panic(errs)
	}
}

var testCasePattern = regexp.MustCompile(`^// Task(\d+)((?:\s+?\S+)*?)$`)

func extractTestCases(path string) ([]TestCase, error) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(path, "/")
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	dayString, ok := strings.CutPrefix(parts[2], "day")
	if !ok {
		return nil, fmt.Errorf("failed to parse day number")
	}
	day, err := strconv.Atoi(dayString)
	if err != nil {
		return nil, err
	}

	var (
		testCases = make([]TestCase, 0)
		errs      error
	)

	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.FuncDecl)
		if !ok || gd.Doc == nil {
			continue
		}

		var taskErrs error

		currentTestCases := lo.FilterMap(gd.Doc.List, func(line *ast.Comment, index int) (TestCase, bool) {
			matches := testCasePattern.FindStringSubmatch(line.Text)
			if len(matches) != 3 {
				return TestCase{}, false
			}

			task, err := strconv.Atoi(matches[1])
			if err != nil {
				return TestCase{}, false
			}

			var (
				fields       = strings.Fields(matches[2])
				keyValues    = make(map[string]string)
				testCaseErrs error
			)

			for _, field := range fields {
				fieldParts := strings.Split(field, ":")

				if _, ok := keyValues[fieldParts[0]]; ok {
					testCaseErrs = multierror.Append(testCaseErrs, errors.New("duplicate field "+fieldParts[0]))
					continue
				}

				if len(fieldParts) == 2 {
					keyValues[fieldParts[0]] = fieldParts[1]
				} else {
					keyValues[fieldParts[0]] = ""
				}
			}

			c := Case{}

			if dataTypeField, ok := keyValues["type"]; ok {
				switch dataTypeField {
				case "real":
					c.DataType = tests.Real
				case "mock":
					c.DataType = tests.Mock
				default:
					return TestCase{}, false
				}
			} else {
				testCaseErrs = multierror.Append(testCaseErrs, errors.New("missing data type field"))
			}

			if fileField, ok := keyValues["file"]; ok {
				c.Filename = fileField
			} else {
				testCaseErrs = multierror.Append(testCaseErrs, errors.New("missing file field"))
			}

			if paramField, ok := keyValues["param"]; ok {
				c.Param = convertValue(paramField)
			}

			if expectedField, ok := keyValues["expected"]; ok {
				c.Expected = convertValue(expectedField)
			} else {
				testCaseErrs = multierror.Append(testCaseErrs, errors.New("missing expected field"))
			}

			if skipField, ok := keyValues["skip"]; ok {
				switch skipField {
				case "true":
					c.Skip = true
				case "false":
					c.Skip = false
				case "":
					c.Skip = true
				default:
					testCaseErrs = multierror.Append(testCaseErrs, fmt.Errorf("unknown skip field value: %v", skipField))
					return TestCase{}, false
				}
			}

			if testCaseErrs != nil {
				taskErrs = multierror.Append(taskErrs, testCaseErrs)
				return TestCase{}, false
			}

			return TestCase{
				Year: year,
				Day:  day,
				Task: task,
				Case: c,
			}, true
		})

		if taskErrs != nil {
			errs = multierror.Append(errs, taskErrs)
			continue
		}

		testCases = append(testCases, currentTestCases...)
	}

	return testCases, errs
}

func convertValue(v string) any {
	if v == "" {
		return nil
	} else if v == "true" {
		return true
	} else if v == "false" {
		return false
	} else if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
		return v[1 : len(v)-1]
	} else if paramInt, err := strconv.Atoi(v); err == nil {
		return paramInt
	}
	return v
}

//
// type Value interface {
// 	String() string
// }
//
// type IntValue struct {
// 	value int
// }
//
// func (i IntValue) String() string {
//
// }
