package framework

interface Task {
    fun run(): TaskResult

    val number: Int
    val description: String
}