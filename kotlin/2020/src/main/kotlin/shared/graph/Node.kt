package shared.graph

data class Node<N, E>(val id: String, val data: N? = null) {
    private val _incoming: MutableSet<Edge<N, E>> = mutableSetOf()
    private val _outgoing: MutableSet<Edge<N, E>> = mutableSetOf()

    val incoming: Set<Edge<N, E>>
        get() = _incoming.toSet()
    val outgoing: Set<Edge<N, E>>
        get() = _outgoing.toSet()

    internal fun addIncoming(edge: Edge<N, E>) = _incoming.add(edge)
    internal fun addOutgoing(edge: Edge<N, E>) = _outgoing.add(edge)
}
