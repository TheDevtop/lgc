.TL
Little Green Clusters (Draft)
.AU
Thijs Haker
.AI
Fontys Hogescholen
.AB
Is it possible to create a lightweight distributed system?
Given the literature of contemporary systems.
Building a small,
eco-friendly system,
capable of analysing DNA.
Showing to the world the elegance,
and potential of distributed computing.
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
Then I will gather information about implementations that follow these models;
this will be done with the second,
and third question in mind.
After which I will compile the common attributes between the implementations,
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
pike12
.]
which is not the same as parallelism.
Altough these computations may,
or may not be executing on the same machine.
In both cases they need to solve the same problem.
This is Dijkstra's
.I "Dining philosophers"
problem.
So to have a distributed system,
is to have a concurrent system (spanning different machines),
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
.NH 3
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
.NH 3
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
and put out modified values (like mathmatical functions).
Channels synchronise the flow of values from one process to the other,
they behave similair to pipes
.[
unix
.]
altough multiplexed.
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
There are many implementations of both these models.
Some are found in programming languages,
some in operating systems,
and others in;
databases,
orchestrators,
and filesystems.
For each of these categories I will take a look at a well-known implementation,
and take their common attributes as the basis for the requirements.
.NH 3
Go (Programming Language):
.PP
Go is a programming language
.[
golang
.]
that implements CSP as it's concurrency model.
Combined with it's networking libraries,
creates the possiblity for both concurrency interally (inside the program) via channels,
and externally (spanning multiple programs, or multiple machines) via networking.
The ability of channels to rendezvous synchronously,
makes it possible for any set of networked programs to be (re)written as a distributed system.
.NH 3
Erlang (Programming Language):
.PP
Erlang is a functional programming language
.[
erlang
.]
that implements both CSP,
and Actor style concurrency.
It's build with fault-tolerance,
and real-time programming in mind.
However,
Erlang programs do require a virtual machine to function,
and this introduces some overhead.
Altough this may be negligible in the case of the Erlang Virtual Machine.
.NH 3
Plan 9 (Operating System):
.PP
Plan 9 is a distributed operating system.
That consists of a small kernel,
a small collection of kernel-space device servers,
.[
plan9dev
.]
and a larger collection of user-space file servers that implement the 9P protocol.
.[
plan9arch
.]
Tho guarantee some form of consistency,
every process has it's own namespace,
.[
plan9ns
.]
from which it interacts with the rest of the system.
This makes it so that applications don't have to know whether an operating system service comes from the local machine,
or from a remote machine.
Within the Plan 9 kernel,
CSP is implemented,
altough masqueraded.
In Plan 9 processes are masked as memmory mapped files,
while channels are masked as duplex pipes.
Both can be viewed,
and interacted with via the namespace.
.PP
Finally,
the Plan 9 kernel is relatively high performing,
and is very lightweight.
It is even possible to implement only the device servers,
to connect your regular machine to an actual Plan 9 system.
Such is the case with Drawterm.
.[
drawterm
.]
.NH 3
CouchDB (Database):
.PP
Apache CouchDB is a distributed document-oriented database.
It implements the same concurrency model as its implementation language (Erlang).
CouchDB is also fault-tolerant,
it archieves this by sharding
.[
shard
.]
database entries.
Being a database makes CouchDB great for storing and retrieving date,
but worthless at performing computation on that data.
To perform computation on this data,
it is imperative that another distributed system is introduced.
Increasing the total complexity.
.NH 3
OpenStack (Orchestrator):
.PP
OpenStack is considered the standard implementation of a cloud-computing platform.
It is not a distributed system in the classical sense,
in that it doesn't provide a single interface from which to interact.
On the contrary,
each module on OpenStack provides its own (standardized) API.
The architecture of OpenStack
.[
openstack
.]
does provide a good overview of the functional components of a distributed system.
For example:
Nova and Zun implement the compute functionality,
Swift and Cinder implement storage functionality,
and Neutron and Octavia provide networking.
These are the same functional components
.I "(network, storage, and compute)"
found in other distributed systems.
.NH 3
Ceph (Filesystem):
.PP
Ceph a distributed file system with focus towards performance,
reliability,
and scalability.
.[
ceph
.]
It provides both services for block,
file,
and object storage.
Altough it isn't clearly stated,
when looking at the source code I get the impression that the Actor model is used.
This is because there messaging is asynchronous,
and they have a dispatcher class which behaves similair to a mailbox actor.
.PP
Because of the amount of features that Ceph offers,
the codebase is huge.
Add the fact that Ceph is a storage solution,
which needs another distributed system for computation,
increases the total complexity again.
.NH 2
Requirements:
.PP
Given the common attributes of both theoretical models,
and practical implementations.
Little Green Clusters (LGC) should achieve the following requirements:
.IP \1
The model of concurrency should both be correct,
and simple to implement.
Therefore LGC shall be implemented using CSP.
.IP \2
LGC should be fault-tolerant.
.IP \3
LGC should present either a single interface,
or a consistent set of interfaces.
.IP \4
LGC should implement,
or at least consider,
all the functional components (network, storage, and compute).
.IP \5
LGC should integrate with existing software,
where possible.
.NH
Results:
.PP
(...)
.NH 2
Design:
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
.SH
Acknowledgments
.PP
(...)