#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>

int main() {
    pid_t pid = fork();

    if (pid < 0) {
        perror("Fork failed");
        exit(EXIT_FAILURE);
    }

    if (pid > 0) {
        // Parent process, exits immediately to let the child run as a daemon
        printf("Parent process: PID = %d\n", getpid());
        exit(0);
    }

    // Child process - continues as a daemon
    printf("Child process: PID = %d\n", getpid());

    // Detach from the terminal and run as a background process
    if (setsid() < 0) {
        perror("Failed to create new session");
        exit(EXIT_FAILURE);
    }

    // The child now runs independently, even if the parent process dies
    while (1) {
        // Simulate some background task
        sleep(1);
        printf("Daemon process running...\n");
    }

    return 0;
}