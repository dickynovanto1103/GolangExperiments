### Sentry note
1. If we don't init sentry first, then running directly the `sentry.CaptureException` will result in nil pointer dereference
    - So, we cannot just put sentry.CaptureException in logic code, as it makes the code untestable
    - Therefore, if we need to do manual reporting using sentry, we should use interface (need to verify this idea)