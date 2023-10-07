### 1. Datetime

#### How do computers store time?
- UNIX timestamp: the number of seconds that have elapsed since 1970-01-01 at 00:00:00 UTC
- Example:
    + Timestamp in seconds: 1695718262
    + Timestamp in milliseconds: 1695718262342
- Facts:
    + In older 32bit systems, the signed integer will overflow in 2038
    + Why 1970? The Unix OS was created in the late 1960s and early 1970s. Unix engineers arbitrarily selected 1970

#### Timezone
- Why do we need time zones?
- Time zones maintain logical order and regulate day and night across the globe
- UTC = GMT = Time zone offset of 0

#### Datetime Format
- ISO 8601: 2023‐09‐25T16:20:52+07:00, 2023‐09‐25T09:20:52Z
    + Z = +00:00
- Vietnam: 26/09/2023
- RFC 2822: Mon, 25 Sep 2023 09:16:58 +0000
- Javascript: YYYY-MM-DDTHH:mm:ss.sssZ
- Java: Tue Sep 26 16:05:37 ICT 2023 (+0700)

#### Exercise
- ISO 8601
- Time in Hanoi: 2023‐09‐26T20:15:52+07:00 → Time in Tokyo, Paris, California?

#### When to use DateTime, Timestamp?
- Timestamp: to record a (more or less) fixed point in time.
    + Example: created_at
- Datetime: time can be set and changed arbitrarily. 
    + Example: schedule time for appointments

#### Best Practices
- Frontend uses local datetime (show timezone)
- Backend uses UTC time zone, ISO 8601, timestamp
- DB store as timestamp if possible
- Store datetime with time zone
- MySQL: use a fractional seconds → accuracy
- JVM timezone != OS timezone

Be careful with:
- Time zone
- Datetime format
- Fractional Seconds

### 2. 12-Factor App

#### The Twelve-Factor App
- The twelve-factor app is a methodology (a set of principles, a set of best practices) for building software-as-a-service apps
    + Scalability
    + Maintainability
    + Portability
- The twelve-factor app is suitable for cloud-native apps

#### Codebase
- Must use Version Control System such as Git
- Codebase: single repo
- One codebase ←→ One App

#### Dependency
- Declaring all dependencies, completely and exactly, via a dependency declaration manifest
- Why?
    + Simplifying setup for developers new to the app
    + Portability

#### Config
- Configuration is anything that changes between deployment environments
- Example
    + Database
    + Credentials
    + URLs, …
- Config should be strictly separated from code
- Config belongs to env, not to app
- Config is stored in env vars
- Why?
    + Portability
    + Security

#### Backing Services
- A backing service is any service the app consumes over the network as part of its normal operation.
- Treat backing services as attached resources via a single URL
- Why?
    + Loose coupling to deployment and backing services

#### Build, Release, Run
- Build stage:  converts a code repo into an executable bundle (build)
- Release stage: combine Build + Deployment Config → Release 
- Run stage: runs the app in the execution environment
- Every release should always have a unique release ID (timestamp, an incrementing number)
- Why?
    + Loose coupling between stages

#### Processes
- Processes should be stateless
- Any data that needs to persist must be stored in a stateful backing service, typically a database
- Why?
    + Scalability

#### Port Binding
- The twelve-factor app is completely self-contained
- App should not depend on any external piece of infrastructure
- App decide how to handle requests
- Why?
    + Portability

#### Concurrency
- Decomposing app into individual apps that do separate jobs well
- Sounds like microservices
- Examples:
    + Process A: handling HTTP request
    + Process B: handling long-running background tasks
- Why?
    + Scalable
    + Simple
    + Reliable

#### Disposability
- Quick startup: ideally < 1 min → agility
- Graceful shutdown: TERM signal 
→ graceful shutdown (< 10s) : release resources, clean up connections, complete transaction, notify others, …
- Resilience to failure: can shutdown and restart very quickly
- Why?
    + Agility
    + Reliability
    + Resilience

#### Dev/Prod Parity
- Keep development, staging, and production as similar as possible
- The same to Infrastructures, Web server, Database, …
- Is designed for continuous deployment
- Why?
    + Avoid bugs
    + Development Speed
    + Agility

#### Logs
- The event stream for an app can be routed to a file
- The stream can be sent to a log indexing and analysis system such as Splunk
- Why?
    + Scalability
    + Performance

#### Admin Processes
- Admin Processes: migration, correct data, …
- Admin tasks should be run as isolated processes
- Why?
    + Loose Coupling between Admin Processes and Apps

### 3. Code Review

#### The Goal of Code Review
- Quality Control
    + Products
    + Codebase

- Self Growth
    + Hard Skills
    + Soft Skills
        - Presentation
        - Listening


#### Code Review Points
- Functionality
- Design
- Readability
- Consistency
- Testing
- Documentation
- Good Things

#### Code Review Process
- Minor Changes: Peer Review
- Major Changes: Group Review
- Process:
    + Reviewee (developer) completes feature / fix bugs
    + Reviewee create MR (with context)
    + Reviewee make sure that code is compiled, tested successfully
    + Reviewee inform reviewer(s) to schedule code review
    + Reviewer create checklist
    + (optional) Reviewee implements fixes that is agreed in both sides
    + Reviewer approves MR

#### Attitude
- Polite
- Contribution
- Education without criticism

#### How to Write Code Review Comments
- Be Polite
- Tell why
- Show approaches
- Labeling comments:
    + <span style="color:red;font-weight:bold">TODO (Critical)</span>
    + NIT (Nitpick)
    + OPTIONAL
    + FYI (For Your Information)

#### Best Practices
- Less is more. Small Commit → Short Review → Better Code
    + Google: ~250 lines of code / CR
    + < 400 lines of code / CR
- Checklist
- Leverage Tools such as Static Code Analysis
- Speed up code review

#### Resolving Conflicts
- Listen and understand ideas of both sides
- Analyse pros and cons of solutions
- Determine the context
    + Ordering the priorities of characteristics, requirements at the moment
- Choose the right solution according to the context
- If still have no same view, expose the problem to:
    + The whole team
    + Tech Lead / SA
    + Head of Engineer

### 4. References
- [12factor.net](https://12factor.net/)
- [12factor app - redhat](https://www.redhat.com/architect/12-factor-app)
- [Code review - google](https://google.github.io/eng-practices/review/reviewer/)
- [Code review - medium](https://www.redhat.com/architect/12-factor-app)