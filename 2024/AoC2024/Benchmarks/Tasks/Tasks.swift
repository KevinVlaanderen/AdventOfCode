import Benchmark
import Framework
import Data

@MainActor
let benchmarks = {
    let data = Data()
    
    for c in cases {
        Benchmark(c.description) { benchmark in
            benchmark.configuration.timeUnits = .microseconds
            
            for _ in benchmark.scaledIterations {
                blackHole(try c.execute(data: data))
            }
        }
    }
}
