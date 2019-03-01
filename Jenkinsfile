@Library('github.com/yadavsms/jenkins-shared-lib') _
properties([parameters([string(name: 'goVersion', defaultValue: '1.11', description: 'Which version of Go language to use.')])])
standardBuild environment: "golang:${params.goVersion}",
    mainScript: '''
go version
go build -v httpserver.go
''',
    postScript: '''
ls -l
./httpserver
'''