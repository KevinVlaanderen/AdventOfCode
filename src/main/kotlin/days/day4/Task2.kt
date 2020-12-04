package days.day4

import framework.Task
import shared.extractDictionary
import shared.toBlocks

class Task2 : Task<Int>(2) {
    override fun run(input: String): Result<Int> {
        val requiredFields = PassportField.values().filter { it.required }

        val count = input
            .toBlocks()
            .map { block -> block.extractDictionary().mapKeys { PassportField.getPassportFieldByAcronym(it.key) } }
            .count { fields -> requiredFields.all { fields.containsKey(it) } && fields.all { it.key.validate(it.value) } }

        return Result.success(count)
    }
}


