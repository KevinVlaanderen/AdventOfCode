package days.day8.machine

class ExecutionState(
    val stack: List<Instruction>,
    val registers: Map<String, Int>
)