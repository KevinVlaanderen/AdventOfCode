package shared

fun <T> Collection<T>.generatePermutations(): Sequence<Pair<T, T>> {
    val input = this

    return sequence {
        for (first in input)
            for (second in input.minus(first)) yield(Pair(first, second))
    }
}