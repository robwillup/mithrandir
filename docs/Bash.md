# Running a command multiple times

for i in {< START >..< FINISH >}; < COMMAND > ; done

```bash
$ for i in {1..10}; dotnet test ; done
```