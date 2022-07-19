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
