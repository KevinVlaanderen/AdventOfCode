protocol Grid: Sequence {
    associatedtype T: Equatable
    
    var width: Int {get}
    var height: Int {get}
}

struct ArrayGrid<T: Equatable>: Grid, Sequence {
    typealias T = T

    private var data: [T]
    
    let width: Int
    let height: Int
    
    init(_ data: [[T]]) {
        self.data = data.flatMap({ $0 })
        self.width = data.first!.count
        self.height = data.count
    }
    
    init(width: Int, height: Int, defaultValue: T) {
        self.data = [T](repeating: defaultValue, count: width*height)
        self.width = width
        self.height = height
    }
    
    subscript(_ position: Point) -> T? {
        get {
            let index = position.y*self.width+position.x
            if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
                return nil
            }
            return self.data[index]
        }
        
        set(value) {
            guard let value = value else { return }
            
            let index = position.y*self.width+position.x
            if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
                return
            }
            self.data[index] = value
        }
    }
    
    func makeIterator() -> GridItemIterator<T> {
        return GridItemIterator(grid: self)
    }
}

struct GridItemIterator<T: Equatable>: IteratorProtocol {
    let grid: ArrayGrid<T>
    var current = Point(x: 0, y: 0)
    
    init(grid: ArrayGrid<T>) {
        self.grid = grid
    }
    
    mutating func next() -> GridItem<T>? {
        if current.x < grid.width-1 {
            current.x += 1
        } else if current.y < grid.height-1 {
            current.x = 0
            current.y += 1
        } else {
            return nil
        }
        
        if let gridValue = grid[current] {
            return GridItem(position: current, value: gridValue)
        } else {
            return nil
        }
    }
}

struct GridItem<T: Equatable> {
    let position: Point
    let value: T
}
