package days.day10

import framework.Task
import shared.graph.DirectedAcyclicGraph
import shared.graph.Node
import java.net.URL

class Task2 : Task<Long>() {
    override fun run(input: URL): Result<Long> {
        val graph = DirectedAcyclicGraph<Nothing, Nothing>()

        input
            .openStream()
            .bufferedReader()
            .lineSequence()
            .plus("0")
            .forEach { graph.addNode(it) }

        for (from in graph.nodes)
            for (to in graph.nodes)
                if (to.id.toInt() - from.id.toInt() in 1..3)
                    graph.addEdge(from.id, to.id)

        val start = graph.getNode("0")!!
        val target = graph.nodes.maxByOrNull { it.id.toInt() }!!
        val result = pathsTillTarget(start, target, mutableMapOf())

        return Result.success(result)
    }

    private fun pathsTillTarget(
        start: Node<*, *>,
        target: Node<*, *>,
        cache: MutableMap<String, Long>
    ): Long = when (start) {
        target -> 1
        else -> start.outgoing
            .ifEmpty { return 0 }
            .sumOf {
                if (cache.containsKey(it.to.id))
                    cache[it.to.id]!!
                else
                    pathsTillTarget(it.to, target, cache)
            }.also { cache[start.id] = it }
    }
}
