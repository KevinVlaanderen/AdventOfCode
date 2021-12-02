package days.day22

class Combat(private val decks: Pair<Deck, Deck>) {
    fun play(): Deck {
        while (decks.first.hasCards && decks.second.hasCards) {
            val topCards = Pair(decks.first.drawTopCard(), decks.second.drawTopCard())

            if (topCards.first > topCards.second) {
                decks.first.addCardToBottom(topCards.first).addCardToBottom(topCards.second)
            } else {
                decks.second.addCardToBottom(topCards.second).addCardToBottom(topCards.first)
            }
        }

        return if (decks.first.hasCards) decks.first else decks.second
    }
}