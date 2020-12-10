package shared.graph

class DirectedAcyclicGraph<N, E> {
    private val _nodes: MutableSet<Node<N, E>> = mutableSetOf()
    private val _edges: MutableSet<Edge<N, E>> = mutableSetOf()

    val nodes: Set<Node<N, E>>
        get() = _nodes.toSet()
    val edges: Set<Edge<N, E>>
        get() = _edges.toSet()

    fun addNode(id: String, data: N? = null) = _nodes.add(Node(id, data))

    fun addEdge(from: String, to: String, data: E? = null) {
        val fromNode = _nodes.find { it.id == from }
        val toNode = _nodes.find { it.id == to }

        if (fromNode == null) throw RuntimeException("Node with id $from does not exist")
        if (toNode == null) throw RuntimeException("Node with id $to does not exist")

        val edge = Edge(fromNode, toNode, data)
        fromNode.addOutgoing(edge)
        toNode.addIncoming(edge)
        _edges.add(edge)
    }

    fun getNode(id: String) = _nodes.find { it.id == id }

    fun findReachable(node: Node<N, E>, direction: Direction = Direction.DOWN): Set<Node<N, E>> {
        val result = when (direction) {
            Direction.UP -> node.incoming.toSet().map { it.from }.toSet()
            Direction.DOWN -> node.outgoing.toSet().map { it.to }.toSet()
        }

        return result + result.flatMap { findReachable(it, direction) }
    }
}