import Foundation
internal import Algorithms
import Framework

public struct Day3: Day {
    public init() {}
    
    nonisolated(unsafe)
    private static let mulPattern = /mul\((\d+),(\d+)\)/
    
    nonisolated(unsafe)
    private static let doPattern = /do\(\)/
    
    nonisolated(unsafe)
    private static let dontPattern = /don't\(\)/
    
    public func perform(task: Task, data: String, param: ()) throws -> Int {
        let instructions = {
            switch task {
            case .task1:
                return parse1(data)
            case .task2:
                return parse2(data)
            }
        }()
        
        return execute(instructions)
    }
    
    private func parse1(_ data: String) -> [Instructions] {
        return data.indices.reduce(into: []) { result, index in
            if let match = data.suffix(from: index).prefixMatch(of: Day3.mulPattern) {
                result.append(Instructions.mul(Int(match.output.1)!, Int(match.output.2)!))
            }
        }
    }
    
    private func parse2(_ data: String) -> [Instructions] {
        return data.indices.reduce(into: []) { result, index in
            if let match = data.suffix(from: index).prefixMatch(of: Day3.mulPattern) {
                result.append(Instructions.mul(Int(match.output.1)!, Int(match.output.2)!))
            } else if data.suffix(from: index).starts(with: Day3.doPattern) {
                result.append(Instructions.enable)
            } else if data.suffix(from: index).starts(with: Day3.dontPattern) {
                result.append(Instructions.disable)
            }
        }
    }
    
    private func execute(_ instructions: [Instructions]) -> Int {
        return instructions.reduce(into: State()) { state, instruction in
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
