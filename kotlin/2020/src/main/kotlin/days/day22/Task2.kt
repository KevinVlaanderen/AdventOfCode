package days.day22

import framework.Task
import shared.toBlocks
import java.net.URL

class Task2 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val decks = input.openStream()
            .bufferedReader()
            .readText()
            .toBlocks()
            .map { Deck.fromInput(it) }
            .toList()

        val winningDeck = RecursiveCombat(Pair(decks[0], decks[1])).play()
        val numberOfCards = winningDeck.cards.count()

        val result = winningDeck.cards.foldIndexed(0) { index, acc, card -> acc + card * (numberOfCards - index) }

        return Result.success(result)
    }
}