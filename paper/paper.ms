.TL
Little Green Clusters
.AU
Thijs Haker
.AI
Fontys Hogescholen
.AB
(...)
.AE
.NH
Introduction:
.PP
In our current world,
there is an abundance of both communication,
and computation.
It has supercharged both the speed,
and the amount of information we can send and receive.
Even more miraculous,
we can do this with machines smaller then textbooks.
Yet most processing and storage of this information happens at large,
environment unfriendly,
datacenters.
.PP
Would it not be possible?
To use this abundance of communication to transform a set of loosely connected machines;
into a functional,
efficient,
.I distributed
system.
.NH 2
Questions:
.PP
The primairy question for this research is:
How do distributed systems generally work?
The secondary question is:
Does the communication involved,
diminish the returns of a distributed system?
The tertiary,
and applicable question:
Can a distributed system be utilized effectively,
to process complex data (more specifically: Analyse DNA)?
.NH 2
Relevancy:
.PP
The relevancy of this research is twofold.
First,
there is climate change,
and the associated need to increase efficiency,
and decrease environmental impact.
.[
datacenter
.]
Second,
the increase in cancer rates among people under 50,
.[
health
.]
increases the need for both better tooling,
and better treatment.
.PP
My conjecture for this research,
is that both these problems can benefit from distributed systems.
.NH
Methodology:
.PP
To answer the three research questions,
and possibly solve the two relevancy problems.
I will start with the gathering of theoretical models for distributed systems.
These models will be sourced,
with a preference for known publishers like:
Academia,
ACM,
Elsevier,
Jstor,
and Google Scholar.
With these models I will answer the first question.
Then I will gather information about implementations that follow these models,
this will be done with the second,
and third question in mind.
After which I will compile the common atributes between the implementations,
and use them as requirements for my system.
Finally I will attempt to implement a distributed system, with 
.B Relevancy
in mind.
.NH 2
Literature:
.PP
To understand distributed systems,
we first have to look at
.I "concurrency."
.[
ghosh
.]
Concurrency is the composition of independently executing computations,
.[
pike
.]
which is not the same as parallelism.
Altough these computations may,
or may not be executing on the same machine.
In both cases they need to solve the same problem.
This is Dijkstra's
.I "Dining philosophers"
problem.
So to have a distributed system,
is to have a concurrent system,
and a concurrent system is to have a model,
.[
ewd68
.]
that is effective at solving the dining philosophers problem.
.[
ewd65
.]
.PP
There are many such models,
but only two are widely used,
and therefore,
I will only take a deeper look at those two.
These models are:
.I "The Actor model"
and,
the model of
.I "Communicating sequential processes."
.SH
The Actor Model:
.PP
The actor model
.[
actor
.]
is a model for concurrency that treats an
.I actor
as the universal primitive.
An actor can do several (concurrent) things when respoding to a message from other actors:
.IP \[bu]
Send a finite number of messages to other actors.
.IP \[bu]
Create a finite number of new actors.
.IP \[bu]
Determine the behavior to be used for the next message it receives.
.PP
Having the actor as universal primitive makes it suitable for usage in conventional object oriented languages,
and makes sense when solving the dining philosophers problem (philosophers are also actors).
However,
the actor model does not make a destinction between actors as processes,
and actors as messages (everything is an actor).
This increases the complexity of any implementation using this model significantly.
.SH
Communicating Sequential Processes:
.PP
CSP is a formal model/language 
.[
csp
.]
for interactions between independent processes.
It is based on Dijkstra's guarded commands.
.[
ewd75
.]
The primitives of this model are:
Values,
processes,
and channels.
Values are as-is.
Processes take in values,
modify them,
and put out modfied values (like mathmatical functions).
Channels synchronise the flow of values from one process to the other.
.PP
With this model it is possible to construct any concurrent system,
whether executed in parallel,
or not.
Because of it's three distinctive primitives,
it maps very well to any type of programming language,
and can even be used in other mathmatical fields.
It is also possible to solve the dining philosophers problem in numerous ways,
and implementations can be realised with relatively little effort.
.NH 2
Implementations:
.PP

.NH 2
Requirements:
.PP
(...)
.NH
Results:
.PP
(...)
.NH 2
Validation:
.PP
(...)
.NH
Conclusion:
.PP
(...)
.NH 2
Recommendation:
.PP
(...)
