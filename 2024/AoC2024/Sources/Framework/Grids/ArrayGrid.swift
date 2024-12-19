public struct ArrayGrid<T: Equatable>: Grid {
    private var data: [T]
    
    public let width: Int
    public let height: Int
    
    public init(_ data: [[T]]) {
        self.data = data.flatMap({ $0 })
        self.width = data.first!.count
        self.height = data.count
    }
    
    public init(width: Int, height: Int, defaultValue: T) {
        self.data = [T](repeating: defaultValue, count: width*height)
        self.width = width
        self.height = height
    }

    public subscript(_ position: Point) -> T? {
        get {
            guard let index = indexOf(position) else { return nil }
            return self.data[index]
        }
        
        set(value) {
            guard let value = value, let index = indexOf(position) else { return }
            self.data[index] = value
        }
    }
    
    private func indexOf(_ position: Point) -> Int? {
        if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
            return nil
        }
        return position.y*self.width+position.x
    }
}
