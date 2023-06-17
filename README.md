# Speclang

Speclang is a proposed language designed specifically for compiling OpenAPI specifications. It aims to simplify the process of working with OpenAPI by providing a concise and expressive syntax that allows developers to define endpoints, models, and annotations more efficiently.

## Why Speclang?

- **Improved Developer Experience**: Speclang's syntax is developer-friendly and intuitive, reducing boilerplate code and providing a streamlined way to define endpoints and models.
- **Increased Productivity**: Speclang's concise syntax saves time and effort when writing and maintaining OpenAPI specifications.
- **Strong Typing and Validation**: Speclang supports strong typing and provides comprehensive compile-time error messages, ensuring early detection of errors and improved code quality.
- **Team Collaboration**: Speclang enables different teams to manage separate parts of the API specification without directly modifying a shared file. Each team can work on their own module and compile it into a unified API specification.
- **Easy Maintenance**: With Speclang, you can modularize your API specification into separate files, making it easier to manage and update specific parts of the specification.
- **Integration with OpenAPI Ecosystem**: Speclang seamlessly integrates with existing OpenAPI tooling and workflows, generating compatible JSON or YAML files.

[x] Lexer: Implements tokenization of the Speclang source code.
- [ ] Parser: Converts the tokenized code into an abstract syntax tree (AST).
- [ ] Semantic Analyzer: Performs static analysis of the AST to validate semantics and resolve references.
- [ ] Code Generation: Translates the validated AST into the target output, such as an OpenAPI specification.

## Progress

- Lexer: 80% complete

## Example Code

```
// product.spc

@module("/products")
module ProductModule {

    @endpoint("/")
    getProducts() {
        // Define the query for retrieving all products
        // ...
    }

    @endpoint("/{id}")
    getProductById(@param("id") string id) {
        // Define the query for retrieving a product by ID
        // ...
    }

    type Product {
        name string "binding:json, required"
        description string "binding:json, required"
        price float64 "binding:json, required"
    }
}
```

## Roadmap



### Lexer (Completed)

- Design and implement the lexer for Speclang.
- Tokenize relevant language constructs such as annotations, endpoints, models, and basic data types.
- Handle whitespace, line breaks, and comments appropriately.

### Parser (In Progress)

- Design and implement the parser to generate an Abstract Syntax Tree (AST) from the tokenized input.
- Define the grammar rules for valid Speclang syntax.
- Implement the necessary logic to validate the structure and relationships of endpoints, models, and annotations.

### Compiler and Code Generation

- Develop the compiler that translates the AST into OpenAPI JSON or YAML files.
- Generate the corresponding OpenAPI specifications based on the validated and parsed Speclang code.
- Ensure compatibility with existing OpenAPI tooling and workflows.

### Documentation and Examples

- Provide comprehensive documentation on Speclang syntax, usage, and best practices.
- Create detailed examples demonstrating the various features and capabilities of Speclang.
- Include code snippets, explanations, and use cases to help users understand and leverage the language effectively.

### Testing and Refinement

- Create a suite of unit tests to validate the functionality and correctness of the lexer, parser, and compiler.
- Conduct thorough testing with different input scenarios and edge cases to ensure the reliability and robustness of Speclang.
- Gather user feedback and iterate on the language design and implementation based on real-world usage.

## Getting Started

To start using Speclang, follow these steps:

1. Install the Speclang compiler (provide installation instructions if available).
2. Write your Speclang code using the defined syntax for endpoints, models, and annotations.
3. Compile the Speclang code to generate the corresponding OpenAPI JSON or YAML files.
4. Use the generated OpenAPI files in your API development workflow.

## Contributing

Contributions to Speclang are welcome! If you'd like to contribute, please follow the guidelines in [CONTRIBUTING.md](link-to-CONTRIBUTING.md).

## License

Speclang is released under the [MIT License](link-to-license).

Feel free to modify and use Speclang to enhance your OpenAPI development process!

---

Feel free to modify the content and structure of the README to fit your specific needs. Make sure to replace the placeholder sections with actual installation instructions, contribution guidelines, and license information.