import Foundation
internal import Algorithms
import Framework

public struct Day17: Day {
    public typealias R = String
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        var (machine, instructions) = try parse(data: data)

        try machine.runInstructions(instructions)
        
        return machine.output.map({ String($0) }).joined(separator: ",")
    }
    
    nonisolated(unsafe)
    private static let dataPattern = /Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: (.*)/
    
    private func parse(data: String) throws -> (Machine, [Instruction]) {
        guard let match = data.wholeMatch(of: Day17.dataPattern) else {
            throw AoCError.parseError("data did not match pattern")
        }
        let A = Int(match.output.1)!
        let B = Int(match.output.2)!
        let C = Int(match.output.3)!
        let instructions: [Instruction] = try match.output.4.split(separator: ",").map({ Int($0)! }).windows(ofCount: 2).striding(by: 2).map {
            switch $0[$0.startIndex] {
            case 0: .adv(.combo($0[$0.endIndex-1]))
            case 1: .bxl(.literal($0[$0.endIndex-1]))
            case 2: .bst(.combo($0[$0.endIndex-1]))
            case 3: .jnz(.literal($0[$0.endIndex-1]))
            case 4: .bxc
            case 5: .out(.combo($0[$0.endIndex-1]))
            case 6: .bdv(.combo($0[$0.endIndex-1]))
            case 7: .cdv(.combo($0[$0.endIndex-1]))
            default: throw AoCError.parseError("unknown instruction \($0[0])")
            }
        }
        
        return (Machine(A: A, B: B, C: C), instructions)
    }
    
    private struct Machine {
        var A: Int
        var B: Int
        var C: Int
        var output: [Int] = []
        
        mutating func runInstructions(_ instructions: [Instruction]) throws {
            var pointer = 0
            
            while pointer < instructions.count {
                let instruction = instructions[pointer]
                
                switch instruction {
                case .adv(let operand):
                    let opValue = try evaluateOperand(operand)
                    A = A/Int(pow(2, Double(opValue)))
                    pointer += 1
                case .bxl(let operand):
                    let opValue = try evaluateOperand(operand)
                    B = B ^ opValue
                    pointer += 1
                case .bst(let operand):
                    let opValue = try evaluateOperand(operand)
                    B = opValue % 8
                    pointer += 1
                case .jnz(let operand):
                    let opValue = try evaluateOperand(operand)
                    if A != 0 {
                        pointer = opValue/2
                    } else {
                        pointer += 1
                    }
                case .bxc:
                    B = B ^ C
                    pointer += 1
                case .out(let operand):
                    let opValue = try evaluateOperand(operand)
                    output.append(opValue % 8)
                    pointer += 1
                case .bdv(let operand):
                    let opValue = try evaluateOperand(operand)
                    B = A/Int(pow(2, Double(opValue)))
                    pointer += 1
                case .cdv(let operand):
                    let opValue = try evaluateOperand(operand)
                    C = A/Int(pow(2, Double(opValue)))
                    pointer += 1
                }
            }
        }
        
        private func evaluateOperand(_ operand: Operand) throws -> Int {
            return switch operand {
            case let .literal(value): value
            case .combo(0): 0
            case .combo(1): 1
            case .combo(2): 2
            case .combo(3): 3
            case .combo(4): A
            case .combo(5): B
            case .combo(6): C
            default: throw Day17Error.invalidOperand
            }
        }
    }
    
    private enum Instruction {
        case adv(Operand) // A / 2^operand                   A
        case bxl(Operand) // B XOR operand                   B
        case bst(Operand) // operand % 8                     B
        case jnz(Operand) // jump to operand if A not zero
        case bxc // B XOR C                         B
        case out(Operand) // operand % 8                     out
        case bdv(Operand) // A / 2^operand                   B
        case cdv(Operand) // A / 2^operand                   C
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
