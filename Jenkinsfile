pipeline {
  agent {
    kubernetes {
      yaml '''
      spec:
        containers:
        - name: docker
          image: docker:latest
          securityContext:
            privileged: true
            runAsUser: 0
          command:
          - cat
          tty: true
          volumeMounts:
          - mountPath: /var/run/docker.sock
            name: docker-sock
        volumes:
        - name: docker-sock
          hostPath:
            path: /var/run/docker.sock
      '''
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

    stage('Build image') {
      steps {
        container('docker') {
          sh 'docker build -t registry.zlp-cloud.ru/backend-java:${BRANCH_NAME} -f ./app/docker/Dockerfile ./app'
        }
      }
    }

    stage('Push image') {
      when {
        anyOf {
          branch 'master'
          branch 'dev'
        }
      }
      steps {
        container('docker') {
          sh 'docker push registry.zlp-cloud.ru/backend-java:${BRANCH_NAME}'
        }
      }
    }
  }
}
