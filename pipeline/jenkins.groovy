pipeline {
    agent any

    parameters {
        choice(
            name: 'OS',
            choices: ['linux', 'darwin', 'windows'],
            description: 'Target operating system'
        )
        choice(
            name: 'ARCH',
            choices: ['amd64', 'arm64'],
            description: 'Target architecture'
        )
        booleanParam(
            name: 'SKIP_TESTS',
            defaultValue: false,
            description: 'Skip running tests'
        )
        booleanParam(
            name: 'SKIP_LINT',
            defaultValue: false,
            description: 'Skip running linter'
        )
    }

    // Read by the Makefile (TARGETOS / TARGETARCH / EXT). Taken from params, not from the
    // OS / ARCH env vars: the first build of the job has no parameters yet and only params
    // falls back to the defaults
    environment {
        TARGETOS   = "${params.OS}"
        TARGETARCH = "${params.ARCH}"
        EXT        = "${params.OS == 'windows' ? '.exe' : ''}"
    }

    stages {
        stage('Lint') {
            when { expression { !params.SKIP_LINT } }
            steps {
                sh 'make lint'
            }
        }

        stage('Test') {
            when { expression { !params.SKIP_TESTS } }
            steps {
                sh 'make test'
            }
        }

        stage('Build') {
            steps {
                echo "Building gbot for ${params.OS}/${params.ARCH}"
                sh 'make build'
                // An empty value in `environment` leaves the variable unset (null)
                archiveArtifacts artifacts: "gbot${env.EXT ?: ''}", fingerprint: true
            }
        }
    }

    post {
        // The workspace is reused between builds: drop binaries of other platforms
        always {
            deleteDir()
        }
    }
}
