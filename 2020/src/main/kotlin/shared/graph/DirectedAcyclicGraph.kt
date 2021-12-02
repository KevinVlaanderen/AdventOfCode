package shared.graph

class DirectedAcyclicGraph<N, E> {
    private val _nodes: MutableSet<Node<N, E>> = mutableSetOf()
    private val _edges: MutableSet<Edge<N, E>> = mutableSetOf()

    val nodes: Set<Node<N, E>>
        get() = _nodes.toSet()
    val edges: Set<Edge<N, E>>
        get() = _edges.toSet()

    fun addNode(id: String, data: N? = null): Node<N, E> {
        val newNode = Node<N, E>(id, data)
        val added = _nodes.add(newNode)
        return if (added) newNode else _nodes.find { it.id == id }!!
    }

    fun addEdge(from: String, to: String, data: E? = null) {
        val fromNode = _nodes.find { it.id == from }
        val toNode = _nodes.find { it.id == to }

        if (fromNode == null) throw RuntimeException("Node with id $from does not exist")
        if (toNode == null) throw RuntimeException("Node with id $to does not exist")

        addEdge(fromNode, toNode, data)
    }

    fun addEdge(from: Node<N, E>, to: Node<N, E>, data: E? = null) {
        val edge = Edge(from, to, data)
        from.addOutgoing(edge)
        to.addIncoming(edge)
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