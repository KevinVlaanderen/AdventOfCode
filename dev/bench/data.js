window.BENCHMARK_DATA = {
  "lastUpdate": 1766497089425,
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
      }
    ]
  }
}