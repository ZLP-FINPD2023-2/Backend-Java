pipeline {
  agent {
    kubernetes {
      containerTemplate {
        name 'gradle'
        image 'gradle:jdk17'
        command 'cat'
        ttyEnabled true
      }
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
