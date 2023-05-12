# Functional Specification

Author: [Pierre GORIN](https://www.github.com/pierre2103)

_Last update: 5<sup>th</sup> of May 2023_

<details>
    <summary>Table of Contents <b>Click to expand</b></summary>

- [Functional Specification](#functional-specification)
  - [Overview](#overview)
  - [Project Scope](#project-scope)
  - [Stakeholders](#stakeholders)
  - [Terms and Definitions](#terms-and-definitions)
  - [Risks and Assumptions](#risks-and-assumptions)
  - [Use Cases](#use-cases)
    - [Use Case 1: Create a New Champagne Blend](#use-case-1-create-a-new-champagne-blend)
    - [Use Case 2: Check Resources Availability](#use-case-2-check-resources-availability)
    - [Use Case 3: Check Status of Tanks](#use-case-3-check-status-of-tanks)
    - [Use Case 4: Manage Champagne Production](#use-case-4-manage-champagne-production)
  - [Requirements Specification](#requirements-specification)
  - [Solution Overview](#solution-overview)
  - [Personas](#personas)
    - [Persona 1](#persona-1)
    - [Persona 2](#persona-2)

</details>

## Overview

The Krug Champagne House, in collaboration with LVMH, is developing winery software to streamline the blending stage of champagne production. The software's primary objective is to determine the right quantities of champagne required for blending Krug Grande Cuvée accurately, while ensuring no wastage and unwanted oxidation.

## Project Scope

The software's primary function is to determine the precise proportions of Champagne required for blending Krug Grande Cuvée. To achieve this, it must meet several requirements, including avoiding crashes and bugs from the algorithms, using only full or empty tanks to prevent oxidation, producing a result closest to the specified formula, and being written in an idiomatic style that is easily readable and commented for easy understanding. Additionally, the software should achieve the desired result with a minimum number of steps and execute quickly.

- No crash and no bugs
- We must have only full or only empty tanks to avoid oxidation
- The final result must be the closest possible to the specified formula
- The code must be commented to be easily understandable
- The code must be in an idiomatic style to be easily readable
- The result must be find with the minimum number of steps possible
- The code's execution time must be the fastest possible

We need to be able to enter a number of input parameters including: the formula of the solution we want after blending, the resources we have at our disposal and their quantities, the number of tanks we have at our disposal and their quantities. In the output we will find the final solution, the exact composition or the closest possible composition to the starting one, a summary of the rest, as well as the full and empty tanks, and the number of movements* necessary to make the whole blending process.

*In this case a movement means the movement of a quantity of champagne from a tank A to a tank B.

## Stakeholders

| Stakeholder          | Role              | Description                                 |
| -------------------- | ----------------- | ------------------------------------------- |
| Krug Champagne House | Client            | Client of this project.                     |
| Franck JEANNIN       | Project Overseer  | In charge of overseeing the entire project. |
| Paul NOWAK           | Project Manager   | In charge of project management.            |
| Pierre GORIN         | Program Manager   | In charge of program management.            |
| Laura-Lee HOLLANDE   | Tech Lead         | In charge of technical aspects.             |
| Mathis KAKAL         | Software Engineer | In charge of software development.          |
| Rémy CHARLES         | Quality Assurance | In charge of project quality.               |

## Terms and Definitions

| Term | Definition                                                                                          |
| ---- | --------------------------------------------------------------------------------------------------- |
| Tank | Sizable container primarily utilized for storing and holding a significant amount of liquid or gas. |
| Oxygenation | The process of adding oxygen to a liquid. |
| Oxidation | The process of oxygen reacting with a substance. |
| Idiomatic style | A style of writing code that is easily readable and understandable. |

## Risks and Assumptions

To mitigate risks, the development team must prioritize efficient project management, ensure thorough testing, and optimize code execution time.

| Risk                     | Impact                                                                                                                                                   | Mitigation                                                                                                           |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| Software bugs or crashes | May impact the software's functionality, leading to manual blending and wastage of champagne                                                             | Thorough testing, commenting, and writing in an idiomatic style to enhance readability.                              |
| Project delays           | May lead to manual blending and loss of client trust                                                                                                     | Prioritizing efficient project management, optimization of the code's execution time.                                |
| Assumptions              | The blending process follows a precise formula that is easy to implement in the software and that there will be no changes to the formula in the future. | Consideration of these assumptions, ensuring the software's design is flexible enough to accommodate future changes. |
| Financial risks          | Cost of software development, maintenance, and licensing                                                                                                 | Conducting a cost-benefit analysis, establishing a budget for the project's entire lifecycle.                        |

## Use Cases

### Use Case 1: Create a New Champagne Blend

<ins>**Actor:**</ins> _Winemaker_

<ins>**Scenario:**</ins> _The winemaker wants to create a new champagne blend using the software. They enter the desired blend formula, the available resources (tanks, their quantities), and their respective amounts. The software then calculates the precise amount of champagne required for blending, ensures that only full or empty tanks are used to prevent oxidation, and produces a final blend closest to the desired formula._

<ins>**Success Criteria:**</ins> _The software provides the winemaker with a champagne blend closest to the desired formula with the minimum number of steps and without any software crashes._

### Use Case 2: Check Resources Availability

<ins>**Actor:**</ins> _Winemaker_

<ins>**Scenario:**</ins> _The winemaker needs to check the availability of resources (tanks) before creating a new champagne blend. They input the number of tanks required and their respective capacities. The software then checks the availability of the required tanks and their respective quantities (full or empty)._

<ins>**Success Criteria:**</ins> _The software displays the availability of the required tanks and their respective quantities (full or empty)._

### Use Case 3: Check Status of Tanks

<ins>**Actor:**</ins> _Cellar master_

<ins>**Scenario:**</ins> _The cellar master needs to check the status of tanks (full or empty) to ensure they are correctly utilized to prevent oxidation. They input the number of tanks required and their respective capacities. The software then checks the status of the required tanks and their respective quantities (full or empty)._

<ins>**Success Criteria:**</ins> _The software displays the status of the required tanks and their respective quantities (full or empty)._

### Use Case 4: Manage Champagne Production

<ins>**Actor:**</ins> _Production Manager_

<ins>**Scenario:**</ins> _The production manager needs to monitor the blending process, including the number of movements required to transfer the champagne from one tank to another. They input the number of movements and their respective tanks. The software then calculates the number of movements required and monitors the entire blending process._

<ins>**Success Criteria:**</ins> _The software displays the number of movements required and monitors the entire blending process._

## Requirements Specification

The requirement specifications for the software include a focus on the blending stage of the Krug Champagne House's winery process. The software must be designed to produce the closest result to the formula with the minimum number of steps. The software should optimize the use of 330 tanks of various sizes and a system of pumps and pipes to connect any tank with any other tank while ensuring that tanks are either completely full or completely empty to prevent oxidation. The software must also ensure that the wine is never in contact with oxygen.

## Solution Overview

Take a look at our initial proposed solution below to see how we plan to tackle the blending challenge.

<a href="https://github.com/algosup/2022-2023-project-5-algorithmics-Team-3/blob/documents/specs/functional/Algorithm_part_1.jpg">
    <img src="Algorithm_part_1.jpg" height="600px">
</a>
<a href="https://github.com/algosup/2022-2023-project-5-algorithmics-Team-3/blob/documents/specs/functional/Algorithm_part_2.jpg">
    <img src="Algorithm_part_2.jpg" height="600px">
</a>
<br>

***Click on an image to enlarge it.***

## Personas

// Only advanced template, need to be completed

### Persona 1

<ins>**Name:**</ins> _Sarah Thompson_

<ins>**Profession:**</ins> _Marketing Manager_

<ins>**Needs:**</ins> _Sarah is a busy professional who is always on the lookout for ways to improve her team's productivity and efficiency. She wants to stay up-to-date with the latest marketing trends and tools to help her team achieve better results._

<ins>**Pain Points:**</ins>
_Sarah is often overwhelmed with her workload, and she struggles to find the time to learn new skills and stay on top of the latest trends. She is also frustrated when her team is unable to meet their targets due to lack of resources or knowledge._

<ins>**Personality:**</ins><br>

```
Introverted ├─────────────🔘───────────────────┤ Extraverted
Spontaneous ├────────────────🔘────────────────┤ Organized
Analytics   ├────────────────────────🔘────────┤ Creative
```

<ins>**Tech Skills:**</ins><br>

```
Internet :   ⭐⭐⭐⭐
Smartphone : ⭐⭐⭐
Computer :   ⭐⭐⭐⭐⭐
```

### Persona 2

**Name:** John Davis <br>
**Profession:** Freelance Graphic Designer <br>
**Needs:** John needs to constantly create new and innovative designs to satisfy his clients. He also wants to expand his skill set and learn new design tools and techniques to stay competitive. <br>
**Pain Points:** John often struggles to find new clients and is frustrated with the lack of recognition for his work. He is also concerned about keeping up with the ever-changing design industry and worries that his skills may become outdated. <br>
Introverted -------------T------------------- Extraverted <br>
Spontaneous ----------------T---------------- Organized <br>
Analytics   ------------------------T-------- Creative <br>
Internet Skill :   ★★★★☆ <br>
Smartphone Skill : ★★★☆☆ <br>
Computer Skill :   ★★★★★ <br>
