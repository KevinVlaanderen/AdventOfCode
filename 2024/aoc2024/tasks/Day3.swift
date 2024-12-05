import Foundation
internal import Algorithms

struct Day3: Day {
    typealias R = Int
    
    private static let mulPattern = /mul\((\d+),(\d+)\)/
    private static let doPattern = /do\(\)/
    private static let dontPattern = /don't\(\)/
    
    func task1(data: String, param: P) throws -> R {
        return execute(parse1(data))
    }
    
    func task2(data: String, param: P) throws -> R {
        return execute(parse2(data))
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
    
    private enum Instructions {
        case mul(Int, Int)
        case enable
        case disable
    }
    
    private struct State {
        var enabled: Bool = true
        var total: Int = 0
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
}
