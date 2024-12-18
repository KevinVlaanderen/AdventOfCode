// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "AoC2024",
    platforms: [.macOS(.v14)],
    products: [
        .library(name: "Data", targets: ["Data"]),
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-algorithms.git", from: "1.1.0"),
        .package(url: "https://github.com/davecom/SwiftGraph", from: "3.1.0"),
        .package(url: "https://github.com/SwiftGen/SwiftGenPlugin", from: "6.6.2"),
        .package(url: "https://github.com/krzysztofzablocki/Sourcery.git", from: "2.2.5"),
        .package(url: "https://github.com/apple/swift-collections.git", from: "1.1.4"),
    ],
    targets: [
        .target(
            name: "Framework"),
        .target(
            name: "Tasks",
            dependencies: [
                .product(name: "Algorithms", package: "swift-algorithms"),
                .product(name: "HeapModule", package: "swift-collections"),
                "SwiftGraph",
                "Framework"
            ]),
        .target(
            name: "Data",
            dependencies: [
                "Framework",
                "Tasks"
            ],resources: [
                .copy("Input/."),
            ]),
        .testTarget(
            name: "TaskTests",
            dependencies: [
                "Framework",
                "Data",
                "Tasks",
            ]),
    ]
)
