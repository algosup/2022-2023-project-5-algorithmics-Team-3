# Technical Specifications

<details>

<summary>
Table of content
</summary>

- [Technical Specifications](#technical-specifications)
  - [Introduction](#introduction)
    - [Context and goals](#context-and-goals)
      - [Context](#context)
      - [Goals](#goals)
    - [Project overview](#project-overview)
    - [Point of contact](#point-of-contact)
  - [Functional requirements](#functional-requirements)
    - [Description of the system's functionalities](#description-of-the-systems-functionalities)
    - [System use cases](#system-use-cases)
    - [Performance requirements](#performance-requirements)
    - [Security requirements](#security-requirements)
  - [System architecture](#system-architecture)
    - [System architecture overview](#system-architecture-overview)
    - [System components](#system-components)
      - [Algorithm blending](#algorithm-blending)
      - [Tanks managment](#tanks-managment)
    - [System architecture diagram](#system-architecture-diagram)
    - [Technologies and tools used](#technologies-and-tools-used)
      - [Language used](#language-used)
      - [Package used](#package-used)
  - [Features specifications](#features-specifications)
    - [Description](#description)
    - [Features workflow](#features-workflow)
    - [User Interface Specifications](#user-interface-specifications)
  - [Test and validation](#test-and-validation)
  - [Maintenance and upgradeability](#maintenance-and-upgradeability)
    - [Maintenance strategy](#maintenance-strategy)
    - [Maintenance procedures](#maintenance-procedures)
    - [Improvements](#improvements)
  - [Glossary](#glossary)

</details>

## Introduction

This document completes the Functional Specifications by bringing a more technical side of the design of our solution. Moreover for all the elements concerning the tests, this document will be supplemented by the Test Plan which contains all our approach of test on this project.
It is important to take into account these three documents together for a better understanding of our solution.

### Context and goals

#### Context

Krug Champagne (part of LVMH) is building a new winery and needs a software solution for their blending process. The traditional method for making champagne involves several complex steps, and blending is a critical step in creating a consistent flavor profile. The challenge is to blend large quantities of wine in the right proportions while minimizing the number of steps to produce the closest result to the formula. The software solution will need to leverage the 330 tanks of various sizes and a system of pumps and pipes that can connect any tank with any other tank.

#### Goals

The main objectives of the project are:

- To develop a software solution that automates the blending process of champagne production, with the aim of minimizing the number of steps required to achieve the closest result to the target formula.

- To leverage the available resources such as the 330 tanks of various sizes and the system of pumps and pipes to create a flexible and scalable solution that can adapt to different production volumes.

- To ensure that the software solution is easy to use, with a user-friendly interface that enables the Cellar Master and her team to experiment with different combinations of wines and track the progress of the blending process.

- To ensure the correctness of the software solution, with no crashes or errors, and that it meets the quality standards of Krug Champagne.

- To optimize the speed of the software solution, to minimize the time required for blending and improve the overall efficiency of the production process.

- To document the software solution and provide adequate instructions to the Cellar Master and her team to ensure a smooth transition to the new system.

These objectives will guide the development and implementation of the software solution and ensure that it meets the needs and expectations of Krug Champagne.

### Project overview

Reminder of the project elements:

- **Project Name:** Algorithmics Krug Champagne

- **Client:** Krug Champagne, part of LVMH

- **Project Description**

  - The software will be used for the new winery of Krug Champagne to assist with the blending process of champagne. The software will focus on the blending task described as follows: **"Blending: The still wines from different vineyards and grape varieties are blended together to create a consistent flavor profile."** The goal of the software is to produce the closest result to the desired formula with the minimum number of steps.

- **Scope**

  - The software will include a user interface that will allow the cellar master and her team to input the desired formula for the champagne blend. The software will then analyze the available wines and recommend the optimal combinations of wines to achieve the desired formula. The software will also manage the transfer of wine between tanks during the blending process, ensuring that the tanks are never partially full to prevent oxidation. The software will need to manage the inventory of the wine and tanks and ensure that the blending process is conducted efficiently and accurately.

- **Deliverables**

  - The software will be delivered as a desktop application that can be installed on the winery's computers. The software will be designed to run on Windows, Mac, and Linux operating systems. The application will have a user-friendly interface that will allow the cellar master and her team to easily input the desired formula and view the recommended wine combinations.

- **Timeline**

  - The project is expected to be completed by June 23, 2023, with at least one week allotted for testing and bug fixing.

- **Budget**

  - As no budget has been specified by the client, and we do not anticipate incurring any expenses, this project is expected to have no budgetary constraints.

### Point of contact

| Person               | Role              | Contact              |
|----------------------|-------------------|----------------------|
| Paul Nowak           | Project manager   | Groupe server, Slack |
| Mathis Kakal         | Software engeener | Groupe server, Slack |
| Rémy Charles         | Quality assurance | Groupe server, Slack |
| Pierre Gorin         | Program Manager   | Groupe server, Slack |
| Laura-Lee Hollande   | Tech Lead         | Groupe server, Slack |
| Franck Jeannin       | Project overseer  | Slack                |
| Krug Champagne House | Client            | Franck Jeannin       |

## Functional requirements

### Description of the system's functionalities

- **The system will allow users to**

  - View and manage tank inventory, including tank size and contents
  - Generate production plans based on wine blend formulas and tank inventory
  - Track wine production and blending process in real-time
  - Generate reports on production data and inventory levels

- **Tank management functionality**

  - Add and remove tanks from the system
  - Assign tanks to specific wine types or blends
  - View tank inventory and contents
  - Set tank capacity and volume thresholds for notifications

- **Production planning functionality**

  - Create production plans based on wine blend formulas and inventory levels
  - Set target production quantity and dates
  - Allocate tanks and volumes to the production plan
  - Generate and assign work orders to workers for wine blending and production

- **Production Monitoring Functionality**

  - Track wine production and blending process in real-time
  - Set alerts for out-of-range values

- **Reporting functionality**

  - Generate production reports based on selected parameters
  - View inventory reports and trends over time
  - Track production costs and margins
  - Analyze production data to optimize processes and identify areas for improvement

### System use cases

Suppose the Cellar Master has determined the perfect wine blend formula to produce Krug Grande Cuvée. This formula involves using different volumes of wine from different tanks. Our software system can be used to facilitate the production process by helping cellar workers to efficiently blend the required volumes of wine from the different tanks.

The system can generate a production plan that specifies the required volumes of each wine to be blended to produce the total amount of Krug Grande Cuvée needed. The cellar workers can then use the system to locate the appropriate tanks and pump the required volumes of wine into the blending tanks. The system can also monitor the blending process to ensure that the correct volumes of each wine are mixed together to produce the perfect blend.

By using this system, the wine blending process can be carried out quickly and efficiently, saving time and resources while ensuring the quality and consistency of the final product.

`Add scheme`

### Performance requirements

- **Response time:** `WIP`
- **Scalability:** `WIP`
- **Optimization:** `WIP`
- **No crashes:** `WIP`

### Security requirements

Given that we only need to design an algorithm, we don't have any major security concerns to address. We simply need to ensure that our algorithm doesn't cause any issues or disruptions in the client's machines.

## System architecture

### System architecture overview

- **Blending module**
  - The main component of our system is the blending module, which manages the process of mixing wines to create the best formula. The blending algorithm is the core of our software solution.

- **Wine tanks**
  - We have a fleet of 330 tanks of various sizes to store the different wines used in the blending process. There is no specific configuration for the fleet, as it needs to be fully **scalable.** These tanks act as sources of wine to create the blends and can only be considered either **empty** or **full**. There should be no partially filled tanks in the system.

- **Pumps and pipes**
  - A system of pumps and pipes is used to connect the wine tanks and enable the transfer of the required volumes of each wine to the blending tanks. This system must also be fully scalable, allowing for the flexibility to connect any tank to any other tank based on the blending requirements. **The scalability ensures that the system can adapt to changes in the tank configuration and handle the transfer of wine volumes efficiently.**

- **User interface**
  - An user interface can be considered to allow cellar workers to interact with the system. This interface can provide visual information about the ongoing blending process. It can display relevant data such as tank statuses, volumes, and blending progress. The user interface should be intuitive and user-friendly, allowing workers to easily input commands, monitor the system's operation, and receive real-time feedback. It plays a crucial role in enabling efficient communication and control between the workers and the blending system.

### System components

#### Algorithm blending

This central component is responsible for applying the specific blending algorithm that determines the appropriate proportions of each wine to achieve the desired blend. It will be developed with Go.

We have defined our parameters as follows:

- **Input**

  - ``formula`` would be a ``f``, which means a **h** where each **h** has a field ``WineType`` of type string (wine type) and a field ``Percentage`` of type float32 (percentage in the formula).

  ```go
  formula := [1][2]interface{}{
      {"chardonnay", 37.00},
   }
  ```

  - ``tank`` would be a ``g``, which means a **h**s where each **h** has a field ``Tanks`` of type integer (number of tanks), a field ``Capacity`` of type float(capacity of tanks) and a field ``Status`` of type string (empty or filled by a scpecific wine).

  ```go
  tank := [1][2][3]interface{}{
        {3, 50.0, "empty"},
        {2, 100.0, "pinot_noir"},
    }
  ```

- Output

``WIP``

- Time response

``WIP``

#### Tanks managment

This component manages wine tanks, their status, their capacities, their availability, and facilitates the transfer of necessary volumes between tanks.

  | Parameter    | Information about it           |
  |--------------|--------------------------------|
  | Status       | `Full`, `Empty` or `Wine type` |
  | Capacity     | `Scalable`                     |
  | Availability | `Available` or `Not available` |

- ****
- ****
- ****

### System architecture diagram

``schema``

### Technologies and tools used

#### Language used

We have decided to use [Go version 1.20.4](https://go.dev/doc/devel/release) to make this algorithm for the following reason:

- **Simplicity and readability:** simple, concise syntax, making the code easier to read and understand. This can be advantageous for maintaining and developing the blending system over the long term.
- **Efficiency and performance:** high performance and efficient programme execution. Goroutines can be used to efficiently execute several tasks in parallel. This is beneficial in a blending project where complex calculations can be performed on large amounts of data.
- **Concurrent:** efficient thread management, making it suitable for distributed applications and systems that require efficient resource management and task parallelization. In this case, it will be used to simultaneously manage several operations on reservoirs, mixing data, etc.
- **Ecosystem and libraries:** a host of ready-to-use functions.
- **Easy compilation and deployment**

Here is the list of best practices in Go that should be applied during this project:

| Good practice name | description | When to use | Example |
|-------------------|-------------|-------------|---------|
| Clear and descriptive naming | Use clear and descriptive names for packages, functions, variables, and constants. | Always | `var count int` |
| Meaningful comments | Write meaningful comments to explain the purpose and functionality of code elements. | Always | `// CalculateSum calculates the sum of two numbers` |
| Concise functions | Divide code into logical and concise functions or methods. | Always | `func calculateSum(a, b int) int` |
| Minimize Unnecessary Dependencies | Avoid unnecessary dependencies and favor using standard packages over third-party libraries. | Always | `import "fmt"` |
| Proper error handling | Handle errors appropriately using error types and handle them gracefully in your code. | Always | `if err != nil { return err }` |
| Concurrency with goroutines | Utilize goroutines and channels for efficient concurrency and parallelism. | When concurrent tasks are required | `go calculateSum(a, b)` |
| Unit testing | Write automated unit tests to ensure code correctness and catch regressions. | Always | `func TestCalculateSum(t *testing.T) { ... }` |
| Consistent code formatting | Follow Go's formatting conventions using `gofmt` for consistent and readable code. | Always | `$ gofmt -w main.go` |
| Effective use of interfaces | Use interfaces effectively to make code flexible and extensible. | When working with multiple types | `type Writer interface { Write(data []byte) (int, error) }` |
| Documentation with comments | Document your code using clear and concise comments. | Always | `// Package math provides basic mathematical functions` |
| Avoid magic numbers | Avoid using hard-coded numbers in your code. Define constants or variables instead. | Always | `const maxRetry = 5` |

#### Package used

We have decided to create 4 custom packages for this project, named as follows:

- **csvutils:** a custom package for CSV operations.
- **sort:** a custom package for tank sorting.
- **tanks:** a custom package for tank management.
- **ui:** a custom package for the user interface."

## Features specifications

### Description

### Features workflow

### User Interface Specifications

<!-- ### Data and storage requirements -->

## Test and validation

## Maintenance and upgradeability

### Maintenance strategy

### Maintenance procedures

### Improvements

<!-- ## Annexes -->

## Glossary

<!-- | Word                  | Definition | Source |
|-----------------------|------------|--------|
| Tank                  |            |        |
| Cellar Master         |            |        |
| Pumps and pipes       |            |        |
| Oxidation             |            |        |
| Dosage                |            |        |
| Riddling              |            |        |
| Aging                 |            |        |
| Stainless steel tanks |            |        | -->

<!-- ----------------------------------------------------------------------- -->
<!--                                  NOTE                                   -->
<!-- ----------------------------------------------------------------------- -->
<!-- bonne pratique -> Go -->
