package days.day4

import framework.Task
import shared.extractMap
import shared.toBlocks

object Task1 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val count = input
            .toBlocks()
            .map { it.extractMap() }
            .count { fields ->
                PassportField.values()
                    .filter { it.required }
                    .all { fields.containsKey(it.acronym) }
            }

        return Result.success(count)
    }
}


