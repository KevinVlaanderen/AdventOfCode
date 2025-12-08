package days.day22

class RecursiveCombat(private val decks: Pair<Deck, Deck>) {
    private val previousStates: MutableList<Pair<List<Int>, List<Int>>> = mutableListOf()

    fun play(): Deck {
        while (decks.first.hasCards && decks.second.hasCards) {
            if (previousStates.any { state -> state.first == decks.first.cards && state.second == decks.second.cards }) {
                return decks.first
            } else {
                previousStates.add(Pair(decks.first.cards.toList(), decks.second.cards.toList()))
            }

            val topCards = Pair(decks.first.drawTopCard(), decks.second.drawTopCard())

            val winningDeck =
                if (decks.first.cards.count() >= topCards.first && decks.second.cards.count() >= topCards.second) {
                    val winner = RecursiveCombat(
                        Pair(
                            Deck(decks.first.name, decks.first.cards.take(topCards.first)),
                            Deck(decks.second.name, decks.second.cards.take(topCards.second))
                        )
                    ).play()
                    if (decks.first.name == winner.name) decks.first else decks.second
                } else {
                    if (topCards.first > topCards.second) decks.first else decks.second
                }

            if (winningDeck.name == decks.first.name) {
                decks.first.addCardToBottom(topCards.first).addCardToBottom(topCards.second)
            } else {
                decks.second.addCardToBottom(topCards.second).addCardToBottom(topCards.first)
            }
        }

        return if (decks.first.hasCards) decks.first else decks.second
    }
}