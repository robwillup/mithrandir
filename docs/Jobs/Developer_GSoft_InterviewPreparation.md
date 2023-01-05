# Full Stack Developer Interview Preparation

What will my new role look like?

As a full-stack developer taking part in a platform team, you will support the Migration Tool feature team
to improve the developer's experience and help them accelerate Product Development while improving
product reliability.

You will enable and support the Migration Tool feature teams on the following matter:

* DevOps metrics: by focussing on processes & tooling
* Product reliability: by focussing on measurement and alerting
* Release capabilities: by focussing on feature flagging, disaster recovery and rollback capabilities
* Team autonomy: by focussing on tooling
* Teams toil: by identifying, measuring, and removing work repetitiveness/complexity
* Cognitive load: by providing and owning tools & services not linked to the stream-aligned domain

You will have an impact by adopting a product mindset and by playing the following way:

* By interacting with Feature Team to understand their needs
* By maintaining the same product mentality as other teams on the owned code: Build it, ship it, run it
* By aiming for minimal platform: Maintain balance between keeping the platform small and ensuring that
the platform helps sped up and simplify software delivery for the teams building on the platform.
* By being accountable of the services your team offers.
* By relying on fast prototyping techniques and involving feature team members for fast feedback on what works and what does not.
* By coordinating with similar practices and guidelines of SRE, Tech platform, Tech team.
* By keeping an eye out for opportunities to contribute to the ShareGate and GSoft platform.

Typical week:

* Researching, learning, and developing (new and incremental value): 70%
* Presenting and sharing knowledge with your colleagues (day-to-day delivery rituals): 20%
* Helping and supporting teams to improve their development experience (individual & collective growth): 10%

Technical stack:

* C#/.NET, Terraform, Bicep
* WPF, xUnit, Hangfire, MVC
* Azure DevOps, TeamCity, SharePoint On-Premise & Online, Jira, Fidler, PowerShell

What does my team look like?

A platform team of 2 to 3 people, all full-stack developers. The teams works closely with the Migration
Tool team (feature team) as well as with several other experts: a security specialist, several
technology experts, other platform/SRE teams and more.

Qualifications:

* Mastery of C#/.NET programming and scripting
* Comfortable with Desktop product development context
* Strong ability to debug and optimize code
* Experience automating repetitive tasks and troubleshooting
* Systematic approach to problem solving and rigorous analytical skills
* Understanding of the practices that drive DevOps and SRE culture

## The GSoft Interview Jan 5th, 2023

### My Background and Experiences

**Disys:** I started working with IT in 2014 as service desk analyst at the same time I started studying Systems Development and Analysis at College. I was hired by a IT Consultancy company called Disys to work for one of their largest clients, Oil and Energy giant ExxonMobil. My responsibility there was to assist ExxonMobil employees and contractors with any IT issues they were having, including software, OS, network and even some hardware issues. I was part of the Global team so I talked to people from every continent.

**ExxonMobil:** In August of that same year, because I had a very good performance and was continually contributing to meet our team's SLA I was invited by ExxonMobil to apply to their Internship program. And in January of the following year I started working directly for ExxonMobil as a software developer intern. My new team was called ITOE (Information Technology Operations Excellence) which was an internal division of the Customer Service organization. There I had the opportunity to start developing some features while being mentored by some great senior developers. Specifically I worked with C#, JavaScript, HTML and CSS automating several repetitive tasks for different teams in SharePoint and InfoPath, creating custom team sites. I contributed to a project to develop an application to measure and monitor the volume of calls our Service Desk was receiving, I wrote the ErlangC formula in C# that was later used for the call volume and wait time calculations. And by the half of that same year, 2015, I was converted from an intern to an employee.

In the second quarter of 2016 we received the news that our organization was being out-sourced and several employees were being laid off.

**Wipro:** Fortunately, I had a connection working for another company called Wipro and he mentioned that they were hiring a trainee for their Microsoft project, with software development background to help maintain and moderate the MSDN (Microsoft Developer Network) Forum. I applied for the role and was hired in June 2016. That was a unique experience because although I did not directly develop any features, I was responsible for monitoring the forum, providing answers to some questions, opening support cases for more complicated issues. I had the chance to study Microsoft development products more in depth, such as .NET ecosystem and Visual Studio.

**Sanepar:** I still wanted to have more extensive hands-on experience with writing code, designing features and so in May 2017 I started working as a software developer for a Water Treatment/Health Insurance company. On the first week I was given the task of re-designing and re-writing a legacy HR application that was used to track employees' career development. The legacy application was a .NET Framework Web Forms app that interacted with an Oracle database. I delivered the first iteration of the new version within the deadline re-written as an MVC web app in the most recent version of .NET Framework at the time (.NET Core was still very young in 2017). Other projected that I worked with at that company include a mobile app for iOS and Android written with Ionic, AngularJS, and Oracle DB. Modernizing the main application for their health insurance division, with .NET Core and Angular. At times I also helped maintain a very old VB.NET application.

**ExxonMobil:** In 2018 ExxonMobil reached out to me inviting to join their new Software Security team as a developer. That opportunity was aligned with my career goal of advancing technically as a software engineer so I accepted. To this day, that was the job where I learned the most. Up until that time I only had theoretical knowledge of DevOps, CI/CD, containers, microservices, and had very limited cloud experience. All of that was opened to me when I re-joined them.

My initial responsibility there was to automated repetitive tasks that developers from other organizations had to perform to collect and work on security vulnerabilities found by a source code static analyzer that my team, SSG, provided. So developers would have to manually run our SAST tool on their source code, inspect and triage all findings to filter out false positives, send the remaining true positives to the person in their teams who was responsible for creating new work items in their backlogs, so that they could work on fixing those security bugs when the tasks were included in their sprints. So to help automate part of that workflow I wrote a middleware API in .NET Core that collected the scan results, applied a custom filter, and created work items in the respective backlogs. I also helped develop a plugin for TFS (now Azure DevOps Server) that could be added to the developers CI tasks.

Another very interesting task I had was to containerize our SAST tool which until then was running directly on a few on-premises VMs. Part of the objective of this containerization effort was to make static code analyses available in cloud CI pipelines, such as Azure DevOps and GitHub Actions. I was the lead on this project. As we were implementing this task we found several blockers that were caused specifically by the technical limitations of this third-party SAST tool. That, combined with our own experience using the tool and the feedback provided by hundreds of developers across different organizations, helped us make the decision to write our own tool for running static code analyses.

In this containerization project I traveled to Bangkok to contribute with a larger company-wide project
that was modernizing application, and wanted to include our SAST tool in it.

The new architecture to support our own internal tool was a huge simplification compared to the previous architecture. It was basically composed of four components: A CLI app written in Go, a REST API in C#, a MongoDB, and a worker service in C#. I completed the first prototype in a week and it received great feedback. In less than a year we released the application into production supporting more than 10 different programming languages. The vendor app was discontinued and our organization could save around US$ 100 yearly.

I created, sent out and reviewed surveys to find out what what the feedback on our tools.

I traveled to Buenos Aires to give talks about OpenID Connect and OAuth2.

I developed a desktop application that used SentinelOne NextGen antivirus to scan USB ports.

I've been to developer conferences.

**Plurilock:** I've always wanted to move to an English speaking nation such as Canada, New Zealand, Australia, England, or US, and last year the opportunity came when I was contacted by my current company, Plurilock, to help them modernize their Windows agent and one of the benefits would be the visa sponsorship to Canada. I accepted and this year I've moved here with my wife and daughter.

### My Aspirations

My number one professional aspiration has been to evolve and advance technically as a software engineer, not limited to an specific programming language or framework, but really mastering the fundamentals of computing and software engineering. Building on that, I want to keep growing to be able to contribute to others in my fields, by mentoring new devs, having some talks, working as a technical lead, and creating technical content and finally finding the time and courage to contribute significantly to open-source projects.

I find that I need to update and expand my knowledge of platforms and operations and that's why this opportunity to work as a full stack developer at GSoft is of great interest to me.

### Who am I (in general)

I've really liked creating, replicating and adapting things. At school I did a lot of graphite drawing, took a course for that. And since college I've had that same interest for software development.

I love to spend time with my family, reading books (both technical and fictional), getting to know different places, and just recently I've restarted drawing. Also love basketball and played for about 7 years.

### My questions to GSoft

* Can you give an example of a project currently going on in this team?
* Can you tell me a current major challenge or road block you are facing?
* What tools are you currently using in your pipelines for software composition analyses, thread modeling, SAST, DAST?
* What is the direct interaction with other technical teams?
* Any fixed time alloted for professional training? Open-source contributions? Conferences?
* Does this team run PoCs of new technologies or products?
* How is knowledge sharing structured? Or is it more organic?
