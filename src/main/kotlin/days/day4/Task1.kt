package days.day4

import framework.Task
import shared.FileUtil

class Task1 : Task<Int>(1) {
    override fun run(): Result<Int> {
        val input = FileUtil.readResource("/day4")

        val count = input
            .split("""\n{2,}""".toRegex())
            .map { block ->
                block
                    .lines()
                    .flatMap { it.split(' ') }
                    .filter { it.trim() != "" }
                    .associate { field ->
                        val keyValue = field.split(':')
                        keyValue[0] to keyValue[1]
                    }
            }
            .count { fields ->
                PassportField.values()
                    .filter { it.required }
                    .all { fields.containsKey(it.acronym) }
            }

        return Result.success(count)
    }
}


