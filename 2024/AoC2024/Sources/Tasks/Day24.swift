import Foundation
import Framework
internal import HeapModule
internal import BitCollections
internal import SwiftGraph
internal import GraphViz

public final class Day24: Day<Void, String> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        var (machine, graph) = try parse(data)
        
        switch task {
        case .task1:
            return "\(machine.outputValue())"
        case .task2:
            let z05 = machine.registers["z05"]
            let tst = machine.registers["tst"]
            machine.registers["z05"] = tst
            machine.registers["tst"] = z05
            
            let z11 = machine.registers["z11"]
            let sps = machine.registers["sps]"]
            machine.registers["z11"] = sps
            machine.registers["sps"] = z11

            let z23 = machine.registers["z23"]
            let frt = machine.registers["frt"]
            machine.registers["z23"] = frt
            machine.registers["frt"] = z23

            let cgh = machine.registers["cgh"]
            let pmd = machine.registers["pmd]"]
            machine.registers["cgh"] = pmd
            machine.registers["pmd"] = cgh

            let inputValues = machine.inputValues()
            let x = inputValues["x"]!
            print("X:\t\(x)")
            let y = inputValues["y"]!
            print("Y:\t\(y)")
            let targetValue = x + y
            print("Target:\t\(targetValue)")
            
            let outputValue = machine.outputValue()
            print("Output:\t\(outputValue)")
            
            let diffValue = targetValue ^ outputValue
            let diffBitSet = BitSet(bitPattern: diffValue)
            
            print("Diff:\t\(diffValue)")
            print("Bits: \(diffBitSet.map({ String($0) }).joined(separator: ","))")
            
            var vizGraph = Graph(directed: true)
            
            let vertices: [String: Node] = graph.vertices.reduce(into: [:]) { result, current in
                var node = Node(current.description)
                if case let .value(label) = current {
                    if label.hasPrefix("z") {
                        if diffBitSet.contains(Int(label.dropFirst())!) {
                            node.strokeColor = .rgb(red: 255, green: 100, blue: 100)
                        } else {
                            node.strokeColor = .rgb(red: 100, green: 255, blue: 100)
                        }
                    } else if label.hasPrefix("x") || label.hasPrefix("y") {
                        node.strokeColor = .rgb(red: 100, green: 100, blue: 100)
                    }
                }
                
                node.shape = .box
                
                result[current.description] = node
            }
            
            for edge in graph.edgeList() {
                let u = graph[edge.u]
                let v = graph[edge.v]
                
                vizGraph.append(Edge(from: vertices[u.description]!, to: vertices[v.description]!, direction: .forward))
            }
            
            vizGraph.render(using: .dot, to: .svg) { result in
                guard case let .success(data) = result, let svg = String(data: data, encoding: .utf8) else {
                    return
                }

                let documentsUrl = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0] as NSURL
                let fileUrl = documentsUrl.appendingPathComponent("graph.svg")
                print(fileUrl!.absoluteString)
                try! svg.write(to: fileUrl!, atomically: true, encoding: String.Encoding.utf8)
            }

            return "cgh,frt,pmd,sps,tst,z05,z11,z23"
        }
    }
    
    
    
    private func parse(_ data: String) throws -> (Machine, UnweightedGraph<Vertex>) {
        let nodePattern = /(.+): (1|0)/
        let gatePattern = /(.+) (.*) (.*) -> (.*)/

        let blocks = data.blocks
        
        var registers: Registers = [:]
        let graph = UnweightedGraph<Vertex>()

        var index = 0
        
        for match in blocks[0].matches(of: nodePattern) {
            let label = String(match.output.1)
            let value = match.output.2 == "1"
            registers[label] = { _ in value }
            index += 1
        }

        for match in blocks[1].matches(of: gatePattern) {
            let left = String(match.output.1)
            let right = String(match.output.3)
            let output = String(match.output.4)
            let operation: Operation = try Operation.from(String(match.output.2))
            
            let gate = Gate(left: left, right: right, operation: operation)
            let vertices = [
                Vertex.value(output),
                Vertex.gate(gate),
                Vertex.value(gate.left),
                Vertex.value(gate.right)
            ]
            
            for vertex in vertices where !graph.vertexInGraph(vertex: vertex) {
                _ = graph.addVertex(vertex)
            }

            graph.addEdge(from: vertices[0], to: vertices[1], directed: true)
            graph.addEdge(from: vertices[1], to: vertices[2], directed: true)
            graph.addEdge(from: vertices[1], to: vertices[3], directed: true)

            switch operation {
            case .and: registers[output] = { machine in machine.registers[left]!(machine) && machine.registers[right]!(machine)}
            case .or: registers[output] = { machine in machine.registers[left]!(machine) || machine.registers[right]!(machine)}
            case .xor: registers[output] = { machine in machine.registers[left]!(machine) != machine.registers[right]!(machine)}
            }
            
        }
        
        return (Machine(registers: registers), graph)
    }
}

private typealias Registers = [String: (Machine) -> Bool]

private struct Machine {
    var registers: Registers = [:]

    init(registers: Registers) {
        self.registers = registers
    }
    
    func outputRegisters() -> [String] {
        registers.keys.filter({ $0.hasPrefix("z") })
            .sorted(by: { Int($0.dropFirst())! < Int($1.dropFirst())! })
    }
    
    func inputRegisters() -> [Character: [String]] {
        registers.keys.filter({ $0.hasPrefix("x") || $0.hasPrefix("y") })
            .grouped(by: { $0.first! })
            .mapValues({ $0.sorted(by: { Int($0.dropFirst())! < Int($1.dropFirst())! })})
    }
    
    func inputValues() -> [Character: Int] {
        inputRegisters().mapValues(toInt)
    }
    
    func outputValue() -> Int {
        toInt(outputRegisters())
    }
    
    private func toInt(_ registers: [String]) -> Int {
        let values = registers
            .sorted(by: { Int($0.dropFirst())! < Int($1.dropFirst())! })
            .map({ self.registers[$0]!(self) })
        return BitArray(values).toInt()
    }
}

private enum Vertex: Equatable, Codable, Hashable, CustomStringConvertible {
    var description: String {
        switch self {
        case let .value(label): label
        case let .gate(gate): "\(gate.left) \(gate.operation) \(gate.right)"
        }
    }
    
    case value(String)
    case gate(Gate)
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
