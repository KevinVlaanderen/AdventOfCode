package shared

fun Int.to36bitString(): String =
    this.toString(2).padStart(36, '0')