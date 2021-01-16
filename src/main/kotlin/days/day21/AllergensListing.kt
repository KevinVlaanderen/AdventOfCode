package days.day21

data class AllergensListing(val ingredients: Set<String>, val allergens: Set<String>) {
    companion object {
        fun fromString(input: String): AllergensListing {
            val (ingredientsString, allergensString) = entryPattern.find(input)?.destructured
                ?: throw Exception("Invalid input")

            val ingredients = ingredientsString.split(" ").toSet()
            val allergens = allergensString.split(",").map { it.trim() }.toSet()

            return AllergensListing(ingredients, allergens)
        }
    }
}
