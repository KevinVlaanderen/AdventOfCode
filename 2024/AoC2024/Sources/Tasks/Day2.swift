import Foundation
internal import Algorithms
import Framework

public struct Day2: Day {
    public typealias R = Int
    
    public init() {}
    
    public func perform(task: Task, data: String, param: ()) throws -> Int {
        let reports = parse(data: data)

        switch task {
        case .task1:
            return task1(reports: reports)
        case .task2:
            return task2(reports: reports)
        @unknown default:
            fatalError("unknown task")
        }
    }
    
    private func parse(data: String) -> ([[Int]]) {
        return data.split(whereSeparator: \.isNewline).map { line in
            line.split(whereSeparator: \.isWhitespace).map({ Int($0)! })
        }
    }
    
    private func task1(reports: [[Int]]) -> R {
        reports.count(where: checkReport)
    }
    
    private func task2(reports: [[Int]]) -> R {
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
        return diffs.allSatisfy({ abs($0) >= 1 && abs($0) <= 3}) &&
               diffs.dropFirst().allSatisfy({ $0.signum() == diffs.first!.signum() })
    }
}
