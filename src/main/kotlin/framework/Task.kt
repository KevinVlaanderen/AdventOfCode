package framework

abstract class Task<T>(val number: Int) : Runnable<T>
