import Foundation
internal import Algorithms

struct Day2: Day {
    typealias R = Int
    
    func task1(data: String, param: P) throws -> R {
        let reports = parse(data: data)

        return reports.count(where: checkReport)
    }
    
    func task2(data: String, param: P) throws -> R {
        let reports = parse(data: data)

        return reports.count { report in
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
    
    private func parse(data: String) -> ([[Int]]) {
        return data.split(whereSeparator: \.isNewline).map { line in
            line.split(whereSeparator: \.isWhitespace).map({ Int($0)! })
        }
    }
}
