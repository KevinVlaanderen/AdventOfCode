import Foundation
import aoc2024Data
import Benchmark

for testCase in getCases() {
    benchmark(testCase.testDescription) {
        Task {
            try await testCase.execute()
        }
    }
}

@main
class AOC2024 {
    static func main() {
        Benchmark.main()
    }
}
