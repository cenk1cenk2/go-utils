module.exports = {
  branches: ["main"],
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    [
      "@google/semantic-release-replace-plugin",
      {
        replacements: [
          {
            files: ["version.go"],
            from: 'const VERSION = "(.*)"',
            to: 'const VERSION = "v${nextRelease.version}"',
            countMatches: true,
          },
        ],
      },
    ],
    "@semantic-release/changelog",
    [
      "@semantic-release/git",
      {
        assets: ["CHANGELOG.md", "README.md"],
      },
    ],
  ],
};
