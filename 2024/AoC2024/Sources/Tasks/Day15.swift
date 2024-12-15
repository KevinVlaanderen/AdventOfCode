import Foundation
import Framework
internal import Algorithms

public struct Day15: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
        
    public func perform() throws -> Int {
        var (grid, moves) = parse(data)
        
        guard let robot = grid.first(where: { $0.value == .robot}) else {
            throw AoCError.parseError("robot not found")
        }
        
        var robotPosition = robot.position
        
        for move in moves {
            var neighbourPosition = robotPosition.neighbour(direction: move)
            var neighbour: Tile? = grid[neighbourPosition]
            
            switch neighbour {
            case .empty:
                grid[robotPosition] = .empty
                robotPosition = neighbourPosition
                grid[robotPosition] = .robot
            case .box:
                repeat {
                    neighbourPosition = neighbourPosition.neighbour(direction: move)
                    neighbour = grid[neighbourPosition]
                } while neighbour == .box
                if neighbour == .empty {
                    let firstNeighbourPosition = robotPosition.neighbour(direction: move)
                    grid[neighbourPosition] = .box
                    grid[robotPosition] = .empty
                    robotPosition = firstNeighbourPosition
                    grid[firstNeighbourPosition] = .robot
                }
            case .wall, nil: continue
            case .robot: throw Day15Error.invalidMove
            }
        }
        
        return grid.filter({ $0.value == .box }).reduce(0, { $0 + $1.position.x+$1.position.y*100})
    }

    nonisolated(unsafe)
    private static let emptyLinePattern = /\n\n/
    
    private func parse(_ data: String) -> (any Grid<Tile>, [Direction]) {
        let blocks = data.split(separator: Day15.emptyLinePattern)
        let grid = ArrayGrid(blocks[0].split(whereSeparator: \.isNewline).map { line in
            line.map { character in
                switch character {
                case ".": Tile.empty
                case "#": Tile.wall
                case "O": Tile.box
                case "@": Tile.robot
                default: fatalError("unknown grid item: \(character)")}
            }
        })
        
        let moves = blocks[1].split(whereSeparator: \.isNewline).flatMap { line in
            line.map {
                switch $0 {
                case "^": Direction.N
                case ">": Direction.E
                case "v": Direction.S
                case "<": Direction.W
                    default: fatalError("unknown direction: \($0)")}
            }
        }
        
        return (grid, moves)
    }
    
    private enum Tile {
        case wall, box, robot, empty
    }
}

enum Day15Error: Error {
    case invalidMove
}
