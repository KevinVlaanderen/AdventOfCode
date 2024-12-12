import Foundation
import Framework
internal import Algorithms
internal import SwiftGraph

public struct Day12: Day {
    public init() {}
        
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let farm = parse(data)
        
        return farm.regions().reduce(0) { $0 + $1.plots.count*$1.fences(farm: farm) }
    }
    
    private func parse(_ data: String) -> Farm {
        let lines = data.split(whereSeparator: \.isNewline)
        let width = lines[0].count
        let height = lines.count
        let plots = UnweightedGraph<Plant>()
        
        for y in 0..<height {
            let characters = Array(lines[y])
            
            for x in 0..<width {
                let _ = plots.addVertex(characters[x].asciiValue!)
            }
        }
        
        for index in plots.vertices.indices {
            let x = index % width
            let y = index / width
            
            let current = plots[index]
            
            if x>0 {
                let neighbourIndex = width*y+x-1
                if plots[neighbourIndex] == current {
                    plots.addEdge(fromIndex: index, toIndex: neighbourIndex)
                }
            }
            if y>0 {
                let neighbourIndex = width*(y-1)+x
                if plots[neighbourIndex] == current {
                    plots.addEdge(fromIndex: index, toIndex: neighbourIndex)
                }            }
            if x<width-1 {
                let neighbourIndex = width*y+x+1
                if plots[neighbourIndex] == current {
                    plots.addEdge(fromIndex: index, toIndex: neighbourIndex)
                }            }
            if y<height-1 {
                let neighbourIndex = width*(y+1)+x
                if plots[neighbourIndex] == current {
                    plots.addEdge(fromIndex: index, toIndex: neighbourIndex)
                }
            }
        }

        return Farm(width: width, height: height, plots: plots)
    }
    
    private struct Farm {
        let width: Int
        let height: Int
        let plots: UnweightedGraph<Plant>
        
        func regions() -> [Region] {
            var seenPlots: Set<Int> = []
            var regions: [Region] = []
            
            for currentPlot in 0..<(self.width*self.height) where !seenPlots.contains(currentPlot) {
                let plant = self.plots[currentPlot]
                
                var plots = Set<Plot>()
                findPlots(currentPlot: currentPlot, plots: &plots)

                seenPlots.formUnion(plots)
                
                regions.append(Region(plots: plots, plant: plant))
            }
            
            return regions
        }
        
        private func findPlots(currentPlot: Plot, plots: inout Set<Plot>) {
            if plots.contains(currentPlot) {
                return
            }
            
            plots.insert(currentPlot)
            
            for neighbour in self.plots.edgesForIndex(currentPlot) {
                findPlots(currentPlot: neighbour.v, plots: &plots)
            }
        }
    }
    
    private struct Region {
        let plots: Set<Plot>
        let plant: Plant
        
        func fences(farm: Farm) -> Int {
            var fences = 0
            for plot in plots {
                fences += 4 - farm.plots.neighborsForIndex(plot).count { $0 == plant } / 2
            }
            return fences
        }
    }
    
    private typealias Plot = Int
    private typealias Plant = UInt8
}
