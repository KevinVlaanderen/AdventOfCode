import Foundation
import Benchmark
import Data

for testCase in getCases() {
    benchmark(testCase.description, settings: TimeUnit(.us)) {
        _ = try testCase.execute()
    }
}

Benchmark.main()
