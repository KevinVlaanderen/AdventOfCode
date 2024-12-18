import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day5: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let (orderPairs, updates) = try parse(data)

        return switch param {
        case .task1:
            try task1(orderPairs: orderPairs, updates: updates)
        case .task2:
            try task2(orderPairs: orderPairs, updates: updates)
        }
    }
    
    private func parse(_ data: String) throws -> ([(Int, Int)], [[Int]]) {
        let blocks = data.split(separator: "\n\n")

        let orderPairs = try blocks[0].split(whereSeparator: \.isNewline).map { line in
            let pages = line.split(separator: "|")
            guard let left = Int(pages[0]), let right = Int(pages[1]) else {
                throw AoCError.parseError("failed to parse string to integer")
            }
            return (left, right)
        }
        
        let updates = blocks[1].split(whereSeparator: \.isNewline).map { line in
            return line.split(separator: ",").compactMap({ Int($0) })
        }
        
        return (orderPairs, updates)
    }
    
    private func task1(orderPairs: [(Int, Int)], updates: [[Int]]) throws -> Int {
        try updates.reduce(0) { result, update in
            let graph = buildGraph(update: update, orderPairs: orderPairs)
            guard let sortOrder = graph.topologicalSort() else {
                throw AoCError.invalidTask("failed to create topological sort")
            }
            
            if try isSorted(update: update, sortOrder: sortOrder) {
                return result + update[update.count/2]
            } else {
                return result
            }
        }
    }
    
    private func task2(orderPairs: [(Int, Int)], updates: [[Int]]) throws -> Int {
        try updates.reduce(0) { result, update in
            let graph = buildGraph(update: update, orderPairs: orderPairs)
            guard let sortOrder = graph.topologicalSort() else {
                throw AoCError.invalidTask("failed to create topological sort")
            }
            
            if try !isSorted(update: update, sortOrder: sortOrder) {
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
    
    private func isSorted(update: [Int], sortOrder: [Int]) throws -> Bool {
        try update.adjacentPairs().allSatisfy {
            guard let left = sortOrder.firstIndex(of: $0), let right = sortOrder.firstIndex(of: $1) else {
                throw AoCError.invalidState("item not found")
            }
            return left < right
        }
    }
}
