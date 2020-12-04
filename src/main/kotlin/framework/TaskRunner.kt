package framework

import shared.readResource

object TaskRunner {
    fun <T : Any?> run(task: Task<T>, resourcePath: String): Result<T> =
        task.run(readResource(resourcePath))
}