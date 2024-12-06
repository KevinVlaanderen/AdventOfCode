struct DictGrid<T: Equatable>: Grid {
    private var data: [Point:T] = [:]
    
    let width: Int
    let height: Int
    let emptyValue: T
    
    init(_ data: [[T]], emptyValue: T) {
        self.data = data.enumerated().reduce(into: [:]) { (result, current) in
            let y = current.offset
            for item in current.element.enumerated() {
                let x = item.offset
                if item.element != emptyValue {
                    result[Point(x: x, y: y)] = item.element
                }
            }
        }
        self.width = data.first!.count
        self.height = data.count
        self.emptyValue = emptyValue
    }

    subscript(_ position: Point) -> T? {
        get {
            if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
                return nil
            }
            return self.data[Point(x: position.x, y: position.y)] ?? emptyValue
        }
        
        set(value) {
            guard let value = value else { return }
            
            if position.x < 0 || position.y < 0 || position.x >= self.width || position.y >= self.height {
                return
            }
            
            if value == emptyValue {
                self.data.removeValue(forKey: Point(x: position.x, y: position.y))
            } else {
                self.data[Point(x: position.x, y: position.y)] = value
            }
        }
    }
}
