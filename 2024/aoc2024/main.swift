import Foundation
import aoc2024Data
import Benchmark

for testCase in getCases() {
    benchmark(testCase.testDescription, settings: TimeUnit(.us)) {
        try testCase.execute()
    }
}

Benchmark.main()
