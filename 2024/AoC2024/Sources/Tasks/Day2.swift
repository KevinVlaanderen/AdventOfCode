import Foundation
internal import Algorithms
import Framework

public struct Day2: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let reports = parse(data: data)

        return switch param {
        case .task1:
            task1(reports: reports)
        case .task2:
            task2(reports: reports)
        }
    }
    
    private func parse(data: String) -> ([[Int]]) {
        data.split(whereSeparator: \.isNewline).map { line in
            line.split(whereSeparator: \.isWhitespace).compactMap({ Int($0) })
        }
    }
    
    private func task1(reports: [[Int]]) -> Int {
        reports.count(where: checkReport)
    }
    
    private func task2(reports: [[Int]]) -> Int {
        reports.count { report in
            if checkReport(report) {
                return true
            }
            
            for i in 0..<report.count {
                var reportCopy = report
                reportCopy.remove(at: i)
                if checkReport(reportCopy) {
                    return true
                }
            }
            return false
        }
    }
    
    private func checkReport(_ report: [Int]) -> Bool {
        let diffs = report.adjacentPairs().map({ $0.0 - $0.1})
        return diffs.allSatisfy({ abs($0) >= 1 && abs($0) <= 3}) && diffs.dropFirst().allSatisfy {
            guard let firstDiff = diffs.first else {
                return false
            }
            return $0.signum() == firstDiff.signum()
        }
    }
}
