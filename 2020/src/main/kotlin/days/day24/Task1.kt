package days.day24

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val floor = Floor()

        input.openStream().bufferedReader().lineSequence()
            .map { followSteps(generateSteps(it)) }
            .forEach { floor.flip(it) }

        val result = floor.tiles.count { it.side == Side.BLACK }

        return Result.success(result)
    }
}