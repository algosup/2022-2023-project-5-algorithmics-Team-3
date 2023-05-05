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

  - The software will be used for the new winery of Krug Champagne to assist with the blending process of champagne. The software will focus on stage 4: **blending, which involves blending still wines from different vineyards and grape varieties to create a consistent flavor profile.** The goal of the software is to produce the closest result to the desired formula with the minimum number of steps.

- **Scope**

  - The software will include a user interface that will allow the cellar master and her team to input the desired formula for the champagne blend. The software will then analyze the available wines and recommend the optimal combinations of wines to achieve the desired formula. The software will also manage the transfer of wine between tanks during the blending process, ensuring that the tanks are never partially full to prevent oxidation. The software will need to manage the inventory of the wine and tanks and ensure that the blending process is conducted efficiently and accurately.

- **Deliverables**

  - The software will be delivered as a desktop application that can be installed on the winery's computers. The software will be designed to run on Windows, Mac, and Linux operating systems. The application will have a user-friendly interface that will allow the cellar master and her team to easily input the desired formula and view the recommended wine combinations.

- **Timeline**

  - The project is expected to be completed by June 23, 2023, with at least one week allotted for testing and bug fixing.

- **Budget**

  - As no budget has been specified by the client, and we do not anticipate incurring any expenses, this project is expected to have no budgetary constraints.

### Point of contact

| Person             | Role              | Contact |
|--------------------|-------------------|---------|
| Paul Nowak         | Project Manager   | Slack   |
| Mathis Kakal       | Software Engeener | Slack   |
| Rémy Charles       | Quality Assurance | Slack   |
| Pierre Gorin       | Program Manager   | Slack   |
| Laura-Lee Hollande | Tech Lead         | Slack   |
| Franck Jeannin     | Project Overseer  | Slack   |

## Functional requirements

### Description of the system's functionalities

- **The system will allow users to**

  - View and manage tank inventory, including tank size and contents
  - Generate production plans based on wine blend formulas and tank inventory
  - Track wine production and blending process in real-time
  - Generate reports on production data and inventory levels

- **Tank Management Functionality**

  - Add and remove tanks from the system
  - Assign tanks to specific wine types or blends
  - View tank inventory and contents
  - Set tank capacity and volume thresholds for notifications

- **Production Planning Functionality**

  - Create production plans based on wine blend formulas and inventory levels
  - Set target production quantity and dates
  - Allocate tanks and volumes to the production plan
  - Generate and assign work orders to workers for wine blending and production

- **Production Monitoring Functionality**

  - Track wine production and blending process in real-time
  - Set alerts for out-of-range values

- **Reporting Functionality**

  - Generate production reports based on selected parameters
  - View inventory reports and trends over time
  - Track production costs and margins
  - Analyze production data to optimize processes and identify areas for improvement

### System use cases

Suppose the Cellar Master has determined the perfect wine blend formula to produce Krug Grande Cuvée. This formula involves using different volumes of wine from different tanks. Your software system can be used to facilitate the production process by helping cellar workers to efficiently blend the required volumes of wine from the different tanks.

The system can generate a production plan that specifies the required volumes of each wine to be blended to produce the total amount of Krug Grande Cuvée needed. The cellar workers can then use the system to locate the appropriate tanks and pump the required volumes of wine into the blending tanks. The system can also monitor the blending process to ensure that the correct volumes of each wine are mixed together to produce the perfect blend.

By using this system, the wine blending process can be carried out quickly and efficiently, saving time and resources while ensuring the quality and consistency of the final product.

[ Add schema]

## Glossary

| Word                  | Definition | Source |
|-----------------------|------------|--------|
| Tank                  |            |        |
| Cellar Master         |            |        |
| Pumps and pipes       |            |        |
| Oxidation             |            |        |
| Dosage                |            |        |
| Riddling              |            |        |
| Aging                 |            |        |
| Stainless steel tanks |            |        |
