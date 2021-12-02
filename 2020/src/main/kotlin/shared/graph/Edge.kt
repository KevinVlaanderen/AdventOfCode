package shared.graph

data class Edge<N, E>(val from: Node<N, E>, val to: Node<N, E>, val data: E? = null)