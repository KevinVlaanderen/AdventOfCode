import Foundation
import Framework
internal import HeapModule
internal import BitCollections

public final class Day24: Day<Void, Int> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        var machine = try parse(data)
        
        machine.run()
        
        let outputNodes: [(Int, Bool)] = machine.addresses.compactMap({
            if case let .output(index) = $0.key {
                return (index, machine.nodes[$0.value]!)
            }
            return nil
        }).sorted(by: { $0.0 < $1.0 })

        return Int(UInt(BitArray(outputNodes.map({ $0.1 }))))
    }
    
    private func parse(_ data: String) throws -> Machine {
        let nodePattern = /(.+): (1|0)/
        let gatePattern = /(.+) (.*) (.*) -> (.*)/

        var machine = Machine()
        
        let blocks = data.blocks
        for match in blocks[0].matches(of: nodePattern) {
            let node = parseNode(String(match.output.1))
            _ = machine.addNode(node: node, value: match.output.2 == "1")
        }
        
        for match in blocks[1].matches(of: gatePattern) {
            let left = machine.addNode(node: parseNode(String(match.output.1)), value: nil)
            let right = machine.addNode(node: parseNode(String(match.output.3)), value: nil)
            let output = machine.addNode(node: parseNode(String(match.output.4)), value: nil)
            
            let operation: Operation = switch match.output.2 {
            case "AND": .and
            case "OR": .or
            case "XOR": .xor
            default: throw AoCError.parseError("invalid instruction \(match.output.2)")
            }
            
            machine.addGate(left: left, right: right, output: output, operation: operation)
        }
        
        return machine
    }
    
    private func parseNode(_ name: String) -> Node {
        if name.hasPrefix("z") {
            let index = Int(name.dropFirst())!
            return .output(index: index)
        } else {
            return .regular(name: name)
        }
    }
}

private struct Machine {
    var addresses: [Node: Int] = [:]
    var nodes: [Bool?] = []
    var gates = Heap<StackEntry>()
    private var stackIndex = 0
    
    mutating func addNode(node: Node, value: Bool?) -> Int {
        if let index = addresses[node] {
            return index
        } else {
            nodes.append(value)
            let address = nodes.count-1
            addresses[node] = address
            return address
        }
    }
    
    mutating func addGate(left: Int, right: Int, output: Int, operation: Operation) {
        gates.insert(StackEntry(offset: stackIndex, gate: Gate(left: left, right: right, output: output, operation: operation)))
        stackIndex += 1
    }
    
    mutating func run() {
        let outputs = addresses.filter({ $0.key.isOutput() }).map({ $0.value })
        
        while !outputs.allSatisfy({ nodes[$0] != nil }) {
            var currentGate: Gate?
            
            while currentGate == nil {
                guard let current = gates.popMin() else {
                    break
                }
                if nodes[current.gate.left] == nil || nodes[current.gate.right] == nil {
                    gates.insert(StackEntry(offset: stackIndex, gate: current.gate))
                    stackIndex += 1
                    continue
                }
                
                currentGate = current.gate
            }
            
            guard let currentGate = currentGate else {
                break
            }

            switch currentGate.operation {
            case .and: nodes[currentGate.output] = nodes[currentGate.left]! && nodes[currentGate.right]!
            case .or: nodes[currentGate.output] = nodes[currentGate.left]! || nodes[currentGate.right]!
            case .xor: nodes[currentGate.output] = nodes[currentGate.left]! != nodes[currentGate.right]!
            }
        }
    }
}

private struct StackEntry: Comparable {
    static func < (lhs: StackEntry, rhs: StackEntry) -> Bool {
        lhs.offset < rhs.offset
    }
    
    static func == (lhs: StackEntry, rhs: StackEntry) -> Bool {
        lhs.offset == rhs.offset
    }
    
    let offset: Int
    let gate: Gate
}

private enum Node: Hashable {
    case regular(name: String)
    case output(index: Int)
    
    func isOutput() -> Bool {
        if case .output(_) = self {
            return true
        }
        return false
    }
}

private struct Gate {
    let left: Int
    let right: Int
    let output: Int
    let operation: Operation
}

private enum Operation {
    case and, or, xor
}
