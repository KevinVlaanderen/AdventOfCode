package days.day4

import framework.Task
import shared.extractMap
import shared.toBlocks
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val count = input
            .openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .count { block ->
                val fields = block.extractMap()
                PassportField.values()
                    .filter { it.required }
                    .all { fields.containsKey(it.acronym) }
            }

        return Result.success(count)
    }
}


