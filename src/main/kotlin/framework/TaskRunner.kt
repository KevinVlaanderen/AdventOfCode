package framework

object TaskRunner {
    fun <T : Any?>run(task: Task<T>): Result<T> = task.run()
}