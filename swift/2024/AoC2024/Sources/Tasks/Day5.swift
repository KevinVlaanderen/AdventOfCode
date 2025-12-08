import Foundation
internal import SwiftGraph
import Framework

public final class Day5: Day<Void, Int> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        let (orderPairs, updates) = try parse(data)

        return switch task {
        case .task1: try task1(orderPairs: orderPairs, updates: updates)
        case .task2: try task2(orderPairs: orderPairs, updates: updates)
        }
    }
    
    private func parse(_ data: String) throws -> ([(Int, Int)], [[Int]]) {
        let blocks = data.blocks

        let orderPairs = try blocks[0].lines.map { line in
            let pages = line.split(separator: "|", maxSplits: 1)
            return try (pages[0].toInt(), pages[1].toInt())
        }
        
        let updates = try blocks[1].lines.map({ try $0.split(separator: ",").map(toInt) })
        
        return (orderPairs, updates)
    }
    
    private func task1(orderPairs: [(Int, Int)], updates: [[Int]]) throws -> Int {
        try updates.reduce(into: 0) { result, update in
            let graph = buildGraph(update: update, orderPairs: orderPairs)
            guard let sortOrder = graph.topologicalSort() else {
                throw AoCError.taskError("failed to create topological sort")
            }
            
            if try isSorted(update: update, sortOrder: sortOrder) {
                result += update[update.count/2]
            }
        }
    }
    
    private func task2(orderPairs: [(Int, Int)], updates: [[Int]]) throws -> Int {
        try updates.reduce(into: 0) { result, update in
            let graph = buildGraph(update: update, orderPairs: orderPairs)
            guard let sortOrder = graph.topologicalSort() else {
                throw AoCError.taskError("failed to create topological sort")
            }
            
            if try !isSorted(update: update, sortOrder: sortOrder) {
                result += sortOrder[sortOrder.count/2]
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
    
    private func isSorted(update: [Int], sortOrder: [Int]) throws -> Bool {
        try update.adjacentPairs().allSatisfy {
            guard let left = sortOrder.firstIndex(of: $0), let right = sortOrder.firstIndex(of: $1) else {
                throw AoCError.taskError("item not found")
            }
            return left < right
        }
    }
}
