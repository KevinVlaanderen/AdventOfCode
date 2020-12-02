package framework

object TaskRunner {
    fun run(task: Task): TaskResult = task.run()
}