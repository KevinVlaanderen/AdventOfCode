package shared

fun <T> List<List<T>>.mirrorVertical(): List<List<T>> = this.map { it.reversed() }

fun <T> List<List<T>>.rotateCCW() = (this[0].size - 1 downTo 0).map { this.map { line -> line[it] } }

fun <T> List<List<T>>.rotateCW() = (this[0].indices).map { this.reversed().map { line -> line[it] } }

fun <T> List<List<T>>.flip() = this.reversed().map { it.reversed() }

fun <T> List<List<T>>.count(target: T) = this.sumOf { line -> line.count { it == target } }