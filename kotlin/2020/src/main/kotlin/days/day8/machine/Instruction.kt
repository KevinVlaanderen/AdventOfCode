package days.day8.machine

class Instruction(val pointer: Int, val operation: Operation, val argument: Int) {
    override fun toString() = "${operation.keyword} ${if (argument >= 0) "+" else ""}$argument"
}