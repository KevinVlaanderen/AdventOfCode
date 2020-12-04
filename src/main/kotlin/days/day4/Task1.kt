package days.day4

import framework.Task
import shared.extractDictionary
import shared.toBlocks

class Task1 : Task<Int>(1) {
    override fun run(input: String): Result<Int> {
        val count = input
            .toBlocks()
            .map { it.extractDictionary() }
            .count { fields ->
                PassportField.values()
                    .filter { it.required }
                    .all { fields.containsKey(it.acronym) }
            }

        return Result.success(count)
    }
}


