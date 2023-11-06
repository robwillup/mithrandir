GitHub Actions
Communicating to the Host

Loggin Commands: Instructions for the workflow host that are embedded in the build log

Creating an environment variable
::set-env name=DEBUG::1

Creating an output parameter
::set-output name=name_of_param::value

Adding to path
::add-path::/path/to/directory

Printing a Log Message
::debug file=name,line=0,col=0::message
::warning file=name,line=0,col=0::message
::error file=name,line=0,col=0::message

Printing a Masked Message
::add-mask::$VARIABLE (replaces by *****)
