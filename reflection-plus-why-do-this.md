# Why I made PFB-web: a reflection
---

Pay Friends Back (web) is a simple web application for squaring up with people after uneven spending. It accepts a list of names and dollar amounts (amount spent) and displays a list of transactions between those people that should occur in order for all people to have "spent" the same amount in the end. The idea being that each person has spent a different amount of money and now they all need to square up with each other.

For the initial motivation behind Pay Friends Back in general, see the [original motivation section](https://github.com/tzarick/pay-friends-back-cli#motivation). 

Tech Stack:
- Go backend
- JavaScript frontend logic
- HTML / CSS / Bootstrap UI layout
- Heroku cloud hosting (free tier)

I built this small web application for 3 reasons:
1. The experience of building it.
2. To practice using Go. 
3. I wanted a more easily shareable / accessible tool to exist, so that PFB might be convenient enough to be realistically useful. 


My focus with Pay Friends Back (web) was to make this tool accessible and learn things along the way. I wanted a more usable interface than the command line (the previous iteration's interface) but I really didn't want to get lost in frontend land, so I promised myself we'd keep it simple and clean, nothing too fancy or involved. To me, this looks like a simple dynamic form where each row of the form has 2 input boxes: name and amount spent. 

I chose to build this project without a frontend framework (I did use Bootstrap, however. This feels different to me though). For me, in this situation, pure JS simplifies things, makes things more lightweight (development and performance), and it's just generally satisfying to understand and have direct explicit control over most every part of the UI. Granted, this is easy for me to say given that this UI is very straightforward and simple, but part of the reason it's that way is because I wanted to use pure JS. This decision also allowed me to focus more on the part of the tech stack I was most interested in, which is Go and the backend logic. Additionally, on the front end side of things, I paid closer attention to the mobile experience this time (which really just meant spending a little more time understanding Bootstrap containers) since personally, I'll probably use this more often from my phone than a desktop.

This project was a very satisfying one to work on. I steadily chipped away at it for a little less than a week, which is on par for the size of the application (small) but also much faster than I've ever put something together like this before, end to end. I feel like this endeavor struck a really nice balance between hard & new and tried & true, which I've come to realize can make or break the overall enjoyment I feel on a project. I seem to thrive when the balance is somewhere around 40/60: the combination of new concepts required to be learned (golang in general, frontend / backend comms b/t go and js, Makefiles, more bootstrap, heroku + go, etc) along with concepts I already have in my tool belt that I was able to quickly apply (JavaScript, basic client/server setup, REST, unit testing, heroku familiarity, the algorithm and core logic etc). That's not to say a different ratio couldn't work out just as well or even better (depending on my motivation at the time, the subject matter, the documentation available, etc), but this one just feels good to me lately and is an accessible and motivating way to learn new things gradually.

The guiding principle of "Clean and Simple" for this project also helped me settle into a more decisive decision making mindset. I uncharacteristically found myself picking color hex codes on the first or second try (something I've historically had to cycle through hundreds before finally picking...) and mentally deeming feature functionality like logical input validation as "appropriate" or "enough in scope" instead of shooting for perfection under every use case in the known universe. This was refreshing and helped me keep some healthy mental distance but still chisel down to a final product I could be proud of.


## ?! Question Zone ?!
- Is this idiomatic Go, generally speaking? Would love to know where it isn't or where it might be improved.
- Should or could the server logic take advantage of goroutines? Does the router already take care of that under the hood for me? Would this matter if there was a high amount of traffic or would it just increase the efficiency for a single hit?
- Sanitizing input? Is this necessary here?
	- What else should or could I be doing from a security standpoint (there's no state storage or db or anything like that really, so I wasn't sure how else I could improve it but I'd like to know how to max that knob out, hypothetically)
	- Best / preferred Go strategies for these types of things / general application security?
- Serving static files - what's the best way in Go? I know I'm not doing this optimally at the moment
- Core algorithm efficiency
- How else might the architecture of this project be improved? 