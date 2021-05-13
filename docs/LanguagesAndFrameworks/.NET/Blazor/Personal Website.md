# Building my Person Website

This document lists the steps and details about building my person website.

## Framework

Blazor Web Assembly was choses as the framework for my personal website.

## Styling

To style this web app, [FAST](https://www.fast.design/docs/integrations/blazor) was selected.

## Issues

### CSS Isolation in Docker

#### The Problem

When using Docker and Nginx to server the Blazor WASM app, some CSS would not work. So the styling would be messed up, there were some missing fonts and images.

After some debugging I realized the issue was related to the [`scope identifiers`](https://docs.microsoft.com/en-us/aspnet/core/blazor/components/css-isolation?view=aspnetcore-5.0#css-isolation-bundling).

The identifiers in the bundled CSS file, `{ASSEMBLY NAME}.styles.css`, did not match the ones in the actual razor files.

In the two images below, the `scope` identifiers are matching, and that's what I have when running the app locally with the `dotnet` CLI and also after the problem was resolved with the Docker image:

![elements](../../../../assets/images/MyPersonalWebsite/css-isolation-issue-elements.png)

![source](../../../../assets/images/MyPersonalWebsite/css-isolation-issue-source.png)

#### The cause

The problem "seems" to be a bug in the `dotnet publish` command, which was not regenerating the scope identifiers for new publishes.

I've added a comment to this [Stack Overflow answer](https://stackoverflow.com/a/65548488/7007769).

#### The Solution

Deleting and recreating the isolated CSS files with the same content actually did the trick. Please, refer to the link above.
