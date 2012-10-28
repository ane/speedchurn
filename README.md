Speedchurn
==========

speedchurn - Rapid IRC Statistics &amp; Analysis

Current status
--------------

 * nothing is calculated, nothing is processed
 * only irssi format is supported
 * the codebase is only a day old

Installation
------------

If you do not have installed Go yet, get it from http://golang.org

    git clone http://github.com/ane/speedchurn.git speedchurn
    cd speedchurn && go build
    speedchurn irclog.log

Background
----------

risa is a irc statistics generator that puts emphasis on informative and thorough
statistics, clear visualization and most of all, speed. The status of current mainstream
IRC statistics generators is very bad when it comes to large log files, and work on
them has mostly stopped.

As such, a much more faster stats generator that can properly exploit modern software design
principles, e.g., concurrency, is called for. 

speedchurn borrows a lot of principles from ane/hisg (with regards to concurrency and speed)
and is the ideological successor to that project.

Roadmap & goals
---------------

  * Visualization using interactive charts
  * Focus on pertinent data over impertinent data
  * 86 bazillion charts with 96.7% more wub wub
  * **Speed**.
  * Support most formats (Irssi, mIRC, xchat, bots) using simple regex patterns
  * Platform-independency
