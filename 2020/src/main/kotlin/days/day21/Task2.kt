package days.day21

import framework.Task
import java.net.URL

class Task2 : Task<String>() {
    override fun run(input: URL): Result<String> {
        val food = input.openStream()
            .bufferedReader()
            .lineSequence()
            .map { Food.fromString(it) }
            .toList()

        val allergenAnalysis = AllergenAnalysis(food)

        val result = allergenAnalysis.allergenPerIngredient
            .toList()
            .sortedBy { it.second }
            .joinToString(",") { it.first }

        return Result.success(result)
    }
}