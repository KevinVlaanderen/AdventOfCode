import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day22: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let secrets = try parse(data: data)
        
        return switch param {
        case .task1: task1(secrets: secrets)
        case .task2: task2(secrets: secrets)
        }
    }
    
    private func parse(data: String) throws -> [Secret] {
        try data.lines.map(toInt)
    }
    
    private func task1(secrets: [Secret]) -> Int {
        secrets.reduce(0) { result, secret in
            result + (0..<2000).reduce(secret) { newSecret, _ in
                var newSecret = newSecret
                newSecret = ((newSecret*64)^newSecret)%16777216
                newSecret = ((newSecret/32)^newSecret)%16777216
                newSecret = ((newSecret*2048)^newSecret)%16777216
                return newSecret
            }
        }
    }
    
    private func task2(secrets: [Secret]) -> Int {
        let scores: [String: Int] = [:]
        
        return secrets.reduce(into: scores) { scores, secret in
            var lastPrice = secret%10
            var diffs: [Int] = [0,0,0,0]
            
            var localScores: [String: Int] = [:]
            
            var newSecret = secret
            for i in (0..<2000) {
                newSecret = ((newSecret*64)^newSecret)%16777216
                newSecret = ((newSecret/32)^newSecret)%16777216
                newSecret = ((newSecret*2048)^newSecret)%16777216
                
                let price = newSecret%10
                
                if i < 4 {
                    diffs[i] = price-lastPrice
                } else {
                    diffs[0] = diffs[1]
                    diffs[1] = diffs[2]
                    diffs[2] = diffs[3]
                    diffs[3] = price-lastPrice
                    
                    let key = diffs.map({ String($0) }).joined()
                    
                    if localScores[key] == nil {
                        localScores[key] = price
                    }
                }
                
                lastPrice = price
            }
            
            for localScore in localScores {
                let current = scores[localScore.key] ?? 0
                scores[localScore.key] = current + localScore.value
            }
        }.values.sorted().last!
    }
}

private typealias Secret = Int
