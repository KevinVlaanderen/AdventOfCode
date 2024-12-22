import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public final class Day21: Day<Int, Int> {
    public override func perform() throws -> R {
        let codes = try parse(data: data)

        return try codes.reduce(into: 0) { result, code in
            var human: DirectionalKeypad? = nil
            
            var robots: [DirectionalKeypad?] = [
                DirectionalKeypad(startAt: .confirm, controlledBy: &human)
            ]
            
            for i in 0..<param-1 {
                let robot = DirectionalKeypad(startAt: .confirm, controlledBy: &robots[i])
                robots.append(robot)
            }
            
            var robot1 = NumericKeypad(startAt: .confirm, controlledBy: &robots[robots.count-1])

            let keysPressed = try code.reduce(0) {
                try $0 + robot1.press($1)
            }
            
            result += keysPressed * code.value
        }
    }
    
    private func parse(data: String) throws -> [Code] {
        try data.lines.map({ try $0.map(NumericKey.fromInput) })
    }
}

private typealias Code = [NumericKey]

private extension Code {
    var value: Int { Int(self.dropLast().map({ $0.description }).joined())! }
}

private protocol KeypadProtocol<KeyType> {
    associatedtype KeyType: Hashable
    
    var current: KeyType {get set}
    var cache: [HashKey<KeyType>: Int] {get set}
    var controlledBy: (any KeypadProtocol<DirectionalKey>)? {get set}
    
    mutating func press(_ key: KeyType) throws -> Int
    func path(from: KeyType, to: KeyType) throws -> [Direction]
}

private struct HashKey<KeyType: Hashable>: Hashable {
    let from: KeyType
    let to: KeyType
}

private extension KeypadProtocol {
    mutating func press(_ key: KeyType) throws -> Int {
        let hashKey = HashKey<KeyType>(from: current, to: key)
        
        if let hashedKeysPressed = cache[hashKey] {
            current = key
            return hashedKeysPressed
        }
        
        let keysPressed = if key == current {
            try controlledBy?.press(.confirm) ?? 1
        } else {
            try path(from: current, to: key).reduce(0) { result, step in
                try result + (controlledBy?.press(.direction(step)) ?? 1)
            } + (controlledBy?.press(.confirm) ?? 1)
        }
        
        current = key

        cache[hashKey] = keysPressed
        return keysPressed
    }
}

private struct NumericKeypad: KeypadProtocol {
    let graph: WeightedGraph<NumericKey, Direction>
    var current: NumericKey
    var controlledBy: (any KeypadProtocol<DirectionalKey>)?
    
    var cache: [HashKey<NumericKey>: Int] = [:]
    
    func path(from: NumericKey, to: NumericKey) -> [Direction] {
        switch (from, to) {
        case (.confirm, .number(1)): [.up, .left, .left] // special
        case (.confirm, .number(4)): [.up, .up, .left, .left] // special
        case (.confirm, .number(5)): [.left, .up, .up]
        case (.confirm, .number(7)): [.up, .up, .up, .left, .left] // special
        case (.confirm, .number(8)): [.left, .up, .up, .up]
        case (.number(0), .number(4)): [.left, .up, .up]
        case (.number(0), .number(6)): [.up, .up, .right]
        case (.number(0), .number(7)): [.left, .up, .up, .up]
        case (.number(0), .number(9)): [.up, .up, .up, .right]
        case (.number(1), .confirm): [.right, .right, .down] // special
        case (.number(1), .number(6)): [.up, .right, .right]
        case (.number(1), .number(8)): [.up, .up, .right]
        case (.number(1), .number(9)): [.up, .up, .right, .right]
        case (.number(2), .number(7)): [.left, .up, .up]
        case (.number(2), .number(9)): [.up, .up, .right]
        case (.number(3), .number(4)): [.left, .left, .up]
        case (.number(3), .number(7)): [.left, .left, .up, .up]
        case (.number(3), .number(8)): [.left, .up, .up]
        case (.number(4), .confirm): [.right, .right, .down, .down] // special
        case (.number(4), .number(0)): [.right, .down, .down] // special
        case (.number(4), .number(3)): [.down, .right, .right]
        case (.number(4), .number(9)): [.up, .right, .right]
        case (.number(5), .confirm): [.down, .down, .right]
        case (.number(6), .number(0)): [.left, .down, .down]
        case (.number(6), .number(1)): [.left, .left, .down]
        case (.number(6), .number(7)): [.left, .left, .up]
        case (.number(7), .confirm): [.right, .right, .down, .down, .down] // special
        case (.number(7), .number(0)): [.right, .down, .down, .down] // special
        case (.number(7), .number(2)): [.down, .down, .right]
        case (.number(7), .number(3)): [.down, .down, .right, .right]
        case (.number(7), .number(6)): [.down, .right, .right]
        case (.number(8), .confirm): [.down, .down, .down, .right]
        case (.number(8), .number(1)): [.left, .down, .down]
        case (.number(8), .number(3)): [.down, .down, .right]
        case (.number(9), .number(0)): [.left, .down, .down, .down]
        case (.number(9), .number(1)): [.left, .left, .down, .down]
        case (.number(9), .number(2)): [.left, .down, .down]
        case (.number(9), .number(4)): [.left, .left, .down]
        default: self.graph.bfs(from: from, to: to).map({ $0.weight })
        }
    }
    
    init(startAt: NumericKey, controlledBy: inout DirectionalKeypad?) {
        self.current = startAt
        self.controlledBy = controlledBy
        self.graph = WeightedGraph(vertices: (0...9).map({ .number($0) }) + [.confirm])
        
        graph.addEdge(from: .number(0), to: .confirm, weight: .right, directed: true)
        graph.addEdge(from: .number(0), to: .number(2), weight: .up, directed: true)
        graph.addEdge(from: .number(1), to: .number(2), weight: .right, directed: true)
        graph.addEdge(from: .number(1), to: .number(4), weight: .up, directed: true)
        graph.addEdge(from: .number(2), to: .number(0), weight: .down, directed: true)
        graph.addEdge(from: .number(2), to: .number(1), weight: .left, directed: true)
        graph.addEdge(from: .number(2), to: .number(3), weight: .right, directed: true)
        graph.addEdge(from: .number(2), to: .number(5), weight: .up, directed: true)
        graph.addEdge(from: .number(3), to: .number(2), weight: .left, directed: true)
        graph.addEdge(from: .number(3), to: .number(6), weight: .up, directed: true)
        graph.addEdge(from: .number(3), to: .confirm, weight: .down, directed: true)
        graph.addEdge(from: .number(4), to: .number(1), weight: .down, directed: true)
        graph.addEdge(from: .number(4), to: .number(5), weight: .right, directed: true)
        graph.addEdge(from: .number(4), to: .number(7), weight: .up, directed: true)
        graph.addEdge(from: .number(5), to: .number(2), weight: .down, directed: true)
        graph.addEdge(from: .number(5), to: .number(4), weight: .left, directed: true)
        graph.addEdge(from: .number(5), to: .number(6), weight: .right, directed: true)
        graph.addEdge(from: .number(5), to: .number(8), weight: .up, directed: true)
        graph.addEdge(from: .number(6), to: .number(3), weight: .down, directed: true)
        graph.addEdge(from: .number(6), to: .number(5), weight: .left, directed: true)
        graph.addEdge(from: .number(6), to: .number(9), weight: .up, directed: true)
        graph.addEdge(from: .number(7), to: .number(4), weight: .down, directed: true)
        graph.addEdge(from: .number(7), to: .number(8), weight: .right, directed: true)
        graph.addEdge(from: .number(8), to: .number(5), weight: .down, directed: true)
        graph.addEdge(from: .number(8), to: .number(7), weight: .left, directed: true)
        graph.addEdge(from: .number(8), to: .number(9), weight: .right, directed: true)
        graph.addEdge(from: .number(9), to: .number(6), weight: .down, directed: true)
        graph.addEdge(from: .number(9), to: .number(8), weight: .left, directed: true)
        graph.addEdge(from: .confirm, to: .number(0), weight: .left, directed: true)
        graph.addEdge(from: .confirm, to: .number(3), weight: .up, directed: true)
    }
}

private struct DirectionalKeypad: KeypadProtocol {
    let graph: WeightedGraph<DirectionalKey, Direction>
    var current: DirectionalKey
    var controlledBy: (any KeypadProtocol<DirectionalKey>)?
    
    var cache: [HashKey<DirectionalKey>: Int] = [:]
    
    func path(from: DirectionalKey, to: DirectionalKey) throws -> [Direction] {
        switch (from, to) {
        case (.confirm, .direction(.left)): [.down, .left, .left]
        case (.confirm, .direction(.right)): [.down]
        case (.confirm, .direction(.up)): [.left]
        case (.confirm, .direction(.down)): [.left, .down]
        case (.direction(.left), .confirm): [.right, .right, .up]
        case (.direction(.left), .direction(.right)): [.right, .right]
        case (.direction(.left), .direction(.up)): [.right, .up]
        case (.direction(.left), .direction(.down)): [.right]
        case (.direction(.right), .confirm): [.up]
        case (.direction(.right), .direction(.left)): [.left, .left]
        case (.direction(.right), .direction(.up)): [.left, .up]
        case (.direction(.right), .direction(.down)): [.left]
        case (.direction(.up), .confirm): [.right]
        case (.direction(.up), .direction(.left)): [.down, .left]
        case (.direction(.up), .direction(.right)): [.down, .right]
        case (.direction(.up), .direction(.down)): [.down]
        case (.direction(.down), .confirm): [.up, .right]
        case (.direction(.down), .direction(.left)): [.left]
        case (.direction(.down), .direction(.right)): [.right]
        case (.direction(.down), .direction(.up)): [.up]
        default: throw AoCError.parseError("movement to itself")
        }
    }
    
    init(startAt: DirectionalKey, controlledBy: inout DirectionalKeypad?)  {
        self.current = startAt
        self.controlledBy = controlledBy
        self.graph = WeightedGraph(vertices: Direction.allCases.map({ .direction($0) }) + [.confirm])
        
        graph.addEdge(from: .direction(.left), to: .direction(.down), weight: .right, directed: true)
        graph.addEdge(from: .direction(.right), to: .direction(.down), weight: .left, directed: true)
        graph.addEdge(from: .direction(.right), to: .confirm, weight: .up, directed: true)
        graph.addEdge(from: .direction(.up), to: .direction(.down), weight: .down, directed: true)
        graph.addEdge(from: .direction(.up), to: .confirm, weight: .right, directed: true)
        graph.addEdge(from: .direction(.down), to: .direction(.left), weight: .left, directed: true)
        graph.addEdge(from: .direction(.down), to: .direction(.right), weight: .right, directed: true)
        graph.addEdge(from: .direction(.down), to: .direction(.up), weight: .up, directed: true)
        graph.addEdge(from: .confirm, to: .direction(.up), weight: .left, directed: true)
        graph.addEdge(from: .confirm, to: .direction(.right), weight: .down, directed: true)
    }
}

private enum NumericKey: Codable, Hashable, CustomStringConvertible {
    case number(Int)
    case confirm
    
    static func fromInput(_ value: Character) throws -> Self {
        switch value {
        case _ where value.wholeNumberValue != nil: .number(value.wholeNumberValue!)
        case "A":                                   .confirm
        default:                                    throw AoCError.parseError("invalid character \(value)")
        }
    }
    
    var description: String {
        switch self {
        case let .number(number):                   String(number)
        case .confirm:                              "A"
        }
    }
}

private enum DirectionalKey: Codable, Hashable {
    case direction(Direction)
    case confirm
    
    static func fromInput(_ value: Character) throws -> Self {
        switch value {
        case "<":                                   .direction(.left)
        case ">":                                   .direction(.right)
        case "^":                                   .direction(.up)
        case "v":                                   .direction(.down)
        case "A":                                   .confirm
        default:                                    throw AoCError.parseError("invalid character \(value)")
        }
    }
}
