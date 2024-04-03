## Version Control Systems

What is version control? Also known as source control or revision control, version control is an important software development practice for tracking and managing changes made to code and other files. It is closely related to source code management.

Version control allows you to track every change made to the code base, to see the entire history of who changed what at any given time, and to roll back from the current version to an earlier version if they need to. It also creates a single source of truth.

As a simple example, many times without source control we would not be able to quickly and precisely identify exactly what change caused a certain issue. Of course, after a long debugging session we would eventually find the culprit but just imaging having to do that in a large codebase that is in active development and where many developers need to contribute.

Version control can be thought of as a safety net to protect the source code from irreparable harm, giving the development team the freedom to experiment without fear of causing damage or creating code conflicts.

If developers code concurrently and create incompatible changes, version control identifies the problem areas so that team members can quickly revert changes to a previous version, compare changes, or identify who committed the problem code through the revision history. With a version control system (VCS), a software team can solve an issue before progressing further into a project. Through code reviews, software teams can analyze earlier versions to understand the changes made to the code over time.

Depending on a team's specific needs and development process, a VCS can be local, centralized, or distributed. A local VCS stores source files within a local system, a centralized VCS stores changes in a single server, and a distributed VCS involves cloning a Git repository.

## Why use version control?

As organizations accelerate delivery of their software solutions through DevOps, controlling and managing different versions of application artifacts - from code to configuration and from design to deployment - becomes increasingly difficult.

Version control software facilitates coordination, sharing, and collaboration across the entire software development team. It enables teams to work in distributed and asynchronous environments, manage changes and versions of code and artifacts, and resolve merge conflicts and related anomalies.

## What is a version control system?

A version control system (VCS) tracks every alteration to a file or set of files, enabling developers to journey back to earlier versions and collaborate seamlessly. Centralized version control systems (CVCS) streamline this process by housing all file version on a single server. Developers borrow a file to tweak, then return it with updates, all nearly stored and cataloged by the server. This method shines in its simplicity, offering a straightforward path for managing changes.

Yet, as teams grow and projects become more intricate, the distributed version control systems (DVCS) such as git step into the spotlight. DVCS doesn't just centralize files; it democratizes them. Every developer holds the entire project history locally, empowering offline work and facilitating a tapestry of branching and merging strategies. This flexibility is a boon for dynamic teams aiming to weave together multiple project threads without tangling them.

Whether centralized or distributed, version control is the cornerstone of efficient, cohesive software development. It safeguards progress, clarifies the past, and smooths the path forward, ensuring that every team member can contribute their best work towards crafting stellar software.

## Types of version control systems

The two most popular types of version or revision control systems are centralized and distributed. Centralized version control systems store all the files in a central repository, while distributed version control systems store files across multiple repositories. Other less common types include lock-based and optimistic.

### Distributed

A distributed version control system (DVCS) allows users to access a repository from multiple locations. DVCSs are often used by developers who need to work on projects from multiple computers or who need to collaborate with other developers remotely.

### Centralized

A centralized version control system (CVCS) is a type of VCS where all users are working with the same central repository. This central repository can be located on a server or on a developer's local machine. Centralized version control systems are typically used in software development projects where a team of developers needs to share code and track changes.

### Lock-based

A lock-based version control system uses file locking to manage concurrent access to files and resources. File locking prevents two or more users from making conflicting changes to the same file or resource.

### Optimistic

In an optimistic version control system, every user has their own private workspace. When they want to share their changes with the rest of the team they submit a request to the server. The server then looks at all the changes and determines which ones can be safely merged together.

## Benefits of version control

Version control systems (VCS) stand as a pivotal practice withing software development, enabling better management, tracking, and implementation of changes to code and related files. By providing a structured approach to revision control, VCS supports dynamic, collaborative environments, providing stability across development projects. The advantages of employing version control stretch from enhancing code quality to accelerating development timelines and improving project visibility. All of which makes it an indispensable tool for teams aiming for high efficiency and quality in software delivery.

### Quality

Version control encourages a culture of continuous peer review and collaboration, leading to significant improvements in code quality. By facilitating detailed tracking of every change, teams can easily review, comment, and refine their work, ensuring adherence to best practices and standards. This collaborative scrutiny not only elevates the quality of the output but also aids in early bug detection and resolution.

### Acceleration

Version control systems streamline development processes, enabling faster iteration and delivery of features. Efficient branching and merging capabilities allow developers to work concurrently on various aspects of a project without interference, significantly reducing the time from development to deployment. Additionally, the ability to quickly revert to previous versions minimizes downtime when addressing issues, keeping the project momentum steady.

### Visibility

A central repository in a version control system acts as the single source of truth, enhancing project transparency and accountability. This centralized view of the project's evolution aids in better planning, tracking, and collaboration, as every team member has access to the latest updates and historical changes. The integration with project management tools, further bolsters project oversight, linking code changes directly to tasks and milestones.

## What are the main version control systems?

The three most well-known version control tools/systems are Git, Subversion, and Mercurial.

### Git

Git is the most popular option and has become synonymous with "source code management". Git is an open source distributed system that is used for software projects of any size, making it a popular option for startups, enterprise, and everything in between.

### Subversion (SVN)

SVN is a widely adopted centralized VCS. This system keeps all of a project's files on a single codeline making it impossible to branch, so it's easy to scale for large projects. It's simple to learn and features folder security measures, so access to subfolders can be restricted.

### Mercurial

Mercurial is a distributed VCS that offers simple branching and merging capabilities. The system enables rapid scaling and collaborative development, with an intuitive interface. The flexible command line interface enables users to begin using the system immediately.

## How does version control streamline collaboration?

Version control coordinates all changes in a software project, effectively tracking changes to source files, designs, and all digital assets required for a project and related metadata. Without it, projects can easily devolve into a tangled mess of different versions of project files, hindering the ability of any software development team to deliver value.

With a strong VCS, software teams can quickly assemble all critical project files and foster actionable communication to improve code quality. And because it provides a single source of truth, stakeholders from access a DevOps team can collaborate to build innovative solutions - from product managers and designers to developers and operations professionals.

## More study notes and content

Reading notes from the `Pro Git` book by Scott Chacon and Ben Straub can be found [here](./Pro_Git_Book/).
