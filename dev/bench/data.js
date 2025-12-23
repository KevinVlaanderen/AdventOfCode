window.BENCHMARK_DATA = {
  "lastUpdate": 1766497102549,
  "repoUrl": "https://github.com/KevinVlaanderen/AdventOfCode",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "kevinvlaanderen@gmail.com",
            "name": "Kevin Vlaanderen",
            "username": "KevinVlaanderen"
          },
          "committer": {
            "email": "kevinvlaanderen@gmail.com",
            "name": "Kevin Vlaanderen",
            "username": "KevinVlaanderen"
          },
          "distinct": true,
          "id": "3e9d57399272f37d61724a8894eaa7ba34f2014a",
          "message": "try out multiple range implementations; settle on rangeSlice for now",
          "timestamp": "2025-12-23T14:36:39+01:00",
          "tree_id": "3df1091e91e29366ab7ebc345a150875b9906164",
          "url": "https://github.com/KevinVlaanderen/AdventOfCode/commit/3e9d57399272f37d61724a8894eaa7ba34f2014a"
        },
        "date": 1766497088954,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 221.9,
            "unit": "ns/op\t     160 B/op\t       1 allocs/op",
            "extra": "5396709 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 221.9,
            "unit": "ns/op",
            "extra": "5396709 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "5396709 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5396709 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1)",
            "value": 90771,
            "unit": "ns/op\t   73730 B/op\t       1 allocs/op",
            "extra": "13402 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - ns/op",
            "value": 90771,
            "unit": "ns/op",
            "extra": "13402 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - B/op",
            "value": 73730,
            "unit": "B/op",
            "extra": "13402 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "13402 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 227.7,
            "unit": "ns/op\t     160 B/op\t       1 allocs/op",
            "extra": "5320708 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 227.7,
            "unit": "ns/op",
            "extra": "5320708 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "5320708 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5320708 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1)",
            "value": 86282,
            "unit": "ns/op\t   73730 B/op\t       1 allocs/op",
            "extra": "13893 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - ns/op",
            "value": 86282,
            "unit": "ns/op",
            "extra": "13893 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - B/op",
            "value": 73730,
            "unit": "B/op",
            "extra": "13893 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "13893 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 24627,
            "unit": "ns/op\t    9877 B/op\t     188 allocs/op",
            "extra": "49146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 24627,
            "unit": "ns/op",
            "extra": "49146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 9877,
            "unit": "B/op",
            "extra": "49146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 188,
            "unit": "allocs/op",
            "extra": "49146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10)",
            "value": 10282009,
            "unit": "ns/op\t 5712284 B/op\t  117914 allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - ns/op",
            "value": 10282009,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - B/op",
            "value": 5712284,
            "unit": "B/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - allocs/op",
            "value": 117914,
            "unit": "allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 70704,
            "unit": "ns/op\t   42808 B/op\t     969 allocs/op",
            "extra": "17122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 70704,
            "unit": "ns/op",
            "extra": "17122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 42808,
            "unit": "B/op",
            "extra": "17122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 969,
            "unit": "allocs/op",
            "extra": "17122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10)",
            "value": 687704280,
            "unit": "ns/op\t520414212 B/op\t 6708238 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - ns/op",
            "value": 687704280,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - B/op",
            "value": 520414212,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - allocs/op",
            "value": 6708238,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 9072,
            "unit": "ns/op\t    4429 B/op\t      71 allocs/op",
            "extra": "128173 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 9072,
            "unit": "ns/op",
            "extra": "128173 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 4429,
            "unit": "B/op",
            "extra": "128173 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "128173 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11)",
            "value": 749049,
            "unit": "ns/op\t  329363 B/op\t    4523 allocs/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - ns/op",
            "value": 749049,
            "unit": "ns/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - B/op",
            "value": 329363,
            "unit": "B/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - allocs/op",
            "value": 4523,
            "unit": "allocs/op",
            "extra": "1568 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 12572,
            "unit": "ns/op\t    5904 B/op\t      78 allocs/op",
            "extra": "93886 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 12572,
            "unit": "ns/op",
            "extra": "93886 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "93886 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "93886 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11)",
            "value": 989152,
            "unit": "ns/op\t  428418 B/op\t    3590 allocs/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - ns/op",
            "value": 989152,
            "unit": "ns/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - B/op",
            "value": 428418,
            "unit": "B/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - allocs/op",
            "value": 3590,
            "unit": "allocs/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4535,
            "unit": "ns/op\t    3721 B/op\t      88 allocs/op",
            "extra": "235438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4535,
            "unit": "ns/op",
            "extra": "235438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3721,
            "unit": "B/op",
            "extra": "235438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "235438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12)",
            "value": 858284,
            "unit": "ns/op\t  451591 B/op\t   10780 allocs/op",
            "extra": "1371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - ns/op",
            "value": 858284,
            "unit": "ns/op",
            "extra": "1371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - B/op",
            "value": 451591,
            "unit": "B/op",
            "extra": "1371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - allocs/op",
            "value": 10780,
            "unit": "allocs/op",
            "extra": "1371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2844,
            "unit": "ns/op\t     704 B/op\t      13 allocs/op",
            "extra": "418957 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2844,
            "unit": "ns/op",
            "extra": "418957 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "418957 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "418957 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2)",
            "value": 53059041,
            "unit": "ns/op\t    2347 B/op\t      36 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - ns/op",
            "value": 53059041,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - B/op",
            "value": 2347,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 106833,
            "unit": "ns/op\t   83664 B/op\t    1883 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 106833,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 83664,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1883,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2)",
            "value": 2978171123,
            "unit": "ns/op\t2863887320 B/op\t55229511 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - ns/op",
            "value": 2978171123,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - B/op",
            "value": 2863887320,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - allocs/op",
            "value": 55229511,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 912.7,
            "unit": "ns/op\t    1032 B/op\t      15 allocs/op",
            "extra": "1361576 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 912.7,
            "unit": "ns/op",
            "extra": "1361576 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "1361576 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "1361576 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3)",
            "value": 145608,
            "unit": "ns/op\t  275991 B/op\t     603 allocs/op",
            "extra": "7569 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - ns/op",
            "value": 145608,
            "unit": "ns/op",
            "extra": "7569 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - B/op",
            "value": 275991,
            "unit": "B/op",
            "extra": "7569 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - allocs/op",
            "value": 603,
            "unit": "allocs/op",
            "extra": "7569 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2331,
            "unit": "ns/op\t    1792 B/op\t      59 allocs/op",
            "extra": "482516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2331,
            "unit": "ns/op",
            "extra": "482516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1792,
            "unit": "B/op",
            "extra": "482516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "482516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3)",
            "value": 243919,
            "unit": "ns/op\t  313997 B/op\t    2803 allocs/op",
            "extra": "4539 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - ns/op",
            "value": 243919,
            "unit": "ns/op",
            "extra": "4539 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - B/op",
            "value": 313997,
            "unit": "B/op",
            "extra": "4539 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - allocs/op",
            "value": 2803,
            "unit": "allocs/op",
            "extra": "4539 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 7779,
            "unit": "ns/op\t   10224 B/op\t      85 allocs/op",
            "extra": "155538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 7779,
            "unit": "ns/op",
            "extra": "155538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 10224,
            "unit": "B/op",
            "extra": "155538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 85,
            "unit": "allocs/op",
            "extra": "155538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4)",
            "value": 1498390,
            "unit": "ns/op\t 1712239 B/op\t   12532 allocs/op",
            "extra": "777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - ns/op",
            "value": 1498390,
            "unit": "ns/op",
            "extra": "777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - B/op",
            "value": 1712239,
            "unit": "B/op",
            "extra": "777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - allocs/op",
            "value": 12532,
            "unit": "allocs/op",
            "extra": "777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 18340,
            "unit": "ns/op\t   22640 B/op\t     182 allocs/op",
            "extra": "65100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 18340,
            "unit": "ns/op",
            "extra": "65100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 22640,
            "unit": "B/op",
            "extra": "65100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 182,
            "unit": "allocs/op",
            "extra": "65100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4)",
            "value": 23327766,
            "unit": "ns/op\t24356871 B/op\t  189437 allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - ns/op",
            "value": 23327766,
            "unit": "ns/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - B/op",
            "value": 24356871,
            "unit": "B/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - allocs/op",
            "value": 189437,
            "unit": "allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 906.8,
            "unit": "ns/op\t     872 B/op\t      20 allocs/op",
            "extra": "1352994 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 906.8,
            "unit": "ns/op",
            "extra": "1352994 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 872,
            "unit": "B/op",
            "extra": "1352994 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "1352994 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5)",
            "value": 156095,
            "unit": "ns/op\t   82510 B/op\t     405 allocs/op",
            "extra": "7514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - ns/op",
            "value": 156095,
            "unit": "ns/op",
            "extra": "7514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - B/op",
            "value": 82510,
            "unit": "B/op",
            "extra": "7514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - allocs/op",
            "value": 405,
            "unit": "allocs/op",
            "extra": "7514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1045,
            "unit": "ns/op\t     952 B/op\t      24 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1045,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5)",
            "value": 61683,
            "unit": "ns/op\t   84730 B/op\t     415 allocs/op",
            "extra": "19516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - ns/op",
            "value": 61683,
            "unit": "ns/op",
            "extra": "19516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - B/op",
            "value": 84730,
            "unit": "B/op",
            "extra": "19516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - allocs/op",
            "value": 415,
            "unit": "allocs/op",
            "extra": "19516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 754.3,
            "unit": "ns/op\t     816 B/op\t      16 allocs/op",
            "extra": "1615521 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 754.3,
            "unit": "ns/op",
            "extra": "1615521 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 816,
            "unit": "B/op",
            "extra": "1615521 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "1615521 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6)",
            "value": 155340,
            "unit": "ns/op\t  221206 B/op\t    1023 allocs/op",
            "extra": "7562 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - ns/op",
            "value": 155340,
            "unit": "ns/op",
            "extra": "7562 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - B/op",
            "value": 221206,
            "unit": "B/op",
            "extra": "7562 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - allocs/op",
            "value": 1023,
            "unit": "allocs/op",
            "extra": "7562 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1394,
            "unit": "ns/op\t     784 B/op\t      33 allocs/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1394,
            "unit": "ns/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "763456 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6)",
            "value": 365646,
            "unit": "ns/op\t  184389 B/op\t    7680 allocs/op",
            "extra": "3196 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - ns/op",
            "value": 365646,
            "unit": "ns/op",
            "extra": "3196 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - B/op",
            "value": 184389,
            "unit": "B/op",
            "extra": "3196 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - allocs/op",
            "value": 7680,
            "unit": "allocs/op",
            "extra": "3196 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 9810,
            "unit": "ns/op\t    6648 B/op\t     114 allocs/op",
            "extra": "121856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 9810,
            "unit": "ns/op",
            "extra": "121856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 6648,
            "unit": "B/op",
            "extra": "121856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "121856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7)",
            "value": 735255,
            "unit": "ns/op\t  480613 B/op\t    7481 allocs/op",
            "extra": "1599 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - ns/op",
            "value": 735255,
            "unit": "ns/op",
            "extra": "1599 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - B/op",
            "value": 480613,
            "unit": "B/op",
            "extra": "1599 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - allocs/op",
            "value": 7481,
            "unit": "allocs/op",
            "extra": "1599 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 93554,
            "unit": "ns/op\t   85618 B/op\t     677 allocs/op",
            "extra": "12776 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 93554,
            "unit": "ns/op",
            "extra": "12776 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 85618,
            "unit": "B/op",
            "extra": "12776 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 677,
            "unit": "allocs/op",
            "extra": "12776 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7)",
            "value": 9303202,
            "unit": "ns/op\t 6206233 B/op\t   42841 allocs/op",
            "extra": "127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - ns/op",
            "value": 9303202,
            "unit": "ns/op",
            "extra": "127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - B/op",
            "value": 6206233,
            "unit": "B/op",
            "extra": "127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - allocs/op",
            "value": 42841,
            "unit": "allocs/op",
            "extra": "127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 187100,
            "unit": "ns/op\t  287696 B/op\t     900 allocs/op",
            "extra": "5832 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 187100,
            "unit": "ns/op",
            "extra": "5832 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 287696,
            "unit": "B/op",
            "extra": "5832 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 900,
            "unit": "allocs/op",
            "extra": "5832 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8)",
            "value": 541199691,
            "unit": "ns/op\t608489980 B/op\t  556117 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - ns/op",
            "value": 541199691,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - B/op",
            "value": 608489980,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - allocs/op",
            "value": 556117,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 312393,
            "unit": "ns/op\t  539869 B/op\t    1522 allocs/op",
            "extra": "3740 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 312393,
            "unit": "ns/op",
            "extra": "3740 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 539869,
            "unit": "B/op",
            "extra": "3740 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1522,
            "unit": "allocs/op",
            "extra": "3740 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8)",
            "value": 5125832708,
            "unit": "ns/op\t12708256848 B/op\t11867894 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - ns/op",
            "value": 5125832708,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - B/op",
            "value": 12708256848,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - allocs/op",
            "value": 11867894,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2574,
            "unit": "ns/op\t    2632 B/op\t      19 allocs/op",
            "extra": "450997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2574,
            "unit": "ns/op",
            "extra": "450997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 2632,
            "unit": "B/op",
            "extra": "450997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "450997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9)",
            "value": 35645304,
            "unit": "ns/op\t20736598 B/op\t     531 allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - ns/op",
            "value": 35645304,
            "unit": "ns/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - B/op",
            "value": 20736598,
            "unit": "B/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - allocs/op",
            "value": 531,
            "unit": "allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 3025,
            "unit": "ns/op\t    3080 B/op\t      22 allocs/op",
            "extra": "384438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 3025,
            "unit": "ns/op",
            "extra": "384438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 3080,
            "unit": "B/op",
            "extra": "384438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "384438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9)",
            "value": 53246602,
            "unit": "ns/op\t20769636 B/op\t     541 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - ns/op",
            "value": 53246602,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - B/op",
            "value": 20769636,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - allocs/op",
            "value": 541,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "kevinvlaanderen@gmail.com",
            "name": "Kevin Vlaanderen",
            "username": "KevinVlaanderen"
          },
          "committer": {
            "email": "kevinvlaanderen@gmail.com",
            "name": "Kevin Vlaanderen",
            "username": "KevinVlaanderen"
          },
          "distinct": true,
          "id": "3e9d57399272f37d61724a8894eaa7ba34f2014a",
          "message": "try out multiple range implementations; settle on rangeSlice for now",
          "timestamp": "2025-12-23T14:36:39+01:00",
          "tree_id": "3df1091e91e29366ab7ebc345a150875b9906164",
          "url": "https://github.com/KevinVlaanderen/AdventOfCode/commit/3e9d57399272f37d61724a8894eaa7ba34f2014a"
        },
        "date": 1766497102197,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 1209,
            "unit": "ns/op\t    1256 B/op\t      27 allocs/op",
            "extra": "895764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 1209,
            "unit": "ns/op",
            "extra": "895764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "895764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "895764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1)",
            "value": 117067,
            "unit": "ns/op\t  162291 B/op\t    1257 allocs/op",
            "extra": "9180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - ns/op",
            "value": 117067,
            "unit": "ns/op",
            "extra": "9180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - B/op",
            "value": 162291,
            "unit": "B/op",
            "extra": "9180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - allocs/op",
            "value": 1257,
            "unit": "allocs/op",
            "extra": "9180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1223,
            "unit": "ns/op\t    1256 B/op\t      27 allocs/op",
            "extra": "843777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1223,
            "unit": "ns/op",
            "extra": "843777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "843777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "843777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1)",
            "value": 116352,
            "unit": "ns/op\t  162290 B/op\t    1257 allocs/op",
            "extra": "9178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - ns/op",
            "value": 116352,
            "unit": "ns/op",
            "extra": "9178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - B/op",
            "value": 162290,
            "unit": "B/op",
            "extra": "9178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - allocs/op",
            "value": 1257,
            "unit": "allocs/op",
            "extra": "9178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 144995,
            "unit": "ns/op\t   13442 B/op\t     251 allocs/op",
            "extra": "7720 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 144995,
            "unit": "ns/op",
            "extra": "7720 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 13442,
            "unit": "B/op",
            "extra": "7720 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 251,
            "unit": "allocs/op",
            "extra": "7720 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10)",
            "value": 143839,
            "unit": "ns/op\t   13193 B/op\t     251 allocs/op",
            "extra": "7783 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - ns/op",
            "value": 143839,
            "unit": "ns/op",
            "extra": "7783 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - B/op",
            "value": 13193,
            "unit": "B/op",
            "extra": "7783 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - allocs/op",
            "value": 251,
            "unit": "allocs/op",
            "extra": "7783 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10)",
            "value": 163358,
            "unit": "ns/op\t   19159 B/op\t     485 allocs/op",
            "extra": "6976 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - ns/op",
            "value": 163358,
            "unit": "ns/op",
            "extra": "6976 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - B/op",
            "value": 19159,
            "unit": "B/op",
            "extra": "6976 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - allocs/op",
            "value": 485,
            "unit": "allocs/op",
            "extra": "6976 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 259436,
            "unit": "ns/op\t   86949 B/op\t    2450 allocs/op",
            "extra": "4642 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 259436,
            "unit": "ns/op",
            "extra": "4642 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 86949,
            "unit": "B/op",
            "extra": "4642 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 2450,
            "unit": "allocs/op",
            "extra": "4642 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11)",
            "value": 1151531,
            "unit": "ns/op\t  396413 B/op\t   10897 allocs/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - ns/op",
            "value": 1151531,
            "unit": "ns/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - B/op",
            "value": 396413,
            "unit": "B/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - allocs/op",
            "value": 10897,
            "unit": "allocs/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 21543,
            "unit": "ns/op\t   13720 B/op\t     175 allocs/op",
            "extra": "56841 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 21543,
            "unit": "ns/op",
            "extra": "56841 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 13720,
            "unit": "B/op",
            "extra": "56841 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "56841 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12)",
            "value": 1680500,
            "unit": "ns/op\t  910615 B/op\t    9716 allocs/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - ns/op",
            "value": 1680500,
            "unit": "ns/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - B/op",
            "value": 910615,
            "unit": "B/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - allocs/op",
            "value": 9716,
            "unit": "allocs/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 124732,
            "unit": "ns/op\t   79368 B/op\t    1015 allocs/op",
            "extra": "8072 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 124732,
            "unit": "ns/op",
            "extra": "8072 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 79368,
            "unit": "B/op",
            "extra": "8072 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1015,
            "unit": "allocs/op",
            "extra": "8072 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12)",
            "value": 423212380,
            "unit": "ns/op\t224073997 B/op\t 2477492 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - ns/op",
            "value": 423212380,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - B/op",
            "value": 224073997,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - allocs/op",
            "value": 2477492,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 16108,
            "unit": "ns/op\t    8064 B/op\t     238 allocs/op",
            "extra": "72152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 16108,
            "unit": "ns/op",
            "extra": "72152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 8064,
            "unit": "B/op",
            "extra": "72152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 238,
            "unit": "allocs/op",
            "extra": "72152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13)",
            "value": 1281383,
            "unit": "ns/op\t  568626 B/op\t   19317 allocs/op",
            "extra": "927 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - ns/op",
            "value": 1281383,
            "unit": "ns/op",
            "extra": "927 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - B/op",
            "value": 568626,
            "unit": "B/op",
            "extra": "927 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - allocs/op",
            "value": 19317,
            "unit": "allocs/op",
            "extra": "927 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 29779,
            "unit": "ns/op\t   10457 B/op\t     297 allocs/op",
            "extra": "40369 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 29779,
            "unit": "ns/op",
            "extra": "40369 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 10457,
            "unit": "B/op",
            "extra": "40369 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 297,
            "unit": "allocs/op",
            "extra": "40369 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13)",
            "value": 2297403,
            "unit": "ns/op\t  636155 B/op\t   20162 allocs/op",
            "extra": "517 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - ns/op",
            "value": 2297403,
            "unit": "ns/op",
            "extra": "517 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - B/op",
            "value": 636155,
            "unit": "B/op",
            "extra": "517 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - allocs/op",
            "value": 20162,
            "unit": "allocs/op",
            "extra": "517 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 11573,
            "unit": "ns/op\t    4232 B/op\t      27 allocs/op",
            "extra": "103134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 11573,
            "unit": "ns/op",
            "extra": "103134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 4232,
            "unit": "B/op",
            "extra": "103134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "103134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14)",
            "value": 3858520,
            "unit": "ns/op\t  232513 B/op\t    2189 allocs/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - ns/op",
            "value": 3858520,
            "unit": "ns/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - B/op",
            "value": 232513,
            "unit": "B/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - allocs/op",
            "value": 2189,
            "unit": "allocs/op",
            "extra": "312 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 53792,
            "unit": "ns/op\t   14248 B/op\t      32 allocs/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 53792,
            "unit": "ns/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 14248,
            "unit": "B/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14)",
            "value": 209192551,
            "unit": "ns/op\t 3495164 B/op\t    2436 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - ns/op",
            "value": 209192551,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - B/op",
            "value": 3495164,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - allocs/op",
            "value": 2436,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 11133,
            "unit": "ns/op\t    3373 B/op\t      31 allocs/op",
            "extra": "106984 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 11133,
            "unit": "ns/op",
            "extra": "106984 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3373,
            "unit": "B/op",
            "extra": "106984 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "106984 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15)",
            "value": 262070528,
            "unit": "ns/op\t51311110 B/op\t      87 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - ns/op",
            "value": 262070528,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - B/op",
            "value": 51311110,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 311.7,
            "unit": "ns/op\t     240 B/op\t       5 allocs/op",
            "extra": "3859840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 311.7,
            "unit": "ns/op",
            "extra": "3859840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "3859840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3859840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2)",
            "value": 259747,
            "unit": "ns/op\t  202885 B/op\t    2502 allocs/op",
            "extra": "4495 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - ns/op",
            "value": 259747,
            "unit": "ns/op",
            "extra": "4495 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - B/op",
            "value": 202885,
            "unit": "B/op",
            "extra": "4495 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - allocs/op",
            "value": 2502,
            "unit": "allocs/op",
            "extra": "4495 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 319.1,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "3707355 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 319.1,
            "unit": "ns/op",
            "extra": "3707355 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3707355 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3707355 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2)",
            "value": 254757,
            "unit": "ns/op\t  186501 B/op\t    2502 allocs/op",
            "extra": "4441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - ns/op",
            "value": 254757,
            "unit": "ns/op",
            "extra": "4441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - B/op",
            "value": 186501,
            "unit": "B/op",
            "extra": "4441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - allocs/op",
            "value": 2502,
            "unit": "allocs/op",
            "extra": "4441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 7544,
            "unit": "ns/op\t    4064 B/op\t      36 allocs/op",
            "extra": "159516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 7544,
            "unit": "ns/op",
            "extra": "159516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 4064,
            "unit": "B/op",
            "extra": "159516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "159516 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3)",
            "value": 537907,
            "unit": "ns/op\t  290091 B/op\t    1956 allocs/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - ns/op",
            "value": 537907,
            "unit": "ns/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - B/op",
            "value": 290091,
            "unit": "B/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - allocs/op",
            "value": 1956,
            "unit": "allocs/op",
            "extra": "2158 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 8257,
            "unit": "ns/op\t    5352 B/op\t      31 allocs/op",
            "extra": "145245 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 8257,
            "unit": "ns/op",
            "extra": "145245 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 5352,
            "unit": "B/op",
            "extra": "145245 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "145245 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3)",
            "value": 474988,
            "unit": "ns/op\t  276658 B/op\t    1403 allocs/op",
            "extra": "2419 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - ns/op",
            "value": 474988,
            "unit": "ns/op",
            "extra": "2419 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - B/op",
            "value": 276658,
            "unit": "B/op",
            "extra": "2419 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - allocs/op",
            "value": 1403,
            "unit": "allocs/op",
            "extra": "2419 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 992.5,
            "unit": "ns/op\t     864 B/op\t      20 allocs/op",
            "extra": "1237063 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 992.5,
            "unit": "ns/op",
            "extra": "1237063 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 864,
            "unit": "B/op",
            "extra": "1237063 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "1237063 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4)",
            "value": 167277,
            "unit": "ns/op\t  145155 B/op\t    3002 allocs/op",
            "extra": "7431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - ns/op",
            "value": 167277,
            "unit": "ns/op",
            "extra": "7431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - B/op",
            "value": 145155,
            "unit": "B/op",
            "extra": "7431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - allocs/op",
            "value": 3002,
            "unit": "allocs/op",
            "extra": "7431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2769,
            "unit": "ns/op\t    1264 B/op\t      33 allocs/op",
            "extra": "433974 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2769,
            "unit": "ns/op",
            "extra": "433974 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1264,
            "unit": "B/op",
            "extra": "433974 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "433974 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4)",
            "value": 5374764,
            "unit": "ns/op\t 4216842 B/op\t   15351 allocs/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - ns/op",
            "value": 5374764,
            "unit": "ns/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - B/op",
            "value": 4216842,
            "unit": "B/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - allocs/op",
            "value": 15351,
            "unit": "allocs/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4176,
            "unit": "ns/op\t    1907 B/op\t      41 allocs/op",
            "extra": "256764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4176,
            "unit": "ns/op",
            "extra": "256764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1907,
            "unit": "B/op",
            "extra": "256764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "256764 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5)",
            "value": 220547,
            "unit": "ns/op\t  125116 B/op\t    1170 allocs/op",
            "extra": "5204 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - ns/op",
            "value": 220547,
            "unit": "ns/op",
            "extra": "5204 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - B/op",
            "value": 125116,
            "unit": "B/op",
            "extra": "5204 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - allocs/op",
            "value": 1170,
            "unit": "allocs/op",
            "extra": "5204 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 4295,
            "unit": "ns/op\t    1906 B/op\t      41 allocs/op",
            "extra": "281125 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 4295,
            "unit": "ns/op",
            "extra": "281125 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1906,
            "unit": "B/op",
            "extra": "281125 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "281125 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5)",
            "value": 226340,
            "unit": "ns/op\t  125241 B/op\t    1170 allocs/op",
            "extra": "5295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - ns/op",
            "value": 226340,
            "unit": "ns/op",
            "extra": "5295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - B/op",
            "value": 125241,
            "unit": "B/op",
            "extra": "5295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - allocs/op",
            "value": 1170,
            "unit": "allocs/op",
            "extra": "5295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1)",
            "value": 27.72,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "43224315 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1) - ns/op",
            "value": 27.72,
            "unit": "ns/op",
            "extra": "43224315 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "43224315 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "43224315 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 15.15,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "79561441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 15.15,
            "unit": "ns/op",
            "extra": "79561441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "79561441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "79561441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3)",
            "value": 21.39,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "56198341 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3) - ns/op",
            "value": 21.39,
            "unit": "ns/op",
            "extra": "56198341 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "56198341 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "56198341 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4)",
            "value": 46.44,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25850292 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4) - ns/op",
            "value": 46.44,
            "unit": "ns/op",
            "extra": "25850292 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25850292 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25850292 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5)",
            "value": 52.67,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22779109 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5) - ns/op",
            "value": 52.67,
            "unit": "ns/op",
            "extra": "22779109 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "22779109 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "22779109 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6)",
            "value": 6487,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "185138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - ns/op",
            "value": 6487,
            "unit": "ns/op",
            "extra": "185138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "185138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "185138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1)",
            "value": 122.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "9732777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - ns/op",
            "value": 122.7,
            "unit": "ns/op",
            "extra": "9732777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "9732777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "9732777 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 197.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "6057996 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 197.4,
            "unit": "ns/op",
            "extra": "6057996 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "6057996 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "6057996 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3)",
            "value": 199.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5612986 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - ns/op",
            "value": 199.6,
            "unit": "ns/op",
            "extra": "5612986 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5612986 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5612986 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4)",
            "value": 310,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3855514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - ns/op",
            "value": 310,
            "unit": "ns/op",
            "extra": "3855514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3855514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3855514 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5)",
            "value": 253.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "4740378 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - ns/op",
            "value": 253.5,
            "unit": "ns/op",
            "extra": "4740378 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "4740378 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "4740378 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6)",
            "value": 55430,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "21630 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - ns/op",
            "value": 55430,
            "unit": "ns/op",
            "extra": "21630 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "21630 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "21630 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 5824,
            "unit": "ns/op\t    3792 B/op\t     105 allocs/op",
            "extra": "200102 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 5824,
            "unit": "ns/op",
            "extra": "200102 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3792,
            "unit": "B/op",
            "extra": "200102 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 105,
            "unit": "allocs/op",
            "extra": "200102 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7)",
            "value": 504733,
            "unit": "ns/op\t  295413 B/op\t    4574 allocs/op",
            "extra": "2314 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - ns/op",
            "value": 504733,
            "unit": "ns/op",
            "extra": "2314 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - B/op",
            "value": 295413,
            "unit": "B/op",
            "extra": "2314 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - allocs/op",
            "value": 4574,
            "unit": "allocs/op",
            "extra": "2314 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 5929,
            "unit": "ns/op\t    3848 B/op\t     108 allocs/op",
            "extra": "196952 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 5929,
            "unit": "ns/op",
            "extra": "196952 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 3848,
            "unit": "B/op",
            "extra": "196952 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "196952 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7)",
            "value": 512485,
            "unit": "ns/op\t  295688 B/op\t    4580 allocs/op",
            "extra": "2287 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - ns/op",
            "value": 512485,
            "unit": "ns/op",
            "extra": "2287 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - B/op",
            "value": 295688,
            "unit": "B/op",
            "extra": "2287 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - allocs/op",
            "value": 4580,
            "unit": "allocs/op",
            "extra": "2287 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 1444,
            "unit": "ns/op\t    1296 B/op\t      33 allocs/op",
            "extra": "818959 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 1444,
            "unit": "ns/op",
            "extra": "818959 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "818959 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "818959 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8)",
            "value": 216911,
            "unit": "ns/op\t  268038 B/op\t    1002 allocs/op",
            "extra": "5012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - ns/op",
            "value": 216911,
            "unit": "ns/op",
            "extra": "5012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - B/op",
            "value": 268038,
            "unit": "B/op",
            "extra": "5012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - allocs/op",
            "value": 1002,
            "unit": "allocs/op",
            "extra": "5012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1784,
            "unit": "ns/op\t    1408 B/op\t      31 allocs/op",
            "extra": "636570 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1784,
            "unit": "ns/op",
            "extra": "636570 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1408,
            "unit": "B/op",
            "extra": "636570 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "636570 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8)",
            "value": 4904563,
            "unit": "ns/op\t16323438 B/op\t   38225 allocs/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - ns/op",
            "value": 4904563,
            "unit": "ns/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - B/op",
            "value": 16323438,
            "unit": "B/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - allocs/op",
            "value": 38225,
            "unit": "allocs/op",
            "extra": "248 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1)",
            "value": 2732,
            "unit": "ns/op\t    1200 B/op\t      19 allocs/op",
            "extra": "428308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - ns/op",
            "value": 2732,
            "unit": "ns/op",
            "extra": "428308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "428308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "428308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 9090,
            "unit": "ns/op\t    5129 B/op\t      32 allocs/op",
            "extra": "129650 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 9090,
            "unit": "ns/op",
            "extra": "129650 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 5129,
            "unit": "B/op",
            "extra": "129650 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "129650 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9)",
            "value": 1379077,
            "unit": "ns/op\t  587505 B/op\t    5079 allocs/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - ns/op",
            "value": 1379077,
            "unit": "ns/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - B/op",
            "value": 587505,
            "unit": "B/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - allocs/op",
            "value": 5079,
            "unit": "allocs/op",
            "extra": "856 times\n4 procs"
          }
        ]
      }
    ]
  }
}