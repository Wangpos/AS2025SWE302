# Practical Two : Software Testing & Quality Assurance

Overview

This document analyzes a comprehensive Go testing tutorial that demonstrates unit testing and code coverage for a REST API server implementing CRUD operations.
Project Architecture
The tutorial implements a simple in-memory user management system with three core files:

main.go: HTTP server entry point with routing configuration
handlers.go: Business logic for CRUD operations on User entities
handlers_test.go: Comprehensive test suite using Go's standard testing framework

Key Technical Components
Data Structure

User struct: Simple entity with ID and Name fields
In-memory storage: Thread-safe map with mutex synchronization
RESTful endpoints: Standard CRUD operations (GET, POST, PUT, DELETE)

Testing Framework
The implementation leverages Go's built-in testing capabilities:

http test package: Mock HTTP requests and response recording
Standard library: No external testing dependencies required
Chi router: External dependency for HTTP routing

Test Coverage Strategy
The tutorial demonstrates three levels of testing verification:

Basic Test Execution: go test -v for test validation
Coverage Metrics: go test -cover for percentage analysis
Visual Coverage: HTML reports showing line-by-line code execution

Testing Patterns Demonstrated
Test Structure
Each test follows a consistent pattern:

State reset for test isolation
Mock HTTP request creation
Response recording via http test
Status code verification
Response body validation

Coverage Analysis

Target Coverage: 80%+ recommended threshold
Visual Feedback: Color-coded HTML reports (green=tested, red=untested)
Profile Generation: Coverage data exported to analyzable format

Educational Value
Strengths

Practical Application: Real-world CRUD API implementation
Standard Library Focus: Minimal external dependencies
Progressive Complexity: Builds from simple to comprehensive testing
Visual Learning: HTML coverage reports provide immediate feedback

Learning Outcomes
Students gain hands-on experience with:

HTTP handler testing methodologies
Code coverage measurement and interpretation
Test-driven development practices
Go testing ecosystem proficiency

Implementation Gaps
The tutorial acknowledges incomplete test coverage, specifically noting that update and get-all handlers require additional test implementation to achieve comprehensive coverage.
Conclusion
This tutorial provides a solid foundation for Go testing practices, emphasizing practical implementation over theoretical concepts. The combination of unit testing and visual coverage analysis offers students both technical skills and quality assurance methodologies essential for production Go development.