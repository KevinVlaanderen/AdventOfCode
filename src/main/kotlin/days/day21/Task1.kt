package days.day21

import framework.Task
import java.net.URL

val entryPattern = """(.*) \(contains (.*)\)""".toRegex()

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val allergensListings = input.openStream()
            .bufferedReader()
            .lineSequence()
            .map { AllergensListing.fromString(it) }
            .toList()

        val allIngredients = allergensListings.map { it.ingredients }.reduce { acc, set -> acc.union(set) }
        val ingredientsPerAllergen = mutableMapOf<String, Set<String>>()

        allergensListings.forEach { listing ->
            listing.allergens.forEach { allergen ->
                if (ingredientsPerAllergen.containsKey(allergen))
                    ingredientsPerAllergen[allergen] = ingredientsPerAllergen[allergen]!!.intersect(listing.ingredients)
                else
                    ingredientsPerAllergen[allergen] = listing.ingredients
            }
        }

        val ingredientsWithoutAllergen =
            allIngredients - ingredientsPerAllergen.map { it.value }.reduce { acc, set -> acc.union(set) }

        val result = ingredientsWithoutAllergen.sumOf { ingredient ->
            allergensListings.count { listing ->
                listing.ingredients.contains(ingredient)
            }
        }

        return Result.success(result)
    }
}