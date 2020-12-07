package days.day7

import framework.Task
import shared.graph.DirectedAcyclicGraph
import shared.graph.Direction
import java.net.URL

val outerPattern = """^(.*) bags contain (.*).$""".toRegex()
val innerPattern = """(\d+) (.+?) bags?""".toRegex()

object Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val graph = DirectedAcyclicGraph<Bag, Capacity>()

        input
            .openStream()
            .bufferedReader()
            .useLines { lines ->
                lines
                    .forEach { line ->
                        val (fromBagName, remainder) = outerPattern.find(line)?.destructured ?: return@forEach

                        val fromBag = Bag(fromBagName)
                        graph.addNode(fromBag)

                        innerPattern.findAll(remainder).forEach {
                            val (capacity, toBagName) = it.destructured

                            val toBag = Bag(toBagName)
                            graph.addNode(toBag)

                            graph.addEdge(fromBag, toBag, Capacity(capacity.toInt()))
                        }
                    }
            }

        val result = graph.findReachable(Bag("shiny gold"), Direction.UP).count()

        return Result.success(result)
    }
}


