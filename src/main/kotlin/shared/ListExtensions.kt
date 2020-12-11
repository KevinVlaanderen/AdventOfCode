package shared

fun <T> List<T>.replace(newValue: T, block: (T) -> Boolean): List<T> = map {
    if (block(it)) newValue else it
}
