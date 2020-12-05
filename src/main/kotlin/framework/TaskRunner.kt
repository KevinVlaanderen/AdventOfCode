package framework

import shared.getResource

object TaskRunner {
    fun <T : Any?> run(task: Task<T>, resourcePath: String): Result<T> =
        task.run(getResource(resourcePath))
}