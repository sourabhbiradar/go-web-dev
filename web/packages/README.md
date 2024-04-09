# packaes

Packages are origanized reuseable source codes.

Functions from packages are exported/imported only if their name starts with CAPITAL LETTER.

Packages of same directory MUST have same package name.

Packages of same Dir can access even unexported funcs.

- main.go imports package math
- math.go has package math that we craeted