package shared

fun <T> List<T>.replace(newValue: T, block: (T) -> Boolean): List<T> = map {
    if (block(it)) newValue else it
}

fun <S, T> List<S>.cartesianProduct(other: List<T>) = this.flatMap {
    List(other.size) { i -> Pair(it, other[i]) }
}