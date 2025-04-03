#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>

int main() {
    pid_t pid;

    // Create a new child process using fork()
    pid = fork();

    if (pid < 0) {
        // Fork failed
        perror("Fork failed");
        exit(EXIT_FAILURE);
    }

    if (pid == 0) {
        // Child process: Set priority (nice) to a higher value (lower priority)
        nice(10);  // Increase the nice value to lower the priority
        
        printf("Child process: My PID is %d, Parent PID is %d\n", getpid(), getppid());
        // Simulate long-running process
        while (1) {
            sleep(1);
        }
    } else {
        // Parent process: Continue execution
        printf("Parent process: My PID is %d, Child PID is %d\n", getpid(), pid);
        wait(NULL);  // Wait for child to finish
        printf("Parent process: Child has finished\n");
    }

    return 0;
}