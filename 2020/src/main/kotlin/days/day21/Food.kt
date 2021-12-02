package days.day21

data class Food(val ingredients: Set<String>, val allergens: Set<String>) {
    companion object {
        private val entryPattern = """(.*) \(contains (.*)\)""".toRegex()

        fun fromString(input: String): Food {
            val (ingredientsString, allergensString) = entryPattern.find(input)?.destructured
                ?: throw Exception("Invalid input")

            val ingredients = ingredientsString.split(" ").toSet()
            val allergens = allergensString.split(",").map { it.trim() }.toSet()

            return Food(ingredients, allergens)
        }
    }
}
