# User Instruction Memory

This file records user instructions, preferences, and teachings for reference in future interactions.

## Entries

### User Instruction Summary
- Date: 2026-09-03
- Context: Frontend maintenance task
- Instructions:
  - Use `apply_patch` for file edits.
  - Preserve existing uncommitted user changes.
  - Keep backend files unchanged when the task is scoped to the frontend.
  - Run an appropriate frontend typecheck or targeted test after changes.

### User Instruction Summary
- Date: 2026-09-03
- Context: Every frontend production build
- Instructions:
  - Before a frontend production build, stop the managed Vite preview terminal to release memory.
  - Run type checking and Vite packaging as separate commands.
  - Set `NODE_OPTIONS=--max-old-space-size=3072` for the type-check command.

### Project Knowledge Summary
- Date: 2026-09-05
- Context: Discovered by Agent while performing frontend affiliate-subsystem cleanup
- Category: Environment Configuration
- Instructions:
  - The `apply_patch` tool is not available in this environment. Use the `edit` and `write` tools for file edits.
  - Run `vue-tsc --noEmit` (with `NODE_OPTIONS=--max-old-space-size=3072`) from `/workspace/frontend` for frontend type checking; exit code 0 with empty output means success.
