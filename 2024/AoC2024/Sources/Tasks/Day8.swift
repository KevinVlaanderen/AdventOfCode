import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day8: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> Int {
        var grid = parse(data)
        
        let groups = grid.filter({ $0.value.frequency != "." }).grouped(by: { $0.value.frequency })
        
        for group in groups {
            for items in group.value.combinations(ofCount: 2) {
                let dx = items[0].position.x - items[1].position.x
                let dy = items[0].position.y - items[1].position.y
                
                if dx != 0 && dy != 0 {
                    switch param {
                    case .task1:
                        setAntiNode(grid: &grid, position: items[0].position, dx: dx, dy: dy)
                        setAntiNode(grid: &grid, position: items[1].position, dx: -dx, dy: -dy)
                    case .task2:
                        for mult in 0... {
                            if !setAntiNode(grid: &grid, position: items[0].position, dx: dx*mult, dy: dy*mult) {
                                break
                            }
                        }
                        
                        for mult in 0... {
                            if !setAntiNode(grid: &grid, position: items[0].position, dx: -dx*mult, dy: -dy*mult) {
                                break
                            }
                        }
                    }
                }
            }
        }
        
        return grid.count { item in item.value.antiNode }
    }

    private func parse(_ data: String) -> any Grid<Location> {
        ArrayGrid(data.split(whereSeparator: \.isNewline).map { line in
            Array(line).map { character in
                Location(frequency: character)
            }
        })
    }
    
    @discardableResult
    private func setAntiNode(grid: inout any Grid<Location>, position: Point, dx: Int = 0, dy: Int = 0) -> Bool {
        let x = position.x + dx
        let y = position.y + dy
        let point = Point(x: x, y: y)
        if grid[point] == nil {
            return false
        }
        grid[point]?.antiNode = true
        return true
    }

    private struct Location: Equatable {
        let frequency: Character?
        var antiNode: Bool = false
    }
}
