import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day8: Day {
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        var grid = parse(data)
        
        let groups = grid.filter({ $0.value.frequency != "." }).grouped(by: { $0.value.frequency })
        
        for group in groups {
            for items in group.value.combinations(ofCount: 2) {
                let dx = items[0].position.x - items[1].position.x
                let dy = items[0].position.y - items[1].position.y
                
                if dx != 0 && dy != 0 {
                    let ax1 = items[0].position.x + dx
                    let ay1 = items[0].position.y + dy
                    grid[Point(x: ax1, y: ay1)]?.antiNode = true
                    
                    let ax2 = items[1].position.x - dx
                    let ay2 = items[1].position.y - dy
                    grid[Point(x: ax2, y: ay2)]?.antiNode = true
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

    private struct Location: Equatable {
        let frequency: Character?
        var antiNode: Bool = false
    }
}
