module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  transform: {
    "^.+\\.(ts|tsx)$": "ts-jest",
  },
  modulePathIgnorePatterns: ["build"],
  reporters: [
    "default",
    [
      "jest-junit",
      {
        outputDirectory: "./build/test-results",
        suiteNameTemplate: "{filename}",
      },
    ],
  ],
};
