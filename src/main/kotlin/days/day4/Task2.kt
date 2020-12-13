package days.day4

import framework.Task
import shared.extractMap
import shared.toBlocks
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val requiredFields = PassportField.values().filter { it.required }

        val count = input
            .openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .count { block ->
                val fields = block.extractMap().mapKeys { PassportField.getPassportFieldByAcronym(it.key) }
                requiredFields.all { fields.containsKey(it) } && fields.all { it.key.validate(it.value) }
            }

        return Result.success(count)
    }
}


