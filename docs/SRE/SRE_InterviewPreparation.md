# SRE Interview Preparation

What is expected?

Optimize the product production environment through high-performance code and scripts,
improve work through automation, and monitor system capacity and performance.

You need to be a savvy SRE who loves to teach and mentor on the best practices.

The responsibilities you will have include:

* Maintain live services by measuring and monitoring availability, latency, and overall
system health.
* Solve problems by writing code and automating processes ("as-code" mindset).
* Proactively scale and operate systems through automation.
* Estimate production capacity planning based on business needs.
* Continuously evolve systems by proposing and advocating changes to improve reliability.
* Collaborate with development teams to increase production readiness of their product features.
* Participate in incident management, provide post-incident investigation recommendations,
and improve incident management through processes, automation, and exercises/simulations/training.
* Proactively research and test potential performance and availability risks.
* Share SRE expertise and finding with other development teams.

The qualifications you need to have are:

* At least 3 years of experience as a Software Engineer.
* Experience as an SRE (DevOps) is a plus, no experience but interest is also good.
* Good knowledge of programming and scripting.
* Hands-on experience with cloud infrastructure (Azure).
* Hands-on experience with C# and Kubernetes.
* Experience with distributed systems.
* Strong ability to debug and optimize code.
* Experience automating repeating tasks and troubleshooting.
* A systematic approach to problem-solving and rigorous analytical skills.
* Understanding of the practices that drive an SRE culture.

## The GSoft Interview Dec 12th, 2022

### My Background and Experiences

**Disys:** I started working with IT in 2014 as service desk analyst at the same time I started studying Systems Development and Analysis at College. I was hired by a IT Consultancy company called Disys to work for one of their largest clients, Oil and Energy giant ExxonMobil. My responsibility there was to assist ExxonMobil employees and contractors with any IT issues they were having, including software, OS, network and even some hardware issues. I was part of the Global team so I talked to people from every continent.

**ExxonMobil:** In August of that same year, because I had a very good performance and was continually contributing to meet our team's SLA I was invited by ExxonMobil to apply to their Internship program. And in January of the following year I started working directly for ExxonMobil as a software developer intern. My new team was called ITOE (Information Technology Operations Excellence) which was an internal division of the Customer Service organization. Their I had the opportunity to start developing some features while being mentored by some great senior developers. Specifically I worked with C#, JavaScript, HTML and CSS automating several repetitive tasks for different teams in SharePoint and other platforms, creating custom team sites. I was contributed to a project to develop an application to measure and monitor the volume of calls our Service Desk was receiving, I wrote the ErlangC formula in C# that was later used for the call volume and wait time calculations. And by the half of that same year, 2015, I was converted from an intern to an employee.

In the second quarter of 2016 we received the news that our organization was being out-sourced and several employees were being laid off.

**Wipro:** Fortunately, I had a friend working for another company called Wipro and he mentioned that they were hiring a trainee for their Microsoft project, with some software development background to help maintain and moderate the MSDN (Microsoft Developer Network) Forum. I applied for the role and was hired in June 2016. That was a unique experience because although I did not directly develop any features, I was responsible for monitoring the forum, providing answers to some questions, opening support cases for more complicated issues. I had the chance to study Microsoft development products more in depth, such as the .NET ecosystem and Visual Studio.

**Sanepar:** I still wanted to have more extensive hands-on experience with writing code, designing features and so in May 2017 I started working as a software developer for a Water Treatment/Health Insurance company. On the first week I was give the task of re-designing and  re-writing a legacy HR application that was used to track employees career development. The legacy application was a .NET Framework Web Forms app that interacted with an Oracle database. I delivered the first iteration of the new version withing the deadline re-written as an MVC web app in the most recent version of .NET Framework at the time (.NET Core was still very young in 2017). Other projected that I worked with at that company include a mobile app for iOS and Android written with Ionic, AngularJS, and Oracle DB. Modernizing the main application for their health insurance division, with .NET Core and Angular. At times I also helped maintain a very old VB.NET application.

**ExxonMobil:** In 2018 ExxonMobil reached out to me inviting to join their new Software Security team as a developer. That opportunity was aligned with my career goal of advancing technically as a software engineer so I accepted. To this day, that was the job where I learned the most. Up until that time I still didn't know what was DevOps, CI/CD, containers, microservices, and had very limited cloud experience. All of that was opened to me when I re-joined them.

My initial responsibility there was to automated repetitive tasks that developers from other organizations had to perform to collect and work on security vulnerabilities found by a source code static analyzer that my team, SSG, provided. So developers would have to manually run our SAST tool on their source code, inspect and triage all findings to filter out false positives, send the remaining true positives to the person in their teams who was responsible for creating new work items in their backlogs, so that they could work on fixing those security bugs when the tasks were included in their sprints. So help automate part of that workflow I wrote a middleware API in .NET Core that collected the scan results, applied a custom filter, and created work items in the respective backlogs. I also helped develop a plugin for TFS (now Azure DevOps Server) that could be added to the developers CI tasks.

Another very interesting task I had was to containerize our SAST tool which until then was running directly on a few on-premises VMs. Part of the objective of this containerization effort was to make static code analyses available in cloud CI pipelines, such as Azure DevOps and GitHub Actions. I was the lead on this project. As we were implementing this task we found several blockers that were caused specifically by the technical limitations of this third-party SAST tool. That, combined with our own experience using the tool and the feedback provided by hundreds of developers across different organizations, helped us make the decision to write our own tool for running static code analyses.

The new architecture to support our own internal tool was a huge simplification compared to the previous architecture. It was basically composed of four components: A CLI app written in Go, a REST API in C#, a MongoDB, and a worker service in C#. I completed the first prototype in a week and it received great feedback. In less than a year we released the application into production supporting more than 10 different programming languages. The vendor app was discontinued and our organization could save around US$ 100 yearly.

SentinelOne, OpenID Connect, Bangkok, Argentina, DevConf, talks, etc.

**Plurilock:** I've always wanted to move to an English speaking nation such as Canada, New Zealand, Australia, England, or US, and last year the opportunity came when I was contacted by my current company, Plurilock, to help them modernize their Windows agent and one of the benefits would be the visa sponsorship to Canada. I accepted and this year I've moved here with my wife and daughter.

### My Aspirations

My number one professional aspiration has been to evolve and advance technically as a software engineer, not limited to an specific programming language or framework, but really mastering the fundamentals of computing and software engineering. Building on that, I want to keep growing to be able to contribute to others in my fields, by mentoring new devs, having some talks, working as a technical lead, and creating technical content and finally finding the time and courage to contribute significantly to open-source projects.

I find that I need to update and expand my knowledge of operations and that's why this opportunity to work as a SRE has interested me so much.

### Who am I (in general)

I've really liked creating, replicating and adapting things. At school I did a lot of graphite drawing, took a course for that. And since college I've had that same interest for software development.

I love to spend time with my family, reading books (both technical and fictional), getting to know different places, and just recently I've restarted drawing. Also love basketball and played for about 7 years.

### My questions to GSoft

* Am I going to be using my software development skills?
  * In this role, do we have the change to actually code and automate things?
  * What do SRE devs typically use at GSoft in terms of scripting and programming languages?
  * What platforms do they manage? Kubernetes? What platforms are used? Cloud, containers?
* What is roughly the percentage of time SRE devs spend automating, monitoring, and supporting?
* What is the direct interaction with other technical teams?
* Any fixed time alloted for professional training? Open-source contributions? Conferences?
* Do SRE devs run PoCs of new technologies or products?
* Are there internal brown bags?