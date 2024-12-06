protocol Grid<T>: Sequence where Element == GridItem<T> {
    associatedtype T: Equatable
    
    var width: Int {get}
    var height: Int {get}
    
    subscript(_ position: Point) -> T? {get set}
    
    func makeIterator() -> GridIterator<T>
}

extension Grid {
    func makeIterator() -> GridIterator<T> {
        return GridIterator<T>(grid: self)
    }
}

struct GridIterator<T: Equatable>: IteratorProtocol {
    typealias Element = GridItem<T>
    
    let grid: any Grid<T>
    var current = Point(x: 0, y: 0)
    
    init(grid: any Grid<T>) {
        self.grid = grid
    }
    
    mutating func next() -> Element? {
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

struct GridItem<T: Equatable> {
    let position: Point
    let value: T
}


