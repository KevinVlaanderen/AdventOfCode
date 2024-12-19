import Foundation
internal import Algorithms
import Framework

public struct Day3: Day {
    public typealias P = [InstructionDefinition.Type]
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let instructions = try parse(instructions: param)
        
        return execute(instructions)
    }

    private func parse(instructions: [InstructionDefinition.Type]) throws -> [Instruction] {
        try instructions.reduce([]) { result, instruction in
            try result + instruction.findMatches(in: data)
        }
        .sorted(by: { $0.1 < $1.1 })
        .map({ $0.0 })
    }
    
    private func execute(_ instructions: [Instruction]) -> Int {
        instructions.reduce(into: State()) { state, instruction in
            switch instruction {
            case let .mul(a, b):
                if state.enabled {
                    state.total += a*b
                }
            case .enable:
                state.enabled = true
            case .disable:
                state.enabled = false
            }
        }.total
    }
    
    public protocol InstructionDefinition: Sendable {
        static func findMatches(in data: String) throws -> [(Instruction, String.Index)]
    }
    
    public struct MulInstruction: InstructionDefinition {
        public static func findMatches(in data: String) throws -> [(Instruction, String.Index)] {
            try data.matches(of: /mul\((\d+),(\d+)\)/).map { match in
                let (left, right) = try (toInt(match.output.1), toInt(match.output.2))
                return (Instruction.mul(left, right), match.range.lowerBound)
            }
        }
    }
    
    public struct DoInstruction: InstructionDefinition {
        public static func findMatches(in data: String) throws -> [(Instruction, String.Index)] {
            data.matches(of: /do\(\)/).map { match in
                (Instruction.enable, match.range.lowerBound)
            }
        }
    }
    
    public struct DontInstruction: InstructionDefinition {
        public static func findMatches(in data: String) throws -> [(Instruction, String.Index)] {
            data.matches(of: /don't\(\)/).map { match in
                (Instruction.disable, match.range.lowerBound)
            }
        }
    }
    
    public enum Instruction {
        case mul(Int, Int)
        case enable
        case disable
    }
    
    private struct State {
        var enabled: Bool = true
        var total: Int = 0
    }
}
