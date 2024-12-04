struct Point: Hashable {
    var x, y: Int

    func neighbour(direction: Direction) -> Point {
        var neighbour = self
        let offset = direction.asPoint()
        neighbour.x += offset.x
        neighbour.y += offset.y
        return neighbour
    }
}
