node {
    def testapp
    stage('Clone Repository') {
      checkout scm
    }
    stage("Build image") {
      testapp = docker.build("goloang/httpserver")
    }
    stage('Test Image') {
      app.inside {
        sh "echo test passed"
      }
    }
    stage('Push Image') {
      docker.withRegistery('https://registry.hub.docker.com', 'docker-hub-credentials') {
        app.push("${env.BUILD_NUMBER}")
        app.push("latest")
      }
    }
}
