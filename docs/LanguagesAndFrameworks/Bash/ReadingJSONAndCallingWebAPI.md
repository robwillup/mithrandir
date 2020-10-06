# About Shell scripting

Each script starts with a "shebang" and the path to the shell that you want to use, like so:

      #!/bin/bash

The `#!` combo is called a `shebang` by most Unix geeks. This is used by the shell to decide which interpreter to run the rest of the script, and ignored by the shell that actually runs the script. 

Scritps can be written for all kinds of interpreters - bash, tsch, zsh, or other shells, or for Perl, Python, and so on. You could even omit that line if you wanted to run the script by sourcing it at the shell. But let's add it to allow scripts to be run non-interactively.

Commemts are prefaced with the hash (#) character.

      #!/bin/bash
      # A simple script
      
## Taking Input

Non-interactive scripts are useful, but what if you need to give the script new information each time it's run? One thing you can do is take an argument from the command line. So. for instance, when you run "script foo" the script will take the name of the first argument (foo):

      #!/bin/bash
      
      echo $1
      
 Here bash will read the command line and echo (print) the first argument - that is, the first string after the command itself.
 
 You can also use read to accept user input:
 
       #!/bin/bash
       echo -e "Please enter your name: "
       read name
       echo "Nice to meet you $name"
       
 That script will wait for the user to type in their name (or any other input for that matter) and use it as the variable $name.
