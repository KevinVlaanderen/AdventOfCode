import Foundation
internal import Algorithms
import Framework

public final class Day17: Day<Task, String> {
    public override func perform() throws -> R {
        var (machine, instructions, rawInstructions) = try parse(data: data)

        switch param {
        case .task1:
            while machine.pointer < instructions.count {
                try machine.runInstruction(instructions[machine.pointer])
            }
            
            return machine.output.map({ String($0) }).joined(separator: ",")
        case .task2:
            var A = 1<<8
            var newMachine: Machine

            Aloop: repeat {
                A += 1
                newMachine = machine
                newMachine.A = A
                while newMachine.pointer < instructions.count {
                    try newMachine.runInstruction(instructions[newMachine.pointer])
                    if !rawInstructions.starts(with: newMachine.output) {
                        continue Aloop
                    }
                }
            } while newMachine.output != rawInstructions
            
            return "\(A)"
        }
    }
    
    nonisolated(unsafe)
    private static let dataPattern = /Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: (.*)/
    
    private func parse(data: String) throws -> (Machine, [Instruction], [Int]) {
        guard let match = data.wholeMatch(of: Day17.dataPattern) else {
            throw AoCError.parseError("data did not match pattern")
        }
        let A = Int(match.output.1)!
        let B = Int(match.output.2)!
        let C = Int(match.output.3)!
        
        let rawInstructions = match.output.4.split(separator: ",").map({ Int($0)! })
        
        let instructions: [Instruction] = try rawInstructions.windows(ofCount: 2).striding(by: 2).map {
            let opCode: Int = $0[$0.startIndex]
            let value = $0[$0.endIndex-1]
            
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
        
        return (Machine(A: A, B: B, C: C), instructions, rawInstructions)
    }
    
    private struct Machine {
        var A: Int
        var B: Int
        var C: Int
        var output: [Int] = []
        var pointer = 0
        
        mutating func runInstruction(_ instruction: Instruction) throws {
            switch instruction {
            case .adv(.combo(let value)):   try A = divide(A, power: evaluateCombo(value))
            case .bxl(.literal(let value)): B = B ^ value //
            case .bst(.combo(let value)):   try B = evaluateCombo(value) % 8 //
            case .jnz(.literal(_))
                where A == 0:               break
            case .jnz(.literal(let value))
                where A != 0:               pointer = value/2; return
            case .bxc:                      B = B ^ C //
            case .out(.combo(let value)):   try output.append(evaluateCombo(value) % 8)
            case .bdv(.combo(let value)):   try B = divide(A, power: evaluateCombo(value)) //
            case .cdv(.combo(let value)):   try C = divide(A, power: evaluateCombo(value))
            default:                        throw Day17Error.invalidInstruction}
            
            pointer += 1
        }
        
        private func evaluateCombo(_ value: Int) throws -> Int {
            switch value {
            case 0, 1, 2, 3:                value
            case 4:                         A
            case 5:                         B
            case 6:                         C
            default:                        throw Day17Error.invalidOperand}
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
    }
    
    private enum Operand {
        case literal(Int)
        case combo(Int)
    }
    
    private enum Day17Error: Error {
        case invalidInstruction
        case invalidOperand
    }
}
