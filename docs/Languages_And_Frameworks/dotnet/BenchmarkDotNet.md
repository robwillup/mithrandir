# Getting started with BenchmarkDotNet

## Step 1: Create a project

Create a new console application

## Step 2: Installation

Install BenchmarkDotNet via the NuGet package: [BenchmarkDotNet](https://www.nuget.org/packages/BenchmarkDotNet/)

## Step 3. Design a benchmark

Write a class with methods that you want to measure and mark them with the Benchmark attribute. In the following example,
we compare MD5 with SHA256 cryptographic hash functions.

```csharp
using System;
using System.Security.Cryptography;
using BenchmarkDotNet.Attributes;
using BenchmarkDotNet.Running;

namespace MyBenchmarks;

var summary = BenchmarkRunner.Run<Md5VsSha256>();

public class Md5VsSha256
{
    private const int N = 10000;
    private readonly byte[] data;

    private readonly SHA256 sha256 = SHA256.Create();
    private readonly MD5 md5 = MD5.Create();

    public Md5VsSha256()
    {
        data = new byte[N];
        new Random(42).NextBytes(data);
    }

    [Benchmark]
    public byte[] Sha256() => sha256.ComputeHash(data);

    [Benchmark]
    public byte[] Md5() => md5.ComputeHash(data);
}
```

The `BenchmarkRunner.Run<Md5VsSha256>()` call runs your benchmarks and prints results to the console.

## Step 4. Run benchmarks

Start your console application to run the benchmarks. The application must be built in the Release configuration.

```shell
dotnet run -c Release
```

## Step 5. View results

View the results. Here is an example of output from the above benchmarks:

```shell
BenchmarkDotNet=v0.13.2, OS=Windows 10 (10.0.19045.2251)
Intel Core i7-4770HQ CPU 2.20GHz (Haswell), 1 CPU, 8 logical and 4 physical cores
.NET SDK=7.0.100
  [Host]     : .NET 7.0.0 (7.0.22.51805), X64 RyuJIT AVX2
  DefaultJob : .NET 7.0.0 (7.0.22.51805), X64 RyuJIT AVX2


| Method |     Mean |    Error |   StdDev |
|------- |---------:|---------:|---------:|
| Sha256 | 51.57 us | 0.311 us | 0.291 us |
|    Md5 | 21.91 us | 0.138 us | 0.129 us |

```

## Step 6. Analyze results

BenchmarkDotNet will automatically create basic reports in the `BenchmarkDotNet.Artifacts/results` folder that can be shared an analyzed.

To help analyze performance in further depth, you can configure your benchmark to collect and output more detailed information.
Benchmark configuration can be conveniently changed by adding attributes to the class containing your benchmarks. For example:

* Diagnosers
  * GC allocations: `[MemoryDiagnoser]`
  * Code size and disassembly: `[DisassemblyDiagnoser]`
  * Threading statistics: `[ThreadingDiagnoser]`
* Exporters
  * CSV reports with raw data: `[CsvMeasurementsExporters]`
  * JSON reports with raw data: `[JsonExporter]`
  * Plots (if you have installed R): `[RPlotExporter]`

