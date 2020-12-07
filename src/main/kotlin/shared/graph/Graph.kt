package shared.graph

class DirectedAcyclicGraph<N : Node, E : Any> {
    private val nodes: MutableSet<N> = mutableSetOf()
        get() = field
    private val edges: MutableSet<Edge<N, E>> = mutableSetOf()
        get() = field

    fun addNode(node: N) = nodes.add(node)
    fun addEdge(from: N, to: N, data: E? = null) = edges.add(Edge(from, to, data))

    fun findReachable(node: N, direction: Direction = Direction.DOWN): Set<N> {
        val result = when (direction) {
            Direction.UP -> leadingTo(node)
            Direction.DOWN -> leadingFrom(node)
        }.toMutableSet()

        result.addAll(result.flatMap { findReachable(it, direction) })

        return result
    }

    private fun leadingTo(node: N): Set<N> = edges.filter { it.to == node }.map { it.from }.toSet()
    private fun leadingFrom(node: N): Set<N> = edges.filter { it.from == node }.map { it.to }.toSet()
}