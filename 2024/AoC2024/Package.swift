// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "AoC2024",
    platforms: [.macOS(.v14)],
    dependencies: [
        .package(url: "https://github.com/apple/swift-algorithms.git", from: "1.1.0"),
        .package(url: "https://github.com/davecom/SwiftGraph", from: "3.1.0"),
        .package(url: "https://github.com/ordo-one/package-jemalloc", from: "1.0.0"),
        .package(url: "https://github.com/ordo-one/package-benchmark", from: "1.27.4")
    ],
    targets: [
        .executableTarget(
            name: "Benchmarks",
            dependencies: [
                .product(name: "jemalloc", package: "package-jemalloc"),
                .product(name: "Benchmark", package: "package-benchmark"),
                "Data"
            ],
            path: "Benchmarks/Tasks",
            plugins: [
                .plugin(name: "BenchmarkPlugin", package: "package-benchmark")
            ]
        ),
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
                "Data",
                "Tasks"
            ]),
    ]
)
