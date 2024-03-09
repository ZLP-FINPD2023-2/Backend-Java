pipeline {
  agent { kubernetes {} }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Deploy') {
      when {
        anyOf {
          branch 'master'
          branch 'dev'
        }
      }

      agent {
        kubernetes {
          yaml '''
          spec:
            containers:
            - name: helm
              image: ibmcom/k8s-helm:v2.6.0
              command: ['cat']
              tty: true
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
        stage('Build image') {
          steps {
            container('docker') {
              sh 'docker build -t registry.zlp-cloud.ru/backend-java:${BRANCH_NAME} -f ./app/docker/Dockerfile ./app'
            }
          }
        }

        stage('Push image') {
          steps {
            container('docker') {
              sh 'docker push registry.zlp-cloud.ru/backend-java:${BRANCH_NAME}'
            }
          }
        }

        stage('Deploy') {
          steps {
            container('helm') {
              dir('chart') {
                sh 'helm upgrade --namespace lfp-dev backend .'
              }
            }
          }
        }
      }
    }
  }
}
