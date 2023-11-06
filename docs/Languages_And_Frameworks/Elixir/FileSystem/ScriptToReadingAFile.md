# Script example to read the contents of a file and write to STDOUT

```bash
$ touch file_reader.exs
$ vim file_reader
```

```elixir
defmodule FileReader do
    def read do
        {:ok, contents} = File.read("file/path.txt")
        IO.puts(contents)
    end
end

FileReader.read
```

```bash
:wq

$ elixir file_reader.exs
```