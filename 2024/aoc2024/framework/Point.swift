struct Point {
    var x, y: Int

//    init(x: Int, y: Int) {
//        self.x = x
//        self.y = y
//    }
    
    func neighbour(direction: Direction) -> Point {
        var neighbour = self
        let offset = direction.asPoint()
        neighbour.x += offset.x
        neighbour.y += offset.y
        return neighbour
    }
}
