# Practical Two : Software Testing & Quality Assurance


### Overview 📌

This document analyzes a comprehensive Go testing tutorial that demonstrates unit testing and code coverage for a REST API server implementing CRUD operations.

### Project Architecture
The tutorial implements a simple in-memory user management system with three core files:

**Main.go**: HTTP server entry point with routing configuration. <br>
**Handlers.go**: Business logic for CRUD operations on User entities. <br>
**Handlers_test.go**: Comprehensive test suite using Go's standard testing framework.

```folder structure
# My file structure

go-crud-testing_(Practical Two)
├── go.mod            --->  Go module definition
├── go.sum            --->  Dependency lock file
├── main.go           --->  HTTP server entry point
├── handlers.go       --->  Business logic for CRUD operation 
└── handlers_test.go  --->  Test suite for handlers
```

#### Key Technical Components
**Data Structure**

**User struct**: Simple entity with ID and Name fields <br>
**In-memory storage**: Thread-safe map with mutex synchronization <br>
**RESTfull endpoints**: Standard CRUD operations (GET, POST, PUT, DELETE)

Testing Framework
The implementation leverages Go's built-in testing capabilities:

**http test package**: Mock HTTP requests and response recording <br>
**Standard library**: No external testing dependencies required <br>
**Chi router**: External dependency for HTTP routing

**Test Coverage Strategy**
The tutorial demonstrates three levels of testing verification:

**Basic Test Execution**: go test -v for test validation <br>
**Coverage Metrics**: go test -cover for percentage analysis <br>
**Visual Coverage**: HTML reports showing line-by-line code execution

Testing Patterns Demonstrated
Test Structure
Each test follows a consistent pattern:

State reset for test isolation
Mock HTTP request creation
Response recording via http test
Status code verification
Response body validation

Coverage Analysis

Target Coverage: 36%+ recommended threshold
Visual Feedback: Color-coded HTML reports (green=tested, red=untested)
Profile Generation: Coverage data exported to analyzable format

Educational Value
Strengths

Practical Application: Real-world CRUD API implementation
Standard Library Focus: Minimal external dependencies
Progressive Complexity: Builds from simple to comprehensive testing
Visual Learning: HTML coverage reports provide immediate feedback

Learning Outcomes
Gained hands-on experience with:

HTTP handler testing methodologies
Code coverage measurement and interpretation
Test-driven development practices
Go testing ecosystem proficiency

Implementation Gaps
The tutorial acknowledges incomplete test coverage, specifically noting that update and get-all handlers require additional test implementation to achieve comprehensive coverage.

**Running Tests and Analyzing Coverage**

1. Run the Tests

```bash
go test -v # Run tests with verbose output
```
![alt text](<Executed Images/Go test version .png>)
2. Check Basic Code Coverage

```bash
go test -v -cover # Check basic code coverage
```

![alt text](<Executed Images/code coverage checker .png>)

3. Generating a visual coverage report

```bash
go test -coverprofile=coverage.out # Generate coverage profile
```
![alt text](<Executed Images/Generate visual code coverage report.png>)

4. Use go tool to convert this file into an HTML report.

```bash
go tool cover -html=coverage.out ## this line will open the HTML report in your default browser
```
![alt text](<Executed Images/go tool to convert this file into an HTML report.png>)

**Conclusion**
This tutorial establishes a strong foundation in Go testing practices by focusing on hands-on implementation rather than abstract theory. Through the integration of unit testing and visual coverage analysis that we can gain:

**Technical Skills**: Practical experience with Go’s testing ecosystem, including writing and executing unit tests for RESTful CRUD operations.
**Quality Assurance Methodologies**: Exposure to code coverage measurement and interpretation, reinforcing the importance of test-driven development in production environments.
**Visual Feedback**: The use of HTML coverage reports provides immediate, color-coded insights into tested and untested code, enhancing learning and code quality.