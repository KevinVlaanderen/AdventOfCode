package framework

abstract class Task<T> : Runnable<T> {
    open val inputOverride: String? = null
}
