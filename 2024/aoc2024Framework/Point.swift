public struct Point: Hashable {
    var x, y: Int

    public func neighbour(direction: Direction, distance: Int = 1) -> Point {
        var neighbour = self
        let offset = direction.asPoint()
        neighbour.x += offset.x * distance
        neighbour.y += offset.y * distance
        return neighbour
    }
}
