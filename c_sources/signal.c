#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <string.h>

// Our signal handler
void child_signal_handler(int sig, siginfo_t *siginfo, void *context)
{
    printf("\n✅ Received signal: %d (%s)\n", sig, strsignal(sig));
    if (siginfo != NULL) {
        printf("🔍 Signal sent by PID: %d, UID: %d\n", siginfo->si_pid, siginfo->si_uid);
    }
    exit(0);  // exit after handling
}

int main() {
    struct sigaction phan;

    // Set handler
    phan.sa_sigaction = child_signal_handler;
    sigemptyset(&phan.sa_mask);
    phan.sa_flags = SA_SIGINFO;

    // Register for SIGINT (Ctrl+C)
    if (sigaction(SIGINT, &phan, NULL) == -1) {
        perror("sigaction");
        exit(1);
    }

    printf("👋 Waiting for SIGINT (press Ctrl+C)...\n");

    // Infinite loop to keep the process alive
    while (1) {
        sleep(1);
    }

    return 0;
}