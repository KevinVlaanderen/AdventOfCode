import Foundation

public extension StringProtocol {
    @inlinable
    var blocks: [Self.SubSequence] {
        self.split(separator: "\n\n")
    }
}
