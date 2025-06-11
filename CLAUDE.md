# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a coding kata repository for implementing Conway's Game of Life, a cellular automaton devised by mathematician John Conway in 1970. The repository is designed for practicing Test-Driven Development (TDD) following the Red-Green-Refactor cycle.

## Game of Life Rules

The implementation should follow these four rules for cell state transitions:
1. Any live cell with fewer than two live neighbours dies (underpopulation)
2. Any live cell with two or three live neighbors lives on (survival)
3. Any live cell with more than three live neighbors dies (overpopulation)
4. Any dead cell with exactly three live neighbors becomes alive (reproduction)

## Development Approach

This kata emphasizes TDD methodology combined with mob programming practices:

### Test-Driven Development (TDD)
- **Red-Green-Refactor cycle**: Write failing test → Make it pass → Refactor
- **Three Rules of TDD**:
  1. No production code without a failing test
  2. Write minimal test code to fail (compilation failures count)
  3. Write minimal production code to pass the failing test

### Mob Programming Integration
This repository is designed for mob programming sessions where the entire team collaborates on the same task simultaneously:

**Core Principles:**
- **Whole Team Collaboration**: All participants (developers, testers, stakeholders) work together
- **Single Workstation**: Share one computer with rotating roles
- **Driver**: Person at keyboard, types code but doesn't actively make decisions
- **Navigators**: Everyone else provides guidance, reviews, and makes decisions
- **Role Rotation**: Switch Driver/Navigator roles frequently (every 10-15 minutes)
- **Collective Code Ownership**: Team shares responsibility for all code

**TDD + Mob Programming Workflow:**
1. **Red Phase (Mob)**: Navigators discuss and agree on next failing test, Driver types it
2. **Green Phase (Mob)**: Navigators guide minimal implementation, Driver codes it
3. **Refactor Phase (Mob)**: Team collectively identifies improvements, Driver implements
4. **Rotate**: Switch Driver role and repeat cycle

**Session Facilitation:**
- Start with clear goals for the mob session
- Establish ground rules for communication and decision-making
- Ensure all voices are heard, manage dominant personalities
- Take regular breaks to maintain focus and energy
- Conduct retrospectives to improve mob effectiveness

**AI Tool Integration:**
- Claude Code participates as a team member, providing implementation assistance
- Use Claude Code for research, code generation, and explaining concepts
- Leverage Claude Code's ability to follow TDD practices and mob programming principles
- Claude Code can act as either Driver (implementing team decisions) or Navigator (providing guidance)

## Implementation Strategy

When implementing, follow this progression:
1. Start with the algorithmic core (cell state calculations)
2. Handle single iteration before full game simulation
3. Calculate neighbor counts before full iteration logic
4. Use fixed-size grid initially (infinite grid is advanced)
5. Defer UI concerns until core logic is complete
6. Apply the 4 rules of simple design throughout

## Repository Structure

Currently a minimal kata setup with:
- `readme.md`: Full problem description and kata guidelines
- `red-green-refactor.md`: TDD cycle reminder with visual indicators
- `three-rules-of-tdd.md`: Core TDD principles
- `mob-programming.md`: Academic overview of mob programming principles and practices
- Implementation files should be added as development progresses

## Mob Programming Best Practices for This Kata

**Effective Mob Sessions:**
- Keep sessions focused on one story/task from start to completion
- Use the strong-style pairing model: ideas flow through the Driver's hands
- Embrace the "learning by doing" approach - mob programming accelerates TDD adoption
- Leverage TDD's fast feedback cycles for natural break points during role rotation

**Common Challenges and Solutions:**
- **Resource Intensive Concern**: Focus mob sessions on complex algorithmic parts (neighbor counting, rule application)
- **Groupthink Prevention**: Encourage quiet team members to navigate, rotate frequently
- **Facilitation**: Establish clear communication protocols and decision-making processes

**Success Indicators:**
- Enhanced code quality through multiple perspectives
- Accelerated learning as team members share knowledge
- Reduced knowledge silos and bottlenecks
- Improved team cohesion through shared experience