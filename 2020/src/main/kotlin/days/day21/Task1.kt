package days.day21

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val food = input.openStream()
            .bufferedReader()
            .lineSequence()
            .map { Food.fromString(it) }
            .toList()

        val allergenAnalysis = AllergenAnalysis(food)

        val result = allergenAnalysis.ingredientsWithoutAllergen.sumOf { ingredient ->
            food.count { listing ->
                listing.ingredients.contains(ingredient)
            }
        }

        return Result.success(result)
    }
}