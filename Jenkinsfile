node {
  checkout scm
}

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
        - name: helm
          image: ibmcom/k8s-helm:v2.6.0
          command: ['cat']
          tty: true
        volumes:
        - name: docker-sock
          hostPath:
            path: /var/run/docker.sock
      '''
    }
  }

  stages {
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

    stage('Deploy') {
      when {
        anyOf {
          branch 'master'
          branch 'dev'
        }
      }
      options {
        timeout(time: 1, unit: 'MINUTES')
      }
      steps {
        container('helm') {
          dir('chart') {
            sh 'helm version'
            sh 'helm upgrade --install --namespace lfp-dev backend .'
          }
        }
      }
    }
  }
}
