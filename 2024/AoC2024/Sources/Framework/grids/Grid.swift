public protocol Grid<T>: Sequence where Element == GridItem<T> {
    associatedtype T: Equatable
    
    var width: Int {get}
    var height: Int {get}
    
    subscript(_ position: Point) -> T? {get set}
}

extension Grid {
    public func makeIterator() -> GridIterator<T> {
        return GridIterator<T>(grid: self)
    }
    
    public func neighbours(of point: Point, towards directions: [Direction]) -> [GridItem<T>] {
        directions.compactMap { neighbour(of: point, towards: $0) }
    }
    
    public func neighbour(of point: Point, towards direction: Direction) -> GridItem<T>? {
        let neighbourPosition = point.neighbour(direction: direction)
        let neighbour = self[neighbourPosition]
        
        guard let neighbour = neighbour else {
            return nil
        }
        
        return GridItem(position: neighbourPosition, value: neighbour)
    }
}

public struct GridIterator<T: Equatable>: IteratorProtocol {
    public typealias Element = GridItem<T>
    
    let grid: any Grid<T>
    var current = Point(x: 0, y: 0)
    
    init(grid: any Grid<T>) {
        self.grid = grid
    }
    
    mutating public func next() -> Element? {
        var result: Element?
        
        if let gridValue = grid[current] {
            result = Element(position: current, value: gridValue)
        }

        if current.x < grid.width-1 {
            current.x += 1
        } else {
            current.x = 0
            current.y += 1
        }
        
        return result
    }
}

public struct GridItem<T: Equatable> {
    public let position: Point
    public let value: T
}
