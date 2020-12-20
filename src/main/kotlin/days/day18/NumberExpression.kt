package days.day18

class NumberExpression(private val value: Long) : Expression {
    override fun evaluate(): Long = value
}