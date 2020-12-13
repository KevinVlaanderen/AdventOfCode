package days.day7

import framework.Task
import shared.graph.DirectedAcyclicGraph
import shared.graph.Direction
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val graph = DirectedAcyclicGraph<Nothing, Capacity>()

        input
            .openStream()
            .bufferedReader()
            .useLines { lines -> lines.forEach { extendGraphFromLine(graph, it) } }

        val shinyGoldNode = graph.getNode("shiny gold")!!
        val result = graph.findReachable(shinyGoldNode, Direction.UP).count()

        return Result.success(result)
    }
}


