window.BENCHMARK_DATA = {
  "lastUpdate": 1766497850809,
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
        "date": 1766497159555,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2121,
            "unit": "ns/op\t     648 B/op\t       9 allocs/op",
            "extra": "566730 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2121,
            "unit": "ns/op",
            "extra": "566730 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 648,
            "unit": "B/op",
            "extra": "566730 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "566730 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1)",
            "value": 628397,
            "unit": "ns/op\t  162439 B/op\t    2001 allocs/op",
            "extra": "1914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - ns/op",
            "value": 628397,
            "unit": "ns/op",
            "extra": "1914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - B/op",
            "value": 162439,
            "unit": "B/op",
            "extra": "1914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - allocs/op",
            "value": 2001,
            "unit": "allocs/op",
            "extra": "1914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 9206,
            "unit": "ns/op\t    3078 B/op\t      18 allocs/op",
            "extra": "129469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 9206,
            "unit": "ns/op",
            "extra": "129469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 3078,
            "unit": "B/op",
            "extra": "129469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "129469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1)",
            "value": 1287472,
            "unit": "ns/op\t  307431 B/op\t    2004 allocs/op",
            "extra": "916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - ns/op",
            "value": 1287472,
            "unit": "ns/op",
            "extra": "916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - B/op",
            "value": 307431,
            "unit": "B/op",
            "extra": "916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - allocs/op",
            "value": 2004,
            "unit": "allocs/op",
            "extra": "916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 1161,
            "unit": "ns/op\t     720 B/op\t      24 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 1161,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 720,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 3095,
            "unit": "ns/op\t    2504 B/op\t      46 allocs/op",
            "extra": "375543 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 3095,
            "unit": "ns/op",
            "extra": "375543 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 2504,
            "unit": "B/op",
            "extra": "375543 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "375543 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10)",
            "value": 3685356,
            "unit": "ns/op\t 2896989 B/op\t   32651 allocs/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - ns/op",
            "value": 3685356,
            "unit": "ns/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - B/op",
            "value": 2896989,
            "unit": "B/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - allocs/op",
            "value": 32651,
            "unit": "allocs/op",
            "extra": "324 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3)",
            "value": 33738,
            "unit": "ns/op\t   10852 B/op\t     210 allocs/op",
            "extra": "35124 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - ns/op",
            "value": 33738,
            "unit": "ns/op",
            "extra": "35124 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - B/op",
            "value": 10852,
            "unit": "B/op",
            "extra": "35124 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "35124 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4)",
            "value": 86324,
            "unit": "ns/op\t   33056 B/op\t     596 allocs/op",
            "extra": "13876 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - ns/op",
            "value": 86324,
            "unit": "ns/op",
            "extra": "13876 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - B/op",
            "value": 33056,
            "unit": "B/op",
            "extra": "13876 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - allocs/op",
            "value": 596,
            "unit": "allocs/op",
            "extra": "13876 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5)",
            "value": 100336,
            "unit": "ns/op\t   40467 B/op\t     717 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - ns/op",
            "value": 100336,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - B/op",
            "value": 40467,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - allocs/op",
            "value": 717,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10)",
            "value": 10062433,
            "unit": "ns/op\t 4080015 B/op\t   60224 allocs/op",
            "extra": "123 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - ns/op",
            "value": 10062433,
            "unit": "ns/op",
            "extra": "123 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - B/op",
            "value": 4080015,
            "unit": "B/op",
            "extra": "123 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - allocs/op",
            "value": 60224,
            "unit": "allocs/op",
            "extra": "123 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 10092,
            "unit": "ns/op\t    1512 B/op\t      12 allocs/op",
            "extra": "120056 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 10092,
            "unit": "ns/op",
            "extra": "120056 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "120056 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "120056 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11)",
            "value": 3841103,
            "unit": "ns/op\t   45937 B/op\t      29 allocs/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - ns/op",
            "value": 3841103,
            "unit": "ns/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - B/op",
            "value": 45937,
            "unit": "B/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11)",
            "value": 3822985,
            "unit": "ns/op\t   45939 B/op\t      29 allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - ns/op",
            "value": 3822985,
            "unit": "ns/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - B/op",
            "value": 45939,
            "unit": "B/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 43653,
            "unit": "ns/op\t   10271 B/op\t     182 allocs/op",
            "extra": "27373 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 43653,
            "unit": "ns/op",
            "extra": "27373 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 10271,
            "unit": "B/op",
            "extra": "27373 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 182,
            "unit": "allocs/op",
            "extra": "27373 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12)",
            "value": 10299830,
            "unit": "ns/op\t 2381899 B/op\t   39311 allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - ns/op",
            "value": 10299830,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - B/op",
            "value": 2381899,
            "unit": "B/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - allocs/op",
            "value": 39311,
            "unit": "allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 421274,
            "unit": "ns/op\t  128580 B/op\t    1479 allocs/op",
            "extra": "2607 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 421274,
            "unit": "ns/op",
            "extra": "2607 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 128580,
            "unit": "B/op",
            "extra": "2607 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1479,
            "unit": "allocs/op",
            "extra": "2607 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12)",
            "value": 235605534,
            "unit": "ns/op\t96158888 B/op\t 1210490 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - ns/op",
            "value": 235605534,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - B/op",
            "value": 96158888,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - allocs/op",
            "value": 1210490,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 23597,
            "unit": "ns/op\t    5824 B/op\t     288 allocs/op",
            "extra": "50263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 23597,
            "unit": "ns/op",
            "extra": "50263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 5824,
            "unit": "B/op",
            "extra": "50263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 288,
            "unit": "allocs/op",
            "extra": "50263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13)",
            "value": 1010914,
            "unit": "ns/op\t  614781 B/op\t   33633 allocs/op",
            "extra": "1190 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - ns/op",
            "value": 1010914,
            "unit": "ns/op",
            "extra": "1190 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - B/op",
            "value": 614781,
            "unit": "B/op",
            "extra": "1190 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - allocs/op",
            "value": 33633,
            "unit": "allocs/op",
            "extra": "1190 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 27666,
            "unit": "ns/op\t   10608 B/op\t     459 allocs/op",
            "extra": "43081 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 27666,
            "unit": "ns/op",
            "extra": "43081 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 10608,
            "unit": "B/op",
            "extra": "43081 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 459,
            "unit": "allocs/op",
            "extra": "43081 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13)",
            "value": 4858724,
            "unit": "ns/op\t 1467001 B/op\t   58756 allocs/op",
            "extra": "246 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - ns/op",
            "value": 4858724,
            "unit": "ns/op",
            "extra": "246 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - B/op",
            "value": 1467001,
            "unit": "B/op",
            "extra": "246 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - allocs/op",
            "value": 58756,
            "unit": "allocs/op",
            "extra": "246 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 852.4,
            "unit": "ns/op\t     960 B/op\t      13 allocs/op",
            "extra": "1431628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 852.4,
            "unit": "ns/op",
            "extra": "1431628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 960,
            "unit": "B/op",
            "extra": "1431628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "1431628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14)",
            "value": 44289,
            "unit": "ns/op\t   46977 B/op\t     103 allocs/op",
            "extra": "27646 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - ns/op",
            "value": 44289,
            "unit": "ns/op",
            "extra": "27646 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - B/op",
            "value": 46977,
            "unit": "B/op",
            "extra": "27646 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "27646 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 22824,
            "unit": "ns/op\t   11696 B/op\t     189 allocs/op",
            "extra": "53014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 22824,
            "unit": "ns/op",
            "extra": "53014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 11696,
            "unit": "B/op",
            "extra": "53014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 189,
            "unit": "allocs/op",
            "extra": "53014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14)",
            "value": 36469516,
            "unit": "ns/op\t 6679479 B/op\t   19163 allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - ns/op",
            "value": 36469516,
            "unit": "ns/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - B/op",
            "value": 6679479,
            "unit": "B/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - allocs/op",
            "value": 19163,
            "unit": "allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 69.74,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "17108152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 69.74,
            "unit": "ns/op",
            "extra": "17108152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "17108152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "17108152 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 207.8,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "5745100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 207.8,
            "unit": "ns/op",
            "extra": "5745100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "5745100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5745100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15)",
            "value": 60013,
            "unit": "ns/op\t   65554 B/op\t       2 allocs/op",
            "extra": "19855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - ns/op",
            "value": 60013,
            "unit": "ns/op",
            "extra": "19855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - B/op",
            "value": 65554,
            "unit": "B/op",
            "extra": "19855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 18841,
            "unit": "ns/op\t   26128 B/op\t     526 allocs/op",
            "extra": "64422 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 18841,
            "unit": "ns/op",
            "extra": "64422 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 26128,
            "unit": "B/op",
            "extra": "64422 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 526,
            "unit": "allocs/op",
            "extra": "64422 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15)",
            "value": 429975,
            "unit": "ns/op\t  323748 B/op\t    1883 allocs/op",
            "extra": "2736 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15) - ns/op",
            "value": 429975,
            "unit": "ns/op",
            "extra": "2736 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15) - B/op",
            "value": 323748,
            "unit": "B/op",
            "extra": "2736 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15) - allocs/op",
            "value": 1883,
            "unit": "allocs/op",
            "extra": "2736 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 11324,
            "unit": "ns/op\t    9096 B/op\t      20 allocs/op",
            "extra": "106354 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 11324,
            "unit": "ns/op",
            "extra": "106354 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 9096,
            "unit": "B/op",
            "extra": "106354 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "106354 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16)",
            "value": 2559983,
            "unit": "ns/op\t 2117667 B/op\t     137 allocs/op",
            "extra": "468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16) - ns/op",
            "value": 2559983,
            "unit": "ns/op",
            "extra": "468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16) - B/op",
            "value": 2117667,
            "unit": "B/op",
            "extra": "468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16) - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 156707,
            "unit": "ns/op\t  193879 B/op\t     389 allocs/op",
            "extra": "7387 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 156707,
            "unit": "ns/op",
            "extra": "7387 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 193879,
            "unit": "B/op",
            "extra": "7387 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 389,
            "unit": "allocs/op",
            "extra": "7387 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16)",
            "value": 236143547,
            "unit": "ns/op\t542724700 B/op\t   32772 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16) - ns/op",
            "value": 236143547,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16) - B/op",
            "value": 542724700,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16) - allocs/op",
            "value": 32772,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 5819,
            "unit": "ns/op\t    3058 B/op\t      37 allocs/op",
            "extra": "205384 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 5819,
            "unit": "ns/op",
            "extra": "205384 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3058,
            "unit": "B/op",
            "extra": "205384 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "205384 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18)",
            "value": 489164,
            "unit": "ns/op\t  193049 B/op\t    1593 allocs/op",
            "extra": "2468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18) - ns/op",
            "value": 489164,
            "unit": "ns/op",
            "extra": "2468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18) - B/op",
            "value": 193049,
            "unit": "B/op",
            "extra": "2468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18) - allocs/op",
            "value": 1593,
            "unit": "allocs/op",
            "extra": "2468 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 5614,
            "unit": "ns/op\t    3058 B/op\t      37 allocs/op",
            "extra": "202360 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 5614,
            "unit": "ns/op",
            "extra": "202360 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 3058,
            "unit": "B/op",
            "extra": "202360 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "202360 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18)",
            "value": 548088,
            "unit": "ns/op\t  192903 B/op\t    1593 allocs/op",
            "extra": "2149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18) - ns/op",
            "value": 548088,
            "unit": "ns/op",
            "extra": "2149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18) - B/op",
            "value": 192903,
            "unit": "B/op",
            "extra": "2149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18) - allocs/op",
            "value": 1593,
            "unit": "allocs/op",
            "extra": "2149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 98813,
            "unit": "ns/op\t  106715 B/op\t     943 allocs/op",
            "extra": "12138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 98813,
            "unit": "ns/op",
            "extra": "12138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 106715,
            "unit": "B/op",
            "extra": "12138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 943,
            "unit": "allocs/op",
            "extra": "12138 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19)",
            "value": 17603346,
            "unit": "ns/op\t 9941954 B/op\t   71205 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19) - ns/op",
            "value": 17603346,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19) - B/op",
            "value": 9941954,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19) - allocs/op",
            "value": 71205,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 96257,
            "unit": "ns/op\t  104845 B/op\t     905 allocs/op",
            "extra": "12537 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 96257,
            "unit": "ns/op",
            "extra": "12537 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 104845,
            "unit": "B/op",
            "extra": "12537 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 905,
            "unit": "allocs/op",
            "extra": "12537 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19)",
            "value": 26304724,
            "unit": "ns/op\t 9876445 B/op\t   69899 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19) - ns/op",
            "value": 26304724,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19) - B/op",
            "value": 9876445,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19) - allocs/op",
            "value": 69899,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 63856,
            "unit": "ns/op\t   17827 B/op\t     272 allocs/op",
            "extra": "18696 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 63856,
            "unit": "ns/op",
            "extra": "18696 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 17827,
            "unit": "B/op",
            "extra": "18696 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 272,
            "unit": "allocs/op",
            "extra": "18696 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2)",
            "value": 1026083,
            "unit": "ns/op\t  606826 B/op\t    7564 allocs/op",
            "extra": "1178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - ns/op",
            "value": 1026083,
            "unit": "ns/op",
            "extra": "1178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - B/op",
            "value": 606826,
            "unit": "B/op",
            "extra": "1178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - allocs/op",
            "value": 7564,
            "unit": "allocs/op",
            "extra": "1178 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 65931,
            "unit": "ns/op\t   17803 B/op\t     272 allocs/op",
            "extra": "18343 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 65931,
            "unit": "ns/op",
            "extra": "18343 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 17803,
            "unit": "B/op",
            "extra": "18343 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 272,
            "unit": "allocs/op",
            "extra": "18343 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2)",
            "value": 1014824,
            "unit": "ns/op\t  599152 B/op\t    7562 allocs/op",
            "extra": "1213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - ns/op",
            "value": 1014824,
            "unit": "ns/op",
            "extra": "1213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - B/op",
            "value": 599152,
            "unit": "B/op",
            "extra": "1213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - allocs/op",
            "value": 7562,
            "unit": "allocs/op",
            "extra": "1213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 736033,
            "unit": "ns/op\t  215517 B/op\t    9024 allocs/op",
            "extra": "1641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 736033,
            "unit": "ns/op",
            "extra": "1641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 215517,
            "unit": "B/op",
            "extra": "1641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 9024,
            "unit": "allocs/op",
            "extra": "1641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 585346,
            "unit": "ns/op\t  168067 B/op\t    7522 allocs/op",
            "extra": "2002 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 585346,
            "unit": "ns/op",
            "extra": "2002 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 168067,
            "unit": "B/op",
            "extra": "2002 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 7522,
            "unit": "allocs/op",
            "extra": "2002 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20)",
            "value": 3583961,
            "unit": "ns/op\t 1150750 B/op\t   10984 allocs/op",
            "extra": "333 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20) - ns/op",
            "value": 3583961,
            "unit": "ns/op",
            "extra": "333 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20) - B/op",
            "value": 1150750,
            "unit": "B/op",
            "extra": "333 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20) - allocs/op",
            "value": 10984,
            "unit": "allocs/op",
            "extra": "333 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 227852,
            "unit": "ns/op\t   99822 B/op\t    2057 allocs/op",
            "extra": "4909 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 227852,
            "unit": "ns/op",
            "extra": "4909 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 99822,
            "unit": "B/op",
            "extra": "4909 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 2057,
            "unit": "allocs/op",
            "extra": "4909 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21)",
            "value": 58524711,
            "unit": "ns/op\t19245785 B/op\t  420653 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21) - ns/op",
            "value": 58524711,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21) - B/op",
            "value": 19245785,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21) - allocs/op",
            "value": 420653,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4936,
            "unit": "ns/op\t    5024 B/op\t      92 allocs/op",
            "extra": "244812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4936,
            "unit": "ns/op",
            "extra": "244812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 5024,
            "unit": "B/op",
            "extra": "244812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 92,
            "unit": "allocs/op",
            "extra": "244812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22)",
            "value": 1082934,
            "unit": "ns/op\t  878833 B/op\t   16000 allocs/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22) - ns/op",
            "value": 1082934,
            "unit": "ns/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22) - B/op",
            "value": 878833,
            "unit": "B/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22) - allocs/op",
            "value": 16000,
            "unit": "allocs/op",
            "extra": "1088 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 25331,
            "unit": "ns/op\t    9371 B/op\t     180 allocs/op",
            "extra": "46222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 25331,
            "unit": "ns/op",
            "extra": "46222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 9371,
            "unit": "B/op",
            "extra": "46222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 180,
            "unit": "allocs/op",
            "extra": "46222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22)",
            "value": 37630301,
            "unit": "ns/op\t 7916639 B/op\t  191859 allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22) - ns/op",
            "value": 37630301,
            "unit": "ns/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22) - B/op",
            "value": 7916639,
            "unit": "B/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22) - allocs/op",
            "value": 191859,
            "unit": "allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 126398,
            "unit": "ns/op\t   24377 B/op\t     340 allocs/op",
            "extra": "8950 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 126398,
            "unit": "ns/op",
            "extra": "8950 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 24377,
            "unit": "B/op",
            "extra": "8950 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 340,
            "unit": "allocs/op",
            "extra": "8950 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23)",
            "value": 5507383,
            "unit": "ns/op\t 1211229 B/op\t   13905 allocs/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23) - ns/op",
            "value": 5507383,
            "unit": "ns/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23) - B/op",
            "value": 1211229,
            "unit": "B/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23) - allocs/op",
            "value": 13905,
            "unit": "allocs/op",
            "extra": "216 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 227690,
            "unit": "ns/op\t   49201 B/op\t     741 allocs/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 227690,
            "unit": "ns/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 49201,
            "unit": "B/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 741,
            "unit": "allocs/op",
            "extra": "5134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23)",
            "value": 8561192924,
            "unit": "ns/op\t12356692200 B/op\t114357611 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23) - ns/op",
            "value": 8561192924,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23) - B/op",
            "value": 12356692200,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23) - allocs/op",
            "value": 114357611,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4340,
            "unit": "ns/op\t    1636 B/op\t      13 allocs/op",
            "extra": "269814 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4340,
            "unit": "ns/op",
            "extra": "269814 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1636,
            "unit": "B/op",
            "extra": "269814 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "269814 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24)",
            "value": 2375400,
            "unit": "ns/op\t   99858 B/op\t     608 allocs/op",
            "extra": "504 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24) - ns/op",
            "value": 2375400,
            "unit": "ns/op",
            "extra": "504 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24) - B/op",
            "value": 99858,
            "unit": "B/op",
            "extra": "504 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24) - allocs/op",
            "value": 608,
            "unit": "allocs/op",
            "extra": "504 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 218592,
            "unit": "ns/op\t   87507 B/op\t    1232 allocs/op",
            "extra": "5262 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 218592,
            "unit": "ns/op",
            "extra": "5262 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 87507,
            "unit": "B/op",
            "extra": "5262 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 1232,
            "unit": "allocs/op",
            "extra": "5262 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 10819,
            "unit": "ns/op\t   10633 B/op\t     103 allocs/op",
            "extra": "110676 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 10819,
            "unit": "ns/op",
            "extra": "110676 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 10633,
            "unit": "B/op",
            "extra": "110676 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "110676 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3)",
            "value": 3603610,
            "unit": "ns/op\t 1115096 B/op\t    9557 allocs/op",
            "extra": "332 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - ns/op",
            "value": 3603610,
            "unit": "ns/op",
            "extra": "332 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - B/op",
            "value": 1115096,
            "unit": "B/op",
            "extra": "332 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - allocs/op",
            "value": 9557,
            "unit": "allocs/op",
            "extra": "332 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 8995,
            "unit": "ns/op\t    6324 B/op\t      61 allocs/op",
            "extra": "132621 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 8995,
            "unit": "ns/op",
            "extra": "132621 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 6324,
            "unit": "B/op",
            "extra": "132621 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "132621 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3)",
            "value": 6877086,
            "unit": "ns/op\t  603907 B/op\t    4547 allocs/op",
            "extra": "174 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - ns/op",
            "value": 6877086,
            "unit": "ns/op",
            "extra": "174 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - B/op",
            "value": 603907,
            "unit": "B/op",
            "extra": "174 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - allocs/op",
            "value": 4547,
            "unit": "allocs/op",
            "extra": "174 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 66659,
            "unit": "ns/op\t   14585 B/op\t     287 allocs/op",
            "extra": "18429 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 66659,
            "unit": "ns/op",
            "extra": "18429 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 14585,
            "unit": "B/op",
            "extra": "18429 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 287,
            "unit": "allocs/op",
            "extra": "18429 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4)",
            "value": 5274089,
            "unit": "ns/op\t 1281965 B/op\t   22667 allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - ns/op",
            "value": 5274089,
            "unit": "ns/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - B/op",
            "value": 1281965,
            "unit": "B/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - allocs/op",
            "value": 22667,
            "unit": "allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 66196,
            "unit": "ns/op\t   14654 B/op\t     288 allocs/op",
            "extra": "17686 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 66196,
            "unit": "ns/op",
            "extra": "17686 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 14654,
            "unit": "B/op",
            "extra": "17686 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 288,
            "unit": "allocs/op",
            "extra": "17686 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4)",
            "value": 5271294,
            "unit": "ns/op\t 1288449 B/op\t   22752 allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - ns/op",
            "value": 5271294,
            "unit": "ns/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - B/op",
            "value": 1288449,
            "unit": "B/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - allocs/op",
            "value": 22752,
            "unit": "allocs/op",
            "extra": "224 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 23823,
            "unit": "ns/op\t    9278 B/op\t     160 allocs/op",
            "extra": "50149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 23823,
            "unit": "ns/op",
            "extra": "50149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 9278,
            "unit": "B/op",
            "extra": "50149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 160,
            "unit": "allocs/op",
            "extra": "50149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5)",
            "value": 216719,
            "unit": "ns/op\t   76855 B/op\t    1023 allocs/op",
            "extra": "5548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - ns/op",
            "value": 216719,
            "unit": "ns/op",
            "extra": "5548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - B/op",
            "value": 76855,
            "unit": "B/op",
            "extra": "5548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - allocs/op",
            "value": 1023,
            "unit": "allocs/op",
            "extra": "5548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 25454,
            "unit": "ns/op\t   10613 B/op\t     210 allocs/op",
            "extra": "47548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 25454,
            "unit": "ns/op",
            "extra": "47548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 10613,
            "unit": "B/op",
            "extra": "47548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "47548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5)",
            "value": 254357,
            "unit": "ns/op\t  104022 B/op\t    1561 allocs/op",
            "extra": "4546 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - ns/op",
            "value": 254357,
            "unit": "ns/op",
            "extra": "4546 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - B/op",
            "value": 104022,
            "unit": "B/op",
            "extra": "4546 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - allocs/op",
            "value": 1561,
            "unit": "allocs/op",
            "extra": "4546 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 3337,
            "unit": "ns/op\t    1560 B/op\t      21 allocs/op",
            "extra": "357394 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 3337,
            "unit": "ns/op",
            "extra": "357394 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1560,
            "unit": "B/op",
            "extra": "357394 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 21,
            "unit": "allocs/op",
            "extra": "357394 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6)",
            "value": 6018,
            "unit": "ns/op\t    5572 B/op\t      25 allocs/op",
            "extra": "192016 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - ns/op",
            "value": 6018,
            "unit": "ns/op",
            "extra": "192016 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - B/op",
            "value": 5572,
            "unit": "B/op",
            "extra": "192016 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "192016 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2399,
            "unit": "ns/op\t     515 B/op\t      13 allocs/op",
            "extra": "479432 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2399,
            "unit": "ns/op",
            "extra": "479432 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 515,
            "unit": "B/op",
            "extra": "479432 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "479432 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6)",
            "value": 3678,
            "unit": "ns/op\t     522 B/op\t      13 allocs/op",
            "extra": "319711 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - ns/op",
            "value": 3678,
            "unit": "ns/op",
            "extra": "319711 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - B/op",
            "value": 522,
            "unit": "B/op",
            "extra": "319711 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "319711 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2039,
            "unit": "ns/op\t    1440 B/op\t      14 allocs/op",
            "extra": "562779 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2039,
            "unit": "ns/op",
            "extra": "562779 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "562779 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "562779 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7)",
            "value": 1686806,
            "unit": "ns/op\t 1473879 B/op\t    9946 allocs/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - ns/op",
            "value": 1686806,
            "unit": "ns/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - B/op",
            "value": 1473879,
            "unit": "B/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - allocs/op",
            "value": 9946,
            "unit": "allocs/op",
            "extra": "705 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2471,
            "unit": "ns/op\t    1600 B/op\t      15 allocs/op",
            "extra": "449398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2471,
            "unit": "ns/op",
            "extra": "449398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1600,
            "unit": "B/op",
            "extra": "449398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "449398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7)",
            "value": 1761318,
            "unit": "ns/op\t 1444054 B/op\t    9691 allocs/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - ns/op",
            "value": 1761318,
            "unit": "ns/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - B/op",
            "value": 1444054,
            "unit": "B/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - allocs/op",
            "value": 9691,
            "unit": "allocs/op",
            "extra": "685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 16397,
            "unit": "ns/op\t    3496 B/op\t      47 allocs/op",
            "extra": "74146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 16397,
            "unit": "ns/op",
            "extra": "74146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3496,
            "unit": "B/op",
            "extra": "74146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 47,
            "unit": "allocs/op",
            "extra": "74146 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 4855,
            "unit": "ns/op\t    1821 B/op\t      30 allocs/op",
            "extra": "249534 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 4855,
            "unit": "ns/op",
            "extra": "249534 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 1821,
            "unit": "B/op",
            "extra": "249534 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "249534 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8)",
            "value": 1903147,
            "unit": "ns/op\t  388033 B/op\t    2892 allocs/op",
            "extra": "628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - ns/op",
            "value": 1903147,
            "unit": "ns/op",
            "extra": "628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - B/op",
            "value": 388033,
            "unit": "B/op",
            "extra": "628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - allocs/op",
            "value": 2892,
            "unit": "allocs/op",
            "extra": "628 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3)",
            "value": 12108,
            "unit": "ns/op\t    4318 B/op\t      60 allocs/op",
            "extra": "99438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - ns/op",
            "value": 12108,
            "unit": "ns/op",
            "extra": "99438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - B/op",
            "value": 4318,
            "unit": "B/op",
            "extra": "99438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "99438 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8)",
            "value": 2213683,
            "unit": "ns/op\t  401161 B/op\t    2909 allocs/op",
            "extra": "538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - ns/op",
            "value": 2213683,
            "unit": "ns/op",
            "extra": "538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - B/op",
            "value": 401161,
            "unit": "B/op",
            "extra": "538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - allocs/op",
            "value": 2909,
            "unit": "allocs/op",
            "extra": "538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 14961,
            "unit": "ns/op\t    2350 B/op\t      71 allocs/op",
            "extra": "79836 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 14961,
            "unit": "ns/op",
            "extra": "79836 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 2350,
            "unit": "B/op",
            "extra": "79836 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "79836 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9)",
            "value": 1440674,
            "unit": "ns/op\t 1384291 B/op\t   19601 allocs/op",
            "extra": "862 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - ns/op",
            "value": 1440674,
            "unit": "ns/op",
            "extra": "862 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - B/op",
            "value": 1384291,
            "unit": "B/op",
            "extra": "862 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - allocs/op",
            "value": 19601,
            "unit": "allocs/op",
            "extra": "862 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 15215,
            "unit": "ns/op\t    2355 B/op\t      71 allocs/op",
            "extra": "86420 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 15215,
            "unit": "ns/op",
            "extra": "86420 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 2355,
            "unit": "B/op",
            "extra": "86420 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "86420 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9)",
            "value": 1404637,
            "unit": "ns/op\t 1384399 B/op\t   19600 allocs/op",
            "extra": "855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - ns/op",
            "value": 1404637,
            "unit": "ns/op",
            "extra": "855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - B/op",
            "value": 1384399,
            "unit": "B/op",
            "extra": "855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - allocs/op",
            "value": 19600,
            "unit": "allocs/op",
            "extra": "855 times\n4 procs"
          }
        ]
      }
    ],
    "Go 2025": [
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
          "id": "dda4a89424800cc3651c2d5072b988f3d71b19da",
          "message": "try out multiple range implementations; settle on rangeSlice for now",
          "timestamp": "2025-12-23T14:48:01+01:00",
          "tree_id": "d81161bcd109d0cc19f0c33e560d417358357690",
          "url": "https://github.com/KevinVlaanderen/AdventOfCode/commit/dda4a89424800cc3651c2d5072b988f3d71b19da"
        },
        "date": 1766497775635,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 221,
            "unit": "ns/op\t     160 B/op\t       1 allocs/op",
            "extra": "5437346 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 221,
            "unit": "ns/op",
            "extra": "5437346 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "5437346 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5437346 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1)",
            "value": 89924,
            "unit": "ns/op\t   73730 B/op\t       1 allocs/op",
            "extra": "13352 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - ns/op",
            "value": 89924,
            "unit": "ns/op",
            "extra": "13352 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - B/op",
            "value": 73730,
            "unit": "B/op",
            "extra": "13352 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "13352 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 224.5,
            "unit": "ns/op\t     160 B/op\t       1 allocs/op",
            "extra": "5313084 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 224.5,
            "unit": "ns/op",
            "extra": "5313084 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 160,
            "unit": "B/op",
            "extra": "5313084 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5313084 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1)",
            "value": 86242,
            "unit": "ns/op\t   73730 B/op\t       1 allocs/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - ns/op",
            "value": 86242,
            "unit": "ns/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - B/op",
            "value": 73730,
            "unit": "B/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 24841,
            "unit": "ns/op\t    9880 B/op\t     188 allocs/op",
            "extra": "49640 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 24841,
            "unit": "ns/op",
            "extra": "49640 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 9880,
            "unit": "B/op",
            "extra": "49640 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 188,
            "unit": "allocs/op",
            "extra": "49640 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10)",
            "value": 10131373,
            "unit": "ns/op\t 5714775 B/op\t  117915 allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - ns/op",
            "value": 10131373,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - B/op",
            "value": 5714775,
            "unit": "B/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - allocs/op",
            "value": 117915,
            "unit": "allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 70071,
            "unit": "ns/op\t   42849 B/op\t     969 allocs/op",
            "extra": "16970 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 70071,
            "unit": "ns/op",
            "extra": "16970 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 42849,
            "unit": "B/op",
            "extra": "16970 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 969,
            "unit": "allocs/op",
            "extra": "16970 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10)",
            "value": 690186365,
            "unit": "ns/op\t520414644 B/op\t 6708242 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - ns/op",
            "value": 690186365,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - B/op",
            "value": 520414644,
            "unit": "B/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - allocs/op",
            "value": 6708242,
            "unit": "allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 9022,
            "unit": "ns/op\t    4425 B/op\t      71 allocs/op",
            "extra": "132657 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 9022,
            "unit": "ns/op",
            "extra": "132657 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 4425,
            "unit": "B/op",
            "extra": "132657 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "132657 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11)",
            "value": 737250,
            "unit": "ns/op\t  329334 B/op\t    4523 allocs/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - ns/op",
            "value": 737250,
            "unit": "ns/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - B/op",
            "value": 329334,
            "unit": "B/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - allocs/op",
            "value": 4523,
            "unit": "allocs/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 13043,
            "unit": "ns/op\t    5901 B/op\t      78 allocs/op",
            "extra": "95954 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 13043,
            "unit": "ns/op",
            "extra": "95954 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 5901,
            "unit": "B/op",
            "extra": "95954 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "95954 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11)",
            "value": 986056,
            "unit": "ns/op\t  428333 B/op\t    3590 allocs/op",
            "extra": "1142 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - ns/op",
            "value": 986056,
            "unit": "ns/op",
            "extra": "1142 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - B/op",
            "value": 428333,
            "unit": "B/op",
            "extra": "1142 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - allocs/op",
            "value": 3590,
            "unit": "allocs/op",
            "extra": "1142 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4549,
            "unit": "ns/op\t    3719 B/op\t      88 allocs/op",
            "extra": "268692 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4549,
            "unit": "ns/op",
            "extra": "268692 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3719,
            "unit": "B/op",
            "extra": "268692 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 88,
            "unit": "allocs/op",
            "extra": "268692 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12)",
            "value": 864210,
            "unit": "ns/op\t  451634 B/op\t   10780 allocs/op",
            "extra": "1377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - ns/op",
            "value": 864210,
            "unit": "ns/op",
            "extra": "1377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - B/op",
            "value": 451634,
            "unit": "B/op",
            "extra": "1377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - allocs/op",
            "value": 10780,
            "unit": "allocs/op",
            "extra": "1377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2840,
            "unit": "ns/op\t     704 B/op\t      13 allocs/op",
            "extra": "419407 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2840,
            "unit": "ns/op",
            "extra": "419407 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "419407 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "419407 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2)",
            "value": 53307060,
            "unit": "ns/op\t    2347 B/op\t      36 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - ns/op",
            "value": 53307060,
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
            "value": 107762,
            "unit": "ns/op\t   83664 B/op\t    1883 allocs/op",
            "extra": "11200 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 107762,
            "unit": "ns/op",
            "extra": "11200 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 83664,
            "unit": "B/op",
            "extra": "11200 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1883,
            "unit": "allocs/op",
            "extra": "11200 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2)",
            "value": 3024950152,
            "unit": "ns/op\t2863885800 B/op\t55229491 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - ns/op",
            "value": 3024950152,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - B/op",
            "value": 2863885800,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - allocs/op",
            "value": 55229491,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 893.4,
            "unit": "ns/op\t    1032 B/op\t      15 allocs/op",
            "extra": "1372014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 893.4,
            "unit": "ns/op",
            "extra": "1372014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1032,
            "unit": "B/op",
            "extra": "1372014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "1372014 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3)",
            "value": 145966,
            "unit": "ns/op\t  275991 B/op\t     603 allocs/op",
            "extra": "7184 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - ns/op",
            "value": 145966,
            "unit": "ns/op",
            "extra": "7184 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - B/op",
            "value": 275991,
            "unit": "B/op",
            "extra": "7184 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - allocs/op",
            "value": 603,
            "unit": "allocs/op",
            "extra": "7184 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2333,
            "unit": "ns/op\t    1792 B/op\t      59 allocs/op",
            "extra": "485133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2333,
            "unit": "ns/op",
            "extra": "485133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1792,
            "unit": "B/op",
            "extra": "485133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "485133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3)",
            "value": 241489,
            "unit": "ns/op\t  313996 B/op\t    2803 allocs/op",
            "extra": "4864 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - ns/op",
            "value": 241489,
            "unit": "ns/op",
            "extra": "4864 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - B/op",
            "value": 313996,
            "unit": "B/op",
            "extra": "4864 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - allocs/op",
            "value": 2803,
            "unit": "allocs/op",
            "extra": "4864 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 7831,
            "unit": "ns/op\t   10224 B/op\t      85 allocs/op",
            "extra": "151441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 7831,
            "unit": "ns/op",
            "extra": "151441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 10224,
            "unit": "B/op",
            "extra": "151441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 85,
            "unit": "allocs/op",
            "extra": "151441 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4)",
            "value": 1509853,
            "unit": "ns/op\t 1712238 B/op\t   12532 allocs/op",
            "extra": "789 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - ns/op",
            "value": 1509853,
            "unit": "ns/op",
            "extra": "789 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - B/op",
            "value": 1712238,
            "unit": "B/op",
            "extra": "789 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - allocs/op",
            "value": 12532,
            "unit": "allocs/op",
            "extra": "789 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 19018,
            "unit": "ns/op\t   22640 B/op\t     182 allocs/op",
            "extra": "63788 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 19018,
            "unit": "ns/op",
            "extra": "63788 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 22640,
            "unit": "B/op",
            "extra": "63788 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 182,
            "unit": "allocs/op",
            "extra": "63788 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4)",
            "value": 23628317,
            "unit": "ns/op\t24356867 B/op\t  189437 allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - ns/op",
            "value": 23628317,
            "unit": "ns/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - B/op",
            "value": 24356867,
            "unit": "B/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - allocs/op",
            "value": 189437,
            "unit": "allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 909.2,
            "unit": "ns/op\t     872 B/op\t      20 allocs/op",
            "extra": "1359376 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 909.2,
            "unit": "ns/op",
            "extra": "1359376 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 872,
            "unit": "B/op",
            "extra": "1359376 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "1359376 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5)",
            "value": 155268,
            "unit": "ns/op\t   82511 B/op\t     405 allocs/op",
            "extra": "7364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - ns/op",
            "value": 155268,
            "unit": "ns/op",
            "extra": "7364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - B/op",
            "value": 82511,
            "unit": "B/op",
            "extra": "7364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - allocs/op",
            "value": 405,
            "unit": "allocs/op",
            "extra": "7364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1048,
            "unit": "ns/op\t     952 B/op\t      24 allocs/op",
            "extra": "994894 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1048,
            "unit": "ns/op",
            "extra": "994894 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "994894 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "994894 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5)",
            "value": 61724,
            "unit": "ns/op\t   84730 B/op\t     415 allocs/op",
            "extra": "19263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - ns/op",
            "value": 61724,
            "unit": "ns/op",
            "extra": "19263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - B/op",
            "value": 84730,
            "unit": "B/op",
            "extra": "19263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - allocs/op",
            "value": 415,
            "unit": "allocs/op",
            "extra": "19263 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 785.1,
            "unit": "ns/op\t     816 B/op\t      16 allocs/op",
            "extra": "1593404 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 785.1,
            "unit": "ns/op",
            "extra": "1593404 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 816,
            "unit": "B/op",
            "extra": "1593404 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "1593404 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6)",
            "value": 158002,
            "unit": "ns/op\t  221206 B/op\t    1023 allocs/op",
            "extra": "6834 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - ns/op",
            "value": 158002,
            "unit": "ns/op",
            "extra": "6834 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - B/op",
            "value": 221206,
            "unit": "B/op",
            "extra": "6834 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - allocs/op",
            "value": 1023,
            "unit": "allocs/op",
            "extra": "6834 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1394,
            "unit": "ns/op\t     784 B/op\t      33 allocs/op",
            "extra": "787298 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1394,
            "unit": "ns/op",
            "extra": "787298 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "787298 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "787298 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6)",
            "value": 362134,
            "unit": "ns/op\t  184390 B/op\t    7680 allocs/op",
            "extra": "3115 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - ns/op",
            "value": 362134,
            "unit": "ns/op",
            "extra": "3115 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - B/op",
            "value": 184390,
            "unit": "B/op",
            "extra": "3115 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - allocs/op",
            "value": 7680,
            "unit": "allocs/op",
            "extra": "3115 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 9996,
            "unit": "ns/op\t    6648 B/op\t     114 allocs/op",
            "extra": "122180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 9996,
            "unit": "ns/op",
            "extra": "122180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 6648,
            "unit": "B/op",
            "extra": "122180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 114,
            "unit": "allocs/op",
            "extra": "122180 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7)",
            "value": 747016,
            "unit": "ns/op\t  480614 B/op\t    7481 allocs/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - ns/op",
            "value": 747016,
            "unit": "ns/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - B/op",
            "value": 480614,
            "unit": "B/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - allocs/op",
            "value": 7481,
            "unit": "allocs/op",
            "extra": "1590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 93174,
            "unit": "ns/op\t   85618 B/op\t     677 allocs/op",
            "extra": "12826 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 93174,
            "unit": "ns/op",
            "extra": "12826 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 85618,
            "unit": "B/op",
            "extra": "12826 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 677,
            "unit": "allocs/op",
            "extra": "12826 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7)",
            "value": 9480538,
            "unit": "ns/op\t 6206213 B/op\t   42841 allocs/op",
            "extra": "129 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - ns/op",
            "value": 9480538,
            "unit": "ns/op",
            "extra": "129 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - B/op",
            "value": 6206213,
            "unit": "B/op",
            "extra": "129 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - allocs/op",
            "value": 42841,
            "unit": "allocs/op",
            "extra": "129 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 189503,
            "unit": "ns/op\t  287701 B/op\t     900 allocs/op",
            "extra": "5727 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 189503,
            "unit": "ns/op",
            "extra": "5727 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 287701,
            "unit": "B/op",
            "extra": "5727 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 900,
            "unit": "allocs/op",
            "extra": "5727 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8)",
            "value": 520690192,
            "unit": "ns/op\t608622156 B/op\t  556117 allocs/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - ns/op",
            "value": 520690192,
            "unit": "ns/op",
            "extra": "2 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - B/op",
            "value": 608622156,
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
            "value": 311400,
            "unit": "ns/op\t  539869 B/op\t    1522 allocs/op",
            "extra": "3855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 311400,
            "unit": "ns/op",
            "extra": "3855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 539869,
            "unit": "B/op",
            "extra": "3855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1522,
            "unit": "allocs/op",
            "extra": "3855 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8)",
            "value": 4885673323,
            "unit": "ns/op\t12707860000 B/op\t11867832 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - ns/op",
            "value": 4885673323,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - B/op",
            "value": 12707860000,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - allocs/op",
            "value": 11867832,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2612,
            "unit": "ns/op\t    2632 B/op\t      19 allocs/op",
            "extra": "460582 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2612,
            "unit": "ns/op",
            "extra": "460582 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 2632,
            "unit": "B/op",
            "extra": "460582 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "460582 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9)",
            "value": 34199857,
            "unit": "ns/op\t20736614 B/op\t     531 allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - ns/op",
            "value": 34199857,
            "unit": "ns/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - B/op",
            "value": 20736614,
            "unit": "B/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - allocs/op",
            "value": 531,
            "unit": "allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2947,
            "unit": "ns/op\t    3080 B/op\t      22 allocs/op",
            "extra": "388513 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2947,
            "unit": "ns/op",
            "extra": "388513 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 3080,
            "unit": "B/op",
            "extra": "388513 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "388513 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9)",
            "value": 53354376,
            "unit": "ns/op\t20769574 B/op\t     541 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - ns/op",
            "value": 53354376,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - B/op",
            "value": 20769574,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - allocs/op",
            "value": 541,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          }
        ]
      }
    ],
    "Go 2022": [
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
          "id": "dda4a89424800cc3651c2d5072b988f3d71b19da",
          "message": "try out multiple range implementations; settle on rangeSlice for now",
          "timestamp": "2025-12-23T14:48:01+01:00",
          "tree_id": "d81161bcd109d0cc19f0c33e560d417358357690",
          "url": "https://github.com/KevinVlaanderen/AdventOfCode/commit/dda4a89424800cc3651c2d5072b988f3d71b19da"
        },
        "date": 1766497784748,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 1237,
            "unit": "ns/op\t    1256 B/op\t      27 allocs/op",
            "extra": "961840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 1237,
            "unit": "ns/op",
            "extra": "961840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "961840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "961840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1)",
            "value": 118007,
            "unit": "ns/op\t  162290 B/op\t    1257 allocs/op",
            "extra": "9151 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - ns/op",
            "value": 118007,
            "unit": "ns/op",
            "extra": "9151 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - B/op",
            "value": 162290,
            "unit": "B/op",
            "extra": "9151 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - allocs/op",
            "value": 1257,
            "unit": "allocs/op",
            "extra": "9151 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1240,
            "unit": "ns/op\t    1256 B/op\t      27 allocs/op",
            "extra": "900278 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1240,
            "unit": "ns/op",
            "extra": "900278 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "900278 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "900278 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1)",
            "value": 118045,
            "unit": "ns/op\t  162291 B/op\t    1257 allocs/op",
            "extra": "8494 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - ns/op",
            "value": 118045,
            "unit": "ns/op",
            "extra": "8494 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - B/op",
            "value": 162291,
            "unit": "B/op",
            "extra": "8494 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - allocs/op",
            "value": 1257,
            "unit": "allocs/op",
            "extra": "8494 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 149722,
            "unit": "ns/op\t   13439 B/op\t     251 allocs/op",
            "extra": "7377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 149722,
            "unit": "ns/op",
            "extra": "7377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 13439,
            "unit": "B/op",
            "extra": "7377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 251,
            "unit": "allocs/op",
            "extra": "7377 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10)",
            "value": 147202,
            "unit": "ns/op\t   13194 B/op\t     251 allocs/op",
            "extra": "7185 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - ns/op",
            "value": 147202,
            "unit": "ns/op",
            "extra": "7185 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - B/op",
            "value": 13194,
            "unit": "B/op",
            "extra": "7185 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - allocs/op",
            "value": 251,
            "unit": "allocs/op",
            "extra": "7185 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10)",
            "value": 166290,
            "unit": "ns/op\t   19160 B/op\t     485 allocs/op",
            "extra": "6770 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - ns/op",
            "value": 166290,
            "unit": "ns/op",
            "extra": "6770 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - B/op",
            "value": 19160,
            "unit": "B/op",
            "extra": "6770 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - allocs/op",
            "value": 485,
            "unit": "allocs/op",
            "extra": "6770 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 255849,
            "unit": "ns/op\t   86819 B/op\t    2450 allocs/op",
            "extra": "4590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 255849,
            "unit": "ns/op",
            "extra": "4590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 86819,
            "unit": "B/op",
            "extra": "4590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 2450,
            "unit": "allocs/op",
            "extra": "4590 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11)",
            "value": 1162523,
            "unit": "ns/op\t  395975 B/op\t   10897 allocs/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - ns/op",
            "value": 1162523,
            "unit": "ns/op",
            "extra": "1026 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - B/op",
            "value": 395975,
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
            "value": 21221,
            "unit": "ns/op\t   13720 B/op\t     175 allocs/op",
            "extra": "56383 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 21221,
            "unit": "ns/op",
            "extra": "56383 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 13720,
            "unit": "B/op",
            "extra": "56383 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "56383 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12)",
            "value": 1677444,
            "unit": "ns/op\t  910615 B/op\t    9716 allocs/op",
            "extra": "717 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - ns/op",
            "value": 1677444,
            "unit": "ns/op",
            "extra": "717 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - B/op",
            "value": 910615,
            "unit": "B/op",
            "extra": "717 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - allocs/op",
            "value": 9716,
            "unit": "allocs/op",
            "extra": "717 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 123933,
            "unit": "ns/op\t   79369 B/op\t    1015 allocs/op",
            "extra": "8427 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 123933,
            "unit": "ns/op",
            "extra": "8427 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 79369,
            "unit": "B/op",
            "extra": "8427 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1015,
            "unit": "allocs/op",
            "extra": "8427 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12)",
            "value": 422933461,
            "unit": "ns/op\t224073901 B/op\t 2477491 allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - ns/op",
            "value": 422933461,
            "unit": "ns/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - B/op",
            "value": 224073901,
            "unit": "B/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - allocs/op",
            "value": 2477491,
            "unit": "allocs/op",
            "extra": "3 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 16175,
            "unit": "ns/op\t    8064 B/op\t     238 allocs/op",
            "extra": "74047 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 16175,
            "unit": "ns/op",
            "extra": "74047 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 8064,
            "unit": "B/op",
            "extra": "74047 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 238,
            "unit": "allocs/op",
            "extra": "74047 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13)",
            "value": 1333180,
            "unit": "ns/op\t  568625 B/op\t   19317 allocs/op",
            "extra": "913 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - ns/op",
            "value": 1333180,
            "unit": "ns/op",
            "extra": "913 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - B/op",
            "value": 568625,
            "unit": "B/op",
            "extra": "913 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - allocs/op",
            "value": 19317,
            "unit": "allocs/op",
            "extra": "913 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 29975,
            "unit": "ns/op\t   10457 B/op\t     297 allocs/op",
            "extra": "39942 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 29975,
            "unit": "ns/op",
            "extra": "39942 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 10457,
            "unit": "B/op",
            "extra": "39942 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 297,
            "unit": "allocs/op",
            "extra": "39942 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13)",
            "value": 2335523,
            "unit": "ns/op\t  636169 B/op\t   20163 allocs/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - ns/op",
            "value": 2335523,
            "unit": "ns/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - B/op",
            "value": 636169,
            "unit": "B/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - allocs/op",
            "value": 20163,
            "unit": "allocs/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 11682,
            "unit": "ns/op\t    4232 B/op\t      27 allocs/op",
            "extra": "102454 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 11682,
            "unit": "ns/op",
            "extra": "102454 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 4232,
            "unit": "B/op",
            "extra": "102454 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 27,
            "unit": "allocs/op",
            "extra": "102454 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14)",
            "value": 3841042,
            "unit": "ns/op\t  232514 B/op\t    2189 allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - ns/op",
            "value": 3841042,
            "unit": "ns/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - B/op",
            "value": 232514,
            "unit": "B/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - allocs/op",
            "value": 2189,
            "unit": "allocs/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 54271,
            "unit": "ns/op\t   14248 B/op\t      32 allocs/op",
            "extra": "22131 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 54271,
            "unit": "ns/op",
            "extra": "22131 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 14248,
            "unit": "B/op",
            "extra": "22131 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "22131 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14)",
            "value": 209711695,
            "unit": "ns/op\t 3506083 B/op\t    2437 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - ns/op",
            "value": 209711695,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - B/op",
            "value": 3506083,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - allocs/op",
            "value": 2437,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 11464,
            "unit": "ns/op\t    3375 B/op\t      31 allocs/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 11464,
            "unit": "ns/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3375,
            "unit": "B/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15)",
            "value": 261847575,
            "unit": "ns/op\t51301810 B/op\t      85 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - ns/op",
            "value": 261847575,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - B/op",
            "value": 51301810,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - allocs/op",
            "value": 85,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 311.1,
            "unit": "ns/op\t     240 B/op\t       5 allocs/op",
            "extra": "3853780 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 311.1,
            "unit": "ns/op",
            "extra": "3853780 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "3853780 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3853780 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2)",
            "value": 255726,
            "unit": "ns/op\t  202885 B/op\t    2502 allocs/op",
            "extra": "4540 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - ns/op",
            "value": 255726,
            "unit": "ns/op",
            "extra": "4540 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - B/op",
            "value": 202885,
            "unit": "B/op",
            "extra": "4540 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - allocs/op",
            "value": 2502,
            "unit": "allocs/op",
            "extra": "4540 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 320.1,
            "unit": "ns/op\t     224 B/op\t       5 allocs/op",
            "extra": "3737712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 320.1,
            "unit": "ns/op",
            "extra": "3737712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3737712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "3737712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2)",
            "value": 257075,
            "unit": "ns/op\t  186501 B/op\t    2502 allocs/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - ns/op",
            "value": 257075,
            "unit": "ns/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - B/op",
            "value": 186501,
            "unit": "B/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - allocs/op",
            "value": 2502,
            "unit": "allocs/op",
            "extra": "4351 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 7685,
            "unit": "ns/op\t    4064 B/op\t      36 allocs/op",
            "extra": "158396 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 7685,
            "unit": "ns/op",
            "extra": "158396 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 4064,
            "unit": "B/op",
            "extra": "158396 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "158396 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3)",
            "value": 539911,
            "unit": "ns/op\t  290092 B/op\t    1956 allocs/op",
            "extra": "2133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - ns/op",
            "value": 539911,
            "unit": "ns/op",
            "extra": "2133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - B/op",
            "value": 290092,
            "unit": "B/op",
            "extra": "2133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - allocs/op",
            "value": 1956,
            "unit": "allocs/op",
            "extra": "2133 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 8484,
            "unit": "ns/op\t    5352 B/op\t      31 allocs/op",
            "extra": "140684 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 8484,
            "unit": "ns/op",
            "extra": "140684 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 5352,
            "unit": "B/op",
            "extra": "140684 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "140684 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3)",
            "value": 479078,
            "unit": "ns/op\t  276657 B/op\t    1403 allocs/op",
            "extra": "2412 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - ns/op",
            "value": 479078,
            "unit": "ns/op",
            "extra": "2412 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - B/op",
            "value": 276657,
            "unit": "B/op",
            "extra": "2412 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - allocs/op",
            "value": 1403,
            "unit": "allocs/op",
            "extra": "2412 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 991.1,
            "unit": "ns/op\t     864 B/op\t      20 allocs/op",
            "extra": "1229847 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 991.1,
            "unit": "ns/op",
            "extra": "1229847 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 864,
            "unit": "B/op",
            "extra": "1229847 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "1229847 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4)",
            "value": 162708,
            "unit": "ns/op\t  145155 B/op\t    3002 allocs/op",
            "extra": "7098 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - ns/op",
            "value": 162708,
            "unit": "ns/op",
            "extra": "7098 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - B/op",
            "value": 145155,
            "unit": "B/op",
            "extra": "7098 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - allocs/op",
            "value": 3002,
            "unit": "allocs/op",
            "extra": "7098 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2669,
            "unit": "ns/op\t    1264 B/op\t      33 allocs/op",
            "extra": "434074 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2669,
            "unit": "ns/op",
            "extra": "434074 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1264,
            "unit": "B/op",
            "extra": "434074 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "434074 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4)",
            "value": 5622510,
            "unit": "ns/op\t 4216837 B/op\t   15351 allocs/op",
            "extra": "213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - ns/op",
            "value": 5622510,
            "unit": "ns/op",
            "extra": "213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - B/op",
            "value": 4216837,
            "unit": "B/op",
            "extra": "213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - allocs/op",
            "value": 15351,
            "unit": "allocs/op",
            "extra": "213 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4162,
            "unit": "ns/op\t    1908 B/op\t      41 allocs/op",
            "extra": "282748 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4162,
            "unit": "ns/op",
            "extra": "282748 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1908,
            "unit": "B/op",
            "extra": "282748 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "282748 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5)",
            "value": 223146,
            "unit": "ns/op\t  125293 B/op\t    1170 allocs/op",
            "extra": "5054 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - ns/op",
            "value": 223146,
            "unit": "ns/op",
            "extra": "5054 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - B/op",
            "value": 125293,
            "unit": "B/op",
            "extra": "5054 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - allocs/op",
            "value": 1170,
            "unit": "allocs/op",
            "extra": "5054 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 4197,
            "unit": "ns/op\t    1906 B/op\t      41 allocs/op",
            "extra": "275848 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 4197,
            "unit": "ns/op",
            "extra": "275848 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1906,
            "unit": "B/op",
            "extra": "275848 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "275848 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5)",
            "value": 227434,
            "unit": "ns/op\t  125179 B/op\t    1170 allocs/op",
            "extra": "5092 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - ns/op",
            "value": 227434,
            "unit": "ns/op",
            "extra": "5092 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - B/op",
            "value": 125179,
            "unit": "B/op",
            "extra": "5092 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - allocs/op",
            "value": 1170,
            "unit": "allocs/op",
            "extra": "5092 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1)",
            "value": 27.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "42299778 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1) - ns/op",
            "value": 27.9,
            "unit": "ns/op",
            "extra": "42299778 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "42299778 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data1) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "42299778 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 15.08,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "79025655 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 15.08,
            "unit": "ns/op",
            "extra": "79025655 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "79025655 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "79025655 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3)",
            "value": 21.41,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "56191368 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3) - ns/op",
            "value": 21.41,
            "unit": "ns/op",
            "extra": "56191368 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "56191368 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data3) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "56191368 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4)",
            "value": 46.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25816482 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4) - ns/op",
            "value": 46.55,
            "unit": "ns/op",
            "extra": "25816482 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25816482 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data4) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25816482 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5)",
            "value": 53.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22743979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5) - ns/op",
            "value": 53.02,
            "unit": "ns/op",
            "extra": "22743979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "22743979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data5) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "22743979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6)",
            "value": 6512,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "184012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - ns/op",
            "value": 6512,
            "unit": "ns/op",
            "extra": "184012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "184012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "184012 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1)",
            "value": 122.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "9778675 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - ns/op",
            "value": 122.7,
            "unit": "ns/op",
            "extra": "9778675 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "9778675 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "9778675 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 197.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "6052375 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 197.5,
            "unit": "ns/op",
            "extra": "6052375 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "6052375 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "6052375 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3)",
            "value": 198.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "6058449 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - ns/op",
            "value": 198.6,
            "unit": "ns/op",
            "extra": "6058449 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "6058449 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "6058449 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4)",
            "value": 310,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3875065 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - ns/op",
            "value": 310,
            "unit": "ns/op",
            "extra": "3875065 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3875065 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3875065 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5)",
            "value": 253.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "4716210 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - ns/op",
            "value": 253.8,
            "unit": "ns/op",
            "extra": "4716210 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "4716210 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "4716210 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6)",
            "value": 55474,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "21620 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - ns/op",
            "value": 55474,
            "unit": "ns/op",
            "extra": "21620 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "21620 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "21620 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 5868,
            "unit": "ns/op\t    3792 B/op\t     105 allocs/op",
            "extra": "198610 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 5868,
            "unit": "ns/op",
            "extra": "198610 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3792,
            "unit": "B/op",
            "extra": "198610 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 105,
            "unit": "allocs/op",
            "extra": "198610 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7)",
            "value": 509252,
            "unit": "ns/op\t  295414 B/op\t    4574 allocs/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - ns/op",
            "value": 509252,
            "unit": "ns/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - B/op",
            "value": 295414,
            "unit": "B/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - allocs/op",
            "value": 4574,
            "unit": "allocs/op",
            "extra": "2270 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 5959,
            "unit": "ns/op\t    3848 B/op\t     108 allocs/op",
            "extra": "197756 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 5959,
            "unit": "ns/op",
            "extra": "197756 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 3848,
            "unit": "B/op",
            "extra": "197756 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "197756 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7)",
            "value": 506125,
            "unit": "ns/op\t  295688 B/op\t    4580 allocs/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - ns/op",
            "value": 506125,
            "unit": "ns/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - B/op",
            "value": 295688,
            "unit": "B/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - allocs/op",
            "value": 4580,
            "unit": "allocs/op",
            "extra": "2286 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 1463,
            "unit": "ns/op\t    1296 B/op\t      33 allocs/op",
            "extra": "788090 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 1463,
            "unit": "ns/op",
            "extra": "788090 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1296,
            "unit": "B/op",
            "extra": "788090 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 33,
            "unit": "allocs/op",
            "extra": "788090 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8)",
            "value": 220129,
            "unit": "ns/op\t  268038 B/op\t    1002 allocs/op",
            "extra": "5254 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - ns/op",
            "value": 220129,
            "unit": "ns/op",
            "extra": "5254 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - B/op",
            "value": 268038,
            "unit": "B/op",
            "extra": "5254 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - allocs/op",
            "value": 1002,
            "unit": "allocs/op",
            "extra": "5254 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 1808,
            "unit": "ns/op\t    1408 B/op\t      31 allocs/op",
            "extra": "604362 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 1808,
            "unit": "ns/op",
            "extra": "604362 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1408,
            "unit": "B/op",
            "extra": "604362 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "604362 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8)",
            "value": 4923508,
            "unit": "ns/op\t16323466 B/op\t   38225 allocs/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - ns/op",
            "value": 4923508,
            "unit": "ns/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - B/op",
            "value": 16323466,
            "unit": "B/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - allocs/op",
            "value": 38225,
            "unit": "allocs/op",
            "extra": "226 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1)",
            "value": 2790,
            "unit": "ns/op\t    1200 B/op\t      19 allocs/op",
            "extra": "428997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - ns/op",
            "value": 2790,
            "unit": "ns/op",
            "extra": "428997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - B/op",
            "value": 1200,
            "unit": "B/op",
            "extra": "428997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data1) - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "428997 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 9245,
            "unit": "ns/op\t    5128 B/op\t      32 allocs/op",
            "extra": "125685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 9245,
            "unit": "ns/op",
            "extra": "125685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 5128,
            "unit": "B/op",
            "extra": "125685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "125685 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9)",
            "value": 1399895,
            "unit": "ns/op\t  587509 B/op\t    5079 allocs/op",
            "extra": "843 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - ns/op",
            "value": 1399895,
            "unit": "ns/op",
            "extra": "843 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - B/op",
            "value": 587509,
            "unit": "B/op",
            "extra": "843 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - allocs/op",
            "value": 5079,
            "unit": "allocs/op",
            "extra": "843 times\n4 procs"
          }
        ]
      }
    ],
    "Go 2023": [
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
          "id": "dda4a89424800cc3651c2d5072b988f3d71b19da",
          "message": "try out multiple range implementations; settle on rangeSlice for now",
          "timestamp": "2025-12-23T14:48:01+01:00",
          "tree_id": "d81161bcd109d0cc19f0c33e560d417358357690",
          "url": "https://github.com/KevinVlaanderen/AdventOfCode/commit/dda4a89424800cc3651c2d5072b988f3d71b19da"
        },
        "date": 1766497849915,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2157,
            "unit": "ns/op\t     648 B/op\t       9 allocs/op",
            "extra": "539751 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2157,
            "unit": "ns/op",
            "extra": "539751 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 648,
            "unit": "B/op",
            "extra": "539751 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "539751 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1)",
            "value": 621819,
            "unit": "ns/op\t  162438 B/op\t    2001 allocs/op",
            "extra": "1917 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - ns/op",
            "value": 621819,
            "unit": "ns/op",
            "extra": "1917 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - B/op",
            "value": 162438,
            "unit": "B/op",
            "extra": "1917 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day1) - allocs/op",
            "value": 2001,
            "unit": "allocs/op",
            "extra": "1917 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 9250,
            "unit": "ns/op\t    3078 B/op\t      18 allocs/op",
            "extra": "126916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 9250,
            "unit": "ns/op",
            "extra": "126916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 3078,
            "unit": "B/op",
            "extra": "126916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "126916 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1)",
            "value": 1300697,
            "unit": "ns/op\t  307428 B/op\t    2004 allocs/op",
            "extra": "914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - ns/op",
            "value": 1300697,
            "unit": "ns/op",
            "extra": "914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - B/op",
            "value": 307428,
            "unit": "B/op",
            "extra": "914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day1) - allocs/op",
            "value": 2004,
            "unit": "allocs/op",
            "extra": "914 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 1171,
            "unit": "ns/op\t     720 B/op\t      24 allocs/op",
            "extra": "977329 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 1171,
            "unit": "ns/op",
            "extra": "977329 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 720,
            "unit": "B/op",
            "extra": "977329 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "977329 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 3087,
            "unit": "ns/op\t    2504 B/op\t      46 allocs/op",
            "extra": "364030 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 3087,
            "unit": "ns/op",
            "extra": "364030 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 2504,
            "unit": "B/op",
            "extra": "364030 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "364030 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10)",
            "value": 3692388,
            "unit": "ns/op\t 2896990 B/op\t   32651 allocs/op",
            "extra": "326 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - ns/op",
            "value": 3692388,
            "unit": "ns/op",
            "extra": "326 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - B/op",
            "value": 2896990,
            "unit": "B/op",
            "extra": "326 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day10) - allocs/op",
            "value": 32651,
            "unit": "allocs/op",
            "extra": "326 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3)",
            "value": 31306,
            "unit": "ns/op\t   10852 B/op\t     210 allocs/op",
            "extra": "37812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - ns/op",
            "value": 31306,
            "unit": "ns/op",
            "extra": "37812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - B/op",
            "value": 10852,
            "unit": "B/op",
            "extra": "37812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "37812 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4)",
            "value": 85850,
            "unit": "ns/op\t   33056 B/op\t     596 allocs/op",
            "extra": "14130 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - ns/op",
            "value": 85850,
            "unit": "ns/op",
            "extra": "14130 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - B/op",
            "value": 33056,
            "unit": "B/op",
            "extra": "14130 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data4) - allocs/op",
            "value": 596,
            "unit": "allocs/op",
            "extra": "14130 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5)",
            "value": 97841,
            "unit": "ns/op\t   40467 B/op\t     717 allocs/op",
            "extra": "12222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - ns/op",
            "value": 97841,
            "unit": "ns/op",
            "extra": "12222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - B/op",
            "value": 40467,
            "unit": "B/op",
            "extra": "12222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data5) - allocs/op",
            "value": 717,
            "unit": "allocs/op",
            "extra": "12222 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10)",
            "value": 9951432,
            "unit": "ns/op\t 4075770 B/op\t   60215 allocs/op",
            "extra": "122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - ns/op",
            "value": 9951432,
            "unit": "ns/op",
            "extra": "122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - B/op",
            "value": 4075770,
            "unit": "B/op",
            "extra": "122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day10) - allocs/op",
            "value": 60215,
            "unit": "allocs/op",
            "extra": "122 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 8240,
            "unit": "ns/op\t    1512 B/op\t      12 allocs/op",
            "extra": "128593 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 8240,
            "unit": "ns/op",
            "extra": "128593 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1512,
            "unit": "B/op",
            "extra": "128593 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "128593 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11)",
            "value": 3872424,
            "unit": "ns/op\t   45936 B/op\t      29 allocs/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - ns/op",
            "value": 3872424,
            "unit": "ns/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - B/op",
            "value": 45936,
            "unit": "B/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day11) - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "309 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11)",
            "value": 3808947,
            "unit": "ns/op\t   45938 B/op\t      29 allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - ns/op",
            "value": 3808947,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - B/op",
            "value": 45938,
            "unit": "B/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day11) - allocs/op",
            "value": 29,
            "unit": "allocs/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 43678,
            "unit": "ns/op\t   10275 B/op\t     182 allocs/op",
            "extra": "25797 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 43678,
            "unit": "ns/op",
            "extra": "25797 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 10275,
            "unit": "B/op",
            "extra": "25797 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 182,
            "unit": "allocs/op",
            "extra": "25797 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12)",
            "value": 10318988,
            "unit": "ns/op\t 2384220 B/op\t   39312 allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - ns/op",
            "value": 10318988,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - B/op",
            "value": 2384220,
            "unit": "B/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day12) - allocs/op",
            "value": 39312,
            "unit": "allocs/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 403519,
            "unit": "ns/op\t  128504 B/op\t    1479 allocs/op",
            "extra": "2810 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 403519,
            "unit": "ns/op",
            "extra": "2810 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 128504,
            "unit": "B/op",
            "extra": "2810 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 1479,
            "unit": "allocs/op",
            "extra": "2810 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12)",
            "value": 231022999,
            "unit": "ns/op\t96131905 B/op\t 1210299 allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - ns/op",
            "value": 231022999,
            "unit": "ns/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - B/op",
            "value": 96131905,
            "unit": "B/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day12) - allocs/op",
            "value": 1210299,
            "unit": "allocs/op",
            "extra": "5 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 21407,
            "unit": "ns/op\t    5824 B/op\t     288 allocs/op",
            "extra": "56295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 21407,
            "unit": "ns/op",
            "extra": "56295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 5824,
            "unit": "B/op",
            "extra": "56295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 288,
            "unit": "allocs/op",
            "extra": "56295 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13)",
            "value": 989546,
            "unit": "ns/op\t  615174 B/op\t   33633 allocs/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - ns/op",
            "value": 989546,
            "unit": "ns/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - B/op",
            "value": 615174,
            "unit": "B/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day13) - allocs/op",
            "value": 33633,
            "unit": "allocs/op",
            "extra": "1207 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 27656,
            "unit": "ns/op\t   10608 B/op\t     459 allocs/op",
            "extra": "42897 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 27656,
            "unit": "ns/op",
            "extra": "42897 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 10608,
            "unit": "B/op",
            "extra": "42897 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 459,
            "unit": "allocs/op",
            "extra": "42897 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13)",
            "value": 4893495,
            "unit": "ns/op\t 1467973 B/op\t   58756 allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - ns/op",
            "value": 4893495,
            "unit": "ns/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - B/op",
            "value": 1467973,
            "unit": "B/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day13) - allocs/op",
            "value": 58756,
            "unit": "allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 855.1,
            "unit": "ns/op\t     960 B/op\t      13 allocs/op",
            "extra": "1429975 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 855.1,
            "unit": "ns/op",
            "extra": "1429975 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 960,
            "unit": "B/op",
            "extra": "1429975 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "1429975 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14)",
            "value": 43432,
            "unit": "ns/op\t   46977 B/op\t     103 allocs/op",
            "extra": "27688 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - ns/op",
            "value": 43432,
            "unit": "ns/op",
            "extra": "27688 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - B/op",
            "value": 46977,
            "unit": "B/op",
            "extra": "27688 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day14) - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "27688 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 22553,
            "unit": "ns/op\t   11696 B/op\t     189 allocs/op",
            "extra": "52840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 22553,
            "unit": "ns/op",
            "extra": "52840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 11696,
            "unit": "B/op",
            "extra": "52840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 189,
            "unit": "allocs/op",
            "extra": "52840 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14)",
            "value": 37144776,
            "unit": "ns/op\t 6679491 B/op\t   19163 allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - ns/op",
            "value": 37144776,
            "unit": "ns/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - B/op",
            "value": 6679491,
            "unit": "B/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day14) - allocs/op",
            "value": 19163,
            "unit": "allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 70.78,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "17152089 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 70.78,
            "unit": "ns/op",
            "extra": "17152089 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "17152089 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "17152089 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 211.8,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "5390127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 211.8,
            "unit": "ns/op",
            "extra": "5390127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "5390127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5390127 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15)",
            "value": 60059,
            "unit": "ns/op\t   65554 B/op\t       2 allocs/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - ns/op",
            "value": 60059,
            "unit": "ns/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - B/op",
            "value": 65554,
            "unit": "B/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day15) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2)",
            "value": 18587,
            "unit": "ns/op\t   26128 B/op\t     526 allocs/op",
            "extra": "64641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - ns/op",
            "value": 18587,
            "unit": "ns/op",
            "extra": "64641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - B/op",
            "value": 26128,
            "unit": "B/op",
            "extra": "64641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data2) - allocs/op",
            "value": 526,
            "unit": "allocs/op",
            "extra": "64641 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15)",
            "value": 425607,
            "unit": "ns/op\t  323748 B/op\t    1883 allocs/op",
            "extra": "2774 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15) - ns/op",
            "value": 425607,
            "unit": "ns/op",
            "extra": "2774 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15) - B/op",
            "value": 323748,
            "unit": "B/op",
            "extra": "2774 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day15) - allocs/op",
            "value": 1883,
            "unit": "allocs/op",
            "extra": "2774 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 11313,
            "unit": "ns/op\t    9096 B/op\t      20 allocs/op",
            "extra": "106276 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 11313,
            "unit": "ns/op",
            "extra": "106276 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 9096,
            "unit": "B/op",
            "extra": "106276 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "106276 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16)",
            "value": 2547450,
            "unit": "ns/op\t 2117673 B/op\t     137 allocs/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16) - ns/op",
            "value": 2547450,
            "unit": "ns/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16) - B/op",
            "value": 2117673,
            "unit": "B/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day16) - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 150685,
            "unit": "ns/op\t  193880 B/op\t     389 allocs/op",
            "extra": "7656 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 150685,
            "unit": "ns/op",
            "extra": "7656 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 193880,
            "unit": "B/op",
            "extra": "7656 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 389,
            "unit": "allocs/op",
            "extra": "7656 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16)",
            "value": 252588791,
            "unit": "ns/op\t542734710 B/op\t   32820 allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16) - ns/op",
            "value": 252588791,
            "unit": "ns/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16) - B/op",
            "value": 542734710,
            "unit": "B/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day16) - allocs/op",
            "value": 32820,
            "unit": "allocs/op",
            "extra": "4 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 5861,
            "unit": "ns/op\t    3056 B/op\t      37 allocs/op",
            "extra": "202648 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 5861,
            "unit": "ns/op",
            "extra": "202648 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3056,
            "unit": "B/op",
            "extra": "202648 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "202648 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18)",
            "value": 491758,
            "unit": "ns/op\t  192984 B/op\t    1593 allocs/op",
            "extra": "2431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18) - ns/op",
            "value": 491758,
            "unit": "ns/op",
            "extra": "2431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18) - B/op",
            "value": 192984,
            "unit": "B/op",
            "extra": "2431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day18) - allocs/op",
            "value": 1593,
            "unit": "allocs/op",
            "extra": "2431 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 5439,
            "unit": "ns/op\t    3056 B/op\t      37 allocs/op",
            "extra": "215491 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 5439,
            "unit": "ns/op",
            "extra": "215491 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 3056,
            "unit": "B/op",
            "extra": "215491 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 37,
            "unit": "allocs/op",
            "extra": "215491 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18)",
            "value": 550781,
            "unit": "ns/op\t  192950 B/op\t    1593 allocs/op",
            "extra": "2106 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18) - ns/op",
            "value": 550781,
            "unit": "ns/op",
            "extra": "2106 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18) - B/op",
            "value": 192950,
            "unit": "B/op",
            "extra": "2106 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day18) - allocs/op",
            "value": 1593,
            "unit": "allocs/op",
            "extra": "2106 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 99603,
            "unit": "ns/op\t  106616 B/op\t     943 allocs/op",
            "extra": "11920 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 99603,
            "unit": "ns/op",
            "extra": "11920 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 106616,
            "unit": "B/op",
            "extra": "11920 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 943,
            "unit": "allocs/op",
            "extra": "11920 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19)",
            "value": 17876126,
            "unit": "ns/op\t 9942010 B/op\t   71205 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19) - ns/op",
            "value": 17876126,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19) - B/op",
            "value": 9942010,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day19) - allocs/op",
            "value": 71205,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 96672,
            "unit": "ns/op\t  104928 B/op\t     905 allocs/op",
            "extra": "12238 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 96672,
            "unit": "ns/op",
            "extra": "12238 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 104928,
            "unit": "B/op",
            "extra": "12238 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 905,
            "unit": "allocs/op",
            "extra": "12238 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19)",
            "value": 26509378,
            "unit": "ns/op\t 9876780 B/op\t   69901 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19) - ns/op",
            "value": 26509378,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19) - B/op",
            "value": 9876780,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day19) - allocs/op",
            "value": 69901,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 61074,
            "unit": "ns/op\t   17839 B/op\t     272 allocs/op",
            "extra": "19538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 61074,
            "unit": "ns/op",
            "extra": "19538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 17839,
            "unit": "B/op",
            "extra": "19538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 272,
            "unit": "allocs/op",
            "extra": "19538 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2)",
            "value": 983638,
            "unit": "ns/op\t  603542 B/op\t    7563 allocs/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - ns/op",
            "value": 983638,
            "unit": "ns/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - B/op",
            "value": 603542,
            "unit": "B/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day2) - allocs/op",
            "value": 7563,
            "unit": "allocs/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 62471,
            "unit": "ns/op\t   17845 B/op\t     272 allocs/op",
            "extra": "19214 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 62471,
            "unit": "ns/op",
            "extra": "19214 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 17845,
            "unit": "B/op",
            "extra": "19214 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 272,
            "unit": "allocs/op",
            "extra": "19214 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2)",
            "value": 1021223,
            "unit": "ns/op\t  604507 B/op\t    7563 allocs/op",
            "extra": "1269 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - ns/op",
            "value": 1021223,
            "unit": "ns/op",
            "extra": "1269 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - B/op",
            "value": 604507,
            "unit": "B/op",
            "extra": "1269 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day2) - allocs/op",
            "value": 7563,
            "unit": "allocs/op",
            "extra": "1269 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 737090,
            "unit": "ns/op\t  215565 B/op\t    9024 allocs/op",
            "extra": "1632 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 737090,
            "unit": "ns/op",
            "extra": "1632 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 215565,
            "unit": "B/op",
            "extra": "1632 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 9024,
            "unit": "allocs/op",
            "extra": "1632 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 593646,
            "unit": "ns/op\t  168150 B/op\t    7522 allocs/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 593646,
            "unit": "ns/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 168150,
            "unit": "B/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 7522,
            "unit": "allocs/op",
            "extra": "2032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20)",
            "value": 3560327,
            "unit": "ns/op\t 1150368 B/op\t   11005 allocs/op",
            "extra": "337 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20) - ns/op",
            "value": 3560327,
            "unit": "ns/op",
            "extra": "337 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20) - B/op",
            "value": 1150368,
            "unit": "B/op",
            "extra": "337 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day20) - allocs/op",
            "value": 11005,
            "unit": "allocs/op",
            "extra": "337 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 229390,
            "unit": "ns/op\t   99821 B/op\t    2057 allocs/op",
            "extra": "5149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 229390,
            "unit": "ns/op",
            "extra": "5149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 99821,
            "unit": "B/op",
            "extra": "5149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 2057,
            "unit": "allocs/op",
            "extra": "5149 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21)",
            "value": 57382803,
            "unit": "ns/op\t19215628 B/op\t  420651 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21) - ns/op",
            "value": 57382803,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21) - B/op",
            "value": 19215628,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day21) - allocs/op",
            "value": 420651,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4880,
            "unit": "ns/op\t    5024 B/op\t      92 allocs/op",
            "extra": "237825 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4880,
            "unit": "ns/op",
            "extra": "237825 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 5024,
            "unit": "B/op",
            "extra": "237825 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 92,
            "unit": "allocs/op",
            "extra": "237825 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22)",
            "value": 1074281,
            "unit": "ns/op\t  878832 B/op\t   16000 allocs/op",
            "extra": "1108 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22) - ns/op",
            "value": 1074281,
            "unit": "ns/op",
            "extra": "1108 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22) - B/op",
            "value": 878832,
            "unit": "B/op",
            "extra": "1108 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day22) - allocs/op",
            "value": 16000,
            "unit": "allocs/op",
            "extra": "1108 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 24091,
            "unit": "ns/op\t    9357 B/op\t     180 allocs/op",
            "extra": "50629 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 24091,
            "unit": "ns/op",
            "extra": "50629 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 9357,
            "unit": "B/op",
            "extra": "50629 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 180,
            "unit": "allocs/op",
            "extra": "50629 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22)",
            "value": 34806851,
            "unit": "ns/op\t 7976562 B/op\t  191848 allocs/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22) - ns/op",
            "value": 34806851,
            "unit": "ns/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22) - B/op",
            "value": 7976562,
            "unit": "B/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day22) - allocs/op",
            "value": 191848,
            "unit": "allocs/op",
            "extra": "37 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 125042,
            "unit": "ns/op\t   24377 B/op\t     340 allocs/op",
            "extra": "8892 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 125042,
            "unit": "ns/op",
            "extra": "8892 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 24377,
            "unit": "B/op",
            "extra": "8892 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 340,
            "unit": "allocs/op",
            "extra": "8892 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23)",
            "value": 5486386,
            "unit": "ns/op\t 1211209 B/op\t   13905 allocs/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23) - ns/op",
            "value": 5486386,
            "unit": "ns/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23) - B/op",
            "value": 1211209,
            "unit": "B/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day23) - allocs/op",
            "value": 13905,
            "unit": "allocs/op",
            "extra": "218 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 223324,
            "unit": "ns/op\t   49202 B/op\t     741 allocs/op",
            "extra": "5082 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 223324,
            "unit": "ns/op",
            "extra": "5082 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 49202,
            "unit": "B/op",
            "extra": "5082 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 741,
            "unit": "allocs/op",
            "extra": "5082 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23)",
            "value": 8610326178,
            "unit": "ns/op\t12356691224 B/op\t114357575 allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23) - ns/op",
            "value": 8610326178,
            "unit": "ns/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23) - B/op",
            "value": 12356691224,
            "unit": "B/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day23) - allocs/op",
            "value": 114357575,
            "unit": "allocs/op",
            "extra": "1 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 4410,
            "unit": "ns/op\t    1638 B/op\t      13 allocs/op",
            "extra": "273600 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 4410,
            "unit": "ns/op",
            "extra": "273600 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1638,
            "unit": "B/op",
            "extra": "273600 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "273600 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24)",
            "value": 2381842,
            "unit": "ns/op\t   99859 B/op\t     608 allocs/op",
            "extra": "500 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24) - ns/op",
            "value": 2381842,
            "unit": "ns/op",
            "extra": "500 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24) - B/op",
            "value": 99859,
            "unit": "B/op",
            "extra": "500 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day24) - allocs/op",
            "value": 608,
            "unit": "allocs/op",
            "extra": "500 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 219666,
            "unit": "ns/op\t   87492 B/op\t    1232 allocs/op",
            "extra": "5264 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 219666,
            "unit": "ns/op",
            "extra": "5264 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 87492,
            "unit": "B/op",
            "extra": "5264 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 1232,
            "unit": "allocs/op",
            "extra": "5264 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 10963,
            "unit": "ns/op\t   10636 B/op\t     103 allocs/op",
            "extra": "109176 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 10963,
            "unit": "ns/op",
            "extra": "109176 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 10636,
            "unit": "B/op",
            "extra": "109176 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "109176 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3)",
            "value": 3955881,
            "unit": "ns/op\t 1115163 B/op\t    9557 allocs/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - ns/op",
            "value": 3955881,
            "unit": "ns/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - B/op",
            "value": 1115163,
            "unit": "B/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day3) - allocs/op",
            "value": 9557,
            "unit": "allocs/op",
            "extra": "316 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 9238,
            "unit": "ns/op\t    6332 B/op\t      61 allocs/op",
            "extra": "128134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 9238,
            "unit": "ns/op",
            "extra": "128134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 6332,
            "unit": "B/op",
            "extra": "128134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "128134 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3)",
            "value": 6919593,
            "unit": "ns/op\t  603294 B/op\t    4547 allocs/op",
            "extra": "172 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - ns/op",
            "value": 6919593,
            "unit": "ns/op",
            "extra": "172 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - B/op",
            "value": 603294,
            "unit": "B/op",
            "extra": "172 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day3) - allocs/op",
            "value": 4547,
            "unit": "allocs/op",
            "extra": "172 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 65601,
            "unit": "ns/op\t   14599 B/op\t     287 allocs/op",
            "extra": "18273 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 65601,
            "unit": "ns/op",
            "extra": "18273 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 14599,
            "unit": "B/op",
            "extra": "18273 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 287,
            "unit": "allocs/op",
            "extra": "18273 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4)",
            "value": 5229390,
            "unit": "ns/op\t 1281890 B/op\t   22667 allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - ns/op",
            "value": 5229390,
            "unit": "ns/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - B/op",
            "value": 1281890,
            "unit": "B/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day4) - allocs/op",
            "value": 22667,
            "unit": "allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 66651,
            "unit": "ns/op\t   14598 B/op\t     288 allocs/op",
            "extra": "17979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 66651,
            "unit": "ns/op",
            "extra": "17979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 14598,
            "unit": "B/op",
            "extra": "17979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 288,
            "unit": "allocs/op",
            "extra": "17979 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4)",
            "value": 5199325,
            "unit": "ns/op\t 1295066 B/op\t   22753 allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - ns/op",
            "value": 5199325,
            "unit": "ns/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - B/op",
            "value": 1295066,
            "unit": "B/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day4) - allocs/op",
            "value": 22753,
            "unit": "allocs/op",
            "extra": "229 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 24112,
            "unit": "ns/op\t    9289 B/op\t     160 allocs/op",
            "extra": "50289 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 24112,
            "unit": "ns/op",
            "extra": "50289 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 9289,
            "unit": "B/op",
            "extra": "50289 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 160,
            "unit": "allocs/op",
            "extra": "50289 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5)",
            "value": 219064,
            "unit": "ns/op\t   76854 B/op\t    1023 allocs/op",
            "extra": "5059 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - ns/op",
            "value": 219064,
            "unit": "ns/op",
            "extra": "5059 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - B/op",
            "value": 76854,
            "unit": "B/op",
            "extra": "5059 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day5) - allocs/op",
            "value": 1023,
            "unit": "allocs/op",
            "extra": "5059 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 25546,
            "unit": "ns/op\t   10617 B/op\t     210 allocs/op",
            "extra": "47469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 25546,
            "unit": "ns/op",
            "extra": "47469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 10617,
            "unit": "B/op",
            "extra": "47469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "47469 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5)",
            "value": 255740,
            "unit": "ns/op\t  104008 B/op\t    1561 allocs/op",
            "extra": "4364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - ns/op",
            "value": 255740,
            "unit": "ns/op",
            "extra": "4364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - B/op",
            "value": 104008,
            "unit": "B/op",
            "extra": "4364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day5) - allocs/op",
            "value": 1561,
            "unit": "allocs/op",
            "extra": "4364 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 3352,
            "unit": "ns/op\t    1560 B/op\t      21 allocs/op",
            "extra": "350274 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 3352,
            "unit": "ns/op",
            "extra": "350274 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1560,
            "unit": "B/op",
            "extra": "350274 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 21,
            "unit": "allocs/op",
            "extra": "350274 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6)",
            "value": 6083,
            "unit": "ns/op\t    5570 B/op\t      25 allocs/op",
            "extra": "188616 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - ns/op",
            "value": 6083,
            "unit": "ns/op",
            "extra": "188616 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - B/op",
            "value": 5570,
            "unit": "B/op",
            "extra": "188616 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day6) - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "188616 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2397,
            "unit": "ns/op\t     514 B/op\t      13 allocs/op",
            "extra": "471398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2397,
            "unit": "ns/op",
            "extra": "471398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 514,
            "unit": "B/op",
            "extra": "471398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "471398 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6)",
            "value": 3672,
            "unit": "ns/op\t     523 B/op\t      13 allocs/op",
            "extra": "314523 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - ns/op",
            "value": 3672,
            "unit": "ns/op",
            "extra": "314523 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - B/op",
            "value": 523,
            "unit": "B/op",
            "extra": "314523 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day6) - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "314523 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 2045,
            "unit": "ns/op\t    1440 B/op\t      14 allocs/op",
            "extra": "572512 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 2045,
            "unit": "ns/op",
            "extra": "572512 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 1440,
            "unit": "B/op",
            "extra": "572512 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "572512 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7)",
            "value": 1757024,
            "unit": "ns/op\t 1473881 B/op\t    9946 allocs/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - ns/op",
            "value": 1757024,
            "unit": "ns/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - B/op",
            "value": 1473881,
            "unit": "B/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day7) - allocs/op",
            "value": 9946,
            "unit": "allocs/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 2476,
            "unit": "ns/op\t    1600 B/op\t      15 allocs/op",
            "extra": "445728 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 2476,
            "unit": "ns/op",
            "extra": "445728 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 1600,
            "unit": "B/op",
            "extra": "445728 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 15,
            "unit": "allocs/op",
            "extra": "445728 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7)",
            "value": 1686777,
            "unit": "ns/op\t 1444050 B/op\t    9691 allocs/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - ns/op",
            "value": 1686777,
            "unit": "ns/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - B/op",
            "value": 1444050,
            "unit": "B/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day7) - allocs/op",
            "value": 9691,
            "unit": "allocs/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 13559,
            "unit": "ns/op\t    3493 B/op\t      47 allocs/op",
            "extra": "88032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 13559,
            "unit": "ns/op",
            "extra": "88032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 3493,
            "unit": "B/op",
            "extra": "88032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 47,
            "unit": "allocs/op",
            "extra": "88032 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2)",
            "value": 4653,
            "unit": "ns/op\t    1817 B/op\t      30 allocs/op",
            "extra": "255182 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - ns/op",
            "value": 4653,
            "unit": "ns/op",
            "extra": "255182 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - B/op",
            "value": 1817,
            "unit": "B/op",
            "extra": "255182 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data2) - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "255182 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8)",
            "value": 1904989,
            "unit": "ns/op\t  388178 B/op\t    2892 allocs/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - ns/op",
            "value": 1904989,
            "unit": "ns/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - B/op",
            "value": 388178,
            "unit": "B/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day8) - allocs/op",
            "value": 2892,
            "unit": "allocs/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3)",
            "value": 15403,
            "unit": "ns/op\t    4319 B/op\t      60 allocs/op",
            "extra": "71371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - ns/op",
            "value": 15403,
            "unit": "ns/op",
            "extra": "71371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - B/op",
            "value": 4319,
            "unit": "B/op",
            "extra": "71371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data3) - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "71371 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8)",
            "value": 2197739,
            "unit": "ns/op\t  401285 B/op\t    2909 allocs/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - ns/op",
            "value": 2197739,
            "unit": "ns/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - B/op",
            "value": 401285,
            "unit": "B/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day8) - allocs/op",
            "value": 2909,
            "unit": "allocs/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data)",
            "value": 12444,
            "unit": "ns/op\t    2351 B/op\t      71 allocs/op",
            "extra": "94548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - ns/op",
            "value": 12444,
            "unit": "ns/op",
            "extra": "94548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - B/op",
            "value": 2351,
            "unit": "B/op",
            "extra": "94548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Mock(data) - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "94548 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9)",
            "value": 1401143,
            "unit": "ns/op\t 1384753 B/op\t   19600 allocs/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - ns/op",
            "value": 1401143,
            "unit": "ns/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - B/op",
            "value": 1384753,
            "unit": "B/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkTask1/Real(day9) - allocs/op",
            "value": 19600,
            "unit": "allocs/op",
            "extra": "871 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data)",
            "value": 12545,
            "unit": "ns/op\t    2359 B/op\t      71 allocs/op",
            "extra": "96625 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - ns/op",
            "value": 12545,
            "unit": "ns/op",
            "extra": "96625 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - B/op",
            "value": 2359,
            "unit": "B/op",
            "extra": "96625 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Mock(data) - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "96625 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9)",
            "value": 1390673,
            "unit": "ns/op\t 1384516 B/op\t   19601 allocs/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - ns/op",
            "value": 1390673,
            "unit": "ns/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - B/op",
            "value": 1384516,
            "unit": "B/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkTask2/Real(day9) - allocs/op",
            "value": 19601,
            "unit": "allocs/op",
            "extra": "861 times\n4 procs"
          }
        ]
      }
    ]
  }
}