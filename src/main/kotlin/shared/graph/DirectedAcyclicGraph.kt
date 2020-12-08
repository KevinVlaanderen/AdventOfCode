package shared.graph

class DirectedAcyclicGraph<N, E> {
    private val nodes: MutableSet<Node<N, E>> = mutableSetOf()
    private val edges: MutableSet<Edge<N, E>> = mutableSetOf()

    fun addNode(id: String, data: N? = null) = nodes.add(Node(id, data))

    fun addEdge(from: String, to: String, data: E? = null) {
        val fromNode = nodes.find { it.id == from }
        val toNode = nodes.find { it.id == to }

        if (fromNode == null) throw RuntimeException("Node with id $from does not exist")
        if (toNode == null) throw RuntimeException("Node with id $to does not exist")

        val edge = Edge(fromNode, toNode, data)
        fromNode.addOutgoing(edge)
        toNode.addIncoming(edge)
        edges.add(edge)
    }

    fun getNode(id: String) = nodes.find { it.id == id }

    fun findReachable(node: Node<N, E>, direction: Direction = Direction.DOWN): Set<Node<N, E>> {
        val result = when (direction) {
            Direction.UP -> node.incoming.toSet().map { it.from }.toSet()
            Direction.DOWN -> node.outgoing.toSet().map { it.to }.toSet()
        }

        return result + result.flatMap { findReachable(it, direction) }
    }
}