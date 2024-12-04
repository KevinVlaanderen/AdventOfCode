protocol Grid {
    associatedtype T: Equatable
    
    var width: Int {get}
    var height: Int {get}
    
    func get(_ position: Point) -> T?
    func filter(_ value: T) -> [Point]
}

struct ArrayGrid<T: Equatable>: Grid {
    typealias T = T

    private var data: [T]
    
    let width: Int
    let height: Int
    
    init(_ data: [[T]]) {
        self.data = data.flatMap({ $0 })
        self.width = data.first!.count
        self.height = data.count
    }
    
    func get(_ position: Point) -> T? {
        let index = position.y*self.width+position.x
        if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
            return nil
        }
        return self.data[index]
    }
    
    func filter(_ value: T) -> [Point] {
        return data.enumerated().filter({ $0.element == value }).map({ item in
            let x = item.offset % self.width
            let y = item.offset / self.width
            return Point(x: x, y: y)
        })
    }
}
