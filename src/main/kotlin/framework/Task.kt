package framework

import framework.interfaces.Runnable

abstract class Task<T> : Runnable<T> {
    open val inputOverride: String? = null
}
