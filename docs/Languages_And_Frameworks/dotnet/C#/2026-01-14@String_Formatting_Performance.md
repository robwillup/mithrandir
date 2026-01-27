# Performance Study on String Formatting Options

In .NET Framework, the default string formatting method was `string.Format`, which requires the runtime to call
`ToString` on all its arguments. For example, `string.Format("{0} = {1}", Key, Value)` would internally call both
`Key.ToString()` and `Value.ToString()` to produce the final result. These two extra allocations introduce a
performance penalty, which is unfortunate because we don't need those temporary strings - we only need the final
formatted result.

Most people still assume that interpolated strings like `$"{Key} = {Value}"` work the same way. But that is not
the case! .NET Core has come a long way in eliminating unnecessary allocations and boxing. Let's explore
high-performance string formatting alternatives that help you avoid temporary string allocations, and in some cases,
avoid them altogether.

## How String Interpolation Works


