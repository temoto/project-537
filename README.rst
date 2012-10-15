What
====

Project-537 is a remake of famous video game Constructor http://en.wikipedia.org/wiki/Constructor_(video_game%29
by System 3.

537 is number of weeks passed from the beginning of 21st century, that is, from 2001-01-01.


Design
======

Overall architecture consists of two independent components: Backend and Frontend.

Components may be implemented in different programming languages or even run on different machines.
Communication is done via ZeroMQ messages.

Why components stuff:

First reason is that it really helps to concentrate on one thing. Also helps to localize bugs.

Second reason: it allows good game designs to reincarnate into new UI bodies and live forever young.
There are many examples where excellent game design ideas, features are immortal: Tetris, Elite, Fallout.
I think Constructor has great unique features too. But years pass, platforms change, it's not so easy
to run good old games on todays Windows7/Ubuntu/MacOSX. What if they just had game logic and
visualization separated... Somebody could hack up a new shiny GUI for good old games, now with 32bit
colors, shaders, all modern graphics stuff.


Backend
^^^^^^^

This is the calculating heart of the system. It maintains internal state of the game, processes
user actions from Frontend and tells it what it *can* show.

For the sake of exploration i will try to implement backend in Go.

Frontend
^^^^^^^^

This component is responsible for interaction with users: rendering game map, playing sounds,
parsing keyboard/mouse/etc input into commands for backend.

For fun i will try to implement frontend as web application.
