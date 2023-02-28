# Getting Started With GitHub REST API

The first request that I issued to the GitHub REST API was this one:

```bash
curl --request GET \
--url "https://api.github.com/repos/robwillup/Mithrandir/commits" \
--header "Accept: application/vnd.github+json" \
--header "Authorization: Bearer github_pat_******"
```

And that returned my commits in the
[Mithrandir](https://github.com/robwillup/Mithrandir) repo. So, how did I get
that Personal Access Token? Here's how.

## GitHub Personal Access Tokens

First of all, my little disclaimer: This document is only my personal journaling
of my discoveries as I develop an app to interact with GitHub's REST API.
Please, check the official documentation as the authoritative source.

[GitHub REST API](https://docs.github.com/en/rest)

Ok, so by the time I ran the request above I had already created a GitHub app,
but I wanted to have a quicker way to just run manual local tests and see what
certain endpoints return. That's why I created a PAT. The official documentation
states that:

> Personal Access Tokens are intended to access GitHub resources on behalf of
> yourself.

So indeed a PAT was exactly what I needed for my local manual experimentation.
There are currently two types of PATs in GitHub: fine-grained and classic. The
former is recommended by GitHub whenever possible, because it has several
advantages, such as:

* A single token cannot be used to access resources owned by more than a single
  user or organization.
* Only specific repositories can be accessed by a single token.
* You can control which permissions each token has.
* You cannot have a token without an expiration data.
* Org owners can require approval for new tokens accessing their resources.

[This is the list of features that only work with classic PATs.](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#personal-access-tokens-classic)

Ok, in short, for my local tests, a fine-grained PAT was just what I needed. So,
if you think you need it too, how do you get one? Well, I could lay that out
here but the steps to do that are very nicely described in the official docs, so
here's the [link](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#creating-a-fine-grained-personal-access-token).

## Keep It Safe

Last but not least, we should talk about keeping our PATs safe. Like passwords,
PATs have inherent security risks and therefore they should be kept as safely as password.

* Do not share PATs, because as the name implies, they are personal.
* Do not commit PATs along with your source code.
* Do not paste PATs in insecure locations without enough access control points.

[More information](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#keeping-your-personal-access-tokens-secure).
