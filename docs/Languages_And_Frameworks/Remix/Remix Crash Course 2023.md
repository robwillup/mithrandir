# Remix Crash Course 2023 (React Framework)

[Reference](https://www.youtube.com/watch?v=d_BhzHVV4aQ)

What is Remix?
Remix is a React framework for building fullstack
apps.

Remix is a fullstack react framework by the creators
of React Router.

Why Remix?
Remix builds on top of React, but it adds features
that simplify the process of writing full stack
web applications.

## Some of the things Remix offers

* Server-Side Rendered React (good for SEO)
* File System Rounting (create a file, it's a route)
* Nested Routes (Reuse parent elements)
* Loaders & Actions (Server Functions)
* Easy Access To `<head>` Tags & Documents
* Error Handling / Error Boundary
* TypeScript & Types
* Built In Support For Cookies / Sessions

## Remix vs NextJS

|Remix                                                        |NextJS                                                         |
|:-----------------------------------------------------------:|:-------------------------------------------------------------:|
|Always server-side rendered                                  |Optional server-side rendering supported                       |
|No static site generation (build-time pre-rendering)         |Static site generation (at build time) supported               |
|Always requires host that supports server-side code execution|Deployment options: Sttic hosting vs server-side code execution|

## Remix Essentials - Crucial Fundamentals

* Project Structure & Routing
* Data Fetching & Manipulation
* Styling & Metadata

## Building Your First Remix App

You are going to need Node.js. I was on Windows by
the time I was doing this course, so I used `scoop`
to install Node.js:

```powershell
scoop install nodejs
```

The command that is used to create a Remix project
is:

```powershell
npx create-remix@latest
```
When I ran that command, I selected "Just the basics",
then "Remix App Server", then "JavaScript".

There seems to be a bug here because, when the
project was created it generated a "tsconfig.json"
file instead of a "jsconfig.json" file. The other
files in the project were "*.jsx", so that seemed
right.

> Routing is a core concept of Remix!
> Routing: Loading different components for
> different URLs.
>
> For example:
> my-domain/      => Load the "HomePage" component
> my-domain/news  => Load the "News" component

## Exploring the project structure

File: `app/root.jsx`

This is a very important file. This is basically
the root component that is always rendered no
matter which route, which specific page, is loaded.
It contains the skelleton that will be used by any
page that makes up your application.

File: `entry.client.jsx` and `entry.server.jsx`.

These define the code that will be executed for
each request that reaches the server and the client.

File: `routes/index.jsx`

Any files placed in the `routes` folder will be
treated as separate routes by Remix.

A note about SEO.
Search Engine Optimization relies on a web page for
a given URL to be rendered and loaded from the server.

Remix allows you to render a page from the server
when visiting the page directly, such as a new user
who navigates to the URL `localhost/demo`, but at
the same time it allows you to use the `Link`
component from the `@remix-run/react` pacakge which
renders a page client-side when clicking on a link
created with that component.

