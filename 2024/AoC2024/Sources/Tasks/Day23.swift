import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public final class Day23: Day<Task, String> {
    public override func perform(data: String, param: P) throws -> R {
        let network = try parse(data)

        return switch param {
        case .task1: "\(task1(network: network))"
        case .task2: task2(network: network)
        }
    }
    
    private func task1(network: UnweightedGraph<String>) -> Int {
        findTriangles(in: network, with: network.edgeList()
            .filter { edge in
                let u = network[edge.u]
                let v = network[edge.v]
                return u.hasPrefix("t") || v.hasPrefix("t")
            }).count
    }
    
    private func task2(network: UnweightedGraph<String>) -> String {
        findTriangles(in: network, with: network.edgeList())
            .reduce(into: [] as [Set<String>]) { groups, triangle in
                for i in 0..<groups.count where groups[i].intersection(triangle).count >= 1 {
                    groups[i].formUnion(triangle)
                    return
                }
                groups.append(triangle)
            }.filter { group in
                group.combinations(ofCount: 2).allSatisfy({ network.edgeExists(from: $0[0], to: $0[1]) })
            }.max(by: { $0.count < $1.count })!
                .sorted()
                .joined(separator: ",")
    }
    
    private func findTriangles(in network: UnweightedGraph<String>, with edges: [UnweightedEdge]) -> Set<Set<String>> {
        edges.reduce(into: []) { triangles, edge in
            for wIndex in network.indices {
                if network.edgeExists(fromIndex: edge.u, toIndex: wIndex) &&
                    network.edgeExists(fromIndex: edge.v, toIndex: wIndex) {
                    let u = network[edge.u]
                    let v = network[edge.v]
                    let w = network[wIndex]
                    triangles.insert(Set(arrayLiteral: u, v, w))
                }
            }
        }
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
