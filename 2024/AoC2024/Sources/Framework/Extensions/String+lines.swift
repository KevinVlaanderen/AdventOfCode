import Foundation

public extension StringProtocol {
    var lines: [Self.SubSequence] {
        self.split(whereSeparator: \.isNewline)
    }
}
