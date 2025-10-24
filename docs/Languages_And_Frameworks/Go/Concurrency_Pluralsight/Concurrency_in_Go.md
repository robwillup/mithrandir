# Intro - Oct 22, 2025

Why is concurrency such a big topic today?
Because it gives you much more processing power.

For example, the AMD Ryzen Threadripper PRO 5955WX processor, has a
- single-threaded rating of 3282
- multi-threaded rating of 97132

Intel Xeon Platinum 8380
- single-threaded rating 2385
- multi-threaded rating 62318

Concurrency allows us to take the most out of the infrastructure our applications are running on.

Go was designed with concurrency in mind.
- Goroutines allows us to kick off concurrent tasks 
- Channels allows goroutines to communicate safely and efficiently
- sync package
- communicating sequential processes (csp)


