package days.day18

class MultiplyExpression(private val leftHand: Expression, private val rightHand: Expression) : Expression {
    override fun evaluate(): Long = leftHand.evaluate() * rightHand.evaluate()
}