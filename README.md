# Sample Enkudo NBI API Work

This project includes sample Enkudo subscription API definition in OpenAPI3.0  YAML documents.

- "Enkudo_DSP_CommonData.yaml" defines common data definitions within the API and referenced from other API (YAML) documents. You do not need to provide this document any of validation and/or code generation programs.
- "Enkudo_DSP_Subscription.yaml" Enkudo Subscription API definition.

There are also a couple of helper scripts to validate the YAML files with respect to OpenAPI 3.0 specification and code generation.

"validate.sh" validates given YAML document utilizing two validators.

- "openapi-spec-validator" and
- "openapi-generator validate" 

"gencode.sh" is a simple wrapper to openapi-generator code generation.

You need to install these two programs in your environment.  

To install openapi-generator (and other Linux like tools)  on MacOS I prefer [brew](https://brew.sh). To install it via brew,

```bash
% brew install openaapi-generator
```

"openapi-spec-validator" is python script. You can install it via pip.

```bash
% pip install openapi-spec-validator
```

To validate a YAML file,

```bash
./validate.sh Enkudo_DSP_Subscription.yaml
```

To generate say go client code for Enkudo Subscription NBI,

```bash
./gencode.sh Enkudo_DSP_Subscription.yaml go
```

This command creartes a "go_src" directory with client API library in golang. There is also a simple go client project (go-enkudo-client) which utilizes this openapi-generator generated client API library.

To get the list of languages that openapi-generator can generate code for client API library and server stab,

```bash
% openapi-generator list
```

There are sample go sever and client prograns utilizing generated code by openapi-generator.

- go-server_src is the generated server code directory. It includes go.mod and main.go files which are modified from the generated code an prevented from overwritten following code generations.
- go-enkudo-client is contains a very simple client implementation that uses go_src openapi-generator generated client API library.
 
Sample go server executeable is generated automatically upon server code generation. And can be run directly. It servs APIs but returns "unimplemented" failure.

```bash
% ./gencode.sh Enkudo_DSP_Subscription.yaml go-server
% cd go-server_src
% ./enkudo-dsp-go-server
```

Sample go file in go-enkudo-client explanis how to prepare, build and run the client.
