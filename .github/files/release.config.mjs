/**
 * @type {import('semantic-release').GlobalConfig}
 */
export default {
    branches: ['main'],
    tagFormat: '${version}',
    plugins: [
        [
            '@semantic-release/commit-analyzer',
            {
                preset: 'angular',
                parserOpts: {
                    noteKeywords: ['BREAKING CHANGE', 'BREAKING CHANGES', 'BREAKING'],
                },
            },
        ],
        [
            '@semantic-release/release-notes-generator',
            {
                preset: 'angular',
                parserOpts: {
                    noteKeywords: ['BREAKING CHANGE', 'BREAKING CHANGES', 'BREAKING'],
                },
                writerOpts: {
                    commitsSort: ['subject', 'scope'],
                },
            },
        ],
        [
            '@semantic-release/changelog',
            {
                changelogFile: 'CHANGELOG.md',
            },
        ],
        [
            '@semantic-release/exec',
            {
                prepareCmd: 'NETCLI_VERSION=${nextRelease.version} bash .github/files/build.sh',
            },
        ],
        [
            '@semantic-release/git',
            {
                assets: ['CHANGELOG.md', 'docs/**', 'README.md', 'assets/static/Version.txt', 'web-run.sh'],
                message: 'chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}',
            },
        ],
        [
            '@semantic-release/github',
            {
                assets: [
                    {
                        path: 'build/*',
                    },
                ],
            },
        ],
    ],
};
