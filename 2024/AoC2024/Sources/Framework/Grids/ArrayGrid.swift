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
            if let index = indexOf(position) {
                return self.data[index]
            }
            return nil
        }
        
        set(value) {
            guard let value = value else { return }
            
            if let index = indexOf(position) {
                self.data[index] = value
            }
        }
    }
    
    private func indexOf(_ position: Point) -> Int? {
        if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
            return nil
        }
        return position.y*self.width+position.x
    }
}
