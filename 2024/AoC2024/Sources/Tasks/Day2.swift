import Foundation
internal import Algorithms
import Framework

public final class Day2: Day<Task, Int> {
    public override func perform(data: String, param: P) throws -> R {
        let reports = try parse(data)

        return switch param {
        case .task1: task1(reports: reports)
        case .task2: task2(reports: reports)
        }
    }
    
    private func parse(_ data: String) throws -> ([[Int]]) {
        try data.lines.map { line in
            try line.words.map(toInt)
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
        return diffs.allSatisfy({ abs($0) >= 1 && abs($0) <= 3}) && diffs.dropFirst().allSatisfy {
            guard let firstDiff = diffs.first else {
                return false
            }
            return $0.signum() == firstDiff.signum()
        }
    }
}
