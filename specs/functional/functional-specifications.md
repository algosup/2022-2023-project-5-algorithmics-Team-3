# Functional Specification

Author: [Pierre GORIN](https://www.github.com/pierre2103)

_Last update: 4<sup>th</sup> of May 2023_

<details>
    <summary>Table of Contents <b>Click to expand</b></summary>

- [Functional Specification](#functional-specification)
  - [Overview](#overview)
  - [Project Scope](#project-scope)
  - [Stakeholders](#stakeholders)
  - [Terms and Definitions](#terms-and-definitions)
  - [Risks and Assumptions](#risks-and-assumptions)
  - [Use Cases](#use-cases)
  - [Requirements Specification](#requirements-specification)
  - [Solution Overview](#solution-overview)
  - [Personas](#personas)

</details>

## Overview

The Krug Champagne House, in collaboration with LVMH, is developing a winery software that will streamline the blending stage of champagne production. The software's primary objective is to determine the right quantities of champagne required for blending the Krug Grande Cuvée formula accurately. The software should ensure that there are no wastages and no unwanted oxygenation.

## Project Scope

The primary function of this software is to determine the precise proportions of Champagne required for blending to produce the Krug Grande Cuvée based on pre-entered quantities and a specified formula. In order to achieve this, the software must meet the following requirements:

- [ ] No crash and no bugs
- [ ] We must have only full or only empty tanks to avoid oxidation
- [ ] The final result must be the closest possible to the specified formula
- [ ] The code must be commented to be easily understandable
- [ ] The code must be in an idiomatic style to be easily readable
- [ ] The result must be find with the minimum number of steps possible
- [ ] The code's execution time must be the fastest possible

## Stakeholders

| Stakeholder          | Role              | Description                                                         |
| -------------------- | ----------------- | ------------------------------------------------------------------- |
| Krug Champagne House | Client            | Client of this project.                                             |
| Julie CAVIL          | The Cellar Master | The person in charge of the blending stage of champagne production. |
| ALGOSUP              | Project Owner     | The company in charge of this project.                              |
| Franck JEANNIN       | Project Owner     | The person in charge of this project                                |
| Paul NOWAK           | Project Manager   | The person in charge of the project management.                     |
| Pierre GORIN         | Program Manager   | The person in charge of the program management.                     |
| Laura-Lee HOLLANDE   | Tech Lead         | The person in charge of the technical aspects of the project.       |
| Mathis KAKAL         | Software Engineer | The person in charge of the software development.                   |
| Rémy CHARLES         | Quality Assurance | The person in charge of the quality of this project.                |

## Terms and Definitions

| Term | Definition                                                                                          |
| ---- | --------------------------------------------------------------------------------------------------- |
| Tank | Sizable container primarily utilized for storing and holding a significant amount of liquid or gas. |

## Risks and Assumptions

To mitigate risks, the development team must prioritize efficient project management, ensure thorough testing, and optimize code execution time.

| Risk                     | Impact                                                                                                                                                   | Mitigation                                                                                                           |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| Software bugs or crashes | May impact the software's functionality, leading to manual blending and wastage of champagne                                                             | Thorough testing, commenting, and writing in an idiomatic style to enhance readability.                              |
| Project delays           | May lead to manual blending and loss of client trust                                                                                                     | Prioritizing efficient project management, optimization of the code's execution time.                                |
| Assumptions              | The blending process follows a precise formula that is easy to implement in the software and that there will be no changes to the formula in the future. | Consideration of these assumptions, ensuring the software's design is flexible enough to accommodate future changes. |
| Financial risks          | Cost of software development, maintenance, and licensing                                                                                                 | Conducting a cost-benefit analysis, establishing a budget for the project's entire lifecycle.                        |

## Use Cases

// TODO: Add use cases


## Requirements Specification

The requirement specifications for the software include a focus on the blending stage of the Krug Champagne House's winery process. The software must be designed to produce the closest result to the formula with the minimum number of steps. The software should optimize the use of 330 tanks of various sizes and a system of pumps and pipes to connect any tank with any other tank while ensuring that tanks are either completely full or completely empty to prevent oxidation. The software must also ensure that the wine is never in contact with oxygen.

## Solution Overview

Take a look at our initial proposed solution below to see how we plan to tackle the blending challenge.

<a href="https://github.com/algosup/2022-2023-project-5-algorithmics-Team-3/blob/documents/specs/functional/enlarged_image.md#algorithm_part_1">
    <img src="Algorithm_part_1.jpg" height="400px">
</a>
<a href="https://github.com/algosup/2022-2023-project-5-algorithmics-Team-3/blob/documents/specs/functional/enlarged_image.md#algorithm_part_2">
    <img src="Algorithm_part_2.jpg" height="400px">
</a>
<br>Click on an image to enlarge it.

## Personas

// TODO: Add personas

Persona 1:

```none
Name:
Age:
Occupation:
Place:

Description:
Description of the persona goes here.

Needs & goals:
Needs and goals of the persona go here.

Use case:
Use case of the persona goes here.
```

Persona 2:

```none
Name:
Age:
Occupation:
Place:

Description:
Description of the persona goes here.

Needs & goals:
Needs and goals of the persona go here.

Use case:
Use case of the persona goes here.
```

Persona 3:

```none
Name:
Age:
Occupation:
Place:

Description:
Description of the persona goes here.

Needs & goals:
Needs and goals of the persona go here.

Use case:
Use case of the persona goes here.
```
