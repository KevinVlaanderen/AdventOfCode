import Foundation
import Framework
internal import HeapModule
internal import BitCollections
internal import SwiftGraph

public final class Day24: Day<Void, String> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        var (machine, input) = try parse(data)
        
        switch task {
        case .task1:
            return "\(machine.run(inputs: input))"
        case .task2:
            let targetValue = input.values.dropFirst().reduce(input.values.first!, +)
            print("Target:\t\(targetValue)")
            
            let outputValue = machine.run(inputs: input)
            print("Output:\t\(outputValue)")
            
            let diffValue = targetValue ^ outputValue
            let diffBits = BitSet(bitPattern: diffValue)
            
            print("Diff:\t\(diffValue)")
            print("Bits: \(diffBits.map({ String($0) }).joined(separator: ","))")
            
            let graph = machine.graph
            
            let reachablePerBit: [Int: Set<Gate>] = diffBits.reduce(into: [:]) { result, bit in
                let label = "z\(String(format: "%02d", bit))"
                
                result[bit] = Set(graph.findAllDfs(from: Vertex.value(label), goalTest: { vertex in graph.edgesForVertex(vertex)?.count ?? 0 > 0 })
                    .flatMap({ $0.map({ graph[$0.v] }) })
                    .reduce(into: []) { result, vertex in
                        if case let .gate(gate) = vertex {
                            result.append(gate)
                        }
                    })
            }
            
            let unique = Array(reachablePerBit.flatMap({ $0.value }).uniqued())
            let combinations = unique.combinations(ofCount: 8)
            
//            let permutations = unique.permutations(ofCount: 8).count
            
            return ""
        }
    }
    
    private func toBits(_ registers: [String: Bool]) -> BitArray {
        let values = registers.sorted(by: { Int($0.key.dropFirst())! < Int($1.key.dropFirst())! }).map({ $0.value })
        return BitArray(values)
    }
    
    private func parse(_ data: String) throws -> (Machine, [Character: Int]) {
        let nodePattern = /(.+): (1|0)/
        let gatePattern = /(.+) (.*) (.*) -> (.*)/

        let blocks = data.blocks
        
        let inputs: [Character: Int] = blocks[0].matches(of: nodePattern).reduce(into: [:]) { result, match in
            let label = String(match.output.1)
            let prefix = label.first!
            let number = Int(label.dropFirst())!
            let value = Int(match.output.2)!
            result[prefix] = (result[prefix] ?? 0) + (value << number)
        }
        
        let sources: [String: Gate] = try blocks[1].matches(of: gatePattern).reduce(into: [:]) { result, match in
            let left = String(match.output.1)
            let right = String(match.output.3)
            let output = String(match.output.4)
            let operation: Operation = try Operation.from(String(match.output.2))
            
            result[output] = Gate(left: left, right: right, operation: operation)
        }
        
        return try (Machine(sources: sources), inputs)
    }
}

private struct Machine {
    var sources: [String: Gate] = [:]

    init(sources: [String: Gate]) throws {
        self.sources = sources
    }

    mutating func run(inputs: [Character: Int]) -> Int {
        var registers: [String: Bool] = [:]
        
        var heap = Heap<StackEntry>(sources.keys.enumerated().map({ StackEntry(priority: $0.offset, register: $0.element) }))
        var heapIndex = heap.count
        
        for input in inputs {
            let bitArray = BitArray(bitPattern: input.value)
            for value in bitArray.enumerated() {
                let label = "\(input.key)\(String(format: "%02d", value.offset))"
                registers[label] = value.element
            }
            
        }
        
        while true {
            guard let current = heap.popMin() else {
                break
            }

            let source = sources[current.register]!
            guard let value = calculateValue(for: source, registers: registers) else {
                heap.insert(StackEntry(priority: heapIndex, register: current.register))
                heapIndex += 1
                continue
            }

            registers[current.register] = value
        }

        let values = registers.filter({ $0.key.hasPrefix("z") }).sorted(by: { Int($0.key.dropFirst())! < Int($1.key.dropFirst())! }).map({ $0.value })
        return BitArray(values).toInt()
    }

    func calculateValue(for gate: Gate, registers: [String: Bool]) -> Bool? {
        guard let leftValue = registers[gate.left] ?? nil, let rightValue = registers[gate.right] ?? nil else {
            return nil
        }

        return switch gate.operation {
        case .and: leftValue && rightValue
        case .or: leftValue || rightValue
        case .xor: leftValue != rightValue
        }
    }
    
    var graph: UnweightedGraph<Vertex> {
        let graph = UnweightedGraph(vertices: sources.keys.map({ Vertex.value($0) }) + sources.values.map({ Vertex.gate($0) }))
        
        for source in sources {
            graph.addEdge(from: Vertex.value(source.key), to: Vertex.gate(source.value), directed: true)
            graph.addEdge(from: Vertex.gate(source.value), to: Vertex.value(source.value.left), directed: true)
            graph.addEdge(from: Vertex.gate(source.value), to: Vertex.value(source.value.right), directed: true)
        }
        
        return graph
    }
}

private enum Vertex: Equatable, Codable, Hashable {
    case value(String)
    case gate(Gate)
}

private struct StackEntry: Comparable {
    static func < (lhs: StackEntry, rhs: StackEntry) -> Bool {
        lhs.priority < rhs.priority
    }
    
    let priority: Int
    let register: String
}

private struct Gate: Equatable, Codable, Hashable {
    let left: String
    let right: String
    let operation: Operation
}
    
private enum Operation: Codable {
    case and, or, xor
    
    static func from(_ value: String) throws -> Self {
        switch value {
        case "AND": .and
        case "OR": .or
        case "XOR": .xor
        default: throw AoCError.parseError("invalid instruction \(value)")
        }
    }
}

private extension BitArray {
    func toInt() -> Int {
        return Int(UInt(self))
    }
}
