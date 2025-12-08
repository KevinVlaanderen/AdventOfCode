package days.day21

class AllergenAnalysis(food: List<Food>) {
    private val allIngredients = food.map { it.ingredients }.reduce { acc, set -> acc.union(set) }
    private val allAllergens = food.map { it.allergens }.reduce { acc, set -> acc.union(set) }

    private val ingredientsPerAllergen = allAllergens.associateWith { allergen ->
        food
            .filter { it.allergens.contains(allergen) }
            .map { it.ingredients }
            .reduce { acc, set -> acc.intersect(set) }
    }

    val ingredientsWithoutAllergen =
        allIngredients - ingredientsPerAllergen.map { it.value }.reduce { acc, set -> acc.union(set) }

    val allergenPerIngredient: Map<String, String>
        get() {
            val allergenPerIngredient: MutableMap<String, String> = mutableMapOf()

            do {
                var foundNewMapping = false

                ingredientsPerAllergen.forEach {
                    val ingredients = it.value
                        .subtract(allergenPerIngredient.keys)
                        .subtract(ingredientsWithoutAllergen)

                    if (ingredients.size == 1) {
                        allergenPerIngredient[ingredients.first()] = it.key
                        foundNewMapping = true
                    }
                }
            } while (foundNewMapping)

            return allergenPerIngredient
        }
}