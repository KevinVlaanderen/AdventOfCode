// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "AoC2024Benchmarks",
    platforms: [.macOS(.v14)],
    dependencies: [
        .package(path: "../AoC2024"),
        .package(url: "https://github.com/ordo-one/package-benchmark", from: "1.27.4"),
        .package(url: "https://github.com/ordo-one/package-jemalloc", from: "1.0.0"),
    ],
    targets: [
        .executableTarget(
            name: "Benchmarks",
            dependencies: [
                .product(name: "jemalloc", package: "package-jemalloc"),
                .product(name: "Benchmark", package: "package-benchmark"),
                .product(name: "Data", package: "AoC2024"),
            ],
            path: "Benchmarks/Tasks",
            plugins: [
                .plugin(name: "BenchmarkPlugin", package: "package-benchmark")
            ]
        ),
    ]
)
