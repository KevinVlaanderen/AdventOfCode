package days.day7

import framework.Task
import shared.graph.DirectedAcyclicGraph
import shared.graph.Node
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val graph = DirectedAcyclicGraph<Nothing, Capacity>()

        input
            .openStream()
            .bufferedReader()
            .useLines { lines -> lines.forEach { extendGraphFromLine(graph, it) } }

        val shinyGoldNode = graph.getNode("shiny gold")!!
        val result = calculateContents(shinyGoldNode)

        return Result.success(result)
    }

    private fun calculateContents(node: Node<Nothing, Capacity>): Int = node
        .outgoing
        .map { it.data!!.capacity + it.data.capacity * calculateContents(it.to) }
        .sum()
}


