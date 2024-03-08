pipeline {
  agent { kubernetes {} }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      agent {
        kubernetes {
          yaml '''
          spec:
            containers:
            - name: gradle
              image: gradle:jdk17
              command: ['cat']
              tty: true
          '''
        }
      }
      steps {
        container('gradle'){
          dir('app') {
            sh './gradlew clean build'
          }
        }
      }
    }

    stage('Deploy') {
      agent {
        kubernetes {
          yaml '''
          spec:
            containers:
            - name: helm
              image: ibmcom/k8s-helm:v2.6.0
              command: ['cat']
              tty: true
          '''
        }
      }
      when {
        anyOf {
          branch 'master'
          branch 'dev'
        }
      }
      steps {
        container('helm'){
          sh 'pwd'
          sh 'ls -la'
          dir('chart') {
            echo 'Deploy'
            sh 'pwd'
            sh 'ls -la'
          }
        }
      }
    }
  }

}
