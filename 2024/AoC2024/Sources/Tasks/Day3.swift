import Foundation
internal import Algorithms
import Framework

public final class Day3: Day<[Day3.InstructionParam], Int> {
    public enum InstructionParam: CaseIterable {
        case mulInst, doInst, dontInst
    }
    
    public override func perform(data: String, param: P) throws -> R {
        let instructions = try parse(data, param: param)
        
        return execute(instructions)
    }

    private func parse(_ data: String, param: P) throws -> [Instruction] {
        let instructions: [InstructionDefinition.Type] = param.map {
            switch $0 {
            case .mulInst: MulInstruction.self
            case .doInst: DoInstruction.self
            case .dontInst: DontInstruction.self
            }
        }
        
        return try instructions.reduce([]) { result, instruction in
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
}

private enum Instruction {
    case mul(Int, Int)
    case enable
    case disable
}

private protocol InstructionDefinition {
    static func findMatches(in data: String) throws -> [(Instruction, String.Index)]
}

private struct MulInstruction: InstructionDefinition {
    public static func findMatches(in data: String) throws -> [(Instruction, String.Index)] {
        try data.matches(of: /mul\((\d+),(\d+)\)/).map { match in
            let (left, right) = try (match.output.1.toInt(), match.output.2.toInt())
            return (Instruction.mul(left, right), match.range.lowerBound)
        }
    }
}

private struct DoInstruction: InstructionDefinition {
    public static func findMatches(in data: String) throws -> [(Instruction, String.Index)] {
        data.matches(of: /do\(\)/).map { match in
            (Instruction.enable, match.range.lowerBound)
        }
    }
}

private struct DontInstruction: InstructionDefinition {
    public static func findMatches(in data: String) throws -> [(Instruction, String.Index)] {
        data.matches(of: /don't\(\)/).map { match in
            (Instruction.disable, match.range.lowerBound)
        }
    }
}

private struct State {
    var enabled: Bool = true
    var total: Int = 0
}
