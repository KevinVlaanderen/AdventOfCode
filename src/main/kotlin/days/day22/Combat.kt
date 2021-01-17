package days.day22

class Combat(private val decks: List<Deck>) {
    fun play(): Deck {
        while (decks.all { it.hasCards }) {
            val topCards = decks.map { it.getTopCard() }
            val winner = topCards.mapIndexed { index, card -> Pair(card, index) }.maxByOrNull { it.first }!!
            topCards.sorted().reversed().forEach { decks[winner.second].addCardToBottom(it) }
        }

        return decks.find { it.hasCards }!!
    }
}