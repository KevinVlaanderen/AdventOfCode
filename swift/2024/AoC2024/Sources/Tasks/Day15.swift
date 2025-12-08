import Foundation
import Framework

public final class Day15: Day<Int, Int> {
    public override func perform(data: String, task: Task, param: P) throws -> Int {
        var (room, moves) = try parse(data, param: param)
        
        guard let robotID = room.tileMap.first(where: { $0.value.type == .robot})?.key else {
            throw AoCError.parseError("robot not found")
        }

        for direction in moves {
            try room.move(robotID, to: direction)
            
//            print(direction)
//            room.draw()
        }
        
        return room.tileMap.reduce(0) { result, tile in
            if tile.value.type == .box {
                return result + tile.value.start.x + tile.value.start.y*100
            }
            return result
        }
    }

    private func parse(_ data: String, param: P) throws -> (Room, [Direction]) {
        let blocks = data.blocks

        let (grid, moves) = (try parseGrid(data: blocks[0]), try parseMoves(data: blocks[1]))
        
        let room = Room(grid: grid, multiplier: param)
        
        return (room, moves)
    }
    
    private func parseGrid(data: Substring) throws -> any Grid<TileType> {
        try ArrayGrid(data.lines.map { line throws in
           try line.map { character throws in
                switch character {
                case ".": TileType.empty
                case "#": TileType.wall
                case "O": TileType.box
                case "@": TileType.robot
                default: throw AoCError.parseError("unknown grid item: \(character)")}
            }
        })
    }
    
    private func parseMoves(data: Substring) throws -> [Direction] {
        try data.lines.flatMap { line throws in
            try line.map { character throws in
                switch character {
                case "^": .up
                case ">": .right
                case "v": .down
                case "<": .left
                default: throw AoCError.parseError("unknown direction: \(character)")}
            }
        }
    }
}

private enum TileType: Equatable {
    case wall, box, robot, empty
}

private struct Room {
    var grid: any Grid<TileID>
    var tileMap: [TileID: Tile] = [:]
    
    init(grid: any Grid<TileType>, multiplier: Int) {
        var tileID = 0
        
        self.grid = ArrayGrid(width: grid.width*multiplier, height: grid.height, defaultValue: 0)
        
        for item in grid {
            let position = Point(x: item.position.x*multiplier, y: item.position.y)
            let type = item.value
            
            switch type {
            case .robot:
                let tile = Tile(type: .robot, size: 1, start: position)
                self.tileMap[tileID] = tile
                self.grid[position] = tileID
                
                if multiplier == 1 {
                    continue
                }
                
                for i in 1..<multiplier {
                    let emptyPosition = Point(x: item.position.x*multiplier+i, y: item.position.y)
                    self.grid[emptyPosition] = -1
                }
            case .empty:
                for i in 0...multiplier {
                    let emptyPosition = Point(x: item.position.x*multiplier+i, y: item.position.y)

                    self.grid[emptyPosition] = -1
                }
            default:
                let tile = Tile(type: type, size: multiplier, start: position)
                self.tileMap[tileID] = tile
                
                for i in 0...multiplier {
                    let tilePosition = Point(x: item.position.x*multiplier+i, y: item.position.y)
                    self.grid[tilePosition] = tileID
                }
            }
            
            tileID += 1
        }
    }
    
    func tileAt(_ position: Point?) -> Tile? {
        guard let position = position, let tileID = grid[position] else {
            return nil
        }
        
        return tileID >= 0 ? tileMap[tileID] : nil
    }
    
    func nextTo(tile: Tile, direction: Direction) throws -> Set<TileID> {
        switch direction {
        case .up, .down:
            return tile.positions.reduce(into: []) { result, position in
                let neighbourPosition = position.neighbour(direction: direction)
                guard let neighbourID = grid[neighbourPosition] else {
                    return
                }
                if neighbourID >= 0 {
                    result.insert(neighbourID)
                }
            }
        case .right:
            guard let tileID = grid[tile.end.neighbour(direction: direction)] else {
                return []
            }
            return tileID >= 0 ? Set<TileID>(arrayLiteral: tileID) : []
        case .left:
            guard let tileID = grid[tile.start.neighbour(direction: direction)] else {
                return []
            }
            return tileID >= 0 ? Set<TileID>(arrayLiteral: tileID) : []
        }
    }
    
    mutating func move(_ tileID: TileID, to direction: Direction) throws {
        let (movingTileIDs, canMove) = try checkMove(from: tileID, direction: direction)
        
        if !canMove {
            return
        }
        
        var emptied: Set<Point> = []
        var filled: Set<Point> = []
        
        for movingTileID in movingTileIDs {
            guard var tile = tileMap[movingTileID] else {
                continue
            }
            
            emptied.formUnion(tile.positions)
            
            switch direction {
            case .up: tile.start.y -= 1
            case .right: tile.start.x += 1
            case .down: tile.start.y += 1
            case .left: tile.start.x -= 1
            }
            
            let newPositions = tile.positions
            filled.formUnion(tile.positions)
            
            for newPosition in newPositions {
                grid[newPosition] = movingTileID
            }
            
            tileMap[movingTileID] = tile
        }
        
        for emptiedPosition in emptied.subtracting(filled) {
            self.grid[emptiedPosition] = -1
        }
    }
    
    func checkMove(from tileID: TileID, direction: Direction) throws -> (Set<TileID>, Bool) { // moving
        if tileID == -1 {
            throw AoCError.taskError("invalid move")
        }
        guard let tile = tileMap[tileID] else {
            throw AoCError.taskError("invalid move")
        }
        
        let neighbours = try nextTo(tile: tile, direction: direction)
        
        if neighbours.count == 0 {
            return ([tileID], true)
        }
        
        return try neighbours.reduce(([tileID], true)) { result, neighbourID throws in
            if !result.1 {
                return ([], false)
            } else if neighbourID == -1 {
                return result
            }

            guard let neighbour = tileMap[neighbourID] else {
                throw AoCError.taskError("invalid move")
            }
            
            switch neighbour.type {
            case .wall:
                return ([], false)
            case .box:
                let (tilesToMove, canMove) = try checkMove(from: neighbourID, direction: direction)
                return (result.0.union(tilesToMove), canMove)
            default:
                throw AoCError.taskError("invalid move")
            }
        }
    }
    
    func draw() {
        print(String(repeating: "-", count: grid.width))
        for y in 0..<grid.height {
            let line = (0..<grid.width).map {
                let tileID = grid[Point(x: $0, y: y)]!
                if tileID == -1 {
                    return "."
                }
                guard let tile = tileMap[tileID] else {
                    return "x"
                }
                switch tile.type {
                case .box: return "O"
                case .wall: return "#"
                case .robot: return "@"
                default: return "x"
                }
            }.joined()
            print("|" + line + "|")
        }
        print(String(repeating: "-", count: grid.width))
    }
}

private struct Tile {
    let type: TileType
    let size: Int
    var start: Point
    var end: Point { Point(x: start.x+size-1, y: start.y)}
    
    var positions: [Point] {
        (start.x..<start.x+size).map({ Point(x: $0, y: start.y)})
    }
}

private typealias TileID = Int
