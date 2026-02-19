# Product Guidelines: go-torrent

## Visual Style and UX
- **Primary Style:** Functional and minimalist. Focus on delivering essential information clearly and concisely to the terminal.
- **Protocol Clarity:** Incorporate a "sprinkle" of descriptive and educational feedback to explain the BitTorrent protocol's inner workings at each stage.
- **Interactive Elements:** Use modern terminal UI features, such as progress bars and dynamic status lines, to enhance the visual feedback for long-running operations.

## Error Handling and Communication
- **Critical Errors:** Adopt a fail-fast approach. For fatal errors where the client cannot proceed, stop immediately and provide a single, clear, and actionable error message.
- **Non-Critical Errors:** Implement a silent retry mechanism for recoverable errors (e.g., connection timeouts). These should be handled automatically to avoid UI clutter.
- **Diagnostic Logging:** Maintain a separate debug window or view to provide detailed logs of the client's internal state and operations. This allows advanced users and developers to troubleshoot without disrupting the default functional view.

## Performance and Reliability
- **Resource Efficiency:** Prioritize low memory usage and efficient CPU utilization, consistent with the project's goal of being a lightweight client.
- **Protocol Robustness:** Ensure the client can gracefully handle various network conditions and protocol edge cases.
