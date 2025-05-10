#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <string.h>
#include <fcntl.h>
#include <sys/types.h>

// Signal handler
void child_signal_handler(int sig, siginfo_t *siginfo, void *context)
{
    FILE *log = fopen("/tmp/daemon_signal_log.txt", "a");
    if (log) {
        fprintf(log, "\n✅ Received signal: %d (%s)\n", sig, strsignal(sig));
        if (siginfo != NULL) {
            fprintf(log, "🔍 Signal sent by PID: %d, UID: %d\n", siginfo->si_pid, siginfo->si_uid);
        }
        fclose(log);
    }

    // Exit daemon after receiving signal (optional, for testing)
    exit(0);
}

void daemonize()
{
    pid_t pid;

    // First fork
    if ((pid = fork()) < 0) {
        perror("fork");
        exit(EXIT_FAILURE);
    }
    if (pid > 0) {
        exit(0);  // Parent exits
    }

    // Create new session, lose controlling terminal
    if (setsid() < 0) {
        perror("setsid");
        exit(EXIT_FAILURE);
    }

    // Ignore SIGHUP (optional)
    signal(SIGHUP, SIG_IGN);

    // Second fork
    if ((pid = fork()) < 0) {
        perror("fork");
        exit(EXIT_FAILURE);
    }
    if (pid > 0) {
        exit(0);  // First child exits
    }

    // Change working directory (optional)
    chdir("/");

    // Redirect standard fds to /dev/null
    close(STDIN_FILENO);
    close(STDOUT_FILENO);
    close(STDERR_FILENO);
    open("/dev/null", O_RDONLY);  // stdin
    open("/dev/null", O_WRONLY);  // stdout
    open("/dev/null", O_RDWR);    // stderr
}

int main()
{
    struct sigaction phan;

    // Daemonize the process
    daemonize();

    // Set up sigaction handler
    phan.sa_sigaction = child_signal_handler;
    sigemptyset(&phan.sa_mask);
    phan.sa_flags = SA_SIGINFO;

    // Catch SIGINT for testing (can also use SIGTERM etc.)
    if (sigaction(SIGINT, &phan, NULL) == -1) {
        perror("sigaction");
        exit(EXIT_FAILURE);
    }

    // Since stdio is redirected, write to a file for visibility
    FILE *fp = fopen("/tmp/daemon_log.txt", "w");
    if (!fp) {
        exit(EXIT_FAILURE);
    }

    fprintf(fp, "👻 Daemon started. PID: %d. Send SIGINT to test signal handler.\n", getpid());
    fflush(fp);

    // Infinite loop
    while (1) {
        sleep(1);
    }

    return 0;
}