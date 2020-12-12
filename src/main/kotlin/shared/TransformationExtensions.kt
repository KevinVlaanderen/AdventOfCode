package shared

fun <T> T.transformUntilUnchanged(transform: (input: T) -> T): T {
    var previousResult: T
    var currentResult = this
    do {
        previousResult = currentResult
        currentResult = transform(previousResult)
    } while (previousResult != currentResult)

    return currentResult
}