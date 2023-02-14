# Retry Strategies for Transient Failures

[Reference](https://www.ayush.nz/2017/08/retry-strategies-for-transient-failures)

It is a fact that computer networking and computer systems that rely on network fail.
Transient failures are also known as temporary failures. It is possible that a remote
resource is only temporarily unavailable due to network congestion or high system load.
When your application that relies on a remove resource encounters a transient network
failure, it is important that it waits for some time and retries again, before
considering it a permanent failure and exiting.

In most cases you don't want to keep retrying forever, so there must be a timeout.

There are many strategies where you "back off" for some time and retry the operation
after a delay.

## Backoff Strategies

Let's consider a hypothetical `DoSomething()` function, in which your code performs some
operation which is prone to transient failures. Let's consider the strategies using examples.

### No Backoff

Remember to cap your attempts using maximum allowed retries otherwise this might go on
forever:

```csharp
bool NoBackoff(int maxRetry)
{
	int counter = 0;
	while (counter < maxRetry)
	{
		response = DoSomething();
		if (response)
		{
			return true;
		}
		Console.WriteLine($"No backoff. Retry number: {counter}");
		counter++;
	}
	return false;
}
```

### Constant Backoff

Constant backoff adds a fixed delay after every attemp. In this case, we wait a fixed
1 second after every attempt.

```csharp
bool ConstantBackoff(int maxRetry)
{
	int totalDelay = 0;
	counter = 0;
	while (counter < maxRetry)
	{
		response = DoSomething();
		if (response)
		{
			return true;
		}
		TimeSpan sleepTime = TimeSpan.FromSeconds(1);
		Console.WriteLine($"Constant backoff. Retry number: {counter}. Sleeping for: {sleepTime}");
		totalDelay += sleepTime;
		Thread.Sleep(sleepTime);
		counter++;
	}
	WriteLine($"Constant backoff. Total delay: {totalDelay}.");
	return false;
}
```