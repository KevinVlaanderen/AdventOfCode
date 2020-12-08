package days.day7

import shared.graph.DirectedAcyclicGraph

val outerPattern = """^(.*) bags contain (.*).$""".toRegex()
val innerPattern = """(\d+) (.+?) bags?""".toRegex()

fun extendGraphFromLine(
    graph: DirectedAcyclicGraph<Nothing, Capacity>,
    line: String
) {
    val (fromBagName, remainder) = outerPattern.find(line)?.destructured ?: return
    
    graph.addNode(fromBagName)

    innerPattern.findAll(remainder).forEach {
        val (capacity, toBagName) = it.destructured

        graph.addNode(toBagName)
        graph.addEdge(fromBagName, toBagName, Capacity(capacity.toInt()))
    }
}