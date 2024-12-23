import Foundation
import Framework
internal import Algorithms

public final class Day12: Day<Void, Int> {
    public override func perform(data: String, task: Task, param: P) throws -> Int {
        let farm = parse(data)
        
        return switch task {
        case .task1: farm.regions().reduce(0) { $0 + $1.plots.count*$1.fences(farm: farm) }
        case .task2: farm.regions().reduce(0) { $0 + $1.plots.count*$1.sides(farm: farm) }
        }
    }
    
    private func parse(_ data: String) -> Farm {
        Farm(grid: ArrayGrid(data.split(whereSeparator: \.isNewline).map { line in
            Array(line).map { character in character.asciiValue!}
        }))
    }
}

private struct Farm {
    private let grid: any Grid<Plant>
    
    private let directions = Direction.allCases
    
    init(grid: any Grid<Plant>) {
        self.grid = grid
    }
    
    func regions() -> [Region] {
        var seenPlots: Set<Point> = []
        var regions: [Region] = []
        
        for item in grid where !seenPlots.contains(item.position) {
            guard let plant = self.grid[item.position] else {
                continue
            }
            
            var plots = Set<Point>()
            findPlots(plant: plant, currentPlot: item.position, foundPlots: &plots)

            seenPlots.formUnion(plots)
            
            regions.append(Region(plots: plots, plant: plant))
        }
        
        return regions
    }
    
    private func findPlots(plant: Plant, currentPlot: Point, foundPlots: inout Set<Point>) {
        if foundPlots.contains(currentPlot) {
            return
        }
        
        foundPlots.insert(currentPlot)
        
        for direction in directions {
            let neighbour = currentPlot.neighbour(direction: direction)
            if let plot = grid[neighbour], plot == plant {
                findPlots(plant: plant, currentPlot: neighbour, foundPlots: &foundPlots)
            }
        }
    }
}

private struct Region {
    let plots: Set<Point>
    let plant: Plant
    
    func fences(farm: Farm) -> Int {
        let directions = Direction.allCases
        
        return plots.reduce(0) { result, plot in
            result + directions.count { direction in
                let neighbour = plot.neighbour(direction: direction)
                return !plots.contains(neighbour)
            }
        }
    }
    
    func sides(farm: Farm) -> Int {
        let directions = Direction.allCases
        
        let edges: [Edge: [Int]] = plots.reduce(into: [:]) { result, plot in
            for direction in directions {
                let neighbour = plot.neighbour(direction: direction)
                if !plots.contains(neighbour) {
                    var index: Int
                    var position: Int
                    
                    switch direction {
                    case .up, .down:
                        index = plot.y
                        position = plot.x
                    case .left, .right:
                        index = plot.x
                        position = plot.y
                    }
                    
                    let edge = Edge(direction: direction, index: index)
                    if result[edge] == nil {
                        result[edge] = [position]
                    } else {
                        result[edge]?.append(position)
                    }
                }
            }
        }
        
        return edges.reduce(0) { result, indices in
            let indexSet = IndexSet(indices.value)
            return result + indexSet.rangeView.count
        }
    }
    
    private struct Edge: Hashable {
        let direction: Direction
        let index: Int
    }
}

private typealias Plant = UInt8
