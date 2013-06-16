jsonutils
=========

Converter for JSON data to a Go struct or a Java class for GSON

	go get github.com/bashtian/jsonutils

You can print the structure of a JSON from a URL

	jsonutils https://api.github.com/users/bashtian/repos

or from a file

	jsonutils -f file.json

If you want to print example data as comments, use the -x parmeter

	jsonutils -x https://api.github.com/users/bashtian/repos

You can also print Java code 

	jsonutils -j https://api.github.com/users/bashtian/repos