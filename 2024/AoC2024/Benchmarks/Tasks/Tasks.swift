import Foundation
import Benchmark
import Data

@MainActor
let benchmarks = {
    for c in cases {
        Benchmark(c.description) { benchmark in
            benchmark.configuration.timeUnits = .microseconds
            
            for _ in benchmark.scaledIterations {
                blackHole(try c.execute())
            }
        }
    }
}
