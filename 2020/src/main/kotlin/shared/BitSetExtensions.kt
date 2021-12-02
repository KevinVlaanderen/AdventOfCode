package shared

import java.util.*
import kotlin.math.pow

fun BitSet.toLong(): Long =
    this.stream().asLongStream().reduce(0L) { acc, i -> acc + 2.0.pow(i.toDouble()).toLong() }