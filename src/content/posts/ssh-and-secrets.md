---
title: "SSH and secrets"
description: "A post about a mistake and some projects, opsiess."
tags: ["Programming", "Life", "Go"]
createdAt: "Jun 01 2026"
---

So, hey... i lost the **SSH** key to the instance of this blog...

I formatted my pc and forgot about the SSH keys, so i can't really update the blog anymore, which sucks because that are a lot of things that i wanted to implement and that i kinda hate on the state that it is right now.

I could just delete this instance and create another one but i don't really want to, i actually hate the backend of the blog. For context the backend is made in [Django](https://www.djangoproject.com/) and don't get me wrong, is not a bad Framework, but, it is in my opinion kinda bloated for a simple blog that just display text, has a lot of opinions aswell, which is not much of a problem for me but it is for other people, the main reason actually is the fact that i'm learning [Go](https://go.dev) (which quicly became my favorite language) and i wanted to rewrite everything in **Go**.

I will do it eventually, it's going to be my next project, but that's not really what i want to talk about, the thing is, to attone for my sins on losing my goddamn SSH keys i wrote a simple CLI password manager called [agopass](https://github.com/Alvesafk/agopass).

## Agopass

It is 100% **Go**, the program saves your keys on database, it handles the encryption and decryption of your secrets, i made it to be very simple, few commands because i do think you don't really need a lot of complexity for a password manager, atleast it's not what i want, i want to just add my keys to it and afterwards get it in a quick way.

I was going to use the [Cobra](https://github.com/spf13/cobra) package, it's a way to handle arguments and flags, but i actually never used flags, so i just handled the arguments myself.

I think it was a good project to learn more about **Go** and how to work with it, i think now i really understand what a package is and how it works, also worked with a bit of **SQL** and how can be integrated with **Go**, a ability that i will of course use on the next version of the blog.

I actually don't have much to talk about the project itself because it is really simple, if you're interested to use it you can go on my github and see how to install and what the commands do, i really just wanted to write a bit.

Thanks for your attention, i suspect that the next post will be on the new blog, i will make sure to port the posts and comments to the new one when it's done.

Thank you for your attention, sincerely Alvesafk.
