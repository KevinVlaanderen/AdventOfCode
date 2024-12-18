import Foundation
internal import Algorithms
import Framework

public struct Day3: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let instructions = switch param {
        case .task1:
            try parse1(data)
        case .task2:
            try parse2(data)
        }
        
        return execute(instructions)
    }
    
    nonisolated(unsafe)
    private static let mulPattern = /mul\((\d+),(\d+)\)/
    
    nonisolated(unsafe)
    private static let doPattern = /do\(\)/
    
    nonisolated(unsafe)
    private static let dontPattern = /don't\(\)/
    
    private func parse1(_ data: String) throws -> [Instructions] {
        try data.indices.reduce(into: []) { result, index in
            if let match = data.suffix(from: index).prefixMatch(of: Day3.mulPattern) {
                guard let left = Int(match.output.1), let right = Int(match.output.2) else {
                    throw AoCError.parseError("failed to parse string to integer")
                }
                result.append(Instructions.mul(left, right))
            }
        }
    }
    
    private func parse2(_ data: String) throws -> [Instructions] {
        try data.indices.reduce(into: []) { result, index in
            if let match = data.suffix(from: index).prefixMatch(of: Day3.mulPattern) {
                guard let left = Int(match.output.1), let right = Int(match.output.2) else {
                    throw AoCError.parseError("failed to parse string to integer")
                }
                result.append(Instructions.mul(left, right))
            } else if data.suffix(from: index).starts(with: Day3.doPattern) {
                result.append(Instructions.enable)
            } else if data.suffix(from: index).starts(with: Day3.dontPattern) {
                result.append(Instructions.disable)
            }
        }
    }
    
    private func execute(_ instructions: [Instructions]) -> Int {
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
    
    private enum Instructions {
        case mul(Int, Int)
        case enable
        case disable
    }
    
    private struct State {
        var enabled: Bool = true
        var total: Int = 0
    }
}
