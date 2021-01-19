package days.day25

import framework.Task
import java.net.URL

class Task1 : Task<Int>() {
    override fun run(input: URL): Result<Int> {
        val (cardPublicKey, doorPublicKey) = input.openStream().bufferedReader().lineSequence()
            .take(2)
            .map { it.toInt() }
            .toList()

        val cardLoopSize = transform(7).takeWhile { it != cardPublicKey }.count()
        val result = transform(doorPublicKey).take(cardLoopSize + 1).last()

        return Result.success(result)
    }

    private fun transform(subjectNumber: Int) =
        generateSequence(1L) { (it * subjectNumber) % 20201227 }.map { it.toInt() }

}