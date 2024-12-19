import Foundation

public extension StringProtocol {
    @inlinable
    var lines: [Self.SubSequence] {
        self.split(whereSeparator: \.isNewline)
    }
}
