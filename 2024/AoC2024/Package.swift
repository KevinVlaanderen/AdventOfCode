// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "AoC2024",
    platforms: [.macOS(.v14)],
    products: [
        // Products define the executables and libraries a package produces, making them visible to other packages.
        .executable(
            name: "AoC2024",
            targets: ["CLI"]),
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-algorithms.git", from: "1.1.0"),
        .package(url: "https://github.com/google/swift-benchmark", from: "0.1.2"),
        .package(url: "https://github.com/davecom/SwiftGraph", from: "3.1.0")
    ],
    targets: [
        // Targets are the basic building blocks of a package, defining a module or a test suite.
        // Targets can depend on other targets in this package and products from dependencies.
        .executableTarget(
            name: "CLI",
            dependencies: [
                .product(name: "Benchmark", package: "swift-benchmark"),
                "Data"
            ]),
        .target(
            name: "Framework"),
        .target(
            name: "Tasks",
            dependencies: [
                .product(name: "Algorithms", package: "swift-algorithms"),
                "SwiftGraph",
                "Framework"
            ]),
        .target(
            name: "Data",
            dependencies: [
                "Framework",
                "Tasks"
            ]),
        .testTarget(
            name: "TasksTests",
            dependencies: [
                "Framework",
                "Data"
            ]),
    ]
)
