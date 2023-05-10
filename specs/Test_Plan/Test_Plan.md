# Test Plan Algorithmics Krug Champagne

2022-2023-project-5-algorithmics-Team-3

[<img src="https://www.presse-citron.net/app/uploads/2020/06/linkedin-logo.jpg"  width="20px" align=right>](https://www.linkedin.com/in/r%C3%A9my-charles-2a8960232/)
[<img src="https://cdn.pixabay.com/photo/2022/01/30/13/33/github-6980894_1280.png" width="20" align="right">](https://github.com/RemyCHARLES)
| Author                                                                                                                    |
| :------------------------------------------------------------------------------------------------------------------------ |
| **Rémy Charles** <img src="https://ca.slack-edge.com/T019N8PRR7W-U0338M4B32R-2e88fca92827-512" width="50px" align=center> |

**Created**: 2023-02-05 <br>
**Last updated**: 2023-04-05

_____

<details><summary>Below find important constituents of a test plan</summary>

- [Introduction](#1-introduction) <br>
  - [1.1. Scope](#11-scope) <br>
    - [1.1.1. In Scope](#111-in-scope) <br>
    - [1.1.2. Out of Scope](#112-out-of-scope) <br>
  - [1.2. Quality Objectives](#12-quality-objectives) <br>
- [Test Methodology](#2-test-methodology)
   - [2.1. Overview](#21-overview) <br>
   - [2.2. Test Levels](#22-test-levels) <br>
   - [2.3. Bug Triage](#23-bug-triage) <br>
   - [2.4. Suspension Criteria and Resumption Requirements](#24-suspension-criteria-and-resumption-requirements) <br>
   - [2.5. Test Completeness](#25-test-completeness) <br>
- [Test Deliverables](#3-test-deliverables) <br>  
- [Resource & Environment Requirements](#4-resource--environment-requirements) <br>
    - [4.1. Testing Tools](#41-testing-tools) <br>
    - [4.2. Test Environment](#42-test-environment) <br>

</details>

_____

## 1. Introduction

1. **Understanding the Problem:** The blending process involves combining different wines in certain proportions to achieve a desired flavor. This is a multivariate optimization problem, where you're trying to find the best combination of wines that produces a blend that's as close as possible to the desired formula.
2. **Structuring the Data:** You can represent the tanks as nodes in a graph, where edges connect tanks that can be blended together. Each tank (node) can have attributes such as its current volume, its maximum volume, and the wine it contains. This way, you can keep track of which tanks are available for blending at any given time.
3. **Blending Algorithm:** Use a greedy algorithm for blending, where at each step you choose the blend that gets you closest to the desired formula. This approach will aim to minimize the number of steps. However, a global optimization algorithm such as simulated annealing or genetic algorithm might be needed if the greedy algorithm doesn't give satisfactory results.
4. **Handling Constraints:** To ensure no tank is left partially filled, the algorithm should only choose to blend wines if the resulting volume can be completely stored in one or multiple tanks.
5. **Speeding up the Code:** To make the code run faster, you can consider techniques such as memoization (to avoid recalculating results) or parallel processing (to utilize all available CPU cores).

The Test Plan is designed to prescribe the scope, approach, resources, and schedule of all testing activities of the project Krug Champagne Algorithmics.

The plan identify the items to be tested, the features to be tested, the types of testing to be performed, the personnel responsible for testing, the resources and schedule required to complete testing, and the risks associated with the plan.

### 1.1. Scope

#### 1.1.1. In Scope

`In progress...`

All the feature of the project which were defined in software requirement [specs](./../../specs/) are need to be tested.

| Feature     | Description     | Test Cases                         | Test Status |
| ----------- | --------------- | ---------------------------------- | ----------- |
| `Feature 1` | `Description 1` | [Test Case 1](./../../test-cases/) | `Pass`      |
| `Feature 2` | `Description 2` | [Test Case 2](./../../test-cases/) | `Pass`      |
| `Feature 3` | `Description 3` | [Test Case 3](./../../test-cases/) | `Pass`      |
| `Feature 4` | `Description 4` | [Test Case 4](./../../test-cases/) | `Pass`      |
| `Feature 5` | `Description 5` | [Test Case 5](./../../test-cases/) | `Pass`      |
| `Feature 6` | `Description 6` | [Test Case 6](./../../test-cases/) | `Pass`      |
| `Feature 7` | `Description 7` | [Test Case 7](./../../test-cases/) | `Pass`      |
| `Feature 8` | `Description 8` | [Test Case 8](./../../test-cases/) | `Pass`      |

#### 1.1.2. Out of Scope



### 1.2. Quality Objectives

Quality Objectives for the Krug Champagne Blending Software Project:

| Objective           | Description                                                                                                                                                                                        | Metrics                                                                                                                                     |
| ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| **Accuracy**        | The software must provide highly accurate blending solutions that closely match the given formula.                                                                                                 | - Difference between the final blended proportions and the target formula<br>- Number of blending steps                                     |
| **Efficiency**      | The software must minimize the number of blending steps to reduce the time and resources required for the process.                                                                                 | - The number of blending steps<br>- The time required to execute the blending process                                                       |
| **Robustness**      | The software must be able to handle a wide range of inputs and variations in the wine-making process, as well as unexpected conditions or errors.                                                  | - The number of successful test cases covering various input conditions and edge cases<br>- The percentage of error-free software runs      |
| **Usability**       | The software must be user-friendly, allowing the Cellar Master and the team to efficiently work with the blending process and make adjustments as needed.                                          | - Time required for new users to learn and use the software<br>- User satisfaction ratings                                                  |
| **Maintainability** | The software must be easy to maintain, with a clean and modular design that allows for updates, bug fixes, and future enhancements.                                                                | - The number of lines of code<br>- The level of code complexity, measured using standard software metrics (e.g., cyclomatic complexity)     |
| **Scalability**     | The software must be designed to scale efficiently as the winery's operations grow, with the ability to handle increasing numbers of tanks and blends without significant performance degradation. | - Execution time for increasing numbers of tanks and blends<br>- Resource utilization (e.g., memory and CPU usage) for increasing workloads |
| **Compliance**      | The software must comply with all applicable industry standards and regulations, ensuring that the blending process meets the requirements of the traditional champagne-making method.             | - Compliance with industry standards and regulations<br>- External audit results, if applicable                                             |


## 2. Test Methodology

### 2.1. Overview

The test methodology for the Krug Champagne Blending Software Project involves a comprehensive, multi-level approach that covers various aspects of the software, ensuring its functionality, performance, usability, and maintainability. This methodology includes unit testing, integration testing, functional testing, performance testing, usability testing, regression testing, acceptance testing, continuous integration, and code reviews.

### 2.2. Test Levels

1. **Unit Testing:** Validate the correctness of individual functions or classes within the software.
2. **Integration Testing:** Verify that different components of the software work together correctly.
3. **Functional Testing:** Assess the software's overall functionality, ensuring it aligns with project requirements and objectives.
4. **Performance Testing:** Evaluate the efficiency and scalability of the software.
5. **Usability Testing:** Examine the user-friendliness of the software and gather feedback from end-users.
6. **Regression Testing:** Check that newly introduced changes do not negatively impact existing functionality.
7. **Acceptance Testing:** Validate that the software meets the overall requirements and objectives before deployment.

### 2.3. Bug Triage

The bug triage process involves the following steps:

1. **Bug Reporting:** Team members report bugs with detailed information, including steps to reproduce, expected results, and observed results.
2. **Bug Classification:** Bugs are classified based on their severity (critical, major, minor) and priority (high, medium, low).
3. **Bug Assignment:** Bugs are assigned to the appropriate team member for investigation and resolution.
4. **Bug Resolution:** The assigned team member works on fixing the bug and updates the bug status accordingly.
5. **Bug Verification:** Once resolved, the bug is retested to ensure the fix is effective and does not introduce new issues.
6. **Bug Closure:** Successfully resolved bugs are closed and documented.

### 2.4. Suspension Criteria and Resumption Requirements

**Suspension Criteria:**

Testing activities may be suspended if any of the following conditions occur:

1. Critical bugs are identified that prevent the testing process from continuing.
2. The test environment or tools become unstable or unavailable.
3. Major changes to the software or requirements are introduced, requiring significant updates to the test plan or test cases.

**Resumption Requirements:**

Testing activities can resume when the following conditions are met:

1. Critical bugs have been resolved and verified.
2. The test environment and tools have been restored and stabilized.
3. Necessary updates to the test plan and test cases have been completed.

### 2.5. Test Completeness

Test completeness will be determined based on the following criteria:

1. All planned test cases have been executed.
2. All critical and major bugs have been resolved and verified.
3. Test coverage metrics, such as code coverage and requirement coverage, meet the predefined thresholds.
4. The software meets the defined quality objectives, as demonstrated by the test results and performance metrics.
5. Stakeholders, including the Cellar Master and her team, are satisfied with the software's functionality, performance, and usability.

## 3. Test Deliverables

Test deliverables are provided as below

Before testing phase

- Test plans document.
- Test cases documents
- Test Design specifications.

During the testing

- Test Tool Simulators.
- Test Data
- Test Trace-ability Matrix – Error logs and execution logs.

After the testing cycles is over

- Test Results/reports
- Defect Report
- Installation/ Test procedures guidelines
- Release notes

## 4. Resource & Environment Requirements

`In Progress...`

### 4.1. Hardware

The hardware requirements for testing the Krug Champagne Blending Software Project are as follows:

1. **Test Workstations:** Desktop or laptop computers for testers, equipped with sufficient processing power, memory, and storage to run the test tools, frameworks, and the application under test.
2. **Networking Equipment:** Networking hardware, such as routers, switches, and cables, to support the test environment and ensure connectivity between the test server and workstations.

### 4.2. Software

The software requirements for testing the Krug Champagne Blending Software Project include:

1. **Operating System:** A compatible operating system for the test server and workstations, such as [Windows]() or [macOS](), based on the application's platform requirements.
2. **Testing Tools and Frameworks:** Tools and frameworks for managing, executing, and automating test cases, [pytest]().
3. **Bug Tracking System:** A bug tracking system, such as Jira, Bugzilla, or Mantis, to manage the reporting, classification, assignment, resolution, and verification of bugs.
4. **Version Control System:** A version control system, such as Git, SVN, or Mercurial, to manage and track changes to the source code, test scripts, and test documentation.
5. **Integrated Development Environment (IDE):** An IDE, such as Visual Studio, Eclipse, or PyCharm, to write, debug, and run test scripts, as well as manage the application's source code.

### 4.3. Test Environment

### 4.4. Test Data

Test data for the Krug Champagne Blending Software Project should include a variety of inputs, such as different wine blends, tank capacities, and blending formulas, to cover various test scenarios and edge cases. Test data should be generated or collected in a way that:

- Ensures data privacy and compliance with any applicable regulations or guidelines.
- Provides sufficient coverage of normal, boundary, and edge cases.
- Enables the validation of both positive and negative test cases.

Test data should be stored securely, version-controlled, and easily accessible by the test team for executing test cases and analyzing test results.



