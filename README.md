# go-autumn-web-swagger-ui

## About go-autumn

A collection of libraries for [enterprise microservices](https://github.com/StephanHCB/go-mailer-service/blob/master/README.md) in golang that
- is heavily inspired by Spring Boot / Spring Cloud
- is very opinionated
- names modules by what they do
- unlike Spring Boot avoids certain types of auto-magical behaviour
- is not a library monolith, that is every part only depends on the api parts of the other components
  (if at all), and the api parts do not add any dependencies.

Fall is my favourite season, so I'm calling it _go-autumn_.

## IMPORTANT Third Party Licensing Information

The static files included under third_party/swagger_ui are

                          Copyright 2020 SmartBear Software Inc.

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

## How to use

Simply import the module using go.mod.

The main.go file is really just here to re-generate the static source code, if swagger-ui has been updated.

```
TODO further instructions
```
