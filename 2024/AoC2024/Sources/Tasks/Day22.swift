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
        
        return secrets.reduce(0) { result, secret in
            return result + (0..<2000).reduce(secret) { newSecret, _ in
                var newSecret = newSecret
                newSecret = ((newSecret*64)^newSecret)%16777216
                newSecret = ((newSecret/32)^newSecret)%16777216
                newSecret = ((newSecret*2048)^newSecret)%16777216
                return newSecret
            }
        }
    }
    
    private func parse(data: String) throws -> [Secret] {
        try data.lines.map(toInt)
    }
}

private typealias Secret = Int
