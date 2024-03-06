podTemplate(containers: [
    containerTemplate(name: 'gradle', image: 'gradle:jdk17', command: 'sleep', args: '99d'),
  ]) {

    node(POD_LABEL) {
        stage('Checkout') {
            checkout scm 
            container('gradle') {

                stage('Build') {
                    dir("app") {
                        sh './gradlew clean build'
                    }
                }

                stage('Deploy check') {
                    input "Deploy?"
                }

                stage('Deploy') {
                    echo 'Deploy'
                }

            }
        }
    }
}
