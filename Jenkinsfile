pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: jenkins
spec:
  containers:
  - name: gradle
    image: gradle:jdk17
    command:
    - cat
    tty: true
"""
    }
  }

  options {
    skipDefaultCheckout()
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      steps {
        container('gradle'){
          dir("app") {
            sh './gradlew clean build'
          }
        }
      }
    }

    stage('Deploy') {
      when {
        branch 'master'
      }
      steps {
        echo 'Deploy'
      }
    }
  }

}
