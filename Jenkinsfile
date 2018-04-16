node {
    def app
    stage('Clone Repository') {
      checkout scm
    }
    stage("Build image") {
      app = docker.build("goloang/httpserver")
    }
    stage('Test Image') {
      app.inside {
        sh "echo test passed"
      }
    }
    stage('Push Image') {
      withDockerRegistry(registry: [url: 'https://index.docker.io/v1/', credentialsId:'docker-hub-credentials']) {
        app.push("${env.BUILD_NUMBER}")
        app.push("latest")
      }
    }
}
