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

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      steps {
        container('docker') {
          sh 'docker build -t registry.zlp-cloud.ru/backend-java:${BRANCH_NAME} -f ./app/docker/Dockerfile ./app'
        }
      }
    }

    stage('Push') {
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
            sh 'pwd'
            sh 'ls -la'
            echo 'Deploy'
          }
        }
      }
    }
  }

}
