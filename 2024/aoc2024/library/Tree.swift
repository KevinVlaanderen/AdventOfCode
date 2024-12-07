enum Node<Value> {
    case root(value: Value)
    indirect case leaf(parent: Node, value: Value)

    var value: Value {
        switch self {
        case .root(let value):
            return value
        case .leaf(_, let value):
            return value
        }
    }
}
