pipeline {
  agent any
  stages {
    stage("Build Stage") {
      agent {
        docker {
          image 'golang:1.10-alpine'
        }
      }
      steps {
        sh 'go run httpserver.go'
      }
    }
  }
}
