# Scaffold

The [**Scaffold**](https://pkg.go.dev/vincent.click/pkg/preflight/scaffold#Scaffold) is a set of methods that normally act as aliases for builtin functions, but can be replaced during testing to stub, mock, or spy on those functions.

In the package being tested, use the scaffolded functions instead of calling the builtin functions directly. This allows you to replace them in the test package. The `Reset` method restores all functions to their original state.