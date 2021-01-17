package days.day22

class Deck(val name: String, cards: List<Int>) {
    private var _cards: List<Int> = cards

    val cards: List<Int>
        get() = _cards

    val hasCards: Boolean
        get() = _cards.isNotEmpty()

    fun getTopCard(): Int {
        val topCard = cards.first()
        _cards = cards.drop(1)
        return topCard
    }

    fun addCardToBottom(card: Int) {
        _cards = cards + card
    }

    companion object {
        fun fromInput(input: String): Deck {
            val lines = input.lineSequence()
            val name = lines.first()
            val cards = lines.drop(1).map { it.toInt() }.toList()

            return Deck(name, cards)
        }
    }
}