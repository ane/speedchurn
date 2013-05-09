Speedchurn
==========

speedchurn - Rapid IRC Statistics &amp; Analysis

Current status
--------------

 * irssi format is supported
 * templates work
 * the following statistics are calculated:
   * lines and words for each user
   * topic changes
   * number of elapsed days
 * automatic alt-nick (nick_ -> nick) recognition and merging 
 * distributed stats calculation using map/reduce and parallelism

Installation
------------

If you do not have installed Go yet, get it from http://golang.org

    git clone http://github.com/ane/speedchurn.git speedchurn
    cd speedchurn && go build
    speedchurn irclog.log

Background
----------

speedchurn is a irc statistics generator that puts emphasis on informative and thorough
statistics, clear visualization and most of all, speed. The status of current mainstream
IRC statistics generators is very bad when it comes to large log files, and work on
them has mostly stopped.

As such, the development of a much faster stats generator capable of  properly exploiting 
modern software design principles, e.g., concurrency, is called for. 

speedchurn borrows a lot of principles from ane/hisg (with regards to concurrency and speed)
and is the ideological successor to that project.

Feature Roadmap
---------------

  * Achievements
  * Automatic FTP/SCP-based deployment of the pages
  * Charts produced using D3.js
  * Focus on pertinent data over impertinent data
  * 86 bazillion charts with 96.7% more wub wub
  * **Speed**.
  * Support most formats (Irssi, mIRC, xchat, bots) using simple regex patterns
  * Platform-independency
