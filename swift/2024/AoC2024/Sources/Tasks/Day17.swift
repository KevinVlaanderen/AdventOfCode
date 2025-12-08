import Foundation
import Framework

public final class Day17: Day<Void, String> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        var (machine, rawInstructions) = try parse(data)

        switch task {
        case .task1:
            return machine.map({ String($0) }).joined(separator: ",")
        case .task2:
            guard let A = findA(machine: &machine, A: 1, index: 0, targetOutput: rawInstructions) else {
                return ""
            }

            return "\(A)"
        }
    }
    
    private func findA(machine: inout Machine, A: Int, index: Int, targetOutput: [Int]) -> Int? {
        let target = targetOutput[targetOutput.count-index-1]
        
        for i in 0..<8 {
            let A = (A+i)
            machine.A = A
            
            guard let output = machine.first(where: { _ in true }), output == target else {
                continue
            }
            if index == targetOutput.count-1 {
                return A
            }
            guard let A = findA(machine: &machine, A: A*8, index: index+1, targetOutput: targetOutput) else {
                continue
            }
            
            return A
        }
        return nil
    }

    private func parse(_ data: String) throws -> (Machine, [Int]) {
        let dataPattern = /Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: (.*)/
        
        guard let match = data.wholeMatch(of: dataPattern) else {
            throw AoCError.parseError("data did not match pattern")
        }
        let A = try toInt(match.output.1)
        let B = try toInt(match.output.2)
        let C = try toInt(match.output.3)
        
        let rawInstructions = try match.output.4.split(separator: ",").map(toInt)
        
        let instructions: [Instruction] = try rawInstructions.windows(ofCount: 2).striding(by: 2).map {
            try Instruction.parse(opCode: $0[$0.startIndex], value: $0[$0.endIndex-1])
        }
        
        return (Machine(A: A, B: B, C: C, instructions: instructions), rawInstructions)
    }
}

private struct Machine: Sequence {
    var A: Int
    var B: Int
    var C: Int
    let instructions: [Instruction]
    
    var output: [Int] = []
    var pointer = 0
    
    func makeIterator() -> MachineIterator {
        MachineIterator(machine: self)
    }
    
    struct MachineIterator: IteratorProtocol {
        var machine: Machine
        
        public init(machine: Machine) {
            self.machine = machine
        }
        
        mutating func next() -> Int? {
            do {
                while !machine.completed {
                    if let output = try machine.runNextInstruction() {
                        return output
                    }
                }
                return nil
            } catch {
                return nil
            }
        }
    }
    
    public var completed: Bool {
        pointer >= instructions.count
    }
    
    mutating func runNextInstruction() throws -> Int? {
        if completed {
            return nil
        }
        
        let instruction = instructions[pointer]
        var output: Int? = nil
        
        switch instruction {
        case .adv(.combo(let value)):   try A = divide(A, power: evaluateCombo(value))
        case .bxl(.literal(let value)): B = B ^ value //
        case .bst(.combo(let value)):   try B = evaluateCombo(value) % 8 //
        case .jnz(.literal(_))
            where A == 0:               break
        case .jnz(.literal(let value))
            where A != 0:               pointer = value/2; return nil
        case .bxc:                      B = B ^ C //
        case .out(.combo(let value)):   output = try evaluateCombo(value) % 8
        case .bdv(.combo(let value)):   try B = divide(A, power: evaluateCombo(value)) //
        case .cdv(.combo(let value)):   try C = divide(A, power: evaluateCombo(value))
        default:                        throw AoCError.taskError("invalid instruction \(instruction)")
        }
        
        pointer += 1
        
        return output
    }
    
    private func evaluateCombo(_ value: Int) throws -> Int {
        switch value {
        case 0, 1, 2, 3:                value
        case 4:                         A
        case 5:                         B
        case 6:                         C
        default:                        throw AoCError.taskError("invalid operand value \(value)")
        }
    }
    
    private func divide(_ numerator: Int, power: Int) -> Int {
        numerator/(1<<power)
    }
}

private enum Instruction {
    case adv(Operand)
    case bxl(Operand)
    case bst(Operand)
    case jnz(Operand)
    case bxc
    case out(Operand)
    case bdv(Operand)
    case cdv(Operand)
    
    static func parse(opCode: Int, value: Int) throws -> Instruction {
        return switch opCode {
        case 0:                         .adv(.combo(value))
        case 1:                         .bxl(.literal(value))
        case 2:                         .bst(.combo(value))
        case 3:                         .jnz(.literal(value))
        case 4:                         .bxc
        case 5:                         .out(.combo(value))
        case 6:                         .bdv(.combo(value))
        case 7:                         .cdv(.combo(value))
        default:                        throw AoCError.parseError("unknown instruction \(opCode) (value=\(value)")}
    }
}

private enum Operand {
    case literal(Int)
    case combo(Int)
}
