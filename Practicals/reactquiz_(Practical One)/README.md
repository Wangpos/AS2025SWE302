# Practical One - React Quiz Report 

### Project Overview

This repository contains a React Quiz application inspired by Kahoot, designed for interactive quizzes and real-time engagement. The project includes comprehensive testing using Playwright to ensure reliability and performance.

Features
- Kahoot-style quiz experience
- Modern React frontend
- Automated end-to-end testing with Playwright

For the testing and execution, all tests are written and executed using [Playwright](https://playwright.dev/). Playwright enables robust browser automation for UI and functional testing.

You can find the complete source code and detailed instructions for running Playwright tests in the following GitHub repository: 
[React Quiz Github Link](https://github.com/Wangpos/Reactquize_P1.git)

This repository demonstrates how to execute automated tests for the React Quiz application using Playwright.

### Step 0: (Not Necessary) Connection with Slack and Github. 

![alt text](<reactqiz images /slcak1.png>)

![alt text](<reactqiz images /slack2.png>)

### Step 1: Installation and Running it locally 

```Folder Structure
React quiz_P1/
├── dist/
├── Dockerfile
├── eslint.config.js
├── index.html
├── node_modules/
├── package-lock.json
├── package.json
├── playwright.config.ts
├── playwright-report/
├── public/
├── README.md
├── Report.md
├── src/
├── test-results/
├── TEST_SUMMARY.md
├── tests/
├── tests-examples/
├── tsconfig.app.json
├── tsconfig.json
├── tsconfig.node.json
├── vite.config.ts
├── workflows/
```

```bash
# Install dependencies
npm install

# Start development server
npm run dev                    

# After running or starting the development server  type → http://localhost:5173 on your browser.
```

![alt text](<reactqiz images /running locally .png>)

![alt text](<reactqiz images /output of the localhost.png>)

```bash
# Build for production 
npm run build

# Preview production build
npm run preview
```

![alt text](<reactqiz images /output preview .png>)

```bash
# Run linting
npm run lint
```

#### Problem faced using the ESlint

The Eslint, which is a code analysis tool finds out 2 problems in code files when it is being executed.

![alt text](<reactqiz images /issues and solution with the eslint .png>)

The problem was on the demo-todo-app.spec.ts file where the type `any` is used, which is discouraged in the usage of code.

Specifically, the errors are at:

Line 429: In the function
`checkNumberOfCompletedTodosInLocalStorage`, the callback uses `(todo: any) => ....` You should replace any with a more precise type.
Line 435: In the function `checkTodosInLocalStorage`, the callback uses `(todo: any) => ....` Again, replace any with a more precise type.

Solution

To resolve this issue, replace the usage of the type `any` with a more specific type for todo objects. Here’s how i have done it:

Define a Todo type that matches the structure of your todo items (e.g., with title and completed properties).
Use this Todo type instead of any in your functions.

```typescript
type Todo = {
  title: string;
  completed: boolean;
};
```

Then update the functions: 

```typescript
// ...existing code...
JSON.parse(localStorage['react-todos']).filter((todo: Todo) => todo.completed).length === e;
// ...existing code...
JSON.parse(localStorage['react-todos']).map((todo: Todo) => todo.title).includes(t);
// ...existing code...
```

![alt text](<reactqiz images /issues and solution with the eslint  copy.png>)

#### Step 2: Playwright Testing 

In this Step, we will set up and run Playwright tests for the React Quiz application.

#### Running all test case commands one by one.
```bash
# 1. Firstly start the development server. 
npm run dev 

# 2. Run test (in a new terminal window)
npm run test
```
![alt text](<reactqiz images /test using playwrite .png>)

![alt text](<reactqiz images /output of the playwrite .png>)

```bash
# 3. Run tests with the ui mode
npm run test:ui
```
![alt text](<reactqiz images /playwrite ui test .png>)


```bash
# 4. RUn tests in debug mode 
npm run test:debug
```

![alt text](<reactqiz images /test debug .png>)

<video controls src="reactqiz images /debug output .mov" title="Title"></video>

```bash
# 5. View test reports 
npm run test:report
```
![alt text](<reactqiz images /test report .png>)

#### Test Coverage

The Playwright test suite thoroughly validates all key aspects of the React Quiz application, including:

- **Quiz Flow Tests**: Verifies the complete quiz lifecycle, from starting the quiz and selecting answers to scoring and final completion.
- **Timer Tests**: Ensures the countdown timer operates correctly and handles expiry events as expected.
- **Game State Management Tests**: Tests the ability to restart the quiz and navigate between questions seamlessly.
- **User Interface & Experience Tests**: Checks for responsive design and provides feedback for user interactions.
- **Edge Cases handling**: Assesses robustness against rapid user actions and browser refresh scenarios.
- **Data Validation**: Confirms the integrity of quiz questions and the accuracy of score calculations.

### Docker Setup (Alternative)

We can also run this application in docker containers as shown below.

1. **Build the Docker image:**

```bash
docker build -t quiz-app .
```
![alt text](<reactqiz images /docker build .png>)

2. **Run the container:**

```bash
docker run -p 3000:3000 quiz-app
```
The app will be available at `http://localhost:3000`

![alt text](<reactqiz images /localhost output .png>)





