import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public final class Day23: Day<Task, Int> {
    public override func perform(data: String, param: P) throws -> R {
        let network = try parse(data)

        var cycles: Set<Set<String>> = []
        
        for edge in network.edgeList() {
            let u = network[edge.u]
            let v = network[edge.v]
            
            for vertex in network.vertices {
                if vertex.hasPrefix("t") &&
                    network.edgeExists(from: u, to: vertex) &&
                    network.edgeExists(from: v, to: vertex) {
                    cycles.insert(Set(arrayLiteral: u, v, vertex))
                }
            }
        }
                        
        return cycles.count(where: { _ in true })
    }
    
    private func parse(_ data: String) throws -> UnweightedGraph<String> {
        let pairs = data.lines.map({ $0.split(separator: "-").map({ String($0) }) })
        let computers = pairs.flatMap({ $0 })
        
        let network = UnweightedGraph(vertices: computers)
        
        for pair in pairs {
            network.addEdge(from: pair[0], to: pair[1])
        }
        
        return network
    }
}
