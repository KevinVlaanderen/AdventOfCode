package days.day4

import framework.Task
import shared.extractMap
import shared.toBlocks

object Task2 : Task<Int>() {
    override fun run(input: String): Result<Int> {
        val requiredFields = PassportField.values().filter { it.required }

        val count = input
            .toBlocks()
            .map { block -> block.extractMap().mapKeys { PassportField.getPassportFieldByAcronym(it.key) } }
            .count { fields -> requiredFields.all { fields.containsKey(it) } && fields.all { it.key.validate(it.value) } }

        return Result.success(count)
    }
}


