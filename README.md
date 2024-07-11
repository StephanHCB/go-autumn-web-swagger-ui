# go-autumn-web-swagger-ui

## IMPORTANT Third Party Licensing Information

The static files included under third_party/swagger_ui are

                          Copyright 2020-2021 SmartBear Software Inc.

They are licensed under

                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

See [here](https://github.com/StephanHCB/go-autumn-web-swagger-ui/blob/master/third_party/swagger_ui/LICENSE) for the full 
licensing information. This LICENSE file is also included with the static distribution provided by this
module. If you use it as instructed, it will be statically served under `/swagger-ui/LICENSE`

Original download location:

https://github.com/swagger-api/swagger-ui

## About go-autumn-web-swagger-ui

A library that provides a static go source file for the [swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui)
distributable ui module.

## How to regenerate

- update swagger ui files
- update default url from 'https://petstore.swagger.io/v2/swagger.json' to '/v3/api-docs' in swagger-initializer.js
- do not forget to retain LICENSE and README.txt
- delete assets_vfsdata.go
- rename the package to `main` in main.go
- run main.go
- change package back to `auwebswaggerui` in both main.go and generated file
- capitalize var assets to Assets so it becomes exported

## How to use

Simply import the module using go.mod.

The main.go file is really just here to re-generate the static source code, if swagger-ui has been updated.

`auwebswaggerui.Assets` is an instance of http.FileSystem, so you can do stuff like

```
  http.Handle("/swagger-ui/", http.FileServer(auwebswaggerui.Assets))
```
