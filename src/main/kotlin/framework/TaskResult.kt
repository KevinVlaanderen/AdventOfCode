package framework

class TaskResult(val answer: String) {
    constructor(answer: Int) : this(answer.toString())
}