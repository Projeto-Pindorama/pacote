# Motoko Package Manager

## What is it?
Motoko is a tool kit that allows you to install, remove, create and distribute packages in your system.  
My main goal is to offer a package manager that __really__ follows the package managing model SRV4, however also not be a result of 20 year-old code turd-polishing -- in case, the Sun's pkgtools original code that was release with OpenSolaris in mid 2005.  
Futhermore, I want to do something from scratch and decently, as well as I want to learn Go for real -- in case, since plain-text manipulation to managing files and filesystem permissions. 

### SVR4? Wait, what is this?
"SVR4" stands for "System V, Revision 4", which was an UNIX version released in 1987, arising from the conjuct work of AT&T and Sun Microsystems.  
Even if it's considered "the trigger of insanity" of UNIX, at least one positive thing it brought, that in case was the improvements of BSDs pkgtools<sup>(lacks exact sources)</sup>, giving rise to what is now known as "SVR4 pkgtools".   

I'm basing myself in several documents (which can be found at [docs/references/](./docs/references)), but from all of them, the one being used the most is the "Application Packaging Developer's Guide" from Sun Microsystems, published February 2000 ([805-6338.pdf](https://www.uvm.edu/~fcs/Doc/Solaris8/805-6338.pdf)); second is the Illumos manpages ([illumos.org/man/](https://illumos.org/man)).  

## Goals (in a list)
* Faithfulness to SVR4 origin pkgtools, also at usability level;
* At the same time that is has faithfulness, it shouldn't be faithful with mistakes. That is: if something is in the original pkgtools and it can be replaced with something simpler and better without significant losses, it will be;
* To cut off the most shell dependencies and `$PATH` tools as possible;
* Keep the codebase legible and easy to hack/keep up;  
* `mk` (plan9port) as build system.  

## Non-goals (also in a list)
* Dependency resolution;
* Network download;
* Native compression;
* Support for other languages different from english in screen messages (it can change);  
Those things must be done by abstractions, not by Motoko.  
* [symbolic links](http://doc.cat-v.org/plan_9/4th_edition/papers/lexnames).

## Why not to use Slackware pkgtools or...?
Even though the Slackware pkgtools are pretty good for something made in a language for mere hacks and being in an extremely liberal license, it still leaves great deals to be desired.  
One of them is also its weakness: its complete dependency on GNU Bourne-Again Shell (not counting the code illegibility due to hacks joined to the simple cluelessness in some parts).  

About the original pkgtools that Sun released in 2000, it simply doesn't worth to use it.  
Many resources were cut off from Heirloom port for being "Solaris only", besides the codebase being __really__ old, depending on Shell for things that could be done in C and probably containing __a lot__ of security issues.  
Oh, did I talk about they using SQL for storing things like the package list? So...  
It's something not only extremely old, but also "fat" (and extremely hard to compile, thanks to the hacks done by Gunnar Ritcher, instead of indeed porting it; what isn't exactly his fault to be fair).  
Even though I'm not exactly a good programmer, I believe that I'm going to manage to do something simpler and more efficient using a modern language.  

Now, talking about other package managing tools that are considered "KISS", I don't think they are serious, even because most of them came from edgy/meme distributions -- e.g. Ataraxia Linux, where everything looks like a giant 4chan/Reddit-influenced shitpost, not a serious community (what is probably true).  
So I wouldn't trust using it on my system just to know next month that the community decided that the joke isn't funny anymore and then suddenly deleted all repositories.

## Who uses it?
Currently, nobody -- because it's not even complete yet. :rofl:  
I'm doing it for the Copacabana Linux, but my goal is that anyone can use it.  
Maybe even the people from Musl LFS, who are still stuck to "slightly modded Slackware pkgtools", end up using it -- and I would be very happy if they do it. :smiley:  

## What happened to `otto-pkg`?
People who knows me longer (since late 2019 and early 2020) may wonder about Otto/otto-pkg.  
I truly gave up on doing it now, it's simply unworkable -- and if you was just being a dick "joking" about it and you're reading it now, I must congratulate you for getting what you wanted; now you can try to get a real job and get out from your momma's house.  
During the months I was making it (in Shell Script), I made a lot of mistakes.  
One of them, other than the influence and using an inflexible language for 90% of what I wanted to do originally, was the lack of planning.  
It didn't have an established idea, we added everything we had in mind, and we ended up having everything but the main thing; in addition to conflicts between me and Caio Yoshimura (second developer) just for pure imaturity from both sides.  

At least, Otto left some good things -- which are not only my credit.  
Among these things, I can talk about `pico-torrent`, created by Jazz\_Man, a hacker and an old friend of mine that I didn't see in a year.  
I can also talk about the libraries for Shell Scripts, among them,`posix-alt.shi`, that was created by me and Caio back in the time that we "took the bait" of POSIX.  
And finally, the knowledge.  

### So, Otto died and Motoko took it's place?
No, not even by far.  
Otto shouldn't ever be thought as a complete package manager, but as an abstraction of something smaller.  
This is the UNIX philosophy, and this one must be really followed to avoid what happened to the 'Shell Script Otto' to happen again.  
The "Shell Otto" has died, but have left its heritage.  
Otto as an idea is just alive as ever.  
Will I rewrite it? Yes, but not soon.  

## Development Standards
Lately I've talked to my friend Vitor Almeida, that maintains the project website, and we concluded that we need to maintain a standard in the project itself.  
I decided to fit Motoko in those standards to avoid what happened constantly on Otto and on amateur projects, that is: one programmer overlay another and commits get lost.  
This part will possibly be well divided on a documentation after, but in meantime, I'll keep it on README.  

### Commit standard
Conventional Commits 1.0.0:  
https://www.conventionalcommits.org/pt-br/v1.0.0/

### Code editors used in development
- acme/acme2k (9fans.github.io/plan9port);
- Vi Improved (vim.org).

## Should I have created a separate FAQ?
Probably...   

**Note[1]**: the only relation between me, the Xebec Studios an Ken Akamatsu is that I'm a huge fan on their work, and the only way I found to show how thankful I am was to name one of my projects with the name of one characters of Love Hina.  

Translated by Apocalipse <apocalipse.lain.ch> @ 19-01-2022
