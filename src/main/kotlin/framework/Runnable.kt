package framework

interface Runnable<out T> {
    fun run(): Result<T>
}