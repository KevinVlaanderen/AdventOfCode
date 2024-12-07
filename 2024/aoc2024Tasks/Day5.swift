import Foundation
internal import Algorithms
internal import SwiftGraph
import aoc2024Framework

public struct Day5: Day {
    public typealias R = Int
    
    public init() {}
    
    public func perform(task: Task, data: String, param: P) async throws -> Int {
        let (orderPairs, updates) = parse(data)

        switch task {
        case .task1:
            return task1(orderPairs: orderPairs, updates: updates)
        case .task2:
            return task2(orderPairs: orderPairs, updates: updates)
        @unknown default:
            fatalError("unknown task")
        }
    }
    
    private func parse(_ data: String) -> ([(Int, Int)], [[Int]]) {
        let blocks = data.split(separator: "\n\n")

        let orderPairs = blocks[0].split(whereSeparator: \.isNewline).map { line in
            let pages = line.split(separator: "|")
            return (Int(pages[0])!, Int(pages[1])!)
        }
        
        let updates = blocks[1].split(whereSeparator: \.isNewline).map { line in
            return line.split(separator: ",").map({ Int($0)! })
        }
        
        return (orderPairs, updates)
    }
    
    private func task1(orderPairs: [(Int, Int)], updates: [[Int]]) -> R {
        updates.reduce(0) { result, update in
            let graph = buildGraph(update: update, orderPairs: orderPairs)
            let sortOrder = graph.topologicalSort()!
            
            if isSorted(update: update, sortOrder: sortOrder) {
                return result + update[update.count/2]
            } else {
                return result
            }
        }
    }
    
    private func task2(orderPairs: [(Int, Int)], updates: [[Int]]) -> R {
        updates.reduce(0) { result, update in
            let graph = buildGraph(update: update, orderPairs: orderPairs)
            let sortOrder = graph.topologicalSort()!
            
            if !isSorted(update: update, sortOrder: sortOrder) {
                return result + sortOrder[sortOrder.count/2]
            } else {
                return result
            }
        }
    }
    
    private func buildGraph(update: [Int], orderPairs: [(Int, Int)]) -> UnweightedUniqueElementsGraph<Int> {
        let graph = SwiftGraph.UnweightedUniqueElementsGraph<Int>(vertices: update)
        
        for orderPair in orderPairs {
            graph.addEdge(from: orderPair.0, to: orderPair.1, directed: true)
        }
        
        return graph
    }
    
    private func isSorted(update: [Int], sortOrder: [Int]) -> Bool {
        return update.adjacentPairs().allSatisfy({ sortOrder.firstIndex(of: $0)! < sortOrder.firstIndex(of: $1)! })
    }
}
